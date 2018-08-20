// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CivilianBadgeABI is the input ABI used to generate the binding from.
const CivilianBadgeABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"becomeCivilian\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_civilian\",\"type\":\"address\"}],\"name\":\"getID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CivilianBadgeBin is the compiled bytecode used for deploying new contracts.
const CivilianBadgeBin = `0x608060405234801561001057600080fd5b5061010f806100206000396000f30060806040526004361060485763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663518141308114604d57806399f826a5146061575b600080fd5b348015605857600080fd5b50605f609e565b005b348015606c57600080fd5b50608c73ffffffffffffffffffffffffffffffffffffffff6004351660bb565b60408051918252519081900360200190f35b600080546001908101808355338352602091909152604090912055565b73ffffffffffffffffffffffffffffffffffffffff16600090815260016020526040902054905600a165627a7a7230582093f45e281aa360f88fac86d33e237cb92914211127f0d417b4d5d408bc7c7bd40029`

// DeployCivilianBadge deploys a new Ethereum contract, binding an instance of CivilianBadge to it.
func DeployCivilianBadge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CivilianBadge, error) {
	parsed, err := abi.JSON(strings.NewReader(CivilianBadgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CivilianBadgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CivilianBadge{CivilianBadgeCaller: CivilianBadgeCaller{contract: contract}, CivilianBadgeTransactor: CivilianBadgeTransactor{contract: contract}, CivilianBadgeFilterer: CivilianBadgeFilterer{contract: contract}}, nil
}

// CivilianBadge is an auto generated Go binding around an Ethereum contract.
type CivilianBadge struct {
	CivilianBadgeCaller     // Read-only binding to the contract
	CivilianBadgeTransactor // Write-only binding to the contract
	CivilianBadgeFilterer   // Log filterer for contract events
}

// CivilianBadgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CivilianBadgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CivilianBadgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CivilianBadgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CivilianBadgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CivilianBadgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CivilianBadgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CivilianBadgeSession struct {
	Contract     *CivilianBadge    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CivilianBadgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CivilianBadgeCallerSession struct {
	Contract *CivilianBadgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CivilianBadgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CivilianBadgeTransactorSession struct {
	Contract     *CivilianBadgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CivilianBadgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CivilianBadgeRaw struct {
	Contract *CivilianBadge // Generic contract binding to access the raw methods on
}

// CivilianBadgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CivilianBadgeCallerRaw struct {
	Contract *CivilianBadgeCaller // Generic read-only contract binding to access the raw methods on
}

// CivilianBadgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CivilianBadgeTransactorRaw struct {
	Contract *CivilianBadgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCivilianBadge creates a new instance of CivilianBadge, bound to a specific deployed contract.
func NewCivilianBadge(address common.Address, backend bind.ContractBackend) (*CivilianBadge, error) {
	contract, err := bindCivilianBadge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CivilianBadge{CivilianBadgeCaller: CivilianBadgeCaller{contract: contract}, CivilianBadgeTransactor: CivilianBadgeTransactor{contract: contract}, CivilianBadgeFilterer: CivilianBadgeFilterer{contract: contract}}, nil
}

// NewCivilianBadgeCaller creates a new read-only instance of CivilianBadge, bound to a specific deployed contract.
func NewCivilianBadgeCaller(address common.Address, caller bind.ContractCaller) (*CivilianBadgeCaller, error) {
	contract, err := bindCivilianBadge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CivilianBadgeCaller{contract: contract}, nil
}

// NewCivilianBadgeTransactor creates a new write-only instance of CivilianBadge, bound to a specific deployed contract.
func NewCivilianBadgeTransactor(address common.Address, transactor bind.ContractTransactor) (*CivilianBadgeTransactor, error) {
	contract, err := bindCivilianBadge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CivilianBadgeTransactor{contract: contract}, nil
}

// NewCivilianBadgeFilterer creates a new log filterer instance of CivilianBadge, bound to a specific deployed contract.
func NewCivilianBadgeFilterer(address common.Address, filterer bind.ContractFilterer) (*CivilianBadgeFilterer, error) {
	contract, err := bindCivilianBadge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CivilianBadgeFilterer{contract: contract}, nil
}

// bindCivilianBadge binds a generic wrapper to an already deployed contract.
func bindCivilianBadge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CivilianBadgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CivilianBadge *CivilianBadgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CivilianBadge.Contract.CivilianBadgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CivilianBadge *CivilianBadgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CivilianBadge.Contract.CivilianBadgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CivilianBadge *CivilianBadgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CivilianBadge.Contract.CivilianBadgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CivilianBadge *CivilianBadgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CivilianBadge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CivilianBadge *CivilianBadgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CivilianBadge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CivilianBadge *CivilianBadgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CivilianBadge.Contract.contract.Transact(opts, method, params...)
}

// GetID is a free data retrieval call binding the contract method 0x99f826a5.
//
// Solidity: function getID(_civilian address) constant returns(uint256)
func (_CivilianBadge *CivilianBadgeCaller) GetID(opts *bind.CallOpts, _civilian common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CivilianBadge.contract.Call(opts, out, "getID", _civilian)
	return *ret0, err
}

// GetID is a free data retrieval call binding the contract method 0x99f826a5.
//
// Solidity: function getID(_civilian address) constant returns(uint256)
func (_CivilianBadge *CivilianBadgeSession) GetID(_civilian common.Address) (*big.Int, error) {
	return _CivilianBadge.Contract.GetID(&_CivilianBadge.CallOpts, _civilian)
}

// GetID is a free data retrieval call binding the contract method 0x99f826a5.
//
// Solidity: function getID(_civilian address) constant returns(uint256)
func (_CivilianBadge *CivilianBadgeCallerSession) GetID(_civilian common.Address) (*big.Int, error) {
	return _CivilianBadge.Contract.GetID(&_CivilianBadge.CallOpts, _civilian)
}

// BecomeCivilian is a paid mutator transaction binding the contract method 0x51814130.
//
// Solidity: function becomeCivilian() returns()
func (_CivilianBadge *CivilianBadgeTransactor) BecomeCivilian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CivilianBadge.contract.Transact(opts, "becomeCivilian")
}

// BecomeCivilian is a paid mutator transaction binding the contract method 0x51814130.
//
// Solidity: function becomeCivilian() returns()
func (_CivilianBadge *CivilianBadgeSession) BecomeCivilian() (*types.Transaction, error) {
	return _CivilianBadge.Contract.BecomeCivilian(&_CivilianBadge.TransactOpts)
}

// BecomeCivilian is a paid mutator transaction binding the contract method 0x51814130.
//
// Solidity: function becomeCivilian() returns()
func (_CivilianBadge *CivilianBadgeTransactorSession) BecomeCivilian() (*types.Transaction, error) {
	return _CivilianBadge.Contract.BecomeCivilian(&_CivilianBadge.TransactOpts)
}
