package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Subscription struct {
	EthAddress   string  `json:"ethAddress"`
	Deposit      sdk.Int `json:"deposit"`
	Spend        sdk.Int `json:"spend"`
	Subscribing  bool    `json:"subscribing"`
	RequestCount uint64  `json:"requestCount"`
}

func NewSubscription(ethAddress string) Subscription {
	return Subscription{
		EthAddress:   ethAddress,
		Deposit:      sdk.NewInt(0),
		Spend:        sdk.NewInt(0),
		Subscribing:  false,
		RequestCount: 0,
	}
}

// implement fmt.Stringer
func (s Subscription) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Deposit: %v, Spend: %v, Subscribing: %b, RequestCount: %d`, s.Deposit, s.Spend, s.Subscribing, s.RequestCount))
}
