// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package xz21

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

// XZ21Account is an auto generated low-level Go binding around an user-defined struct.
type XZ21Account struct {
	PubKey []byte
}

// XZ21Para is an auto generated low-level Go binding around an user-defined struct.
type XZ21Para struct {
	Params string
	U      []byte
	G      []byte
}

// XZ21MetaData contains all meta data concerning the XZ21 contract.
var XZ21MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_addrSP\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_addrTPA\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"EnrollAccount\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"GetAccount\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.Account\",\"components\":[{\"name\":\"pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetPara\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.Para\",\"components\":[{\"name\":\"Params\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"U\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"G\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RegisterFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"RegisterPara\",\"inputs\":[{\"name\":\"_params\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_g\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_u\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SearchFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accountIndexTable\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrSM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrSP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrTPA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fileIndexTable\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50604051610c8f380380610c8f833981016040819052602c916082565b600380546001600160a01b03199081163317909155600480546001600160a01b039485169083161790556005805492909316911617905560b0565b80516001600160a01b0381168114607d57600080fd5b919050565b60008060408385031215609457600080fd5b609b836067565b915060a7602084016067565b90509250929050565b610bd0806100bf6000396000f3fe608060405234801561001057600080fd5b50600436106100a85760003560e01c80636b884afb116100715780636b884afb1461014e57806379ef33df1461016157806392f4dfa214610174578063b505ed1e14610187578063b5b247e1146101b5578063d36741481461020857600080fd5b8062b0c2ed146100ad5780631f09c5e1146100d65780632b212658146100eb578063537322e31461010057806358f5823c14610123575b600080fd5b6100c06100bb366004610713565b610228565b6040516100cd919061077b565b60405180910390f35b6100de6102c6565b6040516100cd919061078e565b6100fe6100f936600461089b565b6104b2565b005b61011361010e36600461093a565b6104de565b60405190151581526020016100cd565b600354610136906001600160a01b031681565b6040516001600160a01b0390911681526020016100cd565b600454610136906001600160a01b031681565b600554610136906001600160a01b031681565b6100fe610182366004610953565b610503565b6101a761019536600461093a565b60076020526000908152604090205481565b6040519081526020016100cd565b6100fe6101c33660046109d6565b600082815260076020908152604082209384556001938401805494850181558252902090910180546001600160a01b0319166001600160a01b03909216919091179055565b61021b610216366004610713565b6105c4565b6040516100cd9190610a02565b60066020526000908152604090208054819061024390610a25565b80601f016020809104026020016040519081016040528092919081815260200182805461026f90610a25565b80156102bc5780601f10610291576101008083540402835291602001916102bc565b820191906000526020600020905b81548152906001019060200180831161029f57829003601f168201915b5050505050905081565b6102ea60405180606001604052806060815260200160608152602001606081525090565b600060405180606001604052908160008201805461030790610a25565b80601f016020809104026020016040519081016040528092919081815260200182805461033390610a25565b80156103805780601f1061035557610100808354040283529160200191610380565b820191906000526020600020905b81548152906001019060200180831161036357829003601f168201915b5050505050815260200160018201805461039990610a25565b80601f01602080910402602001604051908101604052809291908181526020018280546103c590610a25565b80156104125780601f106103e757610100808354040283529160200191610412565b820191906000526020600020905b8154815290600101906020018083116103f557829003601f168201915b5050505050815260200160028201805461042b90610a25565b80601f016020809104026020016040519081016040528092919081815260200182805461045790610a25565b80156104a45780601f10610479576101008083540402835291602001916104a4565b820191906000526020600020905b81548152906001019060200180831161048757829003601f168201915b505050505081525050905090565b60006104be8482610ab0565b5060016104cb8282610ab0565b5060026104d88382610ab0565b50505050565b60008181526007602052604081205481036104fb57506000919050565b506001919050565b6003546001600160a01b031633811461051b57600080fd5b61055a6040518060400160405280601e81526020017f456e726f6c6c205355206163636f756e742028416464726573733a25732900008152508561068d565b604080516020601f8501819004810282018301835281018481529091829190869086908190850183828082843760009201829052509390945250506001600160a01b0387168152600660205260409020825190915081906105bb9082610ab0565b50505050505050565b6040805160208082018352606082526001600160a01b03841660009081526006825283902083519182019093528254919290918290829061060490610a25565b80601f016020809104026020016040519081016040528092919081815260200182805461063090610a25565b801561067d5780601f106106525761010080835404028352916020019161067d565b820191906000526020600020905b81548152906001019060200180831161066057829003601f168201915b5050505050815250509050919050565b6106d282826040516024016106a3929190610b70565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b1790526106d6565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b80356001600160a01b038116811461070e57600080fd5b919050565b60006020828403121561072557600080fd5b61072e826106f7565b9392505050565b6000815180845260005b8181101561075b5760208185018101518683018201520161073f565b506000602082860101526020601f19601f83011685010191505092915050565b60208152600061072e6020830184610735565b6020815260008251606060208401526107aa6080840182610735565b90506020840151601f19808584030160408601526107c88383610735565b92506040860151915080858403016060860152506107e68282610735565b95945050505050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff80841115610820576108206107ef565b604051601f8501601f19908116603f01168101908282118183101715610848576108486107ef565b8160405280935085815286868601111561086157600080fd5b858560208301376000602087830101525050509392505050565b600082601f83011261088c57600080fd5b61072e83833560208501610805565b6000806000606084860312156108b057600080fd5b833567ffffffffffffffff808211156108c857600080fd5b818601915086601f8301126108dc57600080fd5b6108eb87833560208501610805565b9450602086013591508082111561090157600080fd5b61090d8783880161087b565b9350604086013591508082111561092357600080fd5b506109308682870161087b565b9150509250925092565b60006020828403121561094c57600080fd5b5035919050565b60008060006040848603121561096857600080fd5b610971846106f7565b9250602084013567ffffffffffffffff8082111561098e57600080fd5b818601915086601f8301126109a257600080fd5b8135818111156109b157600080fd5b8760208285010111156109c357600080fd5b6020830194508093505050509250925092565b600080604083850312156109e957600080fd5b823591506109f9602084016106f7565b90509250929050565b6020815260008251602080840152610a1d6040840182610735565b949350505050565b600181811c90821680610a3957607f821691505b602082108103610a5957634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610aab576000816000526020600020601f850160051c81016020861015610a885750805b601f850160051c820191505b81811015610aa757828155600101610a94565b5050505b505050565b815167ffffffffffffffff811115610aca57610aca6107ef565b610ade81610ad88454610a25565b84610a5f565b602080601f831160018114610b135760008415610afb5750858301515b600019600386901b1c1916600185901b178555610aa7565b600085815260208120601f198616915b82811015610b4257888601518255948401946001909101908401610b23565b5085821015610b605787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b604081526000610b836040830185610735565b905060018060a01b0383166020830152939250505056fea2646970667358221220e7b3fc5d3515af195d25adf1816da87217c1d64467d38728b190858b5ca7004964736f6c63430008190033",
}

