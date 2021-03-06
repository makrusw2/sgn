package singlenode

import (
	"math/big"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn/common"
	tc "github.com/celer-network/sgn/testing/common"
	govtypes "github.com/celer-network/sgn/x/gov/types"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setUpUpgrade() []tc.Killable {
	res := setupNewSGNEnv(nil, "upgrade proposal")
	tc.SleepWithLog(10, "sgn syncing")

	return res
}

func TestE2EUpgrade(t *testing.T) {
	toKill := setUpUpgrade()
	defer tc.TearDown(toKill)

	t.Run("e2e-upgrade", func(t *testing.T) {
		t.Run("upgradeTest", upgradeTest)
	})
}

func upgradeTest(t *testing.T) {
	log.Info("=====================================================================")
	log.Info("======================== Test upgrade ===========================")

	transactor := tc.NewTestTransactor(
		t,
		CLIHome,
		viper.GetString(common.FlagSgnChainID),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetString(common.FlagSgnValidatorAccount),
		viper.GetString(common.FlagSgnPassphrase),
	)

	amt := big.NewInt(5000000000000000000)
	ethAddr, auth, err := tc.GetAuth(tc.ValEthKs[0])
	log.Infof("my eth address %x", ethAddr)
	require.NoError(t, err, "failed to get auth")
	tc.AddCandidateWithStake(t, transactor, ethAddr, auth, tc.ValAccounts[0], amt, big.NewInt(1), big.NewInt(1), big.NewInt(10000), true)
	tc.CheckValidatorNum(t, transactor, 1)

	upgradeHeight := int64(50)
	plan := upgrade.Plan{Name: "test", Height: upgradeHeight}
	content := govtypes.NewUpgradeProposal("Upgrade test", "Upgrade test", plan)
	submitProposalmsg := govtypes.NewMsgSubmitProposal(content, sdk.NewInt(2), transactor.Key.GetAddress())
	transactor.AddTxMsg(submitProposalmsg)

	proposalID := uint64(1)
	byteVoteOption, _ := govtypes.VoteOptionFromString("Yes")
	voteMsg := govtypes.NewMsgVote(transactor.Key.GetAddress(), proposalID, byteVoteOption)
	transactor.AddTxMsg(voteMsg)

	height, err := rpc.GetChainHeight(transactor.CliCtx)
	for err == nil {
		time.Sleep(viper.GetDuration(common.FlagSgnTimeoutCommit) * time.Second)
		prevHeight := height
		height, err = rpc.GetChainHeight(transactor.CliCtx)

		if prevHeight == height {
			break
		}
	}

	require.NoError(t, err, "failed to query block height")
	assert.Equal(t, height, upgradeHeight, "The chain should stop at upgrade height")
}
