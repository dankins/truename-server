// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// ECRecoveryABI is the input ABI used to generate the binding from.
const ECRecoveryABI = "[]"

// ECRecoveryBin is the compiled bytecode used for deploying new contracts.
const ECRecoveryBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820c5fba5ba13098d428c5ffe2ac32794622043a17af1af1d07738d26bab567de0c0029`

// DeployECRecovery deploys a new Ethereum contract, binding an instance of ECRecovery to it.
func DeployECRecovery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECRecovery, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECRecoveryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}, ECRecoveryFilterer: ECRecoveryFilterer{contract: contract}}, nil
}

// ECRecovery is an auto generated Go binding around an Ethereum contract.
type ECRecovery struct {
	ECRecoveryCaller     // Read-only binding to the contract
	ECRecoveryTransactor // Write-only binding to the contract
	ECRecoveryFilterer   // Log filterer for contract events
}

// ECRecoveryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECRecoveryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECRecoveryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECRecoveryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoverySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECRecoverySession struct {
	Contract     *ECRecovery       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECRecoveryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECRecoveryCallerSession struct {
	Contract *ECRecoveryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ECRecoveryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECRecoveryTransactorSession struct {
	Contract     *ECRecoveryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ECRecoveryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECRecoveryRaw struct {
	Contract *ECRecovery // Generic contract binding to access the raw methods on
}

// ECRecoveryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECRecoveryCallerRaw struct {
	Contract *ECRecoveryCaller // Generic read-only contract binding to access the raw methods on
}

// ECRecoveryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECRecoveryTransactorRaw struct {
	Contract *ECRecoveryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECRecovery creates a new instance of ECRecovery, bound to a specific deployed contract.
func NewECRecovery(address common.Address, backend bind.ContractBackend) (*ECRecovery, error) {
	contract, err := bindECRecovery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}, ECRecoveryFilterer: ECRecoveryFilterer{contract: contract}}, nil
}

// NewECRecoveryCaller creates a new read-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryCaller(address common.Address, caller bind.ContractCaller) (*ECRecoveryCaller, error) {
	contract, err := bindECRecovery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryCaller{contract: contract}, nil
}

// NewECRecoveryTransactor creates a new write-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryTransactor(address common.Address, transactor bind.ContractTransactor) (*ECRecoveryTransactor, error) {
	contract, err := bindECRecovery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryTransactor{contract: contract}, nil
}

// NewECRecoveryFilterer creates a new log filterer instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryFilterer(address common.Address, filterer bind.ContractFilterer) (*ECRecoveryFilterer, error) {
	contract, err := bindECRecovery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryFilterer{contract: contract}, nil
}

// bindECRecovery binds a generic wrapper to an already deployed contract.
func bindECRecovery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecovery *ECRecoveryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecovery.Contract.ECRecoveryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecovery *ECRecoveryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecovery.Contract.ECRecoveryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecovery *ECRecoveryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecovery.Contract.ECRecoveryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecovery *ECRecoveryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecovery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecovery *ECRecoveryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecovery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecovery *ECRecoveryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecovery.Contract.contract.Transact(opts, method, params...)
}

// IdentityABI is the input ABI used to generate the binding from.
const IdentityABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedAccounts\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CALL_PREFIX\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_permissions\",\"type\":\"uint8\"},{\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"name\":\"_signedHash\",\"type\":\"bytes\"}],\"name\":\"AuthorizeAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"APPROVEANDCALL_PREFIX\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_baseToken\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"name\":\"_nonce\",\"type\":\"uint256\"},{\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"name\":\"_gasToken\",\"type\":\"address\"}],\"name\":\"approveAndCallGasRelayHash\",\"outputs\":[{\"name\":\"_callGasRelayHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"name\":\"_signedHash\",\"type\":\"bytes\"}],\"name\":\"RevokeAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_nonce\",\"type\":\"uint256\"},{\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"RelayTransaction\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getSignHash\",\"outputs\":[{\"name\":\"signHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"name\":\"_nonce\",\"type\":\"uint256\"},{\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"name\":\"_gasToken\",\"type\":\"address\"}],\"name\":\"callGasRelayHash\",\"outputs\":[{\"name\":\"_callGasRelayHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Read\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Write\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Ping\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"UnknownOperation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Bootstrap\",\"type\":\"event\"}]"

// IdentityBin is the compiled bytecode used for deploying new contracts.
const IdentityBin = `0x608060405234801561001057600080fd5b50604051602080610aab8339810160409081529051600160a060020a03166000908152602081905220805460ff19166001179055610a58806100536000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166324ba588481146100a857806335f894a4146100df57806349048a30146101265780634da3ee83146101725780636e259594146101875780637d61d901146101da578063935b03b31461020b578063affed0e014610254578063b15aa5b714610269578063e27e2e5c14610281575b600080fd5b3480156100b457600080fd5b506100c9600160a060020a03600435166102bc565b6040805160ff9092168252519081900360200190f35b3480156100eb57600080fd5b506100f46102d1565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff199092168252519081900360200190f35b34801561013257600080fd5b5061015e60048035600160a060020a0316906024803560ff16916044359160643590810191013561032c565b604080519115158252519081900360200190f35b34801561017e57600080fd5b506100f461042f565b34801561019357600080fd5b506101c8600160a060020a0360043581169060243581169060443590606435906084359060a4359060c4359060e435166104af565b60408051918252519081900360200190f35b3480156101e657600080fd5b5061015e60048035600160a060020a03169060248035916044359182019101356105bf565b34801561021757600080fd5b5061015e60048035600160a060020a0316906024803591604435808301929082013591606435916084359160a4359160c4359081019101356106bc565b34801561026057600080fd5b506101c8610782565b34801561027557600080fd5b506101c8600435610788565b34801561028d57600080fd5b506101c8600160a060020a036004358116906024359060443590606435906084359060a4359060c435166107c4565b60006020819052908152604090205460ff1681565b604080517f63616c6c47617352656c617928616464726573732c75696e743235362c62797481527f657333322c75696e743235362c75696e743235362c61646472657373290000006020820152905190819003603d01902081565b600080600061036b8686868080601f016020809104026020016040519081016040528093929190818152602001838380828437506108a6945050505050565b600160a060020a03811660009081526020819052604081205491935060ff909116915081116103fb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6163636f756e74206e6f7420617574686f72697a65642e000000000000000000604482015290519081900360640190fd5b600160a060020a0388166000908152602081905260408120805460ff8a1660ff199091161790559250505095945050505050565b604080517f617070726f7665416e6443616c6c47617352656c617928616464726573732c6181527f6464726573732c75696e743235362c627974657333322c75696e743235362c7560208201527f696e74323536290000000000000000000000000000000000000000000000000081830152905190819003604701902081565b604080517f617070726f7665416e6443616c6c47617352656c617928616464726573732c6181527f6464726573732c75696e743235362c627974657333322c75696e743235362c7560208201527f696e7432353629000000000000000000000000000000000000000000000000008183015281519081900360470181206c0100000000000000000000000030810283527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff199091166014830152600160a060020a039a8b1681026018830152988a168902602c820152808201979097526060870195909552608086019390935260a085019190915260c08401529390931690910260e082015290519081900360f401902090565b60008060006105fe8686868080601f016020809104026020016040519081016040528093929190818152602001838380828437506108a6945050505050565b600160a060020a03811660009081526020819052604081205491935060ff9091169150811161068e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6163636f756e74206e6f7420617574686f72697a65642e000000000000000000604482015290519081900360640190fd5b600160a060020a0387166000908152602081905260409020805460ff19169055600192505050949350505050565b60008060006106fa8c8c8c8c6040518083838082843760405192018290039091206002549094508f93508e92508d9150600160a060020a03166107c4565b91506107368286868080601f0160208091040260200160405190810160405280939291908181526020018383808284375061097b945050505050565b6001805481019055604051600160a060020a038d16908c908c908c908083838082843782019150509250505060006040518083038185875af19f9e505050505050505050505050505050565b60015481565b604080517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c8101929092525190819003603c01902090565b604080517f63616c6c47617352656c617928616464726573732c75696e743235362c62797481527f657333322c75696e743235362c75696e743235362c61646472657373290000006020820152815190819003603d0181206c0100000000000000000000000030810283527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff199091166014830152600160a060020a03998a1681026018830152602c820198909852604c810196909652606c860194909452608c85019290925260ac8401529390931690910260cc82015290519081900360e001902090565b600080600080845160411415156108c05760009350610972565b50505060208201516040830151606084015160001a601b60ff821610156108e557601b015b8060ff16601b141580156108fd57508060ff16601c14155b1561090b5760009350610972565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af1158015610965573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b600080600061098985610788565b925061099583856108a6565b600160a060020a03811660009081526020819052604081205491935060ff90911691508111610a2557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6163636f756e74206e6f7420617574686f72697a65642e000000000000000000604482015290519081900360640190fd5b50505050505600a165627a7a723058205a7cb540e02b965251a26d0793b16ec3d4f6066ca7416d85a4e6c9094e1943670029`

// DeployIdentity deploys a new Ethereum contract, binding an instance of Identity to it.
func DeployIdentity(auth *bind.TransactOpts, backend bind.ContractBackend, _account common.Address) (common.Address, *types.Transaction, *Identity, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityBin), backend, _account)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// Identity is an auto generated Go binding around an Ethereum contract.
type Identity struct {
	IdentityCaller     // Read-only binding to the contract
	IdentityTransactor // Write-only binding to the contract
	IdentityFilterer   // Log filterer for contract events
}

// IdentityCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentitySession struct {
	Contract     *Identity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityCallerSession struct {
	Contract *IdentityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IdentityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityTransactorSession struct {
	Contract     *IdentityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IdentityRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRaw struct {
	Contract *Identity // Generic contract binding to access the raw methods on
}

// IdentityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityCallerRaw struct {
	Contract *IdentityCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityTransactorRaw struct {
	Contract *IdentityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentity creates a new instance of Identity, bound to a specific deployed contract.
func NewIdentity(address common.Address, backend bind.ContractBackend) (*Identity, error) {
	contract, err := bindIdentity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// NewIdentityCaller creates a new read-only instance of Identity, bound to a specific deployed contract.
func NewIdentityCaller(address common.Address, caller bind.ContractCaller) (*IdentityCaller, error) {
	contract, err := bindIdentity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityCaller{contract: contract}, nil
}

// NewIdentityTransactor creates a new write-only instance of Identity, bound to a specific deployed contract.
func NewIdentityTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityTransactor, error) {
	contract, err := bindIdentity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityTransactor{contract: contract}, nil
}

// NewIdentityFilterer creates a new log filterer instance of Identity, bound to a specific deployed contract.
func NewIdentityFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityFilterer, error) {
	contract, err := bindIdentity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityFilterer{contract: contract}, nil
}

// bindIdentity binds a generic wrapper to an already deployed contract.
func bindIdentity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.IdentityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transact(opts, method, params...)
}

// APPROVEANDCALLPREFIX is a free data retrieval call binding the contract method 0x4da3ee83.
//
// Solidity: function APPROVEANDCALL_PREFIX() constant returns(bytes4)
func (_Identity *IdentityCaller) APPROVEANDCALLPREFIX(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "APPROVEANDCALL_PREFIX")
	return *ret0, err
}

// APPROVEANDCALLPREFIX is a free data retrieval call binding the contract method 0x4da3ee83.
//
// Solidity: function APPROVEANDCALL_PREFIX() constant returns(bytes4)
func (_Identity *IdentitySession) APPROVEANDCALLPREFIX() ([4]byte, error) {
	return _Identity.Contract.APPROVEANDCALLPREFIX(&_Identity.CallOpts)
}

// APPROVEANDCALLPREFIX is a free data retrieval call binding the contract method 0x4da3ee83.
//
// Solidity: function APPROVEANDCALL_PREFIX() constant returns(bytes4)
func (_Identity *IdentityCallerSession) APPROVEANDCALLPREFIX() ([4]byte, error) {
	return _Identity.Contract.APPROVEANDCALLPREFIX(&_Identity.CallOpts)
}

// CALLPREFIX is a free data retrieval call binding the contract method 0x35f894a4.
//
// Solidity: function CALL_PREFIX() constant returns(bytes4)
func (_Identity *IdentityCaller) CALLPREFIX(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "CALL_PREFIX")
	return *ret0, err
}

// CALLPREFIX is a free data retrieval call binding the contract method 0x35f894a4.
//
// Solidity: function CALL_PREFIX() constant returns(bytes4)
func (_Identity *IdentitySession) CALLPREFIX() ([4]byte, error) {
	return _Identity.Contract.CALLPREFIX(&_Identity.CallOpts)
}

// CALLPREFIX is a free data retrieval call binding the contract method 0x35f894a4.
//
// Solidity: function CALL_PREFIX() constant returns(bytes4)
func (_Identity *IdentityCallerSession) CALLPREFIX() ([4]byte, error) {
	return _Identity.Contract.CALLPREFIX(&_Identity.CallOpts)
}

// ApproveAndCallGasRelayHash is a free data retrieval call binding the contract method 0x6e259594.
//
// Solidity: function approveAndCallGasRelayHash(_baseToken address, _to address, _value uint256, _dataHash bytes32, _nonce uint256, _gasPrice uint256, _gasLimit uint256, _gasToken address) constant returns(_callGasRelayHash bytes32)
func (_Identity *IdentityCaller) ApproveAndCallGasRelayHash(opts *bind.CallOpts, _baseToken common.Address, _to common.Address, _value *big.Int, _dataHash [32]byte, _nonce *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _gasToken common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "approveAndCallGasRelayHash", _baseToken, _to, _value, _dataHash, _nonce, _gasPrice, _gasLimit, _gasToken)
	return *ret0, err
}

// ApproveAndCallGasRelayHash is a free data retrieval call binding the contract method 0x6e259594.
//
// Solidity: function approveAndCallGasRelayHash(_baseToken address, _to address, _value uint256, _dataHash bytes32, _nonce uint256, _gasPrice uint256, _gasLimit uint256, _gasToken address) constant returns(_callGasRelayHash bytes32)
func (_Identity *IdentitySession) ApproveAndCallGasRelayHash(_baseToken common.Address, _to common.Address, _value *big.Int, _dataHash [32]byte, _nonce *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _gasToken common.Address) ([32]byte, error) {
	return _Identity.Contract.ApproveAndCallGasRelayHash(&_Identity.CallOpts, _baseToken, _to, _value, _dataHash, _nonce, _gasPrice, _gasLimit, _gasToken)
}

// ApproveAndCallGasRelayHash is a free data retrieval call binding the contract method 0x6e259594.
//
// Solidity: function approveAndCallGasRelayHash(_baseToken address, _to address, _value uint256, _dataHash bytes32, _nonce uint256, _gasPrice uint256, _gasLimit uint256, _gasToken address) constant returns(_callGasRelayHash bytes32)
func (_Identity *IdentityCallerSession) ApproveAndCallGasRelayHash(_baseToken common.Address, _to common.Address, _value *big.Int, _dataHash [32]byte, _nonce *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _gasToken common.Address) ([32]byte, error) {
	return _Identity.Contract.ApproveAndCallGasRelayHash(&_Identity.CallOpts, _baseToken, _to, _value, _dataHash, _nonce, _gasPrice, _gasLimit, _gasToken)
}

// AuthorizedAccounts is a free data retrieval call binding the contract method 0x24ba5884.
//
// Solidity: function authorizedAccounts( address) constant returns(uint8)
func (_Identity *IdentityCaller) AuthorizedAccounts(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "authorizedAccounts", arg0)
	return *ret0, err
}

// AuthorizedAccounts is a free data retrieval call binding the contract method 0x24ba5884.
//
// Solidity: function authorizedAccounts( address) constant returns(uint8)
func (_Identity *IdentitySession) AuthorizedAccounts(arg0 common.Address) (uint8, error) {
	return _Identity.Contract.AuthorizedAccounts(&_Identity.CallOpts, arg0)
}

// AuthorizedAccounts is a free data retrieval call binding the contract method 0x24ba5884.
//
// Solidity: function authorizedAccounts( address) constant returns(uint8)
func (_Identity *IdentityCallerSession) AuthorizedAccounts(arg0 common.Address) (uint8, error) {
	return _Identity.Contract.AuthorizedAccounts(&_Identity.CallOpts, arg0)
}

// CallGasRelayHash is a free data retrieval call binding the contract method 0xe27e2e5c.
//
// Solidity: function callGasRelayHash(_to address, _value uint256, _dataHash bytes32, _nonce uint256, _gasPrice uint256, _gasLimit uint256, _gasToken address) constant returns(_callGasRelayHash bytes32)
func (_Identity *IdentityCaller) CallGasRelayHash(opts *bind.CallOpts, _to common.Address, _value *big.Int, _dataHash [32]byte, _nonce *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _gasToken common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "callGasRelayHash", _to, _value, _dataHash, _nonce, _gasPrice, _gasLimit, _gasToken)
	return *ret0, err
}

// CallGasRelayHash is a free data retrieval call binding the contract method 0xe27e2e5c.
//
// Solidity: function callGasRelayHash(_to address, _value uint256, _dataHash bytes32, _nonce uint256, _gasPrice uint256, _gasLimit uint256, _gasToken address) constant returns(_callGasRelayHash bytes32)
func (_Identity *IdentitySession) CallGasRelayHash(_to common.Address, _value *big.Int, _dataHash [32]byte, _nonce *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _gasToken common.Address) ([32]byte, error) {
	return _Identity.Contract.CallGasRelayHash(&_Identity.CallOpts, _to, _value, _dataHash, _nonce, _gasPrice, _gasLimit, _gasToken)
}

// CallGasRelayHash is a free data retrieval call binding the contract method 0xe27e2e5c.
//
// Solidity: function callGasRelayHash(_to address, _value uint256, _dataHash bytes32, _nonce uint256, _gasPrice uint256, _gasLimit uint256, _gasToken address) constant returns(_callGasRelayHash bytes32)
func (_Identity *IdentityCallerSession) CallGasRelayHash(_to common.Address, _value *big.Int, _dataHash [32]byte, _nonce *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _gasToken common.Address) ([32]byte, error) {
	return _Identity.Contract.CallGasRelayHash(&_Identity.CallOpts, _to, _value, _dataHash, _nonce, _gasPrice, _gasLimit, _gasToken)
}

// GetSignHash is a free data retrieval call binding the contract method 0xb15aa5b7.
//
// Solidity: function getSignHash(_hash bytes32) constant returns(signHash bytes32)
func (_Identity *IdentityCaller) GetSignHash(opts *bind.CallOpts, _hash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "getSignHash", _hash)
	return *ret0, err
}

// GetSignHash is a free data retrieval call binding the contract method 0xb15aa5b7.
//
// Solidity: function getSignHash(_hash bytes32) constant returns(signHash bytes32)
func (_Identity *IdentitySession) GetSignHash(_hash [32]byte) ([32]byte, error) {
	return _Identity.Contract.GetSignHash(&_Identity.CallOpts, _hash)
}

// GetSignHash is a free data retrieval call binding the contract method 0xb15aa5b7.
//
// Solidity: function getSignHash(_hash bytes32) constant returns(signHash bytes32)
func (_Identity *IdentityCallerSession) GetSignHash(_hash [32]byte) ([32]byte, error) {
	return _Identity.Contract.GetSignHash(&_Identity.CallOpts, _hash)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_Identity *IdentityCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "nonce")
	return *ret0, err
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_Identity *IdentitySession) Nonce() (*big.Int, error) {
	return _Identity.Contract.Nonce(&_Identity.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_Identity *IdentityCallerSession) Nonce() (*big.Int, error) {
	return _Identity.Contract.Nonce(&_Identity.CallOpts)
}

// AuthorizeAccount is a paid mutator transaction binding the contract method 0x49048a30.
//
// Solidity: function AuthorizeAccount(_account address, _permissions uint8, _messageHash bytes32, _signedHash bytes) returns(bool)
func (_Identity *IdentityTransactor) AuthorizeAccount(opts *bind.TransactOpts, _account common.Address, _permissions uint8, _messageHash [32]byte, _signedHash []byte) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "AuthorizeAccount", _account, _permissions, _messageHash, _signedHash)
}

// AuthorizeAccount is a paid mutator transaction binding the contract method 0x49048a30.
//
// Solidity: function AuthorizeAccount(_account address, _permissions uint8, _messageHash bytes32, _signedHash bytes) returns(bool)
func (_Identity *IdentitySession) AuthorizeAccount(_account common.Address, _permissions uint8, _messageHash [32]byte, _signedHash []byte) (*types.Transaction, error) {
	return _Identity.Contract.AuthorizeAccount(&_Identity.TransactOpts, _account, _permissions, _messageHash, _signedHash)
}

// AuthorizeAccount is a paid mutator transaction binding the contract method 0x49048a30.
//
// Solidity: function AuthorizeAccount(_account address, _permissions uint8, _messageHash bytes32, _signedHash bytes) returns(bool)
func (_Identity *IdentityTransactorSession) AuthorizeAccount(_account common.Address, _permissions uint8, _messageHash [32]byte, _signedHash []byte) (*types.Transaction, error) {
	return _Identity.Contract.AuthorizeAccount(&_Identity.TransactOpts, _account, _permissions, _messageHash, _signedHash)
}

// RelayTransaction is a paid mutator transaction binding the contract method 0x935b03b3.
//
// Solidity: function RelayTransaction(_to address, _value uint256, _data bytes, _nonce uint256, _gasPrice uint256, _gasLimit uint256, _signature bytes) returns(bool)
func (_Identity *IdentityTransactor) RelayTransaction(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte, _nonce *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "RelayTransaction", _to, _value, _data, _nonce, _gasPrice, _gasLimit, _signature)
}

// RelayTransaction is a paid mutator transaction binding the contract method 0x935b03b3.
//
// Solidity: function RelayTransaction(_to address, _value uint256, _data bytes, _nonce uint256, _gasPrice uint256, _gasLimit uint256, _signature bytes) returns(bool)
func (_Identity *IdentitySession) RelayTransaction(_to common.Address, _value *big.Int, _data []byte, _nonce *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Identity.Contract.RelayTransaction(&_Identity.TransactOpts, _to, _value, _data, _nonce, _gasPrice, _gasLimit, _signature)
}

// RelayTransaction is a paid mutator transaction binding the contract method 0x935b03b3.
//
// Solidity: function RelayTransaction(_to address, _value uint256, _data bytes, _nonce uint256, _gasPrice uint256, _gasLimit uint256, _signature bytes) returns(bool)
func (_Identity *IdentityTransactorSession) RelayTransaction(_to common.Address, _value *big.Int, _data []byte, _nonce *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Identity.Contract.RelayTransaction(&_Identity.TransactOpts, _to, _value, _data, _nonce, _gasPrice, _gasLimit, _signature)
}

// RevokeAccount is a paid mutator transaction binding the contract method 0x7d61d901.
//
// Solidity: function RevokeAccount(_account address, _messageHash bytes32, _signedHash bytes) returns(bool)
func (_Identity *IdentityTransactor) RevokeAccount(opts *bind.TransactOpts, _account common.Address, _messageHash [32]byte, _signedHash []byte) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "RevokeAccount", _account, _messageHash, _signedHash)
}

// RevokeAccount is a paid mutator transaction binding the contract method 0x7d61d901.
//
// Solidity: function RevokeAccount(_account address, _messageHash bytes32, _signedHash bytes) returns(bool)
func (_Identity *IdentitySession) RevokeAccount(_account common.Address, _messageHash [32]byte, _signedHash []byte) (*types.Transaction, error) {
	return _Identity.Contract.RevokeAccount(&_Identity.TransactOpts, _account, _messageHash, _signedHash)
}

// RevokeAccount is a paid mutator transaction binding the contract method 0x7d61d901.
//
// Solidity: function RevokeAccount(_account address, _messageHash bytes32, _signedHash bytes) returns(bool)
func (_Identity *IdentityTransactorSession) RevokeAccount(_account common.Address, _messageHash [32]byte, _signedHash []byte) (*types.Transaction, error) {
	return _Identity.Contract.RevokeAccount(&_Identity.TransactOpts, _account, _messageHash, _signedHash)
}

// IdentityBootstrapIterator is returned from FilterBootstrap and is used to iterate over the raw logs and unpacked data for Bootstrap events raised by the Identity contract.
type IdentityBootstrapIterator struct {
	Event *IdentityBootstrap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityBootstrapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityBootstrap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityBootstrap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityBootstrapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityBootstrapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityBootstrap represents a Bootstrap event raised by the Identity contract.
type IdentityBootstrap struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBootstrap is a free log retrieval operation binding the contract event 0x83f38dab1f8b68c2e6aa895a461a95a3ff0cb1988db0ecc6e47c15b136c9d55f.
//
// Solidity: e Bootstrap()
func (_Identity *IdentityFilterer) FilterBootstrap(opts *bind.FilterOpts) (*IdentityBootstrapIterator, error) {

	logs, sub, err := _Identity.contract.FilterLogs(opts, "Bootstrap")
	if err != nil {
		return nil, err
	}
	return &IdentityBootstrapIterator{contract: _Identity.contract, event: "Bootstrap", logs: logs, sub: sub}, nil
}

// WatchBootstrap is a free log subscription operation binding the contract event 0x83f38dab1f8b68c2e6aa895a461a95a3ff0cb1988db0ecc6e47c15b136c9d55f.
//
// Solidity: e Bootstrap()
func (_Identity *IdentityFilterer) WatchBootstrap(opts *bind.WatchOpts, sink chan<- *IdentityBootstrap) (event.Subscription, error) {

	logs, sub, err := _Identity.contract.WatchLogs(opts, "Bootstrap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityBootstrap)
				if err := _Identity.contract.UnpackLog(event, "Bootstrap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityPingIterator is returned from FilterPing and is used to iterate over the raw logs and unpacked data for Ping events raised by the Identity contract.
type IdentityPingIterator struct {
	Event *IdentityPing // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityPingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityPing)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityPing)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityPingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityPingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityPing represents a Ping event raised by the Identity contract.
type IdentityPing struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPing is a free log retrieval operation binding the contract event 0xca6e822df923f741dfe968d15d80a18abd25bd1e748bcb9ad81fea5bbb7386af.
//
// Solidity: e Ping()
func (_Identity *IdentityFilterer) FilterPing(opts *bind.FilterOpts) (*IdentityPingIterator, error) {

	logs, sub, err := _Identity.contract.FilterLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return &IdentityPingIterator{contract: _Identity.contract, event: "Ping", logs: logs, sub: sub}, nil
}

// WatchPing is a free log subscription operation binding the contract event 0xca6e822df923f741dfe968d15d80a18abd25bd1e748bcb9ad81fea5bbb7386af.
//
// Solidity: e Ping()
func (_Identity *IdentityFilterer) WatchPing(opts *bind.WatchOpts, sink chan<- *IdentityPing) (event.Subscription, error) {

	logs, sub, err := _Identity.contract.WatchLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityPing)
				if err := _Identity.contract.UnpackLog(event, "Ping", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityReadIterator is returned from FilterRead and is used to iterate over the raw logs and unpacked data for Read events raised by the Identity contract.
type IdentityReadIterator struct {
	Event *IdentityRead // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityReadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRead)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityRead)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityReadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityReadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRead represents a Read event raised by the Identity contract.
type IdentityRead struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRead is a free log retrieval operation binding the contract event 0x97deb47bbb44b0de692e4241b751f167778f2c0a9e6157d119308c91f359dc97.
//
// Solidity: e Read()
func (_Identity *IdentityFilterer) FilterRead(opts *bind.FilterOpts) (*IdentityReadIterator, error) {

	logs, sub, err := _Identity.contract.FilterLogs(opts, "Read")
	if err != nil {
		return nil, err
	}
	return &IdentityReadIterator{contract: _Identity.contract, event: "Read", logs: logs, sub: sub}, nil
}

// WatchRead is a free log subscription operation binding the contract event 0x97deb47bbb44b0de692e4241b751f167778f2c0a9e6157d119308c91f359dc97.
//
// Solidity: e Read()
func (_Identity *IdentityFilterer) WatchRead(opts *bind.WatchOpts, sink chan<- *IdentityRead) (event.Subscription, error) {

	logs, sub, err := _Identity.contract.WatchLogs(opts, "Read")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRead)
				if err := _Identity.contract.UnpackLog(event, "Read", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityUnknownOperationIterator is returned from FilterUnknownOperation and is used to iterate over the raw logs and unpacked data for UnknownOperation events raised by the Identity contract.
type IdentityUnknownOperationIterator struct {
	Event *IdentityUnknownOperation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityUnknownOperationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityUnknownOperation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityUnknownOperation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityUnknownOperationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityUnknownOperationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityUnknownOperation represents a UnknownOperation event raised by the Identity contract.
type IdentityUnknownOperation struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnknownOperation is a free log retrieval operation binding the contract event 0x23f4b3e644a564aaf981ac6fb4a6d76223f5da0c8ad3b15f7056a248ef4c1670.
//
// Solidity: e UnknownOperation()
func (_Identity *IdentityFilterer) FilterUnknownOperation(opts *bind.FilterOpts) (*IdentityUnknownOperationIterator, error) {

	logs, sub, err := _Identity.contract.FilterLogs(opts, "UnknownOperation")
	if err != nil {
		return nil, err
	}
	return &IdentityUnknownOperationIterator{contract: _Identity.contract, event: "UnknownOperation", logs: logs, sub: sub}, nil
}

// WatchUnknownOperation is a free log subscription operation binding the contract event 0x23f4b3e644a564aaf981ac6fb4a6d76223f5da0c8ad3b15f7056a248ef4c1670.
//
// Solidity: e UnknownOperation()
func (_Identity *IdentityFilterer) WatchUnknownOperation(opts *bind.WatchOpts, sink chan<- *IdentityUnknownOperation) (event.Subscription, error) {

	logs, sub, err := _Identity.contract.WatchLogs(opts, "UnknownOperation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityUnknownOperation)
				if err := _Identity.contract.UnpackLog(event, "UnknownOperation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityWriteIterator is returned from FilterWrite and is used to iterate over the raw logs and unpacked data for Write events raised by the Identity contract.
type IdentityWriteIterator struct {
	Event *IdentityWrite // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityWriteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityWrite)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityWrite)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityWriteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityWriteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityWrite represents a Write event raised by the Identity contract.
type IdentityWrite struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWrite is a free log retrieval operation binding the contract event 0xa770741dfa46df8fdae160a7ade69f85bc4341511294e8b14ea83c3eade29a95.
//
// Solidity: e Write()
func (_Identity *IdentityFilterer) FilterWrite(opts *bind.FilterOpts) (*IdentityWriteIterator, error) {

	logs, sub, err := _Identity.contract.FilterLogs(opts, "Write")
	if err != nil {
		return nil, err
	}
	return &IdentityWriteIterator{contract: _Identity.contract, event: "Write", logs: logs, sub: sub}, nil
}

// WatchWrite is a free log subscription operation binding the contract event 0xa770741dfa46df8fdae160a7ade69f85bc4341511294e8b14ea83c3eade29a95.
//
// Solidity: e Write()
func (_Identity *IdentityFilterer) WatchWrite(opts *bind.WatchOpts, sink chan<- *IdentityWrite) (event.Subscription, error) {

	logs, sub, err := _Identity.contract.WatchLogs(opts, "Write")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityWrite)
				if err := _Identity.contract.UnpackLog(event, "Write", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityFactoryABI is the input ABI used to generate the binding from.
const IdentityFactoryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getIdentityForAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_pubKey\",\"type\":\"bytes32\"}],\"name\":\"createIdentity\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IdentityFactoryBin is the compiled bytecode used for deploying new contracts.
const IdentityFactoryBin = `0x608060405234801561001057600080fd5b50610c77806100206000396000f30060806040526004361061004b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663e8f7ba128114610050578063f532a48a1461008d575b600080fd5b34801561005c57600080fd5b50610071600160a060020a03600435166100c5565b60408051600160a060020a039092168252519081900360200190f35b34801561009957600080fd5b506100b1600160a060020a03600435166024356100e3565b604080519115158252519081900360200190f35b600160a060020a039081166000908152602081905260409020541690565b600160a060020a038281166000908152602081905260408120549091829116156101105760009150610189565b83610119610190565b600160a060020a03909116815260405190819003602001906000f080158015610146573d6000803e3d6000fd5b50600160a060020a038581166000908152602081905260409020805473ffffffffffffffffffffffffffffffffffffffff19169183169190911790556001925090505b5092915050565b604051610aab806101a1833901905600608060405234801561001057600080fd5b50604051602080610aab8339810160409081529051600160a060020a03166000908152602081905220805460ff19166001179055610a58806100536000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166324ba588481146100a857806335f894a4146100df57806349048a30146101265780634da3ee83146101725780636e259594146101875780637d61d901146101da578063935b03b31461020b578063affed0e014610254578063b15aa5b714610269578063e27e2e5c14610281575b600080fd5b3480156100b457600080fd5b506100c9600160a060020a03600435166102bc565b6040805160ff9092168252519081900360200190f35b3480156100eb57600080fd5b506100f46102d1565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff199092168252519081900360200190f35b34801561013257600080fd5b5061015e60048035600160a060020a0316906024803560ff16916044359160643590810191013561032c565b604080519115158252519081900360200190f35b34801561017e57600080fd5b506100f461042f565b34801561019357600080fd5b506101c8600160a060020a0360043581169060243581169060443590606435906084359060a4359060c4359060e435166104af565b60408051918252519081900360200190f35b3480156101e657600080fd5b5061015e60048035600160a060020a03169060248035916044359182019101356105bf565b34801561021757600080fd5b5061015e60048035600160a060020a0316906024803591604435808301929082013591606435916084359160a4359160c4359081019101356106bc565b34801561026057600080fd5b506101c8610782565b34801561027557600080fd5b506101c8600435610788565b34801561028d57600080fd5b506101c8600160a060020a036004358116906024359060443590606435906084359060a4359060c435166107c4565b60006020819052908152604090205460ff1681565b604080517f63616c6c47617352656c617928616464726573732c75696e743235362c62797481527f657333322c75696e743235362c75696e743235362c61646472657373290000006020820152905190819003603d01902081565b600080600061036b8686868080601f016020809104026020016040519081016040528093929190818152602001838380828437506108a6945050505050565b600160a060020a03811660009081526020819052604081205491935060ff909116915081116103fb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6163636f756e74206e6f7420617574686f72697a65642e000000000000000000604482015290519081900360640190fd5b600160a060020a0388166000908152602081905260408120805460ff8a1660ff199091161790559250505095945050505050565b604080517f617070726f7665416e6443616c6c47617352656c617928616464726573732c6181527f6464726573732c75696e743235362c627974657333322c75696e743235362c7560208201527f696e74323536290000000000000000000000000000000000000000000000000081830152905190819003604701902081565b604080517f617070726f7665416e6443616c6c47617352656c617928616464726573732c6181527f6464726573732c75696e743235362c627974657333322c75696e743235362c7560208201527f696e7432353629000000000000000000000000000000000000000000000000008183015281519081900360470181206c0100000000000000000000000030810283527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff199091166014830152600160a060020a039a8b1681026018830152988a168902602c820152808201979097526060870195909552608086019390935260a085019190915260c08401529390931690910260e082015290519081900360f401902090565b60008060006105fe8686868080601f016020809104026020016040519081016040528093929190818152602001838380828437506108a6945050505050565b600160a060020a03811660009081526020819052604081205491935060ff9091169150811161068e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6163636f756e74206e6f7420617574686f72697a65642e000000000000000000604482015290519081900360640190fd5b600160a060020a0387166000908152602081905260409020805460ff19169055600192505050949350505050565b60008060006106fa8c8c8c8c6040518083838082843760405192018290039091206002549094508f93508e92508d9150600160a060020a03166107c4565b91506107368286868080601f0160208091040260200160405190810160405280939291908181526020018383808284375061097b945050505050565b6001805481019055604051600160a060020a038d16908c908c908c908083838082843782019150509250505060006040518083038185875af19f9e505050505050505050505050505050565b60015481565b604080517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c8101929092525190819003603c01902090565b604080517f63616c6c47617352656c617928616464726573732c75696e743235362c62797481527f657333322c75696e743235362c75696e743235362c61646472657373290000006020820152815190819003603d0181206c0100000000000000000000000030810283527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff199091166014830152600160a060020a03998a1681026018830152602c820198909852604c810196909652606c860194909452608c85019290925260ac8401529390931690910260cc82015290519081900360e001902090565b600080600080845160411415156108c05760009350610972565b50505060208201516040830151606084015160001a601b60ff821610156108e557601b015b8060ff16601b141580156108fd57508060ff16601c14155b1561090b5760009350610972565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af1158015610965573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b600080600061098985610788565b925061099583856108a6565b600160a060020a03811660009081526020819052604081205491935060ff90911691508111610a2557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6163636f756e74206e6f7420617574686f72697a65642e000000000000000000604482015290519081900360640190fd5b50505050505600a165627a7a723058205a7cb540e02b965251a26d0793b16ec3d4f6066ca7416d85a4e6c9094e1943670029a165627a7a72305820ed766079d59755b5626e61ce37b4cac25f2bc72757b45c6a1f478599c1680cff0029`

// DeployIdentityFactory deploys a new Ethereum contract, binding an instance of IdentityFactory to it.
func DeployIdentityFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IdentityFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityFactory{IdentityFactoryCaller: IdentityFactoryCaller{contract: contract}, IdentityFactoryTransactor: IdentityFactoryTransactor{contract: contract}, IdentityFactoryFilterer: IdentityFactoryFilterer{contract: contract}}, nil
}

// IdentityFactory is an auto generated Go binding around an Ethereum contract.
type IdentityFactory struct {
	IdentityFactoryCaller     // Read-only binding to the contract
	IdentityFactoryTransactor // Write-only binding to the contract
	IdentityFactoryFilterer   // Log filterer for contract events
}

// IdentityFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityFactorySession struct {
	Contract     *IdentityFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityFactoryCallerSession struct {
	Contract *IdentityFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IdentityFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityFactoryTransactorSession struct {
	Contract     *IdentityFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IdentityFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityFactoryRaw struct {
	Contract *IdentityFactory // Generic contract binding to access the raw methods on
}

// IdentityFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityFactoryCallerRaw struct {
	Contract *IdentityFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityFactoryTransactorRaw struct {
	Contract *IdentityFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityFactory creates a new instance of IdentityFactory, bound to a specific deployed contract.
func NewIdentityFactory(address common.Address, backend bind.ContractBackend) (*IdentityFactory, error) {
	contract, err := bindIdentityFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityFactory{IdentityFactoryCaller: IdentityFactoryCaller{contract: contract}, IdentityFactoryTransactor: IdentityFactoryTransactor{contract: contract}, IdentityFactoryFilterer: IdentityFactoryFilterer{contract: contract}}, nil
}

// NewIdentityFactoryCaller creates a new read-only instance of IdentityFactory, bound to a specific deployed contract.
func NewIdentityFactoryCaller(address common.Address, caller bind.ContractCaller) (*IdentityFactoryCaller, error) {
	contract, err := bindIdentityFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityFactoryCaller{contract: contract}, nil
}

// NewIdentityFactoryTransactor creates a new write-only instance of IdentityFactory, bound to a specific deployed contract.
func NewIdentityFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityFactoryTransactor, error) {
	contract, err := bindIdentityFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityFactoryTransactor{contract: contract}, nil
}

// NewIdentityFactoryFilterer creates a new log filterer instance of IdentityFactory, bound to a specific deployed contract.
func NewIdentityFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityFactoryFilterer, error) {
	contract, err := bindIdentityFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityFactoryFilterer{contract: contract}, nil
}

// bindIdentityFactory binds a generic wrapper to an already deployed contract.
func bindIdentityFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityFactory *IdentityFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityFactory.Contract.IdentityFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityFactory *IdentityFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityFactory.Contract.IdentityFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityFactory *IdentityFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityFactory.Contract.IdentityFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityFactory *IdentityFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityFactory *IdentityFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityFactory *IdentityFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityFactory.Contract.contract.Transact(opts, method, params...)
}

// GetIdentityForAccount is a free data retrieval call binding the contract method 0xe8f7ba12.
//
// Solidity: function getIdentityForAccount(_account address) constant returns(address)
func (_IdentityFactory *IdentityFactoryCaller) GetIdentityForAccount(opts *bind.CallOpts, _account common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityFactory.contract.Call(opts, out, "getIdentityForAccount", _account)
	return *ret0, err
}

// GetIdentityForAccount is a free data retrieval call binding the contract method 0xe8f7ba12.
//
// Solidity: function getIdentityForAccount(_account address) constant returns(address)
func (_IdentityFactory *IdentityFactorySession) GetIdentityForAccount(_account common.Address) (common.Address, error) {
	return _IdentityFactory.Contract.GetIdentityForAccount(&_IdentityFactory.CallOpts, _account)
}

// GetIdentityForAccount is a free data retrieval call binding the contract method 0xe8f7ba12.
//
// Solidity: function getIdentityForAccount(_account address) constant returns(address)
func (_IdentityFactory *IdentityFactoryCallerSession) GetIdentityForAccount(_account common.Address) (common.Address, error) {
	return _IdentityFactory.Contract.GetIdentityForAccount(&_IdentityFactory.CallOpts, _account)
}

// CreateIdentity is a paid mutator transaction binding the contract method 0xf532a48a.
//
// Solidity: function createIdentity(_account address, _pubKey bytes32) returns(bool)
func (_IdentityFactory *IdentityFactoryTransactor) CreateIdentity(opts *bind.TransactOpts, _account common.Address, _pubKey [32]byte) (*types.Transaction, error) {
	return _IdentityFactory.contract.Transact(opts, "createIdentity", _account, _pubKey)
}

// CreateIdentity is a paid mutator transaction binding the contract method 0xf532a48a.
//
// Solidity: function createIdentity(_account address, _pubKey bytes32) returns(bool)
func (_IdentityFactory *IdentityFactorySession) CreateIdentity(_account common.Address, _pubKey [32]byte) (*types.Transaction, error) {
	return _IdentityFactory.Contract.CreateIdentity(&_IdentityFactory.TransactOpts, _account, _pubKey)
}

// CreateIdentity is a paid mutator transaction binding the contract method 0xf532a48a.
//
// Solidity: function createIdentity(_account address, _pubKey bytes32) returns(bool)
func (_IdentityFactory *IdentityFactoryTransactorSession) CreateIdentity(_account common.Address, _pubKey [32]byte) (*types.Transaction, error) {
	return _IdentityFactory.Contract.CreateIdentity(&_IdentityFactory.TransactOpts, _account, _pubKey)
}

// RecoveryVerifierABI is the input ABI used to generate the binding from.
const RecoveryVerifierABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"name\":\"_signedHash\",\"type\":\"bytes\"}],\"name\":\"extractAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_pubKey\",\"type\":\"bytes32\"},{\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"name\":\"_signedHash\",\"type\":\"bytes\"}],\"name\":\"verifySignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_pubKey\",\"type\":\"bytes32\"},{\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"name\":\"_signedHash\",\"type\":\"bytes\"}],\"name\":\"verifySignatureNoPrefix\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"name\":\"_signedHash\",\"type\":\"bytes\"}],\"name\":\"extractAccountNoPrefix\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// RecoveryVerifierBin is the compiled bytecode used for deploying new contracts.
const RecoveryVerifierBin = `0x608060405234801561001057600080fd5b506103dc806100206000396000f3006080604052600436106100605763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041662ae32a0811461006557806378efa08a146100a5578063d2f70e75146100ee578063d67745a314610123575b600080fd5b34801561007157600080fd5b50610089600480359060248035908101910135610147565b60408051600160a060020a039092168252519081900360200190f35b3480156100b157600080fd5b506100da60048035600160a060020a031690602480359160443591606435908101910135610198565b604080519115158252519081900360200190f35b3480156100fa57600080fd5b506100da60048035600160a060020a0316906024803591604435916064359081019101356101ff565b34801561012f57600080fd5b50610089600480359060248035908101910135610259565b6000806101538561029d565b905061018f8185858080601f016020809104026020016040519081016040528093929190818152602001838380828437506102db945050505050565b95945050505050565b6000806101a48561029d565b905086600160a060020a03166101ea8286868080601f016020809104026020016040519081016040528093929190818152602001838380828437506102db945050505050565b600160a060020a031614979650505050505050565b600085600160a060020a03166102458585858080601f016020809104026020016040519081016040528093929190818152602001838380828437506102db945050505050565b600160a060020a0316149695505050505050565b60006102958484848080601f016020809104026020016040519081016040528093929190818152602001838380828437506102db945050505050565b949350505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c8101839052905190819003603c019020919050565b600080600080845160411415156102f557600093506103a7565b50505060208201516040830151606084015160001a601b60ff8216101561031a57601b015b8060ff16601b1415801561033257508060ff16601c14155b1561034057600093506103a7565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af115801561039a573d6000803e3d6000fd5b5050506020604051035193505b505050929150505600a165627a7a72305820269d004d7c5b8eb0160dbe804b281a842813f575086b1cbd1026d26f0f8ea0a00029`

// DeployRecoveryVerifier deploys a new Ethereum contract, binding an instance of RecoveryVerifier to it.
func DeployRecoveryVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RecoveryVerifier, error) {
	parsed, err := abi.JSON(strings.NewReader(RecoveryVerifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RecoveryVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RecoveryVerifier{RecoveryVerifierCaller: RecoveryVerifierCaller{contract: contract}, RecoveryVerifierTransactor: RecoveryVerifierTransactor{contract: contract}, RecoveryVerifierFilterer: RecoveryVerifierFilterer{contract: contract}}, nil
}

// RecoveryVerifier is an auto generated Go binding around an Ethereum contract.
type RecoveryVerifier struct {
	RecoveryVerifierCaller     // Read-only binding to the contract
	RecoveryVerifierTransactor // Write-only binding to the contract
	RecoveryVerifierFilterer   // Log filterer for contract events
}

// RecoveryVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecoveryVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoveryVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecoveryVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoveryVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecoveryVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoveryVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecoveryVerifierSession struct {
	Contract     *RecoveryVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecoveryVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecoveryVerifierCallerSession struct {
	Contract *RecoveryVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// RecoveryVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecoveryVerifierTransactorSession struct {
	Contract     *RecoveryVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RecoveryVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecoveryVerifierRaw struct {
	Contract *RecoveryVerifier // Generic contract binding to access the raw methods on
}

// RecoveryVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecoveryVerifierCallerRaw struct {
	Contract *RecoveryVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// RecoveryVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecoveryVerifierTransactorRaw struct {
	Contract *RecoveryVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecoveryVerifier creates a new instance of RecoveryVerifier, bound to a specific deployed contract.
func NewRecoveryVerifier(address common.Address, backend bind.ContractBackend) (*RecoveryVerifier, error) {
	contract, err := bindRecoveryVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RecoveryVerifier{RecoveryVerifierCaller: RecoveryVerifierCaller{contract: contract}, RecoveryVerifierTransactor: RecoveryVerifierTransactor{contract: contract}, RecoveryVerifierFilterer: RecoveryVerifierFilterer{contract: contract}}, nil
}

// NewRecoveryVerifierCaller creates a new read-only instance of RecoveryVerifier, bound to a specific deployed contract.
func NewRecoveryVerifierCaller(address common.Address, caller bind.ContractCaller) (*RecoveryVerifierCaller, error) {
	contract, err := bindRecoveryVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecoveryVerifierCaller{contract: contract}, nil
}

// NewRecoveryVerifierTransactor creates a new write-only instance of RecoveryVerifier, bound to a specific deployed contract.
func NewRecoveryVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*RecoveryVerifierTransactor, error) {
	contract, err := bindRecoveryVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecoveryVerifierTransactor{contract: contract}, nil
}

// NewRecoveryVerifierFilterer creates a new log filterer instance of RecoveryVerifier, bound to a specific deployed contract.
func NewRecoveryVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*RecoveryVerifierFilterer, error) {
	contract, err := bindRecoveryVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecoveryVerifierFilterer{contract: contract}, nil
}

// bindRecoveryVerifier binds a generic wrapper to an already deployed contract.
func bindRecoveryVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RecoveryVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RecoveryVerifier *RecoveryVerifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RecoveryVerifier.Contract.RecoveryVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RecoveryVerifier *RecoveryVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryVerifier.Contract.RecoveryVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RecoveryVerifier *RecoveryVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecoveryVerifier.Contract.RecoveryVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RecoveryVerifier *RecoveryVerifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RecoveryVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RecoveryVerifier *RecoveryVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RecoveryVerifier *RecoveryVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecoveryVerifier.Contract.contract.Transact(opts, method, params...)
}

// ExtractAccount is a free data retrieval call binding the contract method 0x00ae32a0.
//
// Solidity: function extractAccount(_messageHash bytes32, _signedHash bytes) constant returns(address)
func (_RecoveryVerifier *RecoveryVerifierCaller) ExtractAccount(opts *bind.CallOpts, _messageHash [32]byte, _signedHash []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RecoveryVerifier.contract.Call(opts, out, "extractAccount", _messageHash, _signedHash)
	return *ret0, err
}

// ExtractAccount is a free data retrieval call binding the contract method 0x00ae32a0.
//
// Solidity: function extractAccount(_messageHash bytes32, _signedHash bytes) constant returns(address)
func (_RecoveryVerifier *RecoveryVerifierSession) ExtractAccount(_messageHash [32]byte, _signedHash []byte) (common.Address, error) {
	return _RecoveryVerifier.Contract.ExtractAccount(&_RecoveryVerifier.CallOpts, _messageHash, _signedHash)
}

// ExtractAccount is a free data retrieval call binding the contract method 0x00ae32a0.
//
// Solidity: function extractAccount(_messageHash bytes32, _signedHash bytes) constant returns(address)
func (_RecoveryVerifier *RecoveryVerifierCallerSession) ExtractAccount(_messageHash [32]byte, _signedHash []byte) (common.Address, error) {
	return _RecoveryVerifier.Contract.ExtractAccount(&_RecoveryVerifier.CallOpts, _messageHash, _signedHash)
}

// ExtractAccountNoPrefix is a free data retrieval call binding the contract method 0xd67745a3.
//
// Solidity: function extractAccountNoPrefix(_messageHash bytes32, _signedHash bytes) constant returns(address)
func (_RecoveryVerifier *RecoveryVerifierCaller) ExtractAccountNoPrefix(opts *bind.CallOpts, _messageHash [32]byte, _signedHash []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RecoveryVerifier.contract.Call(opts, out, "extractAccountNoPrefix", _messageHash, _signedHash)
	return *ret0, err
}

// ExtractAccountNoPrefix is a free data retrieval call binding the contract method 0xd67745a3.
//
// Solidity: function extractAccountNoPrefix(_messageHash bytes32, _signedHash bytes) constant returns(address)
func (_RecoveryVerifier *RecoveryVerifierSession) ExtractAccountNoPrefix(_messageHash [32]byte, _signedHash []byte) (common.Address, error) {
	return _RecoveryVerifier.Contract.ExtractAccountNoPrefix(&_RecoveryVerifier.CallOpts, _messageHash, _signedHash)
}

// ExtractAccountNoPrefix is a free data retrieval call binding the contract method 0xd67745a3.
//
// Solidity: function extractAccountNoPrefix(_messageHash bytes32, _signedHash bytes) constant returns(address)
func (_RecoveryVerifier *RecoveryVerifierCallerSession) ExtractAccountNoPrefix(_messageHash [32]byte, _signedHash []byte) (common.Address, error) {
	return _RecoveryVerifier.Contract.ExtractAccountNoPrefix(&_RecoveryVerifier.CallOpts, _messageHash, _signedHash)
}

// VerifySignature is a free data retrieval call binding the contract method 0x78efa08a.
//
// Solidity: function verifySignature(_account address, _pubKey bytes32, _messageHash bytes32, _signedHash bytes) constant returns(bool)
func (_RecoveryVerifier *RecoveryVerifierCaller) VerifySignature(opts *bind.CallOpts, _account common.Address, _pubKey [32]byte, _messageHash [32]byte, _signedHash []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RecoveryVerifier.contract.Call(opts, out, "verifySignature", _account, _pubKey, _messageHash, _signedHash)
	return *ret0, err
}

// VerifySignature is a free data retrieval call binding the contract method 0x78efa08a.
//
// Solidity: function verifySignature(_account address, _pubKey bytes32, _messageHash bytes32, _signedHash bytes) constant returns(bool)
func (_RecoveryVerifier *RecoveryVerifierSession) VerifySignature(_account common.Address, _pubKey [32]byte, _messageHash [32]byte, _signedHash []byte) (bool, error) {
	return _RecoveryVerifier.Contract.VerifySignature(&_RecoveryVerifier.CallOpts, _account, _pubKey, _messageHash, _signedHash)
}

// VerifySignature is a free data retrieval call binding the contract method 0x78efa08a.
//
// Solidity: function verifySignature(_account address, _pubKey bytes32, _messageHash bytes32, _signedHash bytes) constant returns(bool)
func (_RecoveryVerifier *RecoveryVerifierCallerSession) VerifySignature(_account common.Address, _pubKey [32]byte, _messageHash [32]byte, _signedHash []byte) (bool, error) {
	return _RecoveryVerifier.Contract.VerifySignature(&_RecoveryVerifier.CallOpts, _account, _pubKey, _messageHash, _signedHash)
}

// VerifySignatureNoPrefix is a free data retrieval call binding the contract method 0xd2f70e75.
//
// Solidity: function verifySignatureNoPrefix(_account address, _pubKey bytes32, _messageHash bytes32, _signedHash bytes) constant returns(bool)
func (_RecoveryVerifier *RecoveryVerifierCaller) VerifySignatureNoPrefix(opts *bind.CallOpts, _account common.Address, _pubKey [32]byte, _messageHash [32]byte, _signedHash []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RecoveryVerifier.contract.Call(opts, out, "verifySignatureNoPrefix", _account, _pubKey, _messageHash, _signedHash)
	return *ret0, err
}

// VerifySignatureNoPrefix is a free data retrieval call binding the contract method 0xd2f70e75.
//
// Solidity: function verifySignatureNoPrefix(_account address, _pubKey bytes32, _messageHash bytes32, _signedHash bytes) constant returns(bool)
func (_RecoveryVerifier *RecoveryVerifierSession) VerifySignatureNoPrefix(_account common.Address, _pubKey [32]byte, _messageHash [32]byte, _signedHash []byte) (bool, error) {
	return _RecoveryVerifier.Contract.VerifySignatureNoPrefix(&_RecoveryVerifier.CallOpts, _account, _pubKey, _messageHash, _signedHash)
}

// VerifySignatureNoPrefix is a free data retrieval call binding the contract method 0xd2f70e75.
//
// Solidity: function verifySignatureNoPrefix(_account address, _pubKey bytes32, _messageHash bytes32, _signedHash bytes) constant returns(bool)
func (_RecoveryVerifier *RecoveryVerifierCallerSession) VerifySignatureNoPrefix(_account common.Address, _pubKey [32]byte, _messageHash [32]byte, _signedHash []byte) (bool, error) {
	return _RecoveryVerifier.Contract.VerifySignatureNoPrefix(&_RecoveryVerifier.CallOpts, _account, _pubKey, _messageHash, _signedHash)
}
