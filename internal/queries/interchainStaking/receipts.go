package queries

type Receipts struct {
	ChainID string `json:"chain_id"`
	Sender  string `json:"sender"`
	TxHash  string `json:"txhash"`
	Amount  []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"amount"`
	FirstSeen string `json:"first_seen"`
	Completed string `json:"completed"`
}

type ReceiptsResponse struct {
	Receipts []Receipts `json:"receipts"`
}
