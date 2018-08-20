package identity

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/dankins/truename-server/pkg/generated/contracts"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IdentityService struct {
	Blockchain *ethclient.Client
	Auth       *bind.TransactOpts
	DGraph     *dgo.Dgraph
}

type ContractInstance struct {
	Uid     string `json:"uid,omitempty"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

func (svc *IdentityService) GetIdentityFactory() ContractInstance {
	q := `{
		me(func: eq(name, "IdentityFactory")) {
			name
			address
		}
	}`

	resp, err := svc.DGraph.NewTxn().Query(context.Background(), q)
	if err != nil {
		log.Fatal(err)
	}

	type Root struct {
		Me []ContractInstance `json:"me"`
	}

	var r Root
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GetIdentityFactory", r.Me)
	if len(r.Me) == 0 {
		return ContractInstance{}
	}

	return r.Me[0]
}

func (svc *IdentityService) DeployIdentityFactory() common.Address {

	address, _, _, _ := contracts.DeployIdentityFactory(svc.Auth, svc.Blockchain)

	op := &api.Operation{}
	op.Schema = `
		name: string @index(exact) .
		address: string .
	`

	contract := ContractInstance{
		Name:    "IdentityFactory",
		Address: address.String(),
	}

	ctx := context.Background()
	err := svc.DGraph.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(contract)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb
	_, err = svc.DGraph.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	return address
}

type IdentityInput struct {
	Account string `json:"account,omitempty"`
	PubKey  string `json:"pubKey,omitempty"`
}

type RelayTransactionInput struct {
	IdentityContractAddress string `json:"identityContract,omitempty"`
	AuthorizedAccount       string `json:"authorizedAccount,omitempty"`
	Gas                     string `json:"gas,omitempty"`
	MessageHash             string `json:"messageHash,omitempty"`
	SignedHash              string `json:"signedHash,omitempty"`
}

func (svc *IdentityService) DeployIdentity(input *IdentityInput) ContractInstance {
	factory := svc.GetIdentityFactory()

	identityFactory, err := contracts.NewIdentityFactory(common.HexToAddress(factory.Address), svc.Blockchain)
	if err != nil || identityFactory == nil {
		log.Fatalf("Could not connect to IdentityFactory contract: %v", err)
	}

	var trans = &bind.TransactOpts{
		From:   svc.Auth.From,
		Signer: svc.Auth.Signer,
		Value:  nil,
	}
	var _pubKey = [32]byte{}
	var _account = common.HexToAddress(input.Account)
	txres, err := identityFactory.CreateIdentity(trans, _account, _pubKey)
	if err != nil {
		log.Fatal(err)
	}

	identityContract := ContractInstance{
		Name:    "Identity",
		Address: txres.Hash().String(),
	}

	return identityContract
}

func (svc *IdentityService) GetIdentityForAccount(account string) string {
	factory := svc.GetIdentityFactory()

	identityFactory, err := contracts.NewIdentityFactory(common.HexToAddress(factory.Address), svc.Blockchain)
	if err != nil || identityFactory == nil {
		log.Fatalf("Could not connect to IdentityFactory contract: %v", err)
	}

	var _account = common.HexToAddress(account)
	identityContract, err := identityFactory.GetIdentityForAccount(&bind.CallOpts{}, _account)

	return identityContract.String()
}

func (svc *IdentityService) GetIdentityABI() string {
	return contracts.IdentityABI
}

func (svc *IdentityService) GetIdentityFactoryABI() string {
	return contracts.IdentityFactoryABI
}

func (svc *IdentityService) RelayTransaction(input *RelayTransactionInput) bool {

	identityContract, err := contracts.NewIdentity(common.HexToAddress(input.IdentityContractAddress), svc.Blockchain)
	if err != nil || identityContract == nil {
		log.Fatalf("Could not connect to IdentityFactory contract: %v", err)
	}

	var trans = &bind.TransactOpts{
		From:     svc.Auth.From,
		Signer:   svc.Auth.Signer,
		Value:    nil,
		GasLimit: 6721975,
	}

	gas := new(big.Int)
	gas, ok := gas.SetString(input.Gas, 10)
	if ok == false {
		log.Fatalf("Failed to call gas.SetString")
	}

	fmt.Println("ready to do it", input.AuthorizedAccount, input.MessageHash, input.SignedHash)
	tx, err := identityContract.EIP1077Request(
		trans,
		common.HexToAddress(input.AuthorizedAccount),
		big.NewInt(11111),
		gas,
		stringToBytes32(input.MessageHash),
		[]byte(input.SignedHash),
	)

	if err != nil || identityContract == nil {
		log.Fatalf("Could not relay transaction", err)
	}

	log.Println("tx: ", tx)

	return true
}

func stringToBytes32(str string) [32]byte {
	var b [32]byte
	copy(b[:], str)

	return b
}
