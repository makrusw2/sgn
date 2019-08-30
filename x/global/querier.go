package global

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryLatestBlock:
			return queryLatestBlock(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("Unknown global query endpoint")
		}
	}
}

func queryLatestBlock(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	latestBlock := keeper.GetLatestBlock(ctx)
	res, err := codec.MarshalJSONIndent(keeper.cdc, latestBlock)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("Could not marshal result to JSON", err.Error()))

	}

	return res, nil
}