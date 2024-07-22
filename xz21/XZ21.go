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

// XZ21Para is an auto generated low-level Go binding around an user-defined struct.
type XZ21Para struct {
	Pairing string
	U       []byte
	G       []byte
}

// XZ21MetaData contains all meta data concerning the XZ21 contract.
var XZ21MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_addrTPA\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_addrSP\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"AccountStatus\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"AppendAccount\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"EnrollAccount\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pubKey\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"GetAddrSM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetAddrSP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetAddrTPA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetPara\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.Para\",\"components\":[{\"name\":\"Pairing\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"U\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"G\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RegisterFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"RegisterPara\",\"inputs\":[{\"name\":\"_pairing\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_g\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_u\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SearchFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrSM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrSP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrTPA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x6080604052348015600f57600080fd5b506040516109f23803806109f2833981016040819052602c916085565b600380546001600160a01b03199081163317909155600480546001600160a01b039384169083161790556005805493909216921691909117905560b3565b80516001600160a01b0381168114608057600080fd5b919050565b60008060408385031215609757600080fd5b609e83606a565b915060aa60208401606a565b90509250929050565b610930806100c26000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80635a4682741161008c5780638a89f3d0116100665780638a89f3d0146101c757806394841488146101d8578063a2fbd776146101e9578063be73cddc1461010557600080fd5b80635a4682741461016a5780636b884afb146101a157806379ef33df146101b457600080fd5b80631f09c5e1146100d45780632b212658146100f2578063398959a314610107578063463da8be1461012c578063537322e31461014457806358f5823c14610157575b600080fd5b6100dc6101fc565b6040516100e99190610530565b60405180910390f35b610105610100366004610644565b6103e8565b005b6004546001600160a01b03165b6040516001600160a01b0390911681526020016100e9565b610134610414565b60405190151581526020016100e9565b6101346101523660046106e3565b610443565b600354610114906001600160a01b031681565b6101056101783660046106fc565b600082815260076020908152604082209384556001938401805494850181558252902090910155565b600454610114906001600160a01b031681565b600554610114906001600160a01b031681565b6005546001600160a01b0316610114565b6003546001600160a01b0316610114565b6101056101f736600461071e565b610468565b61022060405180606001604052806060815260200160608152602001606081525090565b600060405180606001604052908160008201805461023d906107af565b80601f0160208091040260200160405190810160405280929190818152602001828054610269906107af565b80156102b65780601f1061028b576101008083540402835291602001916102b6565b820191906000526020600020905b81548152906001019060200180831161029957829003601f168201915b505050505081526020016001820180546102cf906107af565b80601f01602080910402602001604051908101604052809291908181526020018280546102fb906107af565b80156103485780601f1061031d57610100808354040283529160200191610348565b820191906000526020600020905b81548152906001019060200180831161032b57829003601f168201915b50505050508152602001600282018054610361906107af565b80601f016020809104026020016040519081016040528092919081815260200182805461038d906107af565b80156103da5780601f106103af576101008083540402835291602001916103da565b820191906000526020600020905b8154815290600101906020018083116103bd57829003601f168201915b505050505081525050905090565b60006103f4848261083a565b506001610401828261083a565b50600261040e838261083a565b50505050565b336000908152600660205260408120805461042e906107af565b905060000361043d5750600090565b50600190565b600081815260076020526040812054810361046057506000919050565b506001919050565b6003546001600160a01b031633811461048057600080fd5b604080516020601f8501819004810282018301835281018481529091829190869086908190850183828082843760009201829052509390945250506001600160a01b0387168152600660205260409020825190915081906104e1908261083a565b50505050505050565b6000815180845260005b81811015610510576020818501810151868301820152016104f4565b506000602082860101526020601f19601f83011685010191505092915050565b60208152600082516060602084015261054c60808401826104ea565b90506020840151601f198085840301604086015261056a83836104ea565b925060408601519150808584030160608601525061058882826104ea565b95945050505050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff808411156105c2576105c2610591565b604051601f8501601f19908116603f011681019082821181831017156105ea576105ea610591565b8160405280935085815286868601111561060357600080fd5b858560208301376000602087830101525050509392505050565b600082601f83011261062e57600080fd5b61063d838335602085016105a7565b9392505050565b60008060006060848603121561065957600080fd5b833567ffffffffffffffff8082111561067157600080fd5b818601915086601f83011261068557600080fd5b610694878335602085016105a7565b945060208601359150808211156106aa57600080fd5b6106b68783880161061d565b935060408601359150808211156106cc57600080fd5b506106d98682870161061d565b9150509250925092565b6000602082840312156106f557600080fd5b5035919050565b6000806040838503121561070f57600080fd5b50508035926020909101359150565b60008060006040848603121561073357600080fd5b83356001600160a01b038116811461074a57600080fd5b9250602084013567ffffffffffffffff8082111561076757600080fd5b818601915086601f83011261077b57600080fd5b81358181111561078a57600080fd5b87602082850101111561079c57600080fd5b6020830194508093505050509250925092565b600181811c908216806107c357607f821691505b6020821081036107e357634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610835576000816000526020600020601f850160051c810160208610156108125750805b601f850160051c820191505b818110156108315782815560010161081e565b5050505b505050565b815167ffffffffffffffff81111561085457610854610591565b6108688161086284546107af565b846107e9565b602080601f83116001811461089d57600084156108855750858301515b600019600386901b1c1916600185901b178555610831565b600085815260208120601f198616915b828110156108cc578886015182559484019460019091019084016108ad565b50858210156108ea5787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea264697066735822122086a5e52e73530f4951312152a0ef5da83b48f659919cf98a14ca8fe669054dd364736f6c63430008190033",
}

