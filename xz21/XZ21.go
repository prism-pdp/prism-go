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
	PubKey   []byte
	FileList [][32]byte
}

// XZ21AuditingLog is an auto generated low-level Go binding around an user-defined struct.
type XZ21AuditingLog struct {
	Req    XZ21AuditingReq
	Result bool
	Date   *big.Int
}

// XZ21AuditingReq is an auto generated low-level Go binding around an user-defined struct.
type XZ21AuditingReq struct {
	Chal  []byte
	Proof []byte
	Stage uint8
}

// XZ21FileProperty is an auto generated low-level Go binding around an user-defined struct.
type XZ21FileProperty struct {
	SplitNum uint32
	Creator  common.Address
}

// XZ21Param is an auto generated low-level Go binding around an user-defined struct.
type XZ21Param struct {
	P string
	U []byte
	G []byte
}

// XZ21MetaData contains all meta data concerning the XZ21 contract.
var XZ21MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_addrSP\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"AppendOwner\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"EnrollAccount\",\"inputs\":[{\"name\":\"_type\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"GetAuditingLogs\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structXZ21.AuditingLog[]\",\"components\":[{\"name\":\"req\",\"type\":\"tuple\",\"internalType\":\"structXZ21.AuditingReq\",\"components\":[{\"name\":\"chal\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"stage\",\"type\":\"uint8\",\"internalType\":\"enumXZ21.Stages\"}]},{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"date\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetAuditingReq\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.AuditingReq\",\"components\":[{\"name\":\"chal\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"stage\",\"type\":\"uint8\",\"internalType\":\"enumXZ21.Stages\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetAuditorAddrList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetFileList\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetParam\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.Param\",\"components\":[{\"name\":\"P\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"U\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"G\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetUserAccount\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.Account\",\"components\":[{\"name\":\"pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fileList\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RegisterFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_splitNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"RegisterParam\",\"inputs\":[{\"name\":\"_p\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_g\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_u\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SearchFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.FileProperty\",\"components\":[{\"name\":\"splitNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SetAuditingResult\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SetChal\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_chal\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SetProof\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addrSM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrSP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"auditorAddrList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EventSetAuditingResult\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
	Bin: "0x60c0604052348015600f57600080fd5b506040516121e63803806121e6833981016040819052602c91604a565b336080526001600160a01b031660a0526003805460ff191690556078565b600060208284031215605b57600080fd5b81516001600160a01b0381168114607157600080fd5b9392505050565b60805160a0516121266100c0600039600081816102420152818161066b015281816114fd01526116f80152600081816102060152818161035c015261129301526121266000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80638a04d5e5116100a2578063c1f16b1611610071578063c1f16b16146102ec578063c4e197ec146102ff578063d746cdd01461031f578063debffd9a14610332578063f64b362a1461034557600080fd5b80638a04d5e5146102845780639a2b5e6314610299578063b55957a0146102b9578063b645fd63146102d957600080fd5b806358f5823c116100de57806358f5823c146102015780635a955455146102285780636b884afb1461023d5780637f8382111461026457600080fd5b8063043c90771461011057806322dd2b18146101385780632cc19d971461014d578063537322e314610178575b600080fd5b61012361011e3660046118ec565b610358565b60405190151581526020015b60405180910390f35b61014b610146366004611946565b610669565b005b61016061015b36600461198d565b610760565b6040516001600160a01b03909116815260200161012f565b6101d761018636600461198d565b60408051808201909152600080825260208201525060009081526006602090815260409182902082518084019093525463ffffffff8116835264010000000090046001600160a01b03169082015290565b60408051825163ffffffff1681526020928301516001600160a01b0316928101929092520161012f565b6101607f000000000000000000000000000000000000000000000000000000000000000081565b61023061078a565b60405161012f91906119a6565b6101607f000000000000000000000000000000000000000000000000000000000000000081565b61027761027236600461198d565b6107ec565b60405161012f9190611ab4565b61028c6109de565b60405161012f9190611b39565b6102ac6102a7366004611b9a565b610bca565b60405161012f9190611bf8565b6102cc6102c7366004611b9a565b610cb9565b60405161012f9190611c0b565b61014b6102e7366004611c44565b610ddf565b61014b6102fa366004611d25565b611291565b61031261030d36600461198d565b611367565b60405161012f9190611dc4565b61014b61032d366004611dd7565b6114fb565b61014b610340366004611dd7565b6115ed565b61014b610353366004611e23565b6116f6565b60007f0000000000000000000000000000000000000000000000000000000000000000336001600160a01b038216146103ac5760405162461bcd60e51b81526004016103a390611e4f565b60405180910390fd5b8515806103b95750856001145b6103f45760405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964207479706560a01b60448201526064016103a3565b85600003610529576000805b60045481101561045257866001600160a01b03166004828154811061042757610427611e7d565b6000918252602090912001546001600160a01b03160361044a5760019150610452565b600101610400565b5080156104995760405162461bcd60e51b81526020600482015260156024820152744475706c696361746520545041206164647265737360581b60448201526064016103a3565b6104d86040518060400160405280601f81526020017f456e726f6c6c20545041206163636f756e742028416464726573733a25732900815250876117bd565b50600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180546001600160a01b0319166001600160a01b03871617905561065d565b6001600160a01b0385166000908152600560205260409020805461054c90611e93565b1590506105925760405162461bcd60e51b8152602060048201526014602482015273111d5c1b1a58d85d194814d5481858d8dbdd5b9d60621b60448201526064016103a3565b6105d16040518060400160405280601e81526020017f456e726f6c6c205355206163636f756e742028416464726573733a2573290000815250866117bd565b6040805160606020601f87018190040282018101835291810185815290918291908790879081908501838280828437600092018290525093855250506040805183815260208082018352948501526001600160a01b038a16835260059093525020815181906106409082611f1e565b5060208281015180516106599260018501920190611827565b5050505b50600195945050505050565b7f0000000000000000000000000000000000000000000000000000000000000000336001600160a01b038216146106b25760405162461bcd60e51b81526004016103a390611e4f565b60008363ffffffff16116106fc5760405162461bcd60e51b8152602060048201526011602482015270696e76616c69642073706c6974206e756d60781b60448201526064016103a3565b506000838152600660209081526040808320805463ffffffff969096166001600160c01b0319909616959095176401000000006001600160a01b03959095169485021790945591815260058252918220600190810180549182018155835291200155565b6004818154811061077057600080fd5b6000918252602090912001546001600160a01b0316905081565b606060048054806020026020016040519081016040528092919081815260200182805480156107e257602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116107c4575b5050505050905090565b606060086000838152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156109d3576000848152602090206040805160c081019091526005840290910180548290606082019083908290829061085c90611e93565b80601f016020809104026020016040519081016040528092919081815260200182805461088890611e93565b80156108d55780601f106108aa576101008083540402835291602001916108d5565b820191906000526020600020905b8154815290600101906020018083116108b857829003601f168201915b505050505081526020016001820180546108ee90611e93565b80601f016020809104026020016040519081016040528092919081815260200182805461091a90611e93565b80156109675780601f1061093c57610100808354040283529160200191610967565b820191906000526020600020905b81548152906001019060200180831161094a57829003601f168201915b505050918352505060028281015460209092019160ff169081111561098e5761098e611a39565b600281111561099f5761099f611a39565b9052508152600382015460ff1615156020808301919091526004909201546040909101529082526001929092019101610821565b505050509050919050565b610a0260405180606001604052806060815260200160608152602001606081525090565b6000604051806060016040529081600082018054610a1f90611e93565b80601f0160208091040260200160405190810160405280929190818152602001828054610a4b90611e93565b8015610a985780601f10610a6d57610100808354040283529160200191610a98565b820191906000526020600020905b815481529060010190602001808311610a7b57829003601f168201915b50505050508152602001600182018054610ab190611e93565b80601f0160208091040260200160405190810160405280929190818152602001828054610add90611e93565b8015610b2a5780601f10610aff57610100808354040283529160200191610b2a565b820191906000526020600020905b815481529060010190602001808311610b0d57829003601f168201915b50505050508152602001600282018054610b4390611e93565b80601f0160208091040260200160405190810160405280929190818152602001828054610b6f90611e93565b8015610bbc5780601f10610b9157610100808354040283529160200191610bbc565b820191906000526020600020905b815481529060010190602001808311610b9f57829003601f168201915b505050505081525050905090565b6001600160a01b0381166000908152600560205260408120600101546060918167ffffffffffffffff811115610c0257610c02611c79565b604051908082528060200260200182016040528015610c2b578160200160208202803683370190505b50905060005b6001600160a01b038516600090815260056020526040902060010154811015610cb1576001600160a01b0385166000908152600560205260409020600101805482908110610c8157610c81611e7d565b9060005260206000200154828281518110610c9e57610c9e611e7d565b6020908102919091010152600101610c31565b509392505050565b60408051808201909152606080825260208201526001600160a01b038216600090815260056020526040908190208151808301909252805482908290610cfe90611e93565b80601f0160208091040260200160405190810160405280929190818152602001828054610d2a90611e93565b8015610d775780601f10610d4c57610100808354040283529160200191610d77565b820191906000526020600020905b815481529060010190602001808311610d5a57829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015610dcf57602002820191906000526020600020905b815481526020019060010190808311610dbb575b5050505050815250509050919050565b6000805b600454811015610e3557336001600160a01b031660048281548110610e0a57610e0a611e7d565b6000918252602090912001546001600160a01b031603610e2d5760019150610e35565b600101610de3565b5080610e835760405162461bcd60e51b815260206004820152601860248201527f5450412061757468656e7469636174696f6e206572726f72000000000000000060448201526064016103a3565b600083815260076020526040902060029081015460ff1681811115610eaa57610eaa611a39565b14610eeb5760405162461bcd60e51b8152602060048201526011602482015270139bdd0815d85a5d1a5b99d4995cdd5b1d607a1b60448201526064016103a3565b60008381526008602052604090205415610f9257600083815260086020526040812054610f1a90600190611fde565b60008581526008602052604090208054919250429183908110610f3f57610f3f611e7d565b90600052602060002090600502016004015410610f905760405162461bcd60e51b815260206004820152600f60248201526e3a34b6b2b9ba30b6b81032b93937b960891b60448201526064016103a3565b505b604080516000858152600760205282812060c0830190935282549092829160608301919082908290610fc390611e93565b80601f0160208091040260200160405190810160405280929190818152602001828054610fef90611e93565b801561103c5780601f106110115761010080835404028352916020019161103c565b820191906000526020600020905b81548152906001019060200180831161101f57829003601f168201915b5050505050815260200160018201805461105590611e93565b80601f016020809104026020016040519081016040528092919081815260200182805461108190611e93565b80156110ce5780601f106110a3576101008083540402835291602001916110ce565b820191906000526020600020905b8154815290600101906020018083116110b157829003601f168201915b505050918352505060028281015460209092019160ff16908111156110f5576110f5611a39565b600281111561110657611106611a39565b90525081528415156020808301919091524260409283015260008781526008825291822080546001810182559083529120825180519394508493600590930290910191829081906111579082611f1e565b506020820151600182019061116c9082611f1e565b5060408201518160020160006101000a81548160ff0219169083600281111561119757611197611a39565b0217905550505060208281015160038301805460ff19169115159190911790556040928301516004909201919091558151608081018352600060608201818152825283518084018552818152928201929092529182015260008581526007602052604090208151819061120a9082611f1e565b506020820151600182019061121f9082611f1e565b5060408201518160020160006101000a81548160ff0219169083600281111561124a5761124a611a39565b0217905550506040805186815285151560208201527f961a406eff4d590fdb9e54bdba584518be2e6a7f91b34d5714330f83d571b03192500160405180910390a150505050565b7f0000000000000000000000000000000000000000000000000000000000000000336001600160a01b038216146112da5760405162461bcd60e51b81526004016103a390611e4f565b60035460ff161561132d5760405162461bcd60e51b815260206004820152601e60248201527f446f206e6f74206f7665727772697465205265676973746572506172616d000060448201526064016103a3565b60006113398582611f1e565b5060016113468382611f1e565b5060026113538482611f1e565b50506003805460ff19166001179055505050565b611387604080516060808201835280825260208201529081016000905290565b600082815260076020526040908190208151606081019092528054829082906113af90611e93565b80601f01602080910402602001604051908101604052809291908181526020018280546113db90611e93565b80156114285780601f106113fd57610100808354040283529160200191611428565b820191906000526020600020905b81548152906001019060200180831161140b57829003601f168201915b5050505050815260200160018201805461144190611e93565b80601f016020809104026020016040519081016040528092919081815260200182805461146d90611e93565b80156114ba5780601f1061148f576101008083540402835291602001916114ba565b820191906000526020600020905b81548152906001019060200180831161149d57829003601f168201915b505050918352505060028281015460209092019160ff16908111156114e1576114e1611a39565b60028111156114f2576114f2611a39565b90525092915050565b7f0000000000000000000000000000000000000000000000000000000000000000336001600160a01b038216146115445760405162461bcd60e51b81526004016103a390611e4f565b6001600085815260076020526040902060029081015460ff169081111561156d5761156d611a39565b146115ad5760405162461bcd60e51b815260206004820152601060248201526f2737ba102bb0b4ba34b733a83937b7b360811b60448201526064016103a3565b60008481526007602052604090206001016115c9838583612005565b5050506000918252506007602052604090206002908101805460ff19169091179055565b336000908152600560205260408120805461160790611e93565b9050116116565760405162461bcd60e51b815260206004820152601760248201527f53552061757468656e7469636174696f6e206572726f7200000000000000000060448201526064016103a3565b600083815260076020526040812060029081015460ff169081111561167d5761167d611a39565b146116bc5760405162461bcd60e51b815260206004820152600f60248201526e139bdd0815d85a5d1a5b99d0da185b608a1b60448201526064016103a3565b60008381526007602052604090206116d5828483612005565b5050506000908152600760205260409020600201805460ff19166001179055565b7f0000000000000000000000000000000000000000000000000000000000000000336001600160a01b0382161461173f5760405162461bcd60e51b81526004016103a390611e4f565b60008381526006602052604090205463ffffffff1661178f5760405162461bcd60e51b815260206004820152600c60248201526b696e76616c69642066696c6560a01b60448201526064016103a3565b506001600160a01b031660009081526005602090815260408220600190810180549182018155835291200155565b61180282826040516024016117d39291906120c6565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b179052611806565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b828054828255906000526020600020908101928215611862579160200282015b82811115611862578251825591602001919060010190611847565b5061186e929150611872565b5090565b5b8082111561186e5760008155600101611873565b80356001600160a01b038116811461189e57600080fd5b919050565b60008083601f8401126118b557600080fd5b50813567ffffffffffffffff8111156118cd57600080fd5b6020830191508360208285010111156118e557600080fd5b9250929050565b6000806000806060858703121561190257600080fd5b8435935061191260208601611887565b9250604085013567ffffffffffffffff81111561192e57600080fd5b61193a878288016118a3565b95989497509550505050565b60008060006060848603121561195b57600080fd5b83359250602084013563ffffffff8116811461197657600080fd5b915061198460408501611887565b90509250925092565b60006020828403121561199f57600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b818110156119e75783516001600160a01b0316835292840192918401916001016119c2565b50909695505050505050565b6000815180845260005b81811015611a19576020818501810151868301820152016119fd565b506000602082860101526020601f19601f83011685010191505092915050565b634e487b7160e01b600052602160045260246000fd5b6000815160608452611a6460608501826119f3565b905060208301518482036020860152611a7d82826119f3565b915050604083015160038110611aa357634e487b7160e01b600052602160045260246000fd5b806040860152508091505092915050565b600060208083018184528085518083526040925060408601915060408160051b87010184880160005b83811015611b2b57603f19898403018552815160608151818652611b0382870182611a4f565b838b01511515878c015292890151958901959095525094870194925090860190600101611add565b509098975050505050505050565b602081526000825160606020840152611b5560808401826119f3565b90506020840151601f1980858403016040860152611b7383836119f3565b9250604086015191508085840301606086015250611b9182826119f3565b95945050505050565b600060208284031215611bac57600080fd5b611bb582611887565b9392505050565b60008151808452602080850194506020840160005b83811015611bed57815187529582019590820190600101611bd1565b509495945050505050565b602081526000611bb56020830184611bbc565b602081526000825160406020840152611c2760608401826119f3565b90506020840151601f19848303016040850152611b918282611bbc565b60008060408385031215611c5757600080fd5b8235915060208301358015158114611c6e57600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff80841115611caa57611caa611c79565b604051601f8501601f19908116603f01168101908282118183101715611cd257611cd2611c79565b81604052809350858152868686011115611ceb57600080fd5b858560208301376000602087830101525050509392505050565b600082601f830112611d1657600080fd5b611bb583833560208501611c8f565b600080600060608486031215611d3a57600080fd5b833567ffffffffffffffff80821115611d5257600080fd5b818601915086601f830112611d6657600080fd5b611d7587833560208501611c8f565b94506020860135915080821115611d8b57600080fd5b611d9787838801611d05565b93506040860135915080821115611dad57600080fd5b50611dba86828701611d05565b9150509250925092565b602081526000611bb56020830184611a4f565b600080600060408486031215611dec57600080fd5b83359250602084013567ffffffffffffffff811115611e0a57600080fd5b611e16868287016118a3565b9497909650939450505050565b60008060408385031215611e3657600080fd5b82359150611e4660208401611887565b90509250929050565b60208082526014908201527320baba3432b73a34b1b0ba34b7b71032b93937b960611b604082015260600190565b634e487b7160e01b600052603260045260246000fd5b600181811c90821680611ea757607f821691505b602082108103611ec757634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115611f19576000816000526020600020601f850160051c81016020861015611ef65750805b601f850160051c820191505b81811015611f1557828155600101611f02565b5050505b505050565b815167ffffffffffffffff811115611f3857611f38611c79565b611f4c81611f468454611e93565b84611ecd565b602080601f831160018114611f815760008415611f695750858301515b600019600386901b1c1916600185901b178555611f15565b600085815260208120601f198616915b82811015611fb057888601518255948401946001909101908401611f91565b5085821015611fce5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b81810381811115611fff57634e487b7160e01b600052601160045260246000fd5b92915050565b67ffffffffffffffff83111561201d5761201d611c79565b6120318361202b8354611e93565b83611ecd565b6000601f841160018114612065576000851561204d5750838201355b600019600387901b1c1916600186901b1783556120bf565b600083815260209020601f19861690835b828110156120965786850135825560209485019460019092019101612076565b50868210156120b35760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b6040815260006120d960408301856119f3565b905060018060a01b0383166020830152939250505056fea264697066735822122033aa69f513a25c32dba22b143f5dcafa8ffa4370e2b1fcd791a1b5f779d8afa864736f6c63430008190033",
}

