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
	PubKey string
}

// XZ21Para is an auto generated low-level Go binding around an user-defined struct.
type XZ21Para struct {
	Pairing string
	U       []byte
	G       []byte
}

// XZ21MetaData contains all meta data concerning the XZ21 contract.
var XZ21MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_addrSP\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_addrTPA\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"EnrollAccount\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pubKey\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"GetAccount\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.Account\",\"components\":[{\"name\":\"pubKey\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetPara\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.Para\",\"components\":[{\"name\":\"Pairing\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"U\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"G\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RegisterFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"RegisterPara\",\"inputs\":[{\"name\":\"_pairing\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_g\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_u\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SearchFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accountIndexTable\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"pubKey\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrSM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrSP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrTPA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fileIndexTable\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50604051610cc5380380610cc5833981016040819052602c916082565b600380546001600160a01b03199081163317909155600480546001600160a01b039485169083161790556005805492909316911617905560b0565b80516001600160a01b0381168114607d57600080fd5b919050565b60008060408385031215609457600080fd5b609b836067565b915060a7602084016067565b90509250929050565b610c06806100bf6000396000f3fe608060405234801561001057600080fd5b50600436106100a85760003560e01c80635a468274116100715780635a4682741461014e5780636b884afb1461018557806379ef33df14610198578063a2fbd776146101ab578063b505ed1e146101be578063d3674148146101ec57600080fd5b8062b0c2ed146100ad5780631f09c5e1146100d65780632b212658146100eb578063537322e31461010057806358f5823c14610123575b600080fd5b6100c06100bb366004610714565b61020c565b6040516100cd919061077c565b60405180910390f35b6100de6102aa565b6040516100cd919061078f565b6100fe6100f936600461089c565b610496565b005b61011361010e36600461093b565b6104c2565b60405190151581526020016100cd565b600354610136906001600160a01b031681565b6040516001600160a01b0390911681526020016100cd565b6100fe61015c366004610954565b600082815260076020908152604082209384556001938401805494850181558252902090910155565b600454610136906001600160a01b031681565b600554610136906001600160a01b031681565b6100fe6101b9366004610976565b6104e7565b6101de6101cc36600461093b565b60076020526000908152604090205481565b6040519081526020016100cd565b6101ff6101fa366004610714565b6105c2565b6040516100cd91906109f9565b60066020526000908152604090208054819061022790610a1c565b80601f016020809104026020016040519081016040528092919081815260200182805461025390610a1c565b80156102a05780601f10610275576101008083540402835291602001916102a0565b820191906000526020600020905b81548152906001019060200180831161028357829003601f168201915b5050505050905081565b6102ce60405180606001604052806060815260200160608152602001606081525090565b60006040518060600160405290816000820180546102eb90610a1c565b80601f016020809104026020016040519081016040528092919081815260200182805461031790610a1c565b80156103645780601f1061033957610100808354040283529160200191610364565b820191906000526020600020905b81548152906001019060200180831161034757829003601f168201915b5050505050815260200160018201805461037d90610a1c565b80601f01602080910402602001604051908101604052809291908181526020018280546103a990610a1c565b80156103f65780601f106103cb576101008083540402835291602001916103f6565b820191906000526020600020905b8154815290600101906020018083116103d957829003601f168201915b5050505050815260200160028201805461040f90610a1c565b80601f016020809104026020016040519081016040528092919081815260200182805461043b90610a1c565b80156104885780601f1061045d57610100808354040283529160200191610488565b820191906000526020600020905b81548152906001019060200180831161046b57829003601f168201915b505050505081525050905090565b60006104a28482610aa6565b5060016104af8282610aa6565b5060026104bc8382610aa6565b50505050565b60008181526007602052604081205481036104df57506000919050565b506001919050565b6003546001600160a01b03163381146104ff57600080fd5b6105586040518060600160405280602c8152602001610ba5602c91398585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061068b92505050565b604080516020601f8501819004810282018301835281018481529091829190869086908190850183828082843760009201829052509390945250506001600160a01b0387168152600660205260409020825190915081906105b99082610aa6565b50505050505050565b6040805160208082018352606082526001600160a01b03841660009081526006825283902083519182019093528254919290918290829061060290610a1c565b80601f016020809104026020016040519081016040528092919081815260200182805461062e90610a1c565b801561067b5780601f106106505761010080835404028352916020019161067b565b820191906000526020600020905b81548152906001019060200180831161065e57829003601f168201915b5050505050815250509050919050565b6106d28383836040516024016106a393929190610b66565b60408051601f198184030181529190526020810180516001600160e01b031663e0e9ad4f60e01b1790526106d7565b505050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b80356001600160a01b038116811461070f57600080fd5b919050565b60006020828403121561072657600080fd5b61072f826106f8565b9392505050565b6000815180845260005b8181101561075c57602081850181015186830182015201610740565b506000602082860101526020601f19601f83011685010191505092915050565b60208152600061072f6020830184610736565b6020815260008251606060208401526107ab6080840182610736565b90506020840151601f19808584030160408601526107c98383610736565b92506040860151915080858403016060860152506107e78282610736565b95945050505050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff80841115610821576108216107f0565b604051601f8501601f19908116603f01168101908282118183101715610849576108496107f0565b8160405280935085815286868601111561086257600080fd5b858560208301376000602087830101525050509392505050565b600082601f83011261088d57600080fd5b61072f83833560208501610806565b6000806000606084860312156108b157600080fd5b833567ffffffffffffffff808211156108c957600080fd5b818601915086601f8301126108dd57600080fd5b6108ec87833560208501610806565b9450602086013591508082111561090257600080fd5b61090e8783880161087c565b9350604086013591508082111561092457600080fd5b506109318682870161087c565b9150509250925092565b60006020828403121561094d57600080fd5b5035919050565b6000806040838503121561096757600080fd5b50508035926020909101359150565b60008060006040848603121561098b57600080fd5b610994846106f8565b9250602084013567ffffffffffffffff808211156109b157600080fd5b818601915086601f8301126109c557600080fd5b8135818111156109d457600080fd5b8760208285010111156109e657600080fd5b6020830194508093505050509250925092565b6020815260008251602080840152610a146040840182610736565b949350505050565b600181811c90821680610a3057607f821691505b602082108103610a5057634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156106d2576000816000526020600020601f850160051c81016020861015610a7f5750805b601f850160051c820191505b81811015610a9e57828155600101610a8b565b505050505050565b815167ffffffffffffffff811115610ac057610ac06107f0565b610ad481610ace8454610a1c565b84610a56565b602080601f831160018114610b095760008415610af15750858301515b600019600386901b1c1916600185901b178555610a9e565b600085815260208120601f198616915b82811015610b3857888601518255948401946001909101908401610b19565b5085821015610b565787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b606081526000610b796060830186610736565b6001600160a01b03851660208401528281036040840152610b9a8185610736565b969550505050505056fe456e726f6c6c205355206163636f756e742028416464726573733a25732c205075626c69634b65793a257329a2646970667358221220631fe2d3c109253c167d61bf048fc35f272aa0c080033fc19dd1d7a7d2eddab664736f6c63430008190033",
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
// Solidity: function GetAccount(address _addr) view returns((string))
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
// Solidity: function GetAccount(address _addr) view returns((string))
func (_XZ21 *XZ21Session) GetAccount(_addr common.Address) (XZ21Account, error) {
	return _XZ21.Contract.GetAccount(&_XZ21.CallOpts, _addr)
}

