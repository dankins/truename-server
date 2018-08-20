package civilians

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dankins/truename-server/pkg/generated/contracts"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CivilianBadgeService struct {
	Blockchain *ethclient.Client
	Auth       *bind.TransactOpts
	DGraph     *dgo.Dgraph
}

type ContractInstance struct {
	Uid     string `json:"uid,omitempty"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

func (svc *CivilianBadgeService) GetCivilianBadgeContract() ContractInstance {
	q := `{
		me(func: eq(name, "CivilianBadge")) {
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

	fmt.Println("GetCivilianBadgeContract", r.Me)
	if len(r.Me) == 0 {
		return ContractInstance{}
	}

	return r.Me[0]
}

func (svc *CivilianBadgeService) DeployCivilBadge() common.Address {

	address, _, _, _ := contracts.DeployCivilianBadge(svc.Auth, svc.Blockchain)

	op := &api.Operation{}
	op.Schema = `
		name: string @index(exact) .
		address: string .
	`

	contract := ContractInstance{
		Name:    "CivilianBadge",
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
