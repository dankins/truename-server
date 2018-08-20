package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/dankins/truename-server/pkg/data"
	"github.com/dankins/truename-server/pkg/graphql"
	"github.com/dankins/truename-server/pkg/identity"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

// put private key into file "foo"
// geth account import foo
// puts a keystore file here: ~/Library/Ethereum/keystore/
const key = `{"address":"398af1d3715bb2dcfc6030cf0333dcb1f904195d","crypto":{"cipher":"aes-128-ctr","ciphertext":"d07bb973e8ff8ccf933cc5ed1cc63b6fc79107ada7e4cea46715deacb97c3c12","cipherparams":{"iv":"ce2b86e5ecd989c4da213a68922745aa"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"3c4e1384436da00b7957dd3310bf6defc3f8ed11e171542aa0d1e98696b77347"},"mac":"5f5f533f82792bd8138f0fdc633576f6b2e1caf05f38f214156a11e7200e537a"},"id":"9c193dff-fc1a-4982-ae7a-30e95b570445","version":3}`

func main() {

	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(key), "")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	var identityService = identity.IdentityService{
		Blockchain: blockchain,
		Auth:       auth,
		DGraph:     data.NewClient(),
	}

	//identityService.DeployIdentityFactory()

	var schema = graphql.BuildSchema(&identityService)

	http.Handle("/graphql", &graphql.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":1337", nil))

}
