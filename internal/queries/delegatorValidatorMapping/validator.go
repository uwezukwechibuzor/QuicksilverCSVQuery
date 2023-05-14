package queries

type Validator struct {
	OperatorAddress   string          `json:"operator_address"`
	ConsensusPubKey   ConsensusPubKey `json:"consensus_pubkey"`
	Jailed            bool            `json:"jailed"`
	Status            string          `json:"status"`
	Tokens            string          `json:"tokens"`
	DelegatorShares   string          `json:"delegator_shares"`
	Description       Description     `json:"description"`
	UnbondingHeight   string          `json:"unbonding_height"`
	UnbondingTime     string          `json:"unbonding_time"`
	Commission        Commission      `json:"commission"`
	MinSelfDelegation string          `json:"min_self_delegation"`
}

type ConsensusPubKey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type Commission struct {
	CommissionRates CommissionRates `json:"commission_rates"`
	UpdateTime      string          `json:"update_time"`
}

type CommissionRates struct {
	Rate          string `json:"rate"`
	MaxRate       string `json:"max_rate"`
	MaxChangeRate string `json:"max_change_rate"`
}

type Description struct {
	Moniker         string `json:"moniker"`
	Identity        string `json:"identity"`
	Website         string `json:"website"`
	SecurityContact string `json:"security_contact"`
	Details         string `json:"details"`
}