// XZ21ABI is the input ABI used to generate the binding from.
// Deprecated: Use XZ21MetaData.ABI instead.
var XZ21ABI = XZ21MetaData.ABI

// XZ21Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use XZ21MetaData.Bin instead.
var XZ21Bin = XZ21MetaData.Bin

// DeployXZ21 deploys a new Ethereum contract, binding an instance of XZ21 to it.
func DeployXZ21(auth *bind.TransactOpts, backend bind.ContractBackend, _addrSP common.Address) (common.Address, *types.Transaction, *XZ21, error) {
	parsed, err := XZ21MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(XZ21Bin), backend, _addrSP)
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

// GetAuditingLogs is a free data retrieval call binding the contract method 0x7f838211.
//
// Solidity: function GetAuditingLogs(bytes32 _hash) view returns(((bytes,bytes,uint8),bool,uint256)[])
func (_XZ21 *XZ21Caller) GetAuditingLogs(opts *bind.CallOpts, _hash [32]byte) ([]XZ21AuditingLog, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "GetAuditingLogs", _hash)

	if err != nil {
		return *new([]XZ21AuditingLog), err
	}

	out0 := *abi.ConvertType(out[0], new([]XZ21AuditingLog)).(*[]XZ21AuditingLog)

	return out0, err

}

