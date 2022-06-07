package body

type SessionResponse struct {
	ChannelId    uint64 `json:"channel_id"`
	ChannelToken string `json:"channel_token"`
	Salt         string `json:"salt"`
	WebsocketUrl string `json:"websocket_url"`
	BaseUrl      string `json:"base_url"`
	DohUrl       string `json:"doh_url"`
}