// XZ21ABI is the input ABI used to generate the binding from.
// Deprecated: Use XZ21MetaData.ABI instead.
var XZ21ABI = XZ21MetaData.ABI

// XZ21Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use XZ21MetaData.Bin instead.
var XZ21Bin = XZ21MetaData.Bin

// DeployXZ21 deploys a new Ethereum contract, binding an instance of XZ21 to it.
func DeployXZ21(auth *bind.TransactOpts, backend bind.ContractBackend, _addrSP common.Address, _addrTPA common.Address) (common.Address, *types.Transaction, *XZ21, error) {
	parsed, err := XZ21MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(XZ21Bin), backend, _addrSP, _addrTPA)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XZ21{XZ21Caller: XZ21Caller{contract: contract}, XZ21Transactor: XZ21Transactor{contract: contract}, XZ21Filterer: XZ21Filterer{contract: contract}}, nil
}

// XZ21 is an auto generated Go binding around an Ethereum contract.
type XZ21 struct {
	XZ21Caller     // Read-only binding to the contract
	XZ21Transactor // Write-only binding to the contract
	XZ21Filterer   // Log filterer for contract events
}

// XZ21Caller is an auto generated read-only Go binding around an Ethereum contract.
type XZ21Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XZ21Transactor is an auto generated write-only Go binding around an Ethereum contract.
type XZ21Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XZ21Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XZ21Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XZ21Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XZ21Session struct {
	Contract     *XZ21             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XZ21CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XZ21CallerSession struct {
	Contract *XZ21Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// XZ21TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XZ21TransactorSession struct {
	Contract     *XZ21Transactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XZ21Raw is an auto generated low-level Go binding around an Ethereum contract.
type XZ21Raw struct {
	Contract *XZ21 // Generic contract binding to access the raw methods on
}

// XZ21CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XZ21CallerRaw struct {
	Contract *XZ21Caller // Generic read-only contract binding to access the raw methods on
}

// XZ21TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XZ21TransactorRaw struct {
	Contract *XZ21Transactor // Generic write-only contract binding to access the raw methods on
}

