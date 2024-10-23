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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_addrSP\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"AppendOwner\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"EnrollAccount\",\"inputs\":[{\"name\":\"_type\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"GetAccount\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.Account\",\"components\":[{\"name\":\"pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fileList\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetAddrListTPA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetAuditingLogs\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structXZ21.AuditingLog[]\",\"components\":[{\"name\":\"req\",\"type\":\"tuple\",\"internalType\":\"structXZ21.AuditingReq\",\"components\":[{\"name\":\"chal\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"stage\",\"type\":\"uint8\",\"internalType\":\"enumXZ21.Stages\"}]},{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"date\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetAuditingReqList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structXZ21.AuditingReq[]\",\"components\":[{\"name\":\"chal\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"stage\",\"type\":\"uint8\",\"internalType\":\"enumXZ21.Stages\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetFileList\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetParam\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.Param\",\"components\":[{\"name\":\"P\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"U\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"G\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RegisterFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_splitNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"RegisterParam\",\"inputs\":[{\"name\":\"_p\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_g\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_u\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SearchFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.FileProperty\",\"components\":[{\"name\":\"splitNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SetAuditingResult\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SetChal\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_chal\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SetProof\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addrListTPA\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrSM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrSP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EventSetAuditingResult\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
	Bin: "0x60c0604052348015600f57600080fd5b5060405161206b38038061206b833981016040819052602c916040565b336080526001600160a01b031660a052606e565b600060208284031215605157600080fd5b81516001600160a01b0381168114606757600080fd5b9392505050565b60805160a051611fd161009a600039600061022f0152600081816101db01526103520152611fd16000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80639a2b5e63116100a2578063d746cdd011610071578063d746cdd0146102ec578063debffd9a146102ff578063eaa1e49714610312578063f64b362a14610325578063f65de5991461033857600080fd5b80639a2b5e6314610286578063b645fd63146102a6578063c1f16b16146102b9578063d3674148146102cc57600080fd5b80635dd420df116100de5780635dd420df146102155780636b884afb1461022a5780637f838211146102515780638a04d5e51461027157600080fd5b8063043c90771461011057806322dd2b1814610138578063537322e31461014d57806358f5823c146101d6575b600080fd5b61012361011e366004611735565b61034e565b60405190151581526020015b60405180910390f35b61014b61014636600461178f565b61050c565b005b6101ac61015b3660046117d6565b60408051808201909152600080825260208201525060009081526005602090815260409182902082518084019093525463ffffffff8116835264010000000090046001600160a01b03169082015290565b60408051825163ffffffff1681526020928301516001600160a01b0316928101929092520161012f565b6101fd7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161012f565b61021d6105be565b60405161012f91906117ef565b6101fd7f000000000000000000000000000000000000000000000000000000000000000081565b61026461025f3660046117d6565b610620565b60405161012f91906118fd565b610279610812565b60405161012f9190611982565b6102996102943660046119e3565b6109fe565b60405161012f9190611a41565b61014b6102b4366004611a54565b610aed565b61014b6102c7366004611b35565b610f04565b6102df6102da3660046119e3565b610f30565b60405161012f9190611bd4565b61014b6102fa366004611c0d565b611056565b61014b61030d366004611c0d565b611103565b6101fd6103203660046117d6565b6111d2565b61014b610333366004611c59565b6111fc565b610340611279565b60405161012f929190611c85565b60007f0000000000000000000000000000000000000000000000000000000000000000336001600160a01b0382161461038657600080fd5b8560000361041d576103cd6040518060400160405280601f81526020017f456e726f6c6c20545041206163636f756e742028416464726573733a2573290081525086611529565b600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319166001600160a01b0387161790556104fe565b856001036104f5576104646040518060400160405280601e81526020017f456e726f6c6c205355206163636f756e742028416464726573733a257329000081525086611529565b6040805160606020601f87018190040282018101835291810185815290918291908790879081908501838280828437600092018290525093855250506040805183815260208082018352948501526001600160a01b038a16835260049093525020815181906104d39082611d82565b5060208281015180516104ec9260018501920190611670565b509050506104fe565b60009150610503565b600191505b50949350505050565b60008263ffffffff161161055b5760405162461bcd60e51b8152602060048201526011602482015270696e76616c69642073706c6974206e756d60781b60448201526064015b60405180910390fd5b6000838152600560209081526040808320805463ffffffff969096166001600160c01b0319909616959095176401000000006001600160a01b03959095169485021790945591815260048252918220600190810180549182018155835291200155565b6060600380548060200260200160405190810160405280929190818152602001828054801561061657602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116105f8575b5050505050905090565b606060076000838152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b82821015610807576000848152602090206040805160c081019091526005840290910180548290606082019083908290829061069090611cf7565b80601f01602080910402602001604051908101604052809291908181526020018280546106bc90611cf7565b80156107095780601f106106de57610100808354040283529160200191610709565b820191906000526020600020905b8154815290600101906020018083116106ec57829003601f168201915b5050505050815260200160018201805461072290611cf7565b80601f016020809104026020016040519081016040528092919081815260200182805461074e90611cf7565b801561079b5780601f106107705761010080835404028352916020019161079b565b820191906000526020600020905b81548152906001019060200180831161077e57829003601f168201915b505050918352505060028281015460209092019160ff16908111156107c2576107c2611882565b60028111156107d3576107d3611882565b9052508152600382015460ff1615156020808301919091526004909201546040909101529082526001929092019101610655565b505050509050919050565b61083660405180606001604052806060815260200160608152602001606081525090565b600060405180606001604052908160008201805461085390611cf7565b80601f016020809104026020016040519081016040528092919081815260200182805461087f90611cf7565b80156108cc5780601f106108a1576101008083540402835291602001916108cc565b820191906000526020600020905b8154815290600101906020018083116108af57829003601f168201915b505050505081526020016001820180546108e590611cf7565b80601f016020809104026020016040519081016040528092919081815260200182805461091190611cf7565b801561095e5780601f106109335761010080835404028352916020019161095e565b820191906000526020600020905b81548152906001019060200180831161094157829003601f168201915b5050505050815260200160028201805461097790611cf7565b80601f01602080910402602001604051908101604052809291908181526020018280546109a390611cf7565b80156109f05780601f106109c5576101008083540402835291602001916109f0565b820191906000526020600020905b8154815290600101906020018083116109d357829003601f168201915b505050505081525050905090565b6001600160a01b0381166000908152600460205260408120600101546060918167ffffffffffffffff811115610a3657610a36611a89565b604051908082528060200260200182016040528015610a5f578160200160208202803683370190505b50905060005b6001600160a01b038516600090815260046020526040902060010154811015610ae5576001600160a01b0385166000908152600460205260409020600101805482908110610ab557610ab5611e42565b9060005260206000200154828281518110610ad257610ad2611e42565b6020908102919091010152600101610a65565b509392505050565b600082815260066020526040902060029081015460ff1681811115610b1457610b14611882565b14610b555760405162461bcd60e51b8152602060048201526011602482015270139bdd0815d85a5d1a5b99d4995cdd5b1d607a1b6044820152606401610552565b60008281526007602052604090205415610bfc57600082815260076020526040812054610b8490600190611e6e565b60008481526007602052604090208054919250429183908110610ba957610ba9611e42565b90600052602060002090600502016004015410610bfa5760405162461bcd60e51b815260206004820152600f60248201526e3a34b6b2b9ba30b6b81032b93937b960891b6044820152606401610552565b505b604080516000848152600660205282812060c0830190935282549092829160608301919082908290610c2d90611cf7565b80601f0160208091040260200160405190810160405280929190818152602001828054610c5990611cf7565b8015610ca65780601f10610c7b57610100808354040283529160200191610ca6565b820191906000526020600020905b815481529060010190602001808311610c8957829003601f168201915b50505050508152602001600182018054610cbf90611cf7565b80601f0160208091040260200160405190810160405280929190818152602001828054610ceb90611cf7565b8015610d385780601f10610d0d57610100808354040283529160200191610d38565b820191906000526020600020905b815481529060010190602001808311610d1b57829003601f168201915b505050918352505060028281015460209092019160ff1690811115610d5f57610d5f611882565b6002811115610d7057610d70611882565b9052508152831515602080830191909152426040928301526000868152600782529182208054600181018255908352912082518051939450849360059093029091019182908190610dc19082611d82565b5060208201516001820190610dd69082611d82565b5060408201518160020160006101000a81548160ff02191690836002811115610e0157610e01611882565b0217905550505060208281015160038301805460ff191691151591909117905560409283015160049092019190915581516080810183526000606082018181528252835180840185528181529282019290925291820152600084815260066020526040902081518190610e749082611d82565b5060208201516001820190610e899082611d82565b5060408201518160020160006101000a81548160ff02191690836002811115610eb457610eb4611882565b0217905550905050610ec583611572565b6040805184815283151560208201527f961a406eff4d590fdb9e54bdba584518be2e6a7f91b34d5714330f83d571b031910160405180910390a1505050565b6000610f108482611d82565b506001610f1d8282611d82565b506002610f2a8382611d82565b50505050565b60408051808201909152606080825260208201526001600160a01b038216600090815260046020526040908190208151808301909252805482908290610f7590611cf7565b80601f0160208091040260200160405190810160405280929190818152602001828054610fa190611cf7565b8015610fee5780601f10610fc357610100808354040283529160200191610fee565b820191906000526020600020905b815481529060010190602001808311610fd157829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561104657602002820191906000526020600020905b815481526020019060010190808311611032575b5050505050815250509050919050565b6001600084815260066020526040902060029081015460ff169081111561107f5761107f611882565b146110bf5760405162461bcd60e51b815260206004820152601060248201526f2737ba102bb0b4ba34b733a83937b7b360811b6044820152606401610552565b60008381526006602052604090206001016110db828483611e87565b5060008381526006602052604090206002908101805460ff19166001835b0217905550505050565b600083815260066020526040812060029081015460ff169081111561112a5761112a611882565b146111695760405162461bcd60e51b815260206004820152600f60248201526e139bdd0815d85a5d1a5b99d0da185b608a1b6044820152606401610552565b6000838152600660205260409020611182828483611e87565b506008805460018181019092557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3018490556000848152600660205260409020600201805460ff191682806110f9565b600381815481106111e257600080fd5b6000918252602090912001546001600160a01b0316905081565b60008281526005602052604090205463ffffffff1661124c5760405162461bcd60e51b815260206004820152600c60248201526b696e76616c69642066696c6560a01b6044820152606401610552565b6001600160a01b031660009081526004602090815260408220600190810180549182018155835291200155565b600854606090819060008167ffffffffffffffff81111561129c5761129c611a89565b6040519080825280602002602001820160405280156112c5578160200160208202803683370190505b50905060008267ffffffffffffffff8111156112e3576112e3611a89565b60405190808252806020026020018201604052801561133457816020015b611321604080516060808201835280825260208201529081016000905290565b8152602001906001900390816113015790505b50905060005b8381101561151e5760006008828154811061135757611357611e42565b906000526020600020015490508084838151811061137757611377611e42565b602002602001018181525050600660008281526020019081526020016000206040518060600160405290816000820180546113b190611cf7565b80601f01602080910402602001604051908101604052809291908181526020018280546113dd90611cf7565b801561142a5780601f106113ff5761010080835404028352916020019161142a565b820191906000526020600020905b81548152906001019060200180831161140d57829003601f168201915b5050505050815260200160018201805461144390611cf7565b80601f016020809104026020016040519081016040528092919081815260200182805461146f90611cf7565b80156114bc5780601f10611491576101008083540402835291602001916114bc565b820191906000526020600020905b81548152906001019060200180831161149f57829003601f168201915b505050918352505060028281015460209092019160ff16908111156114e3576114e3611882565b60028111156114f4576114f4611882565b8152505083838151811061150a5761150a611e42565b60209081029190910101525060010161133a565b509094909350915050565b61156e828260405160240161153f929190611f48565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b17905261164f565b5050565b60008060005b6008548110156115ba576008818154811061159557611595611e42565b906000526020600020015484036115b257600192508091506115ba565b600101611578565b50805b6008546115cc90600190611e6e565b8110156116225760086115e0826001611f72565b815481106115f0576115f0611e42565b90600052602060002001546008828154811061160e5761160e611e42565b6000918252602090912001556001016115bd565b50600880548061163457611634611f85565b60019003818190600052602060002001600090559055505050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b8280548282559060005260206000209081019282156116ab579160200282015b828111156116ab578251825591602001919060010190611690565b506116b79291506116bb565b5090565b5b808211156116b757600081556001016116bc565b80356001600160a01b03811681146116e757600080fd5b919050565b60008083601f8401126116fe57600080fd5b50813567ffffffffffffffff81111561171657600080fd5b60208301915083602082850101111561172e57600080fd5b9250929050565b6000806000806060858703121561174b57600080fd5b8435935061175b602086016116d0565b9250604085013567ffffffffffffffff81111561177757600080fd5b611783878288016116ec565b95989497509550505050565b6000806000606084860312156117a457600080fd5b83359250602084013563ffffffff811681146117bf57600080fd5b91506117cd604085016116d0565b90509250925092565b6000602082840312156117e857600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b818110156118305783516001600160a01b03168352928401929184019160010161180b565b50909695505050505050565b6000815180845260005b8181101561186257602081850181015186830182015201611846565b506000602082860101526020601f19601f83011685010191505092915050565b634e487b7160e01b600052602160045260246000fd5b60008151606084526118ad606085018261183c565b9050602083015184820360208601526118c6828261183c565b9150506040830151600381106118ec57634e487b7160e01b600052602160045260246000fd5b806040860152508091505092915050565b600060208083018184528085518083526040925060408601915060408160051b87010184880160005b8381101561197457603f1989840301855281516060815181865261194c82870182611898565b838b01511515878c015292890151958901959095525094870194925090860190600101611926565b509098975050505050505050565b60208152600082516060602084015261199e608084018261183c565b90506020840151601f19808584030160408601526119bc838361183c565b92506040860151915080858403016060860152506119da828261183c565b95945050505050565b6000602082840312156119f557600080fd5b6119fe826116d0565b9392505050565b60008151808452602080850194506020840160005b83811015611a3657815187529582019590820190600101611a1a565b509495945050505050565b6020815260006119fe6020830184611a05565b60008060408385031215611a6757600080fd5b8235915060208301358015158114611a7e57600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff80841115611aba57611aba611a89565b604051601f8501601f19908116603f01168101908282118183101715611ae257611ae2611a89565b81604052809350858152868686011115611afb57600080fd5b858560208301376000602087830101525050509392505050565b600082601f830112611b2657600080fd5b6119fe83833560208501611a9f565b600080600060608486031215611b4a57600080fd5b833567ffffffffffffffff80821115611b6257600080fd5b818601915086601f830112611b7657600080fd5b611b8587833560208501611a9f565b94506020860135915080821115611b9b57600080fd5b611ba787838801611b15565b93506040860135915080821115611bbd57600080fd5b50611bca86828701611b15565b9150509250925092565b602081526000825160406020840152611bf0606084018261183c565b90506020840151601f198483030160408501526119da8282611a05565b600080600060408486031215611c2257600080fd5b83359250602084013567ffffffffffffffff811115611c4057600080fd5b611c4c868287016116ec565b9497909650939450505050565b60008060408385031215611c6c57600080fd5b82359150611c7c602084016116d0565b90509250929050565b604081526000611c986040830185611a05565b6020838203818501528185518084528284019150828160051b85010183880160005b83811015611ce857601f19878403018552611cd6838351611898565b94860194925090850190600101611cba565b50909998505050505050505050565b600181811c90821680611d0b57607f821691505b602082108103611d2b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115611d7d576000816000526020600020601f850160051c81016020861015611d5a5750805b601f850160051c820191505b81811015611d7957828155600101611d66565b5050505b505050565b815167ffffffffffffffff811115611d9c57611d9c611a89565b611db081611daa8454611cf7565b84611d31565b602080601f831160018114611de55760008415611dcd5750858301515b600019600386901b1c1916600185901b178555611d79565b600085815260208120601f198616915b82811015611e1457888601518255948401946001909101908401611df5565b5085821015611e325787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b81810381811115611e8157611e81611e58565b92915050565b67ffffffffffffffff831115611e9f57611e9f611a89565b611eb383611ead8354611cf7565b83611d31565b6000601f841160018114611ee75760008515611ecf5750838201355b600019600387901b1c1916600186901b178355611f41565b600083815260209020601f19861690835b82811015611f185786850135825560209485019460019092019101611ef8565b5086821015611f355760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b604081526000611f5b604083018561183c565b905060018060a01b03831660208301529392505050565b80820180821115611e8157611e81611e58565b634e487b7160e01b600052603160045260246000fdfea2646970667358221220f6c7bed70098969d6ac410b722e384cd72170ddef330b90d9ec9bde3869677bf64736f6c63430008190033",
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

