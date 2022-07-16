package entity

type BarkDevice struct {
	DeviceKey   string `db:"device_key" redis:"device_key"`
	DeviceToken string `db:"device_token" redis:"device_token"`
	Name        string `db:"name"`
}

type RegInfo struct {
	Token    string `form:"token" validate:"uuid4"`
	Name     string `form:"name" validate:"required"`
	DeviceID string `form:"device_id" validate:"required"`
	IsClip   int    `form:"is_clip"`
	Type     string `form:"Type"`
}

type APNsMessage struct {
	DeviceToken string
	Category    string            `json:"category"`
	Title       string            `json:"title"`
	Content     string            `json:"content"` // Title, message, clipBoard
	Sound       string            `json:"sound"`
	Params      map[string]string `json:"ext_params"` // url
}

type HistoryParam struct {
	Offset int `json:"offset" form:"offset"`
	Limit  int `json:"limit" form:"limit"`
}
