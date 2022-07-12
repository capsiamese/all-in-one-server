package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"notification/ent"
	"notification/ent/extensionclient"
	"notification/ent/group"
	"notification/ent/tab"
	"notification/internal/pb"
	"notification/pkg/logger"
	"time"
)

var _ Extension = (*ExtensionUseCase)(nil)

type ExtensionUseCase struct {
	cli *ent.Client
	log logger.Interface
}

func NewExtension(cli *ent.Client, l logger.Interface) *ExtensionUseCase {
	return &ExtensionUseCase{
		cli: cli,
		log: l,
	}
}

func (e *ExtensionUseCase) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := e.cli.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			err := tx.Rollback()
			if err != nil {
				e.log.Errorf("panic rolling back transaction: %w\n", err)
			}
			e.log.Error(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			e.log.Errorln("transaction failed: ", err)
			err = fmt.Errorf("rolling back transaction: %w", rErr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func queryExtensionClient(ctx context.Context, client *ent.Client, name, extensionID string) (*ent.ExtensionClient, error) {
	return client.ExtensionClient.Query().
		Where(extensionclient.And(
			extensionclient.NameEQ(name),
			extensionclient.ExtensionIDEQ(extensionID),
		)).Only(ctx)
}

func createExtensionClient(ctx context.Context, client *ent.Client, name, extensionID string) (*ent.ExtensionClient, error) {
	return client.ExtensionClient.
		Create().
		SetName(name).
		SetExtensionID(extensionID).
		SetClientUID(uuid.NewV4()).
		SetLastAccessTime(time.Now()).
		Save(ctx)
}

func (e *ExtensionUseCase) Register(ctx context.Context, name, extensionID string) (*ent.ExtensionClient, error) {
	ext, err := queryExtensionClient(ctx, e.cli, name, extensionID)
	if err == nil {
		return ext, nil
	} else {
		e.log.Warningln("query client:%v", err)
	}
	return createExtensionClient(ctx, e.cli, name, extensionID)
}

func (e *ExtensionUseCase) Connect(ctx context.Context, uid uuid.UUID, wsConn *websocket.Conn) error {
	return errors.New("")
}

func createGroup(ctx context.Context, client *ent.Client, ext *ent.ExtensionClient, group *pb.Group) (*ent.Group, error) {
	cmd := client.Group.Create().
		SetUID(uuid.NewV4()).
		SetCreatedAt(time.Now()).
		SetClient(ext).
		SetName(group.GetName())
	if group.GetOption() != nil {
		cmd.SetOption(*group.GetOption())
	}
	return cmd.Save(ctx)
}

func createTab(ctx context.Context, client *ent.Client, group *ent.Group, tab *pb.Tab, index int) (*ent.Tab, error) {
	return client.Tab.Create().
		SetName(tab.GetName()).
		SetURL(tab.GetUrl()).
		SetSeq(int32(index + 1)).
		SetFavicon(tab.GetFavicon()).
		SetGroup(group).
		SetUID(uuid.NewV4()).
		Save(ctx)
}

func (e *ExtensionUseCase) Add(ctx context.Context, uid uuid.UUID, group ...*pb.Group) error {
	return e.WithTx(ctx, func(tx *ent.Tx) error {
		var err error
		var client *ent.ExtensionClient
		client, err = tx.ExtensionClient.
			Query().
			Where(extensionclient.ClientUIDEQ(uid)).
			Only(ctx)
		if err != nil {
			return err
		}
		var eg *ent.Group
		for _, g := range group {
			eg, err = createGroup(ctx, tx.Client(), client, g)
			if err != nil {
				return err
			}
			for i, t := range g.GetTabs() {
				_, err = createTab(ctx, tx.Client(), eg, t, i)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (e *ExtensionUseCase) Pull(ctx context.Context, uid uuid.UUID) (*ent.ExtensionClient, error) {
	return e.cli.ExtensionClient.
		Query().
		Where(
			extensionclient.ClientUIDEQ(uid),
		).
		WithGroups(
			func(query *ent.GroupQuery) {
				query.WithTabs()
			},
		).
		Only(ctx)
}

func (e *ExtensionUseCase) RemoveGroup(ctx context.Context, uid, groupUid uuid.UUID) error {
	var groupInfo *ent.Group // return ?
	var tabs []*ent.Tab      // return ?
	return e.WithTx(ctx, func(tx *ent.Tx) error {
		client, err := tx.ExtensionClient.Query().Where(extensionclient.ClientUIDEQ(uid)).Only(ctx)
		if err != nil {
			return err
		}
		groupInfo, err = client.
			QueryGroups().
			Where(group.UIDEQ(groupUid)).
			WithTabs().
			Only(ctx)
		if err != nil {
			return err
		}

		tabs, err = groupInfo.Edges.TabsOrErr()
		if err != nil {
			return err
		}
		err = groupInfo.Update().RemoveTabs(tabs...).Exec(ctx)
		if err != nil {
			return err
		}
		err = tx.Group.DeleteOne(groupInfo).Exec(ctx)
		if err != nil {
			return err
		}
		e.log.Infof("delete group uid:%s groupData:%s\n", uid, groupInfo)
		e.log.Infof("delete group tabs uid:%s group:%s tabs:%v\n", uid, groupUid, tabs)
		return nil
	})
}

func (e *ExtensionUseCase) RemoveTab(ctx context.Context, uid, groupUid, tabUid uuid.UUID) error {
	t, err := e.cli.ExtensionClient.
		Query().
		Where(extensionclient.ClientUIDEQ(uid)).
		QueryGroups().Where(group.UIDEQ(groupUid)).
		QueryTabs().Where(tab.UIDEQ(tabUid)).Only(ctx)
	if err != nil {
		return err
	}
	err = e.cli.Tab.DeleteOne(t).Exec(ctx)
	if err != nil {
		return err
	}
	e.log.Infof("delete tab uid:%s group:%s tabData:%v\n", uid, groupUid, t.String())
	return nil
}

func (e *ExtensionUseCase) SwapTab(ctx context.Context, uid, firstGroupUid, firstGroupTabUid, secondGroupUid, secondGroupTabUid uuid.UUID) error {
	if firstGroupTabUid == secondGroupTabUid {
		return nil
	}
	return e.WithTx(ctx, func(tx *ent.Tx) error {
		tabs, err := tx.ExtensionClient.
			Query().
			Where(extensionclient.ClientUIDEQ(uid)).
			QueryGroups().
			Where(group.Or(group.UIDEQ(firstGroupUid), group.UIDEQ(secondGroupUid))).
			QueryTabs().
			Where(tab.Or(tab.UIDEQ(firstGroupTabUid), tab.UIDEQ(secondGroupTabUid))).All(ctx)
		if err != nil {
			return err
		}
		if len(tabs) != 2 {
			return nil
		}
		first, second := tabs[0], tabs[1]
		fg, err := tx.Tab.QueryGroup(first).Only(ctx)
		if err != nil {
			return err
		}
		sg, err := tx.Tab.QueryGroup(second).Only(ctx)
		if err != nil {
			return err
		}
		err = first.Update().SetSeq(second.Seq).SetGroup(sg).Exec(ctx)
		if err != nil {
			return err
		}
		err = second.Update().SetSeq(first.Seq).SetGroup(fg).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}