// GetAccount is a free data retrieval call binding the contract method 0xd3674148.
//
// Solidity: function GetAccount(address _addr) view returns((string))
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
// Solidity: function accountIndexTable(address ) view returns(string pubKey)
func (_XZ21 *XZ21Caller) AccountIndexTable(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "accountIndexTable", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AccountIndexTable is a free data retrieval call binding the contract method 0x00b0c2ed.
//
// Solidity: function accountIndexTable(address ) view returns(string pubKey)
func (_XZ21 *XZ21Session) AccountIndexTable(arg0 common.Address) (string, error) {
	return _XZ21.Contract.AccountIndexTable(&_XZ21.CallOpts, arg0)
}

// AccountIndexTable is a free data retrieval call binding the contract method 0x00b0c2ed.
//
// Solidity: function accountIndexTable(address ) view returns(string pubKey)
func (_XZ21 *XZ21CallerSession) AccountIndexTable(arg0 common.Address) (string, error) {
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

// EnrollAccount is a paid mutator transaction binding the contract method 0xa2fbd776.
//
// Solidity: function EnrollAccount(address _addr, string _pubKey) returns()
func (_XZ21 *XZ21Transactor) EnrollAccount(opts *bind.TransactOpts, _addr common.Address, _pubKey string) (*types.Transaction, error) {
	return _XZ21.contract.Transact(opts, "EnrollAccount", _addr, _pubKey)
}

// EnrollAccount is a paid mutator transaction binding the contract method 0xa2fbd776.
//
// Solidity: function EnrollAccount(address _addr, string _pubKey) returns()
func (_XZ21 *XZ21Session) EnrollAccount(_addr common.Address, _pubKey string) (*types.Transaction, error) {
	return _XZ21.Contract.EnrollAccount(&_XZ21.TransactOpts, _addr, _pubKey)
}

// EnrollAccount is a paid mutator transaction binding the contract method 0xa2fbd776.
//
// Solidity: function EnrollAccount(address _addr, string _pubKey) returns()
func (_XZ21 *XZ21TransactorSession) EnrollAccount(_addr common.Address, _pubKey string) (*types.Transaction, error) {
	return _XZ21.Contract.EnrollAccount(&_XZ21.TransactOpts, _addr, _pubKey)
}

// RegisterFile is a paid mutator transaction binding the contract method 0x5a468274.
//
// Solidity: function RegisterFile(bytes32 _hash, bytes32 _id) returns()
func (_XZ21 *XZ21Transactor) RegisterFile(opts *bind.TransactOpts, _hash [32]byte, _id [32]byte) (*types.Transaction, error) {
	return _XZ21.contract.Transact(opts, "RegisterFile", _hash, _id)
}

// RegisterFile is a paid mutator transaction binding the contract method 0x5a468274.
//
// Solidity: function RegisterFile(bytes32 _hash, bytes32 _id) returns()
func (_XZ21 *XZ21Session) RegisterFile(_hash [32]byte, _id [32]byte) (*types.Transaction, error) {
	return _XZ21.Contract.RegisterFile(&_XZ21.TransactOpts, _hash, _id)
}

// RegisterFile is a paid mutator transaction binding the contract method 0x5a468274.
//
// Solidity: function RegisterFile(bytes32 _hash, bytes32 _id) returns()
func (_XZ21 *XZ21TransactorSession) RegisterFile(_hash [32]byte, _id [32]byte) (*types.Transaction, error) {
	return _XZ21.Contract.RegisterFile(&_XZ21.TransactOpts, _hash, _id)
}

// RegisterPara is a paid mutator transaction binding the contract method 0x2b212658.
//
// Solidity: function RegisterPara(string _pairing, bytes _g, bytes _u) returns()
func (_XZ21 *XZ21Transactor) RegisterPara(opts *bind.TransactOpts, _pairing string, _g []byte, _u []byte) (*types.Transaction, error) {
	return _XZ21.contract.Transact(opts, "RegisterPara", _pairing, _g, _u)
}

// RegisterPara is a paid mutator transaction binding the contract method 0x2b212658.
//
// Solidity: function RegisterPara(string _pairing, bytes _g, bytes _u) returns()
func (_XZ21 *XZ21Session) RegisterPara(_pairing string, _g []byte, _u []byte) (*types.Transaction, error) {
	return _XZ21.Contract.RegisterPara(&_XZ21.TransactOpts, _pairing, _g, _u)
}

// RegisterPara is a paid mutator transaction binding the contract method 0x2b212658.
//
// Solidity: function RegisterPara(string _pairing, bytes _g, bytes _u) returns()
func (_XZ21 *XZ21TransactorSession) RegisterPara(_pairing string, _g []byte, _u []byte) (*types.Transaction, error) {
	return _XZ21.Contract.RegisterPara(&_XZ21.TransactOpts, _pairing, _g, _u)
}