// GetAuditingLogs is a free data retrieval call binding the contract method 0x7f838211.
//
// Solidity: function GetAuditingLogs(bytes32 _hash) view returns(((bytes,bytes,uint8),bool,uint256)[])
func (_XZ21 *XZ21Session) GetAuditingLogs(_hash [32]byte) ([]XZ21AuditingLog, error) {
	return _XZ21.Contract.GetAuditingLogs(&_XZ21.CallOpts, _hash)
}

// GetAuditingLogs is a free data retrieval call binding the contract method 0x7f838211.
//
// Solidity: function GetAuditingLogs(bytes32 _hash) view returns(((bytes,bytes,uint8),bool,uint256)[])
func (_XZ21 *XZ21CallerSession) GetAuditingLogs(_hash [32]byte) ([]XZ21AuditingLog, error) {
	return _XZ21.Contract.GetAuditingLogs(&_XZ21.CallOpts, _hash)
}

// GetAuditingReq is a free data retrieval call binding the contract method 0xc4e197ec.
//
// Solidity: function GetAuditingReq(bytes32 _hash) view returns((bytes,bytes,uint8))
func (_XZ21 *XZ21Caller) GetAuditingReq(opts *bind.CallOpts, _hash [32]byte) (XZ21AuditingReq, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "GetAuditingReq", _hash)

	if err != nil {
		return *new(XZ21AuditingReq), err
	}

	out0 := *abi.ConvertType(out[0], new(XZ21AuditingReq)).(*XZ21AuditingReq)

	return out0, err

}

