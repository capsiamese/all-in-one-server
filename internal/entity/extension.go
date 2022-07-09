package entity

type GroupOption struct {
	Tags []string `json:"tags"`
}

type GroupInfo struct {
	Name   string       `json:"name,omitempty"`
	Tabs   []*Tab       `json:"tabs"`
	Option *GroupOption `json:"option,omitempty"`
}

type Tab struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Favicon string `json:"favicon"`
}
