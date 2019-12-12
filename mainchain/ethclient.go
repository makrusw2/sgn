package mainchain

import (
	"crypto/ecdsa"
	"io/ioutil"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

type EthClient struct {
	// initialized by SetClient()
	Client *ethclient.Client
	// initialized by SetAuth()
	PrivateKey *ecdsa.PrivateKey
	Address    Addr
	Auth       *bind.TransactOpts
	// initialized by SetContracts()
	GuardAddress  Addr
	Guard         *Guard
	LedgerAddress Addr
	Ledger        *CelerLedger
}

// NewEthClient creates a new eth client and initializes all fields
func NewEthClient(ws, guardAddrStr, ledgerAddrStr, ks, passphrase string) (*EthClient, error) {
	ethClient := &EthClient{}
	err := ethClient.SetClient(ws)
	if err != nil {
		return nil, err
	}

	err = ethClient.SetAuth(ks, passphrase)
	if err != nil {
		return nil, err
	}

	err = ethClient.SetContracts(guardAddrStr, ledgerAddrStr)
	if err != nil {
		return nil, err
	}

	return ethClient, nil
}

func (ethClient *EthClient) SetClient(ws string) error {
	rpcClient, err := ethrpc.Dial(ws)
	if err != nil {
		return err
	}

	ethClient.Client = ethclient.NewClient(rpcClient)
	return nil
}

func (ethClient *EthClient) SetAuth(ks, passphrase string) error {
	keystoreBytes, err := ioutil.ReadFile(ks)
	if err != nil {
		return err
	}

	key, err := keystore.DecryptKey(keystoreBytes, passphrase)
	if err != nil {
		return err
	}

	auth, err := bind.NewTransactor(strings.NewReader(string(keystoreBytes)), passphrase)
	if err != nil {
		return err
	}

	ethClient.PrivateKey = key.PrivateKey
	ethClient.Address = key.Address
	ethClient.Auth = auth
	return nil
}

func (ethClient *EthClient) SetContracts(guardAddrStr, ledgerAddrStr string) error {
	ethClient.GuardAddress = Hex2Addr(guardAddrStr)
	guard, err := NewGuard(ethClient.GuardAddress, ethClient.Client)
	if err != nil {
		return err
	}

	ethClient.LedgerAddress = Hex2Addr(ledgerAddrStr)
	ledger, err := NewCelerLedger(ethClient.LedgerAddress, ethClient.Client)
	if err != nil {
		return err
	}

	ethClient.Guard = guard
	ethClient.Ledger = ledger
	return nil
}
