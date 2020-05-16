package cli

import (
	"fmt"

	"github.com/celer-network/sgn/common"
	"github.com/celer-network/sgn/x/global/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagEpochId = "epochId"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	globalQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the global module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	globalQueryCmd.AddCommand(flags.GetCommands(
		GetCmdEpoch(storeKey, cdc),
		GetCmdQueryParams(storeKey, cdc),
	)...)
	return globalQueryCmd
}

// GetCmdEpoch queries request info
func GetCmdEpoch(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "epoch",
		Short: "query epoch info by epochId",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			epochId := viper.GetInt64(flagEpochId)
			data, err := cdc.MarshalJSON(types.NewQueryEpochParams(epochId))
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			route := fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryEpoch)
			bz, _, err := cliCtx.QueryWithData(route, data)
			if err != nil {
				return err
			}

			var epoch types.Epoch
			cdc.MustUnmarshalJSON(bz, &epoch)
			return cliCtx.PrintOutput(epoch)
		},
	}

	cmd.Flags().Int64(flagEpochId, 0, "Epoch id")

	return cmd
}

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current global parameters information",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			params, err := QueryParams(cliCtx, queryRoute)
			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(params)
		},
	}
}

func QueryParams(cliCtx context.CLIContext, queryRoute string) (params types.Params, err error) {
	route := fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryParameters)
	res, err := common.RobustQuery(cliCtx, route)
	if err != nil {
		return
	}

	err = cliCtx.Codec.UnmarshalJSON(res, &params)
	return
}
