package pb

import (
	"google.golang.org/protobuf/proto"
	"notification/ent"
)

func ClientToPB(ec *ent.ExtensionClient) *Client {
	if ec == nil {
		return nil
	}
	groups := make([]*Group, 0, len(ec.Edges.Groups))
	for _, g := range ec.Edges.Groups {
		groups = append(groups, GroupToPB(g))
	}
	c := &Client{
		Name:           ec.Name,
		Groups:         groups,
		ExtId:          ec.ExtensionID,
		Uid:            ec.ClientUID.String(),
		LastAccessTime: ec.LastAccessTime.Unix(),
	}
	return c
}

func GroupToPB(eg *ent.Group) *Group {
	if eg == nil {
		return nil
	}
	var op *GroupOption
	if o, ok := proto.Clone(&eg.Option).(*GroupOption); ok {
		op = o
	}
	tabs := make([]*Tab, 0, len(eg.Edges.Tabs))
	for _, t := range eg.Edges.Tabs {
		tabs = append(tabs, TabToPB(t))
	}
	g := &Group{
		Name:   eg.Name,
		Uid:    eg.UID.String(),
		Index:  eg.Seq,
		Option: op,
	}
	return g
}

func TabToPB(et *ent.Tab) *Tab {
	if et == nil {
		return nil
	}
	t := &Tab{
		Name:    et.Name,
		Url:     et.URL,
		Favicon: et.Favicon,
		Uid:     et.UID.String(),
		Index:   et.Seq,
	}
	return t
}
