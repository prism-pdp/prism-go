// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dpduado

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BaseCounterMetaData contains all meta data concerning the BaseCounter contract.
var BaseCounterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_initVal\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decrement\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increment\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCount\",\"inputs\":[{\"name\":\"_newVal\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x6080604052348015600f57600080fd5b506040516101ba3803806101ba833981016040819052602c916033565b600055604b565b600060208284031215604457600080fd5b5051919050565b6101608061005a6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806306661abd146100515780632baeceb71461006c578063d09de08a14610076578063d14e62b81461007e575b600080fd5b61005a60005481565b60405190815260200160405180910390f35b610074610091565b005b6100746100b0565b61007461008c3660046100cb565b600055565b600054156100ae576000805490806100a8836100fa565b91905055505b565b60001960005410156100ae576000805490806100a883610111565b6000602082840312156100dd57600080fd5b5035919050565b634e487b7160e01b600052601160045260246000fd5b600081610109576101096100e4565b506000190190565b600060018201610123576101236100e4565b506001019056fea264697066735822122001cf12a87c7be98eec6041d2abf9139a9e5a4f7530680c772286666beecbc10f64736f6c63430008190033",
}

// BaseCounterABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseCounterMetaData.ABI instead.
var BaseCounterABI = BaseCounterMetaData.ABI

// BaseCounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BaseCounterMetaData.Bin instead.
var BaseCounterBin = BaseCounterMetaData.Bin

// DeployBaseCounter deploys a new Ethereum contract, binding an instance of BaseCounter to it.
func DeployBaseCounter(auth *bind.TransactOpts, backend bind.ContractBackend, _initVal *big.Int) (common.Address, *types.Transaction, *BaseCounter, error) {
	parsed, err := BaseCounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BaseCounterBin), backend, _initVal)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseCounter{BaseCounterCaller: BaseCounterCaller{contract: contract}, BaseCounterTransactor: BaseCounterTransactor{contract: contract}, BaseCounterFilterer: BaseCounterFilterer{contract: contract}}, nil
}

// BaseCounter is an auto generated Go binding around an Ethereum contract.
type BaseCounter struct {
	BaseCounterCaller     // Read-only binding to the contract
	BaseCounterTransactor // Write-only binding to the contract
	BaseCounterFilterer   // Log filterer for contract events
}

// BaseCounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseCounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseCounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseCounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseCounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseCounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseCounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseCounterSession struct {
	Contract     *BaseCounter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseCounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseCounterCallerSession struct {
	Contract *BaseCounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BaseCounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseCounterTransactorSession struct {
	Contract     *BaseCounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BaseCounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseCounterRaw struct {
	Contract *BaseCounter // Generic contract binding to access the raw methods on
}

// BaseCounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseCounterCallerRaw struct {
	Contract *BaseCounterCaller // Generic read-only contract binding to access the raw methods on
}

// BaseCounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseCounterTransactorRaw struct {
	Contract *BaseCounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseCounter creates a new instance of BaseCounter, bound to a specific deployed contract.
func NewBaseCounter(address common.Address, backend bind.ContractBackend) (*BaseCounter, error) {
	contract, err := bindBaseCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseCounter{BaseCounterCaller: BaseCounterCaller{contract: contract}, BaseCounterTransactor: BaseCounterTransactor{contract: contract}, BaseCounterFilterer: BaseCounterFilterer{contract: contract}}, nil
}

// NewBaseCounterCaller creates a new read-only instance of BaseCounter, bound to a specific deployed contract.
func NewBaseCounterCaller(address common.Address, caller bind.ContractCaller) (*BaseCounterCaller, error) {
	contract, err := bindBaseCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseCounterCaller{contract: contract}, nil
}

// NewBaseCounterTransactor creates a new write-only instance of BaseCounter, bound to a specific deployed contract.
func NewBaseCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseCounterTransactor, error) {
	contract, err := bindBaseCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseCounterTransactor{contract: contract}, nil
}

// NewBaseCounterFilterer creates a new log filterer instance of BaseCounter, bound to a specific deployed contract.
func NewBaseCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseCounterFilterer, error) {
	contract, err := bindBaseCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseCounterFilterer{contract: contract}, nil
}

// bindBaseCounter binds a generic wrapper to an already deployed contract.
func bindBaseCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaseCounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseCounter *BaseCounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseCounter.Contract.BaseCounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseCounter *BaseCounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseCounter.Contract.BaseCounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseCounter *BaseCounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseCounter.Contract.BaseCounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseCounter *BaseCounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseCounter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseCounter *BaseCounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseCounter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseCounter *BaseCounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseCounter.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_BaseCounter *BaseCounterCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseCounter.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_BaseCounter *BaseCounterSession) Count() (*big.Int, error) {
	return _BaseCounter.Contract.Count(&_BaseCounter.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_BaseCounter *BaseCounterCallerSession) Count() (*big.Int, error) {
	return _BaseCounter.Contract.Count(&_BaseCounter.CallOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_BaseCounter *BaseCounterTransactor) Decrement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseCounter.contract.Transact(opts, "decrement")
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_BaseCounter *BaseCounterSession) Decrement() (*types.Transaction, error) {
	return _BaseCounter.Contract.Decrement(&_BaseCounter.TransactOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_BaseCounter *BaseCounterTransactorSession) Decrement() (*types.Transaction, error) {
	return _BaseCounter.Contract.Decrement(&_BaseCounter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_BaseCounter *BaseCounterTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseCounter.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_BaseCounter *BaseCounterSession) Increment() (*types.Transaction, error) {
	return _BaseCounter.Contract.Increment(&_BaseCounter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_BaseCounter *BaseCounterTransactorSession) Increment() (*types.Transaction, error) {
	return _BaseCounter.Contract.Increment(&_BaseCounter.TransactOpts)
}

// SetCount is a paid mutator transaction binding the contract method 0xd14e62b8.
//
// Solidity: function setCount(uint256 _newVal) returns()
func (_BaseCounter *BaseCounterTransactor) SetCount(opts *bind.TransactOpts, _newVal *big.Int) (*types.Transaction, error) {
	return _BaseCounter.contract.Transact(opts, "setCount", _newVal)
}

// SetCount is a paid mutator transaction binding the contract method 0xd14e62b8.
//
// Solidity: function setCount(uint256 _newVal) returns()
func (_BaseCounter *BaseCounterSession) SetCount(_newVal *big.Int) (*types.Transaction, error) {
	return _BaseCounter.Contract.SetCount(&_BaseCounter.TransactOpts, _newVal)
}

// SetCount is a paid mutator transaction binding the contract method 0xd14e62b8.
//
// Solidity: function setCount(uint256 _newVal) returns()
func (_BaseCounter *BaseCounterTransactorSession) SetCount(_newVal *big.Int) (*types.Transaction, error) {
	return _BaseCounter.Contract.SetCount(&_BaseCounter.TransactOpts, _newVal)
}