// GetAuditingReq is a free data retrieval call binding the contract method 0xc4e197ec.
//
// Solidity: function GetAuditingReq(bytes32 _hash) view returns((bytes,bytes,uint8))
func (_XZ21 *XZ21Session) GetAuditingReq(_hash [32]byte) (XZ21AuditingReq, error) {
	return _XZ21.Contract.GetAuditingReq(&_XZ21.CallOpts, _hash)
}

// GetAuditingReq is a free data retrieval call binding the contract method 0xc4e197ec.
//
// Solidity: function GetAuditingReq(bytes32 _hash) view returns((bytes,bytes,uint8))
func (_XZ21 *XZ21CallerSession) GetAuditingReq(_hash [32]byte) (XZ21AuditingReq, error) {
	return _XZ21.Contract.GetAuditingReq(&_XZ21.CallOpts, _hash)
}

// GetAuditorAddrList is a free data retrieval call binding the contract method 0x5a955455.
//
// Solidity: function GetAuditorAddrList() view returns(address[])
func (_XZ21 *XZ21Caller) GetAuditorAddrList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "GetAuditorAddrList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAuditorAddrList is a free data retrieval call binding the contract method 0x5a955455.
//
// Solidity: function GetAuditorAddrList() view returns(address[])
func (_XZ21 *XZ21Session) GetAuditorAddrList() ([]common.Address, error) {
	return _XZ21.Contract.GetAuditorAddrList(&_XZ21.CallOpts)
}

