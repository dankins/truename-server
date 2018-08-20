package eth_helpers

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

type Helper struct {
	Blockchain *backends.SimulatedBackend
	Key        *ecdsa.PrivateKey
	Auth       *bind.TransactOpts
	Accounts   map[string]Account
}

type Account struct {
	Key     *ecdsa.PrivateKey
	Auth    *bind.TransactOpts
	Address common.Address
}

func NewHelper() Helper {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc)

	accounts := make(map[string]Account)
	accounts["genesis"] = Account{Key: key, Auth: auth}

	return Helper{
		Blockchain: blockchain,
		Key:        key,
		Auth:       auth,
		Accounts:   accounts,
	}
}

func (h *Helper) Transact() *bind.TransactOpts {
	return &bind.TransactOpts{
		From:   h.Auth.From,
		Signer: h.Auth.Signer,
		Value:  nil,
	}
}

func (h *Helper) TransactWithGasLimit() *bind.TransactOpts {
	return &bind.TransactOpts{
		From:     h.Auth.From,
		Signer:   h.Auth.Signer,
		Value:    nil,
		GasLimit: 100000,
	}
}

func (h *Helper) Call() *bind.CallOpts {
	return &bind.CallOpts{}
}

func (h *Helper) AddAccount(name string) Account {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)

	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}

	h.Accounts[name] = Account{
		Key:     key,
		Auth:    auth,
		Address: crypto.PubkeyToAddress(key.PublicKey),
	}

	return h.Accounts[name]
}
