package singlenode

import (
	"math/big"
	"testing"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn/common"
	tc "github.com/celer-network/sgn/testing/common"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func setupValidator() []tc.Killable {
	p := &tc.SGNParams{
		CelrAddr:               tc.E2eProfile.CelrAddr,
		GovernProposalDeposit:  big.NewInt(1), // TODO: use a more practical value
		GovernVoteTimeout:      big.NewInt(1), // TODO: use a more practical value
		SlashTimeout:           big.NewInt(10),
		MinValidatorNum:        big.NewInt(1),
		MaxValidatorNum:        big.NewInt(11),
		MinStakingPool:         big.NewInt(1),
		AdvanceNoticePeriod:   big.NewInt(1), // TODO: use a more practical value
		SidechainGoLiveTimeout: big.NewInt(0),
	}
	res := setupNewSGNEnv(p, "validator")
	tc.SleepWithLog(10, "sgn being ready")

	return res
}

func TestE2EValidator(t *testing.T) {
	toKill := setupValidator()
	defer tc.TearDown(toKill)

	t.Run("e2e-validator", func(t *testing.T) {
		t.Run("validatorTest", validatorTest)
	})
}

func validatorTest(t *testing.T) {
	log.Info("===================================================================")
	log.Info("======================== Test validator ===========================")

	transactor := tc.NewTestTransactor(
		t,
		CLIHome,
		viper.GetString(common.FlagSgnChainID),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetStringSlice(common.FlagSgnTransactors)[0],
		viper.GetString(common.FlagSgnPassphrase),
	)
	amt := big.NewInt(1000000000000000000)

	ethAddr, auth, err := tc.GetAuth(tc.ValEthKs[0])
	log.Infof("my eth address %x", ethAddr)
	require.NoError(t, err, "failed to get auth")
	tc.AddCandidateWithStake(t, transactor, ethAddr, auth, tc.ValAccounts[0], amt, big.NewInt(1), big.NewInt(1), big.NewInt(10000), true)
	tc.CheckValidatorNum(t, transactor, 1)
}