// GetAuditorAddrList is a free data retrieval call binding the contract method 0x5a955455.
//
// Solidity: function GetAuditorAddrList() view returns(address[])
func (_XZ21 *XZ21CallerSession) GetAuditorAddrList() ([]common.Address, error) {
	return _XZ21.Contract.GetAuditorAddrList(&_XZ21.CallOpts)
}

// GetFileList is a free data retrieval call binding the contract method 0x9a2b5e63.
//
// Solidity: function GetFileList(address _owner) view returns(bytes32[])
func (_XZ21 *XZ21Caller) GetFileList(opts *bind.CallOpts, _owner common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "GetFileList", _owner)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetFileList is a free data retrieval call binding the contract method 0x9a2b5e63.
//
// Solidity: function GetFileList(address _owner) view returns(bytes32[])
func (_XZ21 *XZ21Session) GetFileList(_owner common.Address) ([][32]byte, error) {
	return _XZ21.Contract.GetFileList(&_XZ21.CallOpts, _owner)
}

// GetFileList is a free data retrieval call binding the contract method 0x9a2b5e63.
//
// Solidity: function GetFileList(address _owner) view returns(bytes32[])
func (_XZ21 *XZ21CallerSession) GetFileList(_owner common.Address) ([][32]byte, error) {
	return _XZ21.Contract.GetFileList(&_XZ21.CallOpts, _owner)
}

// GetParam is a free data retrieval call binding the contract method 0x8a04d5e5.
//
// Solidity: function GetParam() view returns((string,bytes,bytes))
func (_XZ21 *XZ21Caller) GetParam(opts *bind.CallOpts) (XZ21Param, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "GetParam")

	if err != nil {
		return *new(XZ21Param), err
	}

	out0 := *abi.ConvertType(out[0], new(XZ21Param)).(*XZ21Param)

	return out0, err

}

// GetParam is a free data retrieval call binding the contract method 0x8a04d5e5.
//
// Solidity: function GetParam() view returns((string,bytes,bytes))
func (_XZ21 *XZ21Session) GetParam() (XZ21Param, error) {
	return _XZ21.Contract.GetParam(&_XZ21.CallOpts)
}

// GetParam is a free data retrieval call binding the contract method 0x8a04d5e5.
//
// Solidity: function GetParam() view returns((string,bytes,bytes))
func (_XZ21 *XZ21CallerSession) GetParam() (XZ21Param, error) {
	return _XZ21.Contract.GetParam(&_XZ21.CallOpts)
}