// NewXZ21 creates a new instance of XZ21, bound to a specific deployed contract.
func NewXZ21(address common.Address, backend bind.ContractBackend) (*XZ21, error) {
	contract, err := bindXZ21(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XZ21{XZ21Caller: XZ21Caller{contract: contract}, XZ21Transactor: XZ21Transactor{contract: contract}, XZ21Filterer: XZ21Filterer{contract: contract}}, nil
}

// NewXZ21Caller creates a new read-only instance of XZ21, bound to a specific deployed contract.
func NewXZ21Caller(address common.Address, caller bind.ContractCaller) (*XZ21Caller, error) {
	contract, err := bindXZ21(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XZ21Caller{contract: contract}, nil
}

// NewXZ21Transactor creates a new write-only instance of XZ21, bound to a specific deployed contract.
func NewXZ21Transactor(address common.Address, transactor bind.ContractTransactor) (*XZ21Transactor, error) {
	contract, err := bindXZ21(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XZ21Transactor{contract: contract}, nil
}

// NewXZ21Filterer creates a new log filterer instance of XZ21, bound to a specific deployed contract.
func NewXZ21Filterer(address common.Address, filterer bind.ContractFilterer) (*XZ21Filterer, error) {
	contract, err := bindXZ21(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XZ21Filterer{contract: contract}, nil
}

// bindXZ21 binds a generic wrapper to an already deployed contract.
func bindXZ21(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := XZ21MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XZ21 *XZ21Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XZ21.Contract.XZ21Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XZ21 *XZ21Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XZ21.Contract.XZ21Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XZ21 *XZ21Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XZ21.Contract.XZ21Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XZ21 *XZ21CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XZ21.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XZ21 *XZ21TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XZ21.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XZ21 *XZ21TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XZ21.Contract.contract.Transact(opts, method, params...)
}

// GetAccount is a free data retrieval call binding the contract method 0xd3674148.
//
// Solidity: function GetAccount(address _addr) view returns((bytes))
func (_XZ21 *XZ21Caller) GetAccount(opts *bind.CallOpts, _addr common.Address) (XZ21Account, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "GetAccount", _addr)

	if err != nil {
		return *new(XZ21Account), err
	}

	out0 := *abi.ConvertType(out[0], new(XZ21Account)).(*XZ21Account)

	return out0, err

}

// GetAccount is a free data retrieval call binding the contract method 0xd3674148.
//
// Solidity: function GetAccount(address _addr) view returns((bytes))
func (_XZ21 *XZ21Session) GetAccount(_addr common.Address) (XZ21Account, error) {
	return _XZ21.Contract.GetAccount(&_XZ21.CallOpts, _addr)
}

// GetAccount is a free data retrieval call binding the contract method 0xd3674148.
//
// Solidity: function GetAccount(address _addr) view returns((bytes))
func (_XZ21 *XZ21CallerSession) GetAccount(_addr common.Address) (XZ21Account, error) {
	return _XZ21.Contract.GetAccount(&_XZ21.CallOpts, _addr)
}

// GetPara is a free data retrieval call binding the contract method 0x1f09c5e1.
//
// Solidity: function GetPara() view returns((string,bytes,bytes))
func (_XZ21 *XZ21Caller) GetPara(opts *bind.CallOpts) (XZ21Para, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "GetPara")

	if err != nil {
		return *new(XZ21Para), err
	}

	out0 := *abi.ConvertType(out[0], new(XZ21Para)).(*XZ21Para)

	return out0, err

}

// GetPara is a free data retrieval call binding the contract method 0x1f09c5e1.
//
// Solidity: function GetPara() view returns((string,bytes,bytes))
func (_XZ21 *XZ21Session) GetPara() (XZ21Para, error) {
	return _XZ21.Contract.GetPara(&_XZ21.CallOpts)
}

// GetPara is a free data retrieval call binding the contract method 0x1f09c5e1.
//
// Solidity: function GetPara() view returns((string,bytes,bytes))
func (_XZ21 *XZ21CallerSession) GetPara() (XZ21Para, error) {
	return _XZ21.Contract.GetPara(&_XZ21.CallOpts)
}

// SearchFile is a free data retrieval call binding the contract method 0x537322e3.
//
// Solidity: function SearchFile(bytes32 _hash) view returns(bool)
func (_XZ21 *XZ21Caller) SearchFile(opts *bind.CallOpts, _hash [32]byte) (bool, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "SearchFile", _hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SearchFile is a free data retrieval call binding the contract method 0x537322e3.
//
// Solidity: function SearchFile(bytes32 _hash) view returns(bool)
func (_XZ21 *XZ21Session) SearchFile(_hash [32]byte) (bool, error) {
	return _XZ21.Contract.SearchFile(&_XZ21.CallOpts, _hash)
}

// SearchFile is a free data retrieval call binding the contract method 0x537322e3.
//
// Solidity: function SearchFile(bytes32 _hash) view returns(bool)
func (_XZ21 *XZ21CallerSession) SearchFile(_hash [32]byte) (bool, error) {
	return _XZ21.Contract.SearchFile(&_XZ21.CallOpts, _hash)
}

// AccountIndexTable is a free data retrieval call binding the contract method 0x00b0c2ed.
//
// Solidity: function accountIndexTable(address ) view returns(bytes pubKey)
func (_XZ21 *XZ21Caller) AccountIndexTable(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "accountIndexTable", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AccountIndexTable is a free data retrieval call binding the contract method 0x00b0c2ed.
//
// Solidity: function accountIndexTable(address ) view returns(bytes pubKey)
func (_XZ21 *XZ21Session) AccountIndexTable(arg0 common.Address) ([]byte, error) {
	return _XZ21.Contract.AccountIndexTable(&_XZ21.CallOpts, arg0)
}

// AccountIndexTable is a free data retrieval call binding the contract method 0x00b0c2ed.
//
// Solidity: function accountIndexTable(address ) view returns(bytes pubKey)
func (_XZ21 *XZ21CallerSession) AccountIndexTable(arg0 common.Address) ([]byte, error) {
	return _XZ21.Contract.AccountIndexTable(&_XZ21.CallOpts, arg0)
}

// AddrSM is a free data retrieval call binding the contract method 0x58f5823c.
//
// Solidity: function addrSM() view returns(address)
func (_XZ21 *XZ21Caller) AddrSM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "addrSM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddrSM is a free data retrieval call binding the contract method 0x58f5823c.
//
// Solidity: function addrSM() view returns(address)
func (_XZ21 *XZ21Session) AddrSM() (common.Address, error) {
	return _XZ21.Contract.AddrSM(&_XZ21.CallOpts)
}

// AddrSM is a free data retrieval call binding the contract method 0x58f5823c.
//
// Solidity: function addrSM() view returns(address)
func (_XZ21 *XZ21CallerSession) AddrSM() (common.Address, error) {
	return _XZ21.Contract.AddrSM(&_XZ21.CallOpts)
}

// AddrSP is a free data retrieval call binding the contract method 0x6b884afb.
//
// Solidity: function addrSP() view returns(address)
func (_XZ21 *XZ21Caller) AddrSP(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "addrSP")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddrSP is a free data retrieval call binding the contract method 0x6b884afb.
//
// Solidity: function addrSP() view returns(address)
func (_XZ21 *XZ21Session) AddrSP() (common.Address, error) {
	return _XZ21.Contract.AddrSP(&_XZ21.CallOpts)
}

// AddrSP is a free data retrieval call binding the contract method 0x6b884afb.
//
// Solidity: function addrSP() view returns(address)
func (_XZ21 *XZ21CallerSession) AddrSP() (common.Address, error) {
	return _XZ21.Contract.AddrSP(&_XZ21.CallOpts)
}

// AddrTPA is a free data retrieval call binding the contract method 0x79ef33df.
//
// Solidity: function addrTPA() view returns(address)
func (_XZ21 *XZ21Caller) AddrTPA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "addrTPA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddrTPA is a free data retrieval call binding the contract method 0x79ef33df.
//
// Solidity: function addrTPA() view returns(address)
func (_XZ21 *XZ21Session) AddrTPA() (common.Address, error) {
	return _XZ21.Contract.AddrTPA(&_XZ21.CallOpts)
}

// AddrTPA is a free data retrieval call binding the contract method 0x79ef33df.
//
// Solidity: function addrTPA() view returns(address)
func (_XZ21 *XZ21CallerSession) AddrTPA() (common.Address, error) {
	return _XZ21.Contract.AddrTPA(&_XZ21.CallOpts)
}

// FileIndexTable is a free data retrieval call binding the contract method 0xb505ed1e.
//
// Solidity: function fileIndexTable(bytes32 ) view returns(bytes32 hash)
func (_XZ21 *XZ21Caller) FileIndexTable(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "fileIndexTable", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FileIndexTable is a free data retrieval call binding the contract method 0xb505ed1e.
//
// Solidity: function fileIndexTable(bytes32 ) view returns(bytes32 hash)
func (_XZ21 *XZ21Session) FileIndexTable(arg0 [32]byte) ([32]byte, error) {
	return _XZ21.Contract.FileIndexTable(&_XZ21.CallOpts, arg0)
}

// FileIndexTable is a free data retrieval call binding the contract method 0xb505ed1e.
//
// Solidity: function fileIndexTable(bytes32 ) view returns(bytes32 hash)
func (_XZ21 *XZ21CallerSession) FileIndexTable(arg0 [32]byte) ([32]byte, error) {
	return _XZ21.Contract.FileIndexTable(&_XZ21.CallOpts, arg0)
}

// EnrollAccount is a paid mutator transaction binding the contract method 0x92f4dfa2.
//
// Solidity: function EnrollAccount(address _addr, bytes _pubKey) returns()
func (_XZ21 *XZ21Transactor) EnrollAccount(opts *bind.TransactOpts, _addr common.Address, _pubKey []byte) (*types.Transaction, error) {
	return _XZ21.contract.Transact(opts, "EnrollAccount", _addr, _pubKey)
}

// EnrollAccount is a paid mutator transaction binding the contract method 0x92f4dfa2.
//
// Solidity: function EnrollAccount(address _addr, bytes _pubKey) returns()
func (_XZ21 *XZ21Session) EnrollAccount(_addr common.Address, _pubKey []byte) (*types.Transaction, error) {
	return _XZ21.Contract.EnrollAccount(&_XZ21.TransactOpts, _addr, _pubKey)
}

// EnrollAccount is a paid mutator transaction binding the contract method 0x92f4dfa2.
//
// Solidity: function EnrollAccount(address _addr, bytes _pubKey) returns()
func (_XZ21 *XZ21TransactorSession) EnrollAccount(_addr common.Address, _pubKey []byte) (*types.Transaction, error) {
	return _XZ21.Contract.EnrollAccount(&_XZ21.TransactOpts, _addr, _pubKey)
}

// RegisterFile is a paid mutator transaction binding the contract method 0xb5b247e1.
//
// Solidity: function RegisterFile(bytes32 _hash, address _owner) returns()
func (_XZ21 *XZ21Transactor) RegisterFile(opts *bind.TransactOpts, _hash [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _XZ21.contract.Transact(opts, "RegisterFile", _hash, _owner)
}

// RegisterFile is a paid mutator transaction binding the contract method 0xb5b247e1.
//
// Solidity: function RegisterFile(bytes32 _hash, address _owner) returns()
func (_XZ21 *XZ21Session) RegisterFile(_hash [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _XZ21.Contract.RegisterFile(&_XZ21.TransactOpts, _hash, _owner)
}

// RegisterFile is a paid mutator transaction binding the contract method 0xb5b247e1.
//
// Solidity: function RegisterFile(bytes32 _hash, address _owner) returns()
func (_XZ21 *XZ21TransactorSession) RegisterFile(_hash [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _XZ21.Contract.RegisterFile(&_XZ21.TransactOpts, _hash, _owner)
}

// RegisterPara is a paid mutator transaction binding the contract method 0x2b212658.
//
// Solidity: function RegisterPara(string _params, bytes _g, bytes _u) returns()
func (_XZ21 *XZ21Transactor) RegisterPara(opts *bind.TransactOpts, _params string, _g []byte, _u []byte) (*types.Transaction, error) {
	return _XZ21.contract.Transact(opts, "RegisterPara", _params, _g, _u)
}

// RegisterPara is a paid mutator transaction binding the contract method 0x2b212658.
//
// Solidity: function RegisterPara(string _params, bytes _g, bytes _u) returns()
func (_XZ21 *XZ21Session) RegisterPara(_params string, _g []byte, _u []byte) (*types.Transaction, error) {
	return _XZ21.Contract.RegisterPara(&_XZ21.TransactOpts, _params, _g, _u)
}

// RegisterPara is a paid mutator transaction binding the contract method 0x2b212658.
//
// Solidity: function RegisterPara(string _params, bytes _g, bytes _u) returns()
func (_XZ21 *XZ21TransactorSession) RegisterPara(_params string, _g []byte, _u []byte) (*types.Transaction, error) {
	return _XZ21.Contract.RegisterPara(&_XZ21.TransactOpts, _params, _g, _u)
}
