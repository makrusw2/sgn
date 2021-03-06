package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Syncer struct {
	ValidatorIdx  uint           `json:"validator_idx"`
	ValidatorAddr sdk.AccAddress `json:"validator_addr"`
}

func NewSyncer(validatorIdx uint, validatorAddr sdk.AccAddress) Syncer {
	return Syncer{
		ValidatorIdx:  validatorIdx,
		ValidatorAddr: validatorAddr,
	}
}

// implement fmt.Stringer
func (r Syncer) String() string {
	return strings.TrimSpace(fmt.Sprintf(`ValidatorIdx: %d, ValidatorAddr: %x`, r.ValidatorIdx, r.ValidatorAddr))
}