// GetUserAccount is a free data retrieval call binding the contract method 0xb55957a0.
//
// Solidity: function GetUserAccount(address _addr) view returns((bytes,bytes32[]))
func (_XZ21 *XZ21Caller) GetUserAccount(opts *bind.CallOpts, _addr common.Address) (XZ21Account, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "GetUserAccount", _addr)

	if err != nil {
		return *new(XZ21Account), err
	}

	out0 := *abi.ConvertType(out[0], new(XZ21Account)).(*XZ21Account)

	return out0, err

}

// GetUserAccount is a free data retrieval call binding the contract method 0xb55957a0.
//
// Solidity: function GetUserAccount(address _addr) view returns((bytes,bytes32[]))
func (_XZ21 *XZ21Session) GetUserAccount(_addr common.Address) (XZ21Account, error) {
	return _XZ21.Contract.GetUserAccount(&_XZ21.CallOpts, _addr)
}

// GetUserAccount is a free data retrieval call binding the contract method 0xb55957a0.
//
// Solidity: function GetUserAccount(address _addr) view returns((bytes,bytes32[]))
func (_XZ21 *XZ21CallerSession) GetUserAccount(_addr common.Address) (XZ21Account, error) {
	return _XZ21.Contract.GetUserAccount(&_XZ21.CallOpts, _addr)
}

// SearchFile is a free data retrieval call binding the contract method 0x537322e3.
//
// Solidity: function SearchFile(bytes32 _hash) view returns((uint32,address))
func (_XZ21 *XZ21Caller) SearchFile(opts *bind.CallOpts, _hash [32]byte) (XZ21FileProperty, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "SearchFile", _hash)

	if err != nil {
		return *new(XZ21FileProperty), err
	}

	out0 := *abi.ConvertType(out[0], new(XZ21FileProperty)).(*XZ21FileProperty)

	return out0, err

}

// SearchFile is a free data retrieval call binding the contract method 0x537322e3.
//
// Solidity: function SearchFile(bytes32 _hash) view returns((uint32,address))
func (_XZ21 *XZ21Session) SearchFile(_hash [32]byte) (XZ21FileProperty, error) {
	return _XZ21.Contract.SearchFile(&_XZ21.CallOpts, _hash)
}

