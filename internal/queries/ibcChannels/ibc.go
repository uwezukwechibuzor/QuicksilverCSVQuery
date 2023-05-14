package queries

type Channel struct {
	State       string `json:"state"`
	Order       string `json:"order"`
	Connection  string `json:"connection"`
	ChannelID   string `json:"channel_id"`
	Counterpart struct {
		PortID    string `json:"port_id"`
		ChannelID string `json:"channel_id"`
	} `json:"counterparty"`
}

type ClientState struct {
	ClientID    string `json:"client_id"`
	ClientType  string `json:"client_type"`
	ChainID     string `json:"chain_id"`
	LatestBlock struct {
		Height int64  `json:"height"`
		Hash   string `json:"hash"`
	} `json:"latest_block"`
}

type IBCChannel struct {
	Channel      Channel     `json:"channel"`
	ClientState  ClientState `json:"client_state"`
	ConnectionID string      `json:"connection_id"`
}

type IBCChannelsResponse struct {
	Channels []IBCChannel `json:"channels"`
}
