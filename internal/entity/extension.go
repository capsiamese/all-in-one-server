package entity

import uuid "github.com/satori/go.uuid"

type GroupOption struct {
	Tags []string `json:"tags"`
}

type GroupInfo struct {
	Name   string       `json:"name,omitempty"`
	Tabs   []*Tab       `json:"tabs"`
	Option *GroupOption `json:"option,omitempty"`
	Uid    uuid.UUID    `json:"uid"`
	Index  int32        `json:"index"`
}

type Tab struct {
	Name    string    `json:"name"`
	Url     string    `json:"url"`
	Favicon string    `json:"favicon"`
	Uid     uuid.UUID `json:"uid"`
	Index   int32     `json:"index"`
}