// XZ21ABI is the input ABI used to generate the binding from.
// Deprecated: Use XZ21MetaData.ABI instead.
var XZ21ABI = XZ21MetaData.ABI

// XZ21Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use XZ21MetaData.Bin instead.
var XZ21Bin = XZ21MetaData.Bin

// DeployXZ21 deploys a new Ethereum contract, binding an instance of XZ21 to it.
func DeployXZ21(auth *bind.TransactOpts, backend bind.ContractBackend, _addrTPA common.Address, _addrSP common.Address) (common.Address, *types.Transaction, *XZ21, error) {
	parsed, err := XZ21MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(XZ21Bin), backend, _addrTPA, _addrSP)
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

// AccountStatus is a free data retrieval call binding the contract method 0x463da8be.
//
// Solidity: function AccountStatus() view returns(bool)
func (_XZ21 *XZ21Caller) AccountStatus(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "AccountStatus")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AccountStatus is a free data retrieval call binding the contract method 0x463da8be.
//
// Solidity: function AccountStatus() view returns(bool)
func (_XZ21 *XZ21Session) AccountStatus() (bool, error) {
	return _XZ21.Contract.AccountStatus(&_XZ21.CallOpts)
}

// AccountStatus is a free data retrieval call binding the contract method 0x463da8be.
//
// Solidity: function AccountStatus() view returns(bool)
func (_XZ21 *XZ21CallerSession) AccountStatus() (bool, error) {
	return _XZ21.Contract.AccountStatus(&_XZ21.CallOpts)
}

// GetAddrSM is a free data retrieval call binding the contract method 0x94841488.
//
// Solidity: function GetAddrSM() view returns(address)
func (_XZ21 *XZ21Caller) GetAddrSM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "GetAddrSM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddrSM is a free data retrieval call binding the contract method 0x94841488.
//
// Solidity: function GetAddrSM() view returns(address)
func (_XZ21 *XZ21Session) GetAddrSM() (common.Address, error) {
	return _XZ21.Contract.GetAddrSM(&_XZ21.CallOpts)
}

// GetAddrSM is a free data retrieval call binding the contract method 0x94841488.
//
// Solidity: function GetAddrSM() view returns(address)
func (_XZ21 *XZ21CallerSession) GetAddrSM() (common.Address, error) {
	return _XZ21.Contract.GetAddrSM(&_XZ21.CallOpts)
}

// GetAddrSP is a free data retrieval call binding the contract method 0x398959a3.
//
// Solidity: function GetAddrSP() view returns(address)
func (_XZ21 *XZ21Caller) GetAddrSP(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "GetAddrSP")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddrSP is a free data retrieval call binding the contract method 0x398959a3.
//
// Solidity: function GetAddrSP() view returns(address)
func (_XZ21 *XZ21Session) GetAddrSP() (common.Address, error) {
	return _XZ21.Contract.GetAddrSP(&_XZ21.CallOpts)
}

// GetAddrSP is a free data retrieval call binding the contract method 0x398959a3.
//
// Solidity: function GetAddrSP() view returns(address)
func (_XZ21 *XZ21CallerSession) GetAddrSP() (common.Address, error) {
	return _XZ21.Contract.GetAddrSP(&_XZ21.CallOpts)
}

// GetAddrTPA is a free data retrieval call binding the contract method 0x8a89f3d0.
//
// Solidity: function GetAddrTPA() view returns(address)
func (_XZ21 *XZ21Caller) GetAddrTPA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "GetAddrTPA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddrTPA is a free data retrieval call binding the contract method 0x8a89f3d0.
//
// Solidity: function GetAddrTPA() view returns(address)
func (_XZ21 *XZ21Session) GetAddrTPA() (common.Address, error) {
	return _XZ21.Contract.GetAddrTPA(&_XZ21.CallOpts)
}

// GetAddrTPA is a free data retrieval call binding the contract method 0x8a89f3d0.
//
// Solidity: function GetAddrTPA() view returns(address)
func (_XZ21 *XZ21CallerSession) GetAddrTPA() (common.Address, error) {
	return _XZ21.Contract.GetAddrTPA(&_XZ21.CallOpts)
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

// AppendAccount is a paid mutator transaction binding the contract method 0xbe73cddc.
//
// Solidity: function AppendAccount() returns()
func (_XZ21 *XZ21Transactor) AppendAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XZ21.contract.Transact(opts, "AppendAccount")
}

// AppendAccount is a paid mutator transaction binding the contract method 0xbe73cddc.
//
// Solidity: function AppendAccount() returns()
func (_XZ21 *XZ21Session) AppendAccount() (*types.Transaction, error) {
	return _XZ21.Contract.AppendAccount(&_XZ21.TransactOpts)
}

// AppendAccount is a paid mutator transaction binding the contract method 0xbe73cddc.
//
// Solidity: function AppendAccount() returns()
func (_XZ21 *XZ21TransactorSession) AppendAccount() (*types.Transaction, error) {
	return _XZ21.Contract.AppendAccount(&_XZ21.TransactOpts)
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
