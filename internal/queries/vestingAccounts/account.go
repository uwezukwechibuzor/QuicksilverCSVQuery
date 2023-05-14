package queries

type BaseAccount struct {
	Address       string  `json:"address"`
	PubKey        *PubKey `json:"pub_key"`
	AccountNumber string  `json:"account_number"`
	Sequence      string  `json:"sequence"`
}

type PubKey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type VestingAccount struct {
	Type               string             `json:"@type"`
	BaseVestingAccount BaseVestingAccount `json:"base_vesting_account"`
	StartTime          string             `json:"start_time"`
	VestingPeriods     []VestingPeriod    `json:"vesting_periods"`
}

type BaseVestingAccount struct {
	BaseAccount      BaseAccount `json:"base_account"`
	OriginalVesting  []Coin      `json:"original_vesting"`
	DelegatedFree    []Coin      `json:"delegated_free"`
	DelegatedVesting []Coin      `json:"delegated_vesting"`
	EndTime          string      `json:"end_time"`
}

type VestingPeriod struct {
	Length string `json:"length"`
	Amount []Coin `json:"amount"`
}

type Coin struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type AccountsResponse struct {
	Accounts []VestingAccount `json:"accounts"`
}
