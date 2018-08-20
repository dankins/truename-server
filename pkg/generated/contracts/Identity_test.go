package contracts_test

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/dankins/truename-server/pkg/eth_helpers"
	"github.com/dankins/truename-server/pkg/generated/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func deployVerifier(t *testing.T, helper *eth_helpers.Helper) *contracts.RecoveryVerifier {

	var trans = helper.Transact()
	_, _, verifier, err := contracts.DeployRecoveryVerifier(trans, helper.Blockchain)
	if err != nil {
		t.Fatalf("Failed to Deploy RecoveryVerifier: %v", err)
	}
	helper.Blockchain.Commit()

	return verifier
}

func deployIdentity(t *testing.T, helper *eth_helpers.Helper, _authorizedAccount common.Address) (common.Address, *contracts.Identity) {

	var trans = helper.Transact()
	contractAddress, _, identity, err := contracts.DeployIdentity(trans, helper.Blockchain, _authorizedAccount)
	if err != nil {
		t.Fatalf("Failed to Deploy Identity: %v", err)
	}
	helper.Blockchain.Commit()

	return contractAddress, identity
}

func TestVerifyRecover(t *testing.T) {
	helper := eth_helpers.NewHelper()
	verifier := deployVerifier(t, &helper)

	userB := helper.AddAccount("user_b")

	// have userB sign something
	var _messageHash = crypto.Keccak256Hash([]byte("hello"))
	_signedHash, err := crypto.Sign(_messageHash.Bytes(), userB.Key)

	if err != nil {
		t.Fatal("Failed to Sign!")
	}

	extracted, err := verifier.ExtractAccountNoPrefix(
		&bind.CallOpts{},
		_messageHash,
		_signedHash,
	)
	if err != nil {
		t.Fatalf("Failed to VerifySignature: %v", err)
	}

	if userB.Address.String() != extracted.String() {
		t.Fatal("Failed to VerifySignature!")
	}
}

func TestRelayTransaction(t *testing.T) {
	helper := eth_helpers.NewHelper()
	authorizedUser := helper.AddAccount("authorized_user")
	identityAddress, contract := deployIdentity(t, &helper, authorizedUser.Address)

	nonce, _ := contract.Nonce(&bind.CallOpts{})
	fmt.Println("identityAddress: ", identityAddress.String())
	fmt.Println("identity nonce: ", nonce)

	badgeAddr, _, badgeContract, err := contracts.DeployCivilianBadge(helper.Transact(), helper.Blockchain)
	if err != nil {
		t.Fatal("Failed to Deploy Civilian Badge contract!")
	}
	helper.Blockchain.Commit()

	badgeABI, err := abi.JSON(strings.NewReader(contracts.CivilianBadgeABI))
	_data, _ := badgeABI.Pack("becomeCivilian") //, "<<argument1>>", "<<>argument2>")

	var _messageHash = crypto.Keccak256Hash(_data)
	_signature, err := crypto.Sign(_messageHash.Bytes(), authorizedUser.Key)

	tx, err := contract.RelayTransaction(helper.TransactWithGasLimit(),
		badgeAddr,     // _to
		big.NewInt(0), // _value
		_data,         // _data
		nonce,         // _nonce
		big.NewInt(0), // _gasPrice
		big.NewInt(0), // _gasLimit
		_signature,    // _signature
	)
	if err != nil {
		t.Fatalf("Failed to run RelayTransaction: %v", err)
	}

	fmt.Println("got me a hash:", tx.Hash().String())
	helper.Blockchain.Commit()

	id, err := badgeContract.GetID(&bind.CallOpts{}, identityAddress)
	if err != nil {
		t.Fatalf("Failed to run GetID: %v", err)
	}

	if id.String() != "1" {
		t.Fatal("Expecting ID to be 1 but it is " + id.String())
	}
}

func TestRelayTransactionWithInvalidNonce(t *testing.T) {
	// TODO(dankins): implement
}

// Test inbox contract gets deployed correctly
// func TestRelayTransaction(t *testing.T) {

// 	//Setup simulated block chain
// 	key, _ := crypto.GenerateKey()
// 	auth := bind.NewKeyedTransactor(key)
// 	alloc := make(core.GenesisAlloc)
// 	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
// 	blockchain := backends.NewSimulatedBackend(alloc)

// 	//Deploy IdentityFactory
// 	var trans = &bind.TransactOpts{
// 		From:     auth.From,
// 		Signer:   auth.Signer,
// 		GasLimit: 900000,
// 		Value:    nil,
// 	}
// 	var _pubKey = [32]byte{}

// 	var _authorizedAccount = common.HexToAddress("0x18fe781a5c0d2979de06da6c9aa5f9de58d444a6")
// 	identityAddress, _, identityContract, err := contracts.DeployIdentity(
// 		trans,
// 		blockchain,
// 		_authorizedAccount,
// 		_pubKey,
// 	)
// 	if err != nil {
// 		t.Fatalf("Failed to deploy the IdentityFactory contract: %v", err)
// 	}
// 	if len(identityAddress.Bytes()) == 0 {
// 		t.Error("Expected a valid deployment address. Received empty address byte array instead")
// 	}

// 	fmt.Printf("Deployed IdentityFactory contract: %s\n", identityAddress.String())
// 	// commit all pending transactions
// 	blockchain.Commit()

// 	fmt.Printf("Deployed Identity %s\n", identityAddress.String())

// 	messageHashString, err := hex.DecodeString("15d82131750f2a063e91d10aceffe79a3286a358856655fc488d1bb4ab298572")
// 	var _messageHash [32]byte
// 	copy(_messageHash[:], messageHashString)

// 	_signedHash, err := hex.DecodeString("e1a21d0ad43f8babdb93f208b289d46600ea84f8af93da4aece48e2044cc5c572c1bc634413020c35c2d5aa9db3d0e7f674622445791484df73da14828b56a061c")

// 	fmt.Println("debug _messageHash", _messageHash)
// 	fmt.Println("debug _signedHash", _signedHash)

// 	// tx, err := identityContract.EIP1077Request(
// 	// 	trans,
// 	// 	_authorizedAccount,
// 	// 	big.NewInt(11111),
// 	// 	big.NewInt(10000),
// 	// 	_messageHash,
// 	// 	_signedHash,
// 	// )
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// blockchain.Commit()

// 	// lastRecovery, err := identityContract.LastRecovery(&bind.CallOpts{Pending: true})
// 	// fmt.Println("Authorized Account: ", _authorizedAccount.String())
// 	// fmt.Println("last recovery ", lastRecovery.String())

// 	// rcpt, err := blockchain.TransactionReceipt(context.Background(), tx.Hash())
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	fmt.Println("CumulativeGasUsed", rcpt.CumulativeGasUsed)
// 	fmt.Println("Logs", rcpt.Logs)
// 	fmt.Println("Status", rcpt.Status)
// }
