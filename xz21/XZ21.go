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
	Chal   []byte
	Proof  []byte
	Result bool
}

// XZ21AuditingReq is an auto generated low-level Go binding around an user-defined struct.
type XZ21AuditingReq struct {
	Chal  []byte
	Proof []byte
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_addrSP\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_addrTPA\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"AppendOwner\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"EnrollAccount\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"GetAccount\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.Account\",\"components\":[{\"name\":\"pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fileList\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetAuditingLogs\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structXZ21.AuditingLog[]\",\"components\":[{\"name\":\"chal\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetAuditingReqList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structXZ21.AuditingReq[]\",\"components\":[{\"name\":\"chal\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetFileList\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetParam\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.Param\",\"components\":[{\"name\":\"P\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"U\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"G\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RegisterFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_splitNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"RegisterParam\",\"inputs\":[{\"name\":\"_p\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_g\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_u\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SearchFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.FileProperty\",\"components\":[{\"name\":\"splitNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SetAuditingResult\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SetChal\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_chal\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SetProof\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addrSM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrSP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrTPA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EventSetAuditingResult\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600f57600080fd5b50604051611bc7380380611bc7833981016040819052602c916082565b600380546001600160a01b03199081163317909155600480546001600160a01b039485169083161790556005805492909316911617905560b0565b80516001600160a01b0381168114607d57600080fd5b919050565b60008060408385031215609457600080fd5b609b836067565b915060a7602084016067565b90509250929050565b611b08806100bf6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80639a2b5e6311610097578063d746cdd011610066578063d746cdd0146102a7578063debffd9a146102ba578063f64b362a146102cd578063f65de599146102e057600080fd5b80639a2b5e6314610241578063b645fd6314610261578063c1f16b1614610274578063d36741481461028757600080fd5b806379ef33df116100d357806379ef33df146101e65780637f838211146101f95780638a04d5e51461021957806392f4dfa21461022e57600080fd5b806322dd2b1814610105578063537322e31461011a57806358f5823c146101a85780636b884afb146101d3575b600080fd5b6101186101133660046112d7565b6102f6565b005b61017961012836600461131e565b60408051808201909152600080825260208201525060009081526007602090815260409182902082518084019093525463ffffffff8116835264010000000090046001600160a01b03169082015290565b60408051825163ffffffff1681526020928301516001600160a01b031692810192909252015b60405180910390f35b6003546101bb906001600160a01b031681565b6040516001600160a01b03909116815260200161019f565b6004546101bb906001600160a01b031681565b6005546101bb906001600160a01b031681565b61020c61020736600461131e565b6103a8565b60405161019f919061137d565b610221610550565b60405161019f9190611412565b61011861023c3660046114bc565b61073c565b61025461024f36600461150f565b610824565b60405161019f919061156d565b61011861026f366004611580565b610913565b610118610282366004611661565b610bb8565b61029a61029536600461150f565b610be4565b60405161019f9190611700565b6101186102b5366004611739565b610d0a565b6101186102c8366004611739565b610d73565b6101186102db36600461176c565b610e26565b6102e8610ea3565b60405161019f929190611798565b60008263ffffffff16116103455760405162461bcd60e51b8152602060048201526011602482015270696e76616c69642073706c6974206e756d60781b60448201526064015b60405180910390fd5b6000838152600760209081526040808320805463ffffffff969096166001600160c01b0319909616959095176401000000006001600160a01b03959095169485021790945591815260068252918220600190810180549182018155835291200155565b606060096000838152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561054557838290600052602060002090600302016040518060600160405290816000820180546104109061182e565b80601f016020809104026020016040519081016040528092919081815260200182805461043c9061182e565b80156104895780601f1061045e57610100808354040283529160200191610489565b820191906000526020600020905b81548152906001019060200180831161046c57829003601f168201915b505050505081526020016001820180546104a29061182e565b80601f01602080910402602001604051908101604052809291908181526020018280546104ce9061182e565b801561051b5780601f106104f05761010080835404028352916020019161051b565b820191906000526020600020905b8154815290600101906020018083116104fe57829003601f168201915b50505091835250506002919091015460ff16151560209182015290825260019290920191016103dd565b505050509050919050565b61057460405180606001604052806060815260200160608152602001606081525090565b60006040518060600160405290816000820180546105919061182e565b80601f01602080910402602001604051908101604052809291908181526020018280546105bd9061182e565b801561060a5780601f106105df5761010080835404028352916020019161060a565b820191906000526020600020905b8154815290600101906020018083116105ed57829003601f168201915b505050505081526020016001820180546106239061182e565b80601f016020809104026020016040519081016040528092919081815260200182805461064f9061182e565b801561069c5780601f106106715761010080835404028352916020019161069c565b820191906000526020600020905b81548152906001019060200180831161067f57829003601f168201915b505050505081526020016002820180546106b59061182e565b80601f01602080910402602001604051908101604052809291908181526020018280546106e19061182e565b801561072e5780601f106107035761010080835404028352916020019161072e565b820191906000526020600020905b81548152906001019060200180831161071157829003601f168201915b505050505081525050905090565b6003546001600160a01b031633811461075457600080fd5b6107936040518060400160405280601e81526020017f456e726f6c6c205355206163636f756e742028416464726573733a257329000081525085611114565b6040805160606020601f86018190040282018101835291810184815290918291908690869081908501838280828437600092018290525093855250506040805183815260208082018352948501526001600160a01b0389168352600690935250208151819061080290826118b9565b50602082810151805161081b926001850192019061125b565b50505050505050565b6001600160a01b0381166000908152600660205260408120600101546060918167ffffffffffffffff81111561085c5761085c6115b5565b604051908082528060200260200182016040528015610885578160200160208202803683370190505b50905060005b6001600160a01b03851660009081526006602052604090206001015481101561090b576001600160a01b03851660009081526006602052604090206001018054829081106108db576108db611979565b90600052602060002001548282815181106108f8576108f8611979565b602090810291909101015260010161088b565b509392505050565b6000828152600860205260408120805461092c9061182e565b90501161093857600080fd5b600082815260086020526040812060010180546109549061182e565b90501161096057600080fd5b604080516060810182526000848152600860205291822080548291906109859061182e565b80601f01602080910402602001604051908101604052809291908181526020018280546109b19061182e565b80156109fe5780601f106109d3576101008083540402835291602001916109fe565b820191906000526020600020905b8154815290600101906020018083116109e157829003601f168201915b50505050508152602001600860008681526020019081526020016000206001018054610a299061182e565b80601f0160208091040260200160405190810160405280929190818152602001828054610a559061182e565b8015610aa25780601f10610a7757610100808354040283529160200191610aa2565b820191906000526020600020905b815481529060010190602001808311610a8557829003601f168201915b505050918352505083151560209182015260008581526009825260408120805460018101825590825291902082519293508392600390920201908190610ae890826118b9565b5060208201516001820190610afd90826118b9565b50604091820151600291909101805460ff191691151591909117905580516060810182526000818301818152825282516020808201855282825280840191909152868252600890529190912081518190610b5790826118b9565b5060208201516001820190610b6c90826118b9565b50905050610b798361115d565b6040805184815283151560208201527f961a406eff4d590fdb9e54bdba584518be2e6a7f91b34d5714330f83d571b031910160405180910390a1505050565b6000610bc484826118b9565b506001610bd182826118b9565b506002610bde83826118b9565b50505050565b60408051808201909152606080825260208201526001600160a01b038216600090815260066020526040908190208151808301909252805482908290610c299061182e565b80601f0160208091040260200160405190810160405280929190818152602001828054610c559061182e565b8015610ca25780601f10610c7757610100808354040283529160200191610ca2565b820191906000526020600020905b815481529060010190602001808311610c8557829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015610cfa57602002820191906000526020600020905b815481526020019060010190808311610ce6575b5050505050815250509050919050565b60008381526008602052604081208054610d239061182e565b905011610d2f57600080fd5b60008381526008602052604090206001018054610d4b9061182e565b159050610d5757600080fd5b6000838152600860205260409020600101610bde82848361198f565b60008381526008602052604090208054610d8c9061182e565b159050610dd25760405162461bcd60e51b815260206004820152601460248201527331b430b61034b99030b63932b0b23c9039b2ba1760611b604482015260640161033c565b6000838152600860205260409020610deb82848361198f565b5050600a80546001810182556000919091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8019190915550565b60008281526007602052604090205463ffffffff16610e765760405162461bcd60e51b815260206004820152600c60248201526b696e76616c69642066696c6560a01b604482015260640161033c565b6001600160a01b031660009081526006602090815260408220600190810180549182018155835291200155565b600a54606090819060008167ffffffffffffffff811115610ec657610ec66115b5565b604051908082528060200260200182016040528015610eef578160200160208202803683370190505b50905060008267ffffffffffffffff811115610f0d57610f0d6115b5565b604051908082528060200260200182016040528015610f5257816020015b6040805180820190915260608082526020820152815260200190600190039081610f2b5790505b50905060005b83811015611109576000600a8281548110610f7557610f75611979565b9060005260206000200154905080848381518110610f9557610f95611979565b60200260200101818152505060086000828152602001908152602001600020604051806040016040529081600082018054610fcf9061182e565b80601f0160208091040260200160405190810160405280929190818152602001828054610ffb9061182e565b80156110485780601f1061101d57610100808354040283529160200191611048565b820191906000526020600020905b81548152906001019060200180831161102b57829003601f168201915b505050505081526020016001820180546110619061182e565b80601f016020809104026020016040519081016040528092919081815260200182805461108d9061182e565b80156110da5780601f106110af576101008083540402835291602001916110da565b820191906000526020600020905b8154815290600101906020018083116110bd57829003601f168201915b5050505050815250508383815181106110f5576110f5611979565b602090810291909101015250600101610f58565b509094909350915050565b611159828260405160240161112a929190611a50565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b17905261123a565b5050565b60008060005b600a548110156111a557600a818154811061118057611180611979565b9060005260206000200154840361119d57600192508091506111a5565b600101611163565b50805b600a546111b790600190611a90565b81101561120d57600a6111cb826001611aa9565b815481106111db576111db611979565b9060005260206000200154600a82815481106111f9576111f9611979565b6000918252602090912001556001016111a8565b50600a80548061121f5761121f611abc565b60019003818190600052602060002001600090559055505050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b828054828255906000526020600020908101928215611296579160200282015b8281111561129657825182559160200191906001019061127b565b506112a29291506112a6565b5090565b5b808211156112a257600081556001016112a7565b80356001600160a01b03811681146112d257600080fd5b919050565b6000806000606084860312156112ec57600080fd5b83359250602084013563ffffffff8116811461130757600080fd5b9150611315604085016112bb565b90509250925092565b60006020828403121561133057600080fd5b5035919050565b6000815180845260005b8181101561135d57602081850181015186830182015201611341565b506000602082860101526020601f19601f83011685010191505092915050565b600060208083018184528085518083526040925060408601915060408160051b87010184880160005b8381101561140457603f198984030185528151606081518186526113cc82870182611337565b915050888201518582038a8701526113e48282611337565b9289015115159589019590955250948701949250908601906001016113a6565b509098975050505050505050565b60208152600082516060602084015261142e6080840182611337565b90506020840151601f198085840301604086015261144c8383611337565b925060408601519150808584030160608601525061146a8282611337565b95945050505050565b60008083601f84011261148557600080fd5b50813567ffffffffffffffff81111561149d57600080fd5b6020830191508360208285010111156114b557600080fd5b9250929050565b6000806000604084860312156114d157600080fd5b6114da846112bb565b9250602084013567ffffffffffffffff8111156114f657600080fd5b61150286828701611473565b9497909650939450505050565b60006020828403121561152157600080fd5b61152a826112bb565b9392505050565b60008151808452602080850194506020840160005b8381101561156257815187529582019590820190600101611546565b509495945050505050565b60208152600061152a6020830184611531565b6000806040838503121561159357600080fd5b82359150602083013580151581146115aa57600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff808411156115e6576115e66115b5565b604051601f8501601f19908116603f0116810190828211818310171561160e5761160e6115b5565b8160405280935085815286868601111561162757600080fd5b858560208301376000602087830101525050509392505050565b600082601f83011261165257600080fd5b61152a838335602085016115cb565b60008060006060848603121561167657600080fd5b833567ffffffffffffffff8082111561168e57600080fd5b818601915086601f8301126116a257600080fd5b6116b1878335602085016115cb565b945060208601359150808211156116c757600080fd5b6116d387838801611641565b935060408601359150808211156116e957600080fd5b506116f686828701611641565b9150509250925092565b60208152600082516040602084015261171c6060840182611337565b90506020840151601f1984830301604085015261146a8282611531565b60008060006040848603121561174e57600080fd5b83359250602084013567ffffffffffffffff8111156114f657600080fd5b6000806040838503121561177f57600080fd5b8235915061178f602084016112bb565b90509250929050565b600060408083526117ac6040840186611531565b6020848203818601528186518084528284019150828160051b85010183890160005b8381101561181e57868303601f19018552815180518985526117f28a860182611337565b91880151858303868a015291905061180a8183611337565b9688019694505050908501906001016117ce565b50909a9950505050505050505050565b600181811c9082168061184257607f821691505b60208210810361186257634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156118b4576000816000526020600020601f850160051c810160208610156118915750805b601f850160051c820191505b818110156118b05782815560010161189d565b5050505b505050565b815167ffffffffffffffff8111156118d3576118d36115b5565b6118e7816118e1845461182e565b84611868565b602080601f83116001811461191c57600084156119045750858301515b600019600386901b1c1916600185901b1785556118b0565b600085815260208120601f198616915b8281101561194b5788860151825594840194600190910190840161192c565b50858210156119695787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b67ffffffffffffffff8311156119a7576119a76115b5565b6119bb836119b5835461182e565b83611868565b6000601f8411600181146119ef57600085156119d75750838201355b600019600387901b1c1916600186901b178355611a49565b600083815260209020601f19861690835b82811015611a205786850135825560209485019460019092019101611a00565b5086821015611a3d5760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b604081526000611a636040830185611337565b905060018060a01b03831660208301529392505050565b634e487b7160e01b600052601160045260246000fd5b81810381811115611aa357611aa3611a7a565b92915050565b80820180821115611aa357611aa3611a7a565b634e487b7160e01b600052603160045260246000fdfea264697066735822122085d4525a669daa77a449e2684333b4e9c6595e66cb6424798478a7f2516308ec64736f6c63430008190033",
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

// GetAuditingLogs is a free data retrieval call binding the contract method 0x7f838211.
//
// Solidity: function GetAuditingLogs(bytes32 _hash) view returns((bytes,bytes,bool)[])
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
// Solidity: function GetAuditingLogs(bytes32 _hash) view returns((bytes,bytes,bool)[])
func (_XZ21 *XZ21Session) GetAuditingLogs(_hash [32]byte) ([]XZ21AuditingLog, error) {
	return _XZ21.Contract.GetAuditingLogs(&_XZ21.CallOpts, _hash)
}

// GetAuditingLogs is a free data retrieval call binding the contract method 0x7f838211.
//
// Solidity: function GetAuditingLogs(bytes32 _hash) view returns((bytes,bytes,bool)[])
func (_XZ21 *XZ21CallerSession) GetAuditingLogs(_hash [32]byte) ([]XZ21AuditingLog, error) {
	return _XZ21.Contract.GetAuditingLogs(&_XZ21.CallOpts, _hash)
}

// GetAuditingReqList is a free data retrieval call binding the contract method 0xf65de599.
//
// Solidity: function GetAuditingReqList() view returns(bytes32[], (bytes,bytes)[])
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
// Solidity: function GetAuditingReqList() view returns(bytes32[], (bytes,bytes)[])
func (_XZ21 *XZ21Session) GetAuditingReqList() ([][32]byte, []XZ21AuditingReq, error) {
	return _XZ21.Contract.GetAuditingReqList(&_XZ21.CallOpts)
}

// GetAuditingReqList is a free data retrieval call binding the contract method 0xf65de599.
//
// Solidity: function GetAuditingReqList() view returns(bytes32[], (bytes,bytes)[])
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
