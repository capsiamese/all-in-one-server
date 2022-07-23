package pb

import (
	"aio/ent"
	"google.golang.org/protobuf/encoding/prototext"
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
	tabs := make([]*Tab, 0, len(eg.Edges.Tabs))
	for _, t := range eg.Edges.Tabs {
		tabs = append(tabs, TabToPB(t))
	}
	g := &Group{
		Name:   eg.Name,
		Uid:    eg.UID.String(),
		Index:  eg.Seq,
		Option: new(GroupOption),
	}
	_ = prototext.Unmarshal([]byte(eg.Option), g.Option)
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

func AddressToPB(ea *ent.BarkAddress) *BarkDevice {
	if ea == nil {
		return nil
	}
	return &BarkDevice{
		Name: ea.Name,
		Url:  ea.Target,
		Id:   ea.Index,
	}
}