// SearchFile is a free data retrieval call binding the contract method 0x537322e3.
//
// Solidity: function SearchFile(bytes32 _hash) view returns((uint32,address))
func (_XZ21 *XZ21CallerSession) SearchFile(_hash [32]byte) (XZ21FileProperty, error) {
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

// AuditorAddrList is a free data retrieval call binding the contract method 0x2cc19d97.
//
// Solidity: function auditorAddrList(uint256 ) view returns(address)
func (_XZ21 *XZ21Caller) AuditorAddrList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "auditorAddrList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuditorAddrList is a free data retrieval call binding the contract method 0x2cc19d97.
//
// Solidity: function auditorAddrList(uint256 ) view returns(address)
func (_XZ21 *XZ21Session) AuditorAddrList(arg0 *big.Int) (common.Address, error) {
	return _XZ21.Contract.AuditorAddrList(&_XZ21.CallOpts, arg0)
}

// AuditorAddrList is a free data retrieval call binding the contract method 0x2cc19d97.
//
// Solidity: function auditorAddrList(uint256 ) view returns(address)
func (_XZ21 *XZ21CallerSession) AuditorAddrList(arg0 *big.Int) (common.Address, error) {
	return _XZ21.Contract.AuditorAddrList(&_XZ21.CallOpts, arg0)
}

// AppendOwner is a paid mutator transaction binding the contract method 0xf64b362a.
//
// Solidity: function AppendOwner(bytes32 _hash, address _owner) returns()
func (_XZ21 *XZ21Transactor) AppendOwner(opts *bind.TransactOpts, _hash [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _XZ21.contract.Transact(opts, "AppendOwner", _hash, _owner)
}

// AppendOwner is a paid mutator transaction binding the contract method 0xf64b362a.
//
// Solidity: function AppendOwner(bytes32 _hash, address _owner) returns()
func (_XZ21 *XZ21Session) AppendOwner(_hash [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _XZ21.Contract.AppendOwner(&_XZ21.TransactOpts, _hash, _owner)
}

// AppendOwner is a paid mutator transaction binding the contract method 0xf64b362a.
//
// Solidity: function AppendOwner(bytes32 _hash, address _owner) returns()
func (_XZ21 *XZ21TransactorSession) AppendOwner(_hash [32]byte, _owner common.Address) (*types.Transaction, error) {
	return _XZ21.Contract.AppendOwner(&_XZ21.TransactOpts, _hash, _owner)
}

// EnrollAccount is a paid mutator transaction binding the contract method 0x043c9077.
//
// Solidity: function EnrollAccount(int256 _type, address _addr, bytes _pubKey) returns(bool)
func (_XZ21 *XZ21Transactor) EnrollAccount(opts *bind.TransactOpts, _type *big.Int, _addr common.Address, _pubKey []byte) (*types.Transaction, error) {
	return _XZ21.contract.Transact(opts, "EnrollAccount", _type, _addr, _pubKey)
}

// EnrollAccount is a paid mutator transaction binding the contract method 0x043c9077.
//
// Solidity: function EnrollAccount(int256 _type, address _addr, bytes _pubKey) returns(bool)
func (_XZ21 *XZ21Session) EnrollAccount(_type *big.Int, _addr common.Address, _pubKey []byte) (*types.Transaction, error) {
	return _XZ21.Contract.EnrollAccount(&_XZ21.TransactOpts, _type, _addr, _pubKey)
}

// EnrollAccount is a paid mutator transaction binding the contract method 0x043c9077.
//
// Solidity: function EnrollAccount(int256 _type, address _addr, bytes _pubKey) returns(bool)
func (_XZ21 *XZ21TransactorSession) EnrollAccount(_type *big.Int, _addr common.Address, _pubKey []byte) (*types.Transaction, error) {
	return _XZ21.Contract.EnrollAccount(&_XZ21.TransactOpts, _type, _addr, _pubKey)
}

// RegisterFile is a paid mutator transaction binding the contract method 0x22dd2b18.
//
// Solidity: function RegisterFile(bytes32 _hash, uint32 _splitNum, address _owner) returns()
func (_XZ21 *XZ21Transactor) RegisterFile(opts *bind.TransactOpts, _hash [32]byte, _splitNum uint32, _owner common.Address) (*types.Transaction, error) {
	return _XZ21.contract.Transact(opts, "RegisterFile", _hash, _splitNum, _owner)
}

// RegisterFile is a paid mutator transaction binding the contract method 0x22dd2b18.
//
// Solidity: function RegisterFile(bytes32 _hash, uint32 _splitNum, address _owner) returns()
func (_XZ21 *XZ21Session) RegisterFile(_hash [32]byte, _splitNum uint32, _owner common.Address) (*types.Transaction, error) {
	return _XZ21.Contract.RegisterFile(&_XZ21.TransactOpts, _hash, _splitNum, _owner)
}

// RegisterFile is a paid mutator transaction binding the contract method 0x22dd2b18.
//
// Solidity: function RegisterFile(bytes32 _hash, uint32 _splitNum, address _owner) returns()
func (_XZ21 *XZ21TransactorSession) RegisterFile(_hash [32]byte, _splitNum uint32, _owner common.Address) (*types.Transaction, error) {
	return _XZ21.Contract.RegisterFile(&_XZ21.TransactOpts, _hash, _splitNum, _owner)
}

// RegisterParam is a paid mutator transaction binding the contract method 0xc1f16b16.
//
// Solidity: function RegisterParam(string _p, bytes _g, bytes _u) returns()
func (_XZ21 *XZ21Transactor) RegisterParam(opts *bind.TransactOpts, _p string, _g []byte, _u []byte) (*types.Transaction, error) {
	return _XZ21.contract.Transact(opts, "RegisterParam", _p, _g, _u)
}

// RegisterParam is a paid mutator transaction binding the contract method 0xc1f16b16.
//
// Solidity: function RegisterParam(string _p, bytes _g, bytes _u) returns()
func (_XZ21 *XZ21Session) RegisterParam(_p string, _g []byte, _u []byte) (*types.Transaction, error) {
	return _XZ21.Contract.RegisterParam(&_XZ21.TransactOpts, _p, _g, _u)
}

// RegisterParam is a paid mutator transaction binding the contract method 0xc1f16b16.
//
// Solidity: function RegisterParam(string _p, bytes _g, bytes _u) returns()
func (_XZ21 *XZ21TransactorSession) RegisterParam(_p string, _g []byte, _u []byte) (*types.Transaction, error) {
	return _XZ21.Contract.RegisterParam(&_XZ21.TransactOpts, _p, _g, _u)
}

// SetAuditingResult is a paid mutator transaction binding the contract method 0xb645fd63.
//
// Solidity: function SetAuditingResult(bytes32 _hash, bool _result) returns()
func (_XZ21 *XZ21Transactor) SetAuditingResult(opts *bind.TransactOpts, _hash [32]byte, _result bool) (*types.Transaction, error) {
	return _XZ21.contract.Transact(opts, "SetAuditingResult", _hash, _result)
}

// SetAuditingResult is a paid mutator transaction binding the contract method 0xb645fd63.
//
// Solidity: function SetAuditingResult(bytes32 _hash, bool _result) returns()
func (_XZ21 *XZ21Session) SetAuditingResult(_hash [32]byte, _result bool) (*types.Transaction, error) {
	return _XZ21.Contract.SetAuditingResult(&_XZ21.TransactOpts, _hash, _result)
}

// SetAuditingResult is a paid mutator transaction binding the contract method 0xb645fd63.
//
// Solidity: function SetAuditingResult(bytes32 _hash, bool _result) returns()
func (_XZ21 *XZ21TransactorSession) SetAuditingResult(_hash [32]byte, _result bool) (*types.Transaction, error) {
	return _XZ21.Contract.SetAuditingResult(&_XZ21.TransactOpts, _hash, _result)
}

// SetChal is a paid mutator transaction binding the contract method 0xdebffd9a.
//
// Solidity: function SetChal(bytes32 _hash, bytes _chal) returns()
func (_XZ21 *XZ21Transactor) SetChal(opts *bind.TransactOpts, _hash [32]byte, _chal []byte) (*types.Transaction, error) {
	return _XZ21.contract.Transact(opts, "SetChal", _hash, _chal)
}

// SetChal is a paid mutator transaction binding the contract method 0xdebffd9a.
//
// Solidity: function SetChal(bytes32 _hash, bytes _chal) returns()
func (_XZ21 *XZ21Session) SetChal(_hash [32]byte, _chal []byte) (*types.Transaction, error) {
	return _XZ21.Contract.SetChal(&_XZ21.TransactOpts, _hash, _chal)
}

// SetChal is a paid mutator transaction binding the contract method 0xdebffd9a.
//
// Solidity: function SetChal(bytes32 _hash, bytes _chal) returns()
func (_XZ21 *XZ21TransactorSession) SetChal(_hash [32]byte, _chal []byte) (*types.Transaction, error) {
	return _XZ21.Contract.SetChal(&_XZ21.TransactOpts, _hash, _chal)
}

// SetProof is a paid mutator transaction binding the contract method 0xd746cdd0.
//
// Solidity: function SetProof(bytes32 _hash, bytes _proof) returns()
func (_XZ21 *XZ21Transactor) SetProof(opts *bind.TransactOpts, _hash [32]byte, _proof []byte) (*types.Transaction, error) {
	return _XZ21.contract.Transact(opts, "SetProof", _hash, _proof)
}

// SetProof is a paid mutator transaction binding the contract method 0xd746cdd0.
//
// Solidity: function SetProof(bytes32 _hash, bytes _proof) returns()
func (_XZ21 *XZ21Session) SetProof(_hash [32]byte, _proof []byte) (*types.Transaction, error) {
	return _XZ21.Contract.SetProof(&_XZ21.TransactOpts, _hash, _proof)
}

// SetProof is a paid mutator transaction binding the contract method 0xd746cdd0.
//
// Solidity: function SetProof(bytes32 _hash, bytes _proof) returns()
func (_XZ21 *XZ21TransactorSession) SetProof(_hash [32]byte, _proof []byte) (*types.Transaction, error) {
	return _XZ21.Contract.SetProof(&_XZ21.TransactOpts, _hash, _proof)
}

// XZ21EventSetAuditingResultIterator is returned from FilterEventSetAuditingResult and is used to iterate over the raw logs and unpacked data for EventSetAuditingResult events raised by the XZ21 contract.
type XZ21EventSetAuditingResultIterator struct {
	Event *XZ21EventSetAuditingResult // Event containing the contract specifics and raw log

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
func (it *XZ21EventSetAuditingResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XZ21EventSetAuditingResult)
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
		it.Event = new(XZ21EventSetAuditingResult)
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
func (it *XZ21EventSetAuditingResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XZ21EventSetAuditingResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XZ21EventSetAuditingResult represents a EventSetAuditingResult event raised by the XZ21 contract.
type XZ21EventSetAuditingResult struct {
	Hash   [32]byte
	Result bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventSetAuditingResult is a free log retrieval operation binding the contract event 0x961a406eff4d590fdb9e54bdba584518be2e6a7f91b34d5714330f83d571b031.
//
// Solidity: event EventSetAuditingResult(bytes32 _hash, bool _result)
func (_XZ21 *XZ21Filterer) FilterEventSetAuditingResult(opts *bind.FilterOpts) (*XZ21EventSetAuditingResultIterator, error) {

	logs, sub, err := _XZ21.contract.FilterLogs(opts, "EventSetAuditingResult")
	if err != nil {
		return nil, err
	}
	return &XZ21EventSetAuditingResultIterator{contract: _XZ21.contract, event: "EventSetAuditingResult", logs: logs, sub: sub}, nil
}

// WatchEventSetAuditingResult is a free log subscription operation binding the contract event 0x961a406eff4d590fdb9e54bdba584518be2e6a7f91b34d5714330f83d571b031.
//
// Solidity: event EventSetAuditingResult(bytes32 _hash, bool _result)
func (_XZ21 *XZ21Filterer) WatchEventSetAuditingResult(opts *bind.WatchOpts, sink chan<- *XZ21EventSetAuditingResult) (event.Subscription, error) {

	logs, sub, err := _XZ21.contract.WatchLogs(opts, "EventSetAuditingResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XZ21EventSetAuditingResult)
				if err := _XZ21.contract.UnpackLog(event, "EventSetAuditingResult", log); err != nil {
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

// ParseEventSetAuditingResult is a log parse operation binding the contract event 0x961a406eff4d590fdb9e54bdba584518be2e6a7f91b34d5714330f83d571b031.
//
// Solidity: event EventSetAuditingResult(bytes32 _hash, bool _result)
func (_XZ21 *XZ21Filterer) ParseEventSetAuditingResult(log types.Log) (*XZ21EventSetAuditingResult, error) {
	event := new(XZ21EventSetAuditingResult)
	if err := _XZ21.contract.UnpackLog(event, "EventSetAuditingResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