// GetAccount is a free data retrieval call binding the contract method 0xd3674148.
//
// Solidity: function GetAccount(address _addr) view returns((bytes,bytes32[]))
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
// Solidity: function GetAccount(address _addr) view returns((bytes,bytes32[]))
func (_XZ21 *XZ21Session) GetAccount(_addr common.Address) (XZ21Account, error) {
	return _XZ21.Contract.GetAccount(&_XZ21.CallOpts, _addr)
}

// GetAccount is a free data retrieval call binding the contract method 0xd3674148.
//
// Solidity: function GetAccount(address _addr) view returns((bytes,bytes32[]))
func (_XZ21 *XZ21CallerSession) GetAccount(_addr common.Address) (XZ21Account, error) {
	return _XZ21.Contract.GetAccount(&_XZ21.CallOpts, _addr)
}

// GetAddrListTPA is a free data retrieval call binding the contract method 0x5dd420df.
//
// Solidity: function GetAddrListTPA() view returns(address[])
func (_XZ21 *XZ21Caller) GetAddrListTPA(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "GetAddrListTPA")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAddrListTPA is a free data retrieval call binding the contract method 0x5dd420df.
//
// Solidity: function GetAddrListTPA() view returns(address[])
func (_XZ21 *XZ21Session) GetAddrListTPA() ([]common.Address, error) {
	return _XZ21.Contract.GetAddrListTPA(&_XZ21.CallOpts)
}

// GetAddrListTPA is a free data retrieval call binding the contract method 0x5dd420df.
//
// Solidity: function GetAddrListTPA() view returns(address[])
func (_XZ21 *XZ21CallerSession) GetAddrListTPA() ([]common.Address, error) {
	return _XZ21.Contract.GetAddrListTPA(&_XZ21.CallOpts)
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

// GetAuditingReqList is a free data retrieval call binding the contract method 0xf65de599.
//
// Solidity: function GetAuditingReqList() view returns(bytes32[], (bytes,bytes,uint8)[])
func (_XZ21 *XZ21Caller) GetAuditingReqList(opts *bind.CallOpts) ([][32]byte, []XZ21AuditingReq, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "GetAuditingReqList")

	if err != nil {
		return *new([][32]byte), *new([]XZ21AuditingReq), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	out1 := *abi.ConvertType(out[1], new([]XZ21AuditingReq)).(*[]XZ21AuditingReq)

	return out0, out1, err

}

// GetAuditingReqList is a free data retrieval call binding the contract method 0xf65de599.
//
// Solidity: function GetAuditingReqList() view returns(bytes32[], (bytes,bytes,uint8)[])
func (_XZ21 *XZ21Session) GetAuditingReqList() ([][32]byte, []XZ21AuditingReq, error) {
	return _XZ21.Contract.GetAuditingReqList(&_XZ21.CallOpts)
}

// GetAuditingReqList is a free data retrieval call binding the contract method 0xf65de599.
//
// Solidity: function GetAuditingReqList() view returns(bytes32[], (bytes,bytes,uint8)[])
func (_XZ21 *XZ21CallerSession) GetAuditingReqList() ([][32]byte, []XZ21AuditingReq, error) {
	return _XZ21.Contract.GetAuditingReqList(&_XZ21.CallOpts)
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

// AddrListTPA is a free data retrieval call binding the contract method 0xeaa1e497.
//
// Solidity: function addrListTPA(uint256 ) view returns(address)
func (_XZ21 *XZ21Caller) AddrListTPA(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "addrListTPA", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddrListTPA is a free data retrieval call binding the contract method 0xeaa1e497.
//
// Solidity: function addrListTPA(uint256 ) view returns(address)
func (_XZ21 *XZ21Session) AddrListTPA(arg0 *big.Int) (common.Address, error) {
	return _XZ21.Contract.AddrListTPA(&_XZ21.CallOpts, arg0)
}

// AddrListTPA is a free data retrieval call binding the contract method 0xeaa1e497.
//
// Solidity: function addrListTPA(uint256 ) view returns(address)
func (_XZ21 *XZ21CallerSession) AddrListTPA(arg0 *big.Int) (common.Address, error) {
	return _XZ21.Contract.AddrListTPA(&_XZ21.CallOpts, arg0)
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
