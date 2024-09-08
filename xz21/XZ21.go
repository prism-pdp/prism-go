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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_addrSP\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_addrTPA\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"AppendOwner\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"EnrollAccount\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"GetAccount\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.Account\",\"components\":[{\"name\":\"pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fileList\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetAuditingLogs\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structXZ21.AuditingLog[]\",\"components\":[{\"name\":\"chal\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetAuditingReqList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structXZ21.AuditingReq[]\",\"components\":[{\"name\":\"chal\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetFileList\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetParam\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.Param\",\"components\":[{\"name\":\"P\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"U\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"G\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ReadFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.FileProperty\",\"components\":[{\"name\":\"splitNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RegisterFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_splitNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"RegisterParam\",\"inputs\":[{\"name\":\"_p\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_g\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_u\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SearchFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.FileProperty\",\"components\":[{\"name\":\"splitNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SetAuditingResult\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SetChal\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_chal\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SetProof\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addrSM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrSP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrTPA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EventSetAuditingResult\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600f57600080fd5b50604051611bd2380380611bd2833981016040819052602c916082565b600380546001600160a01b03199081163317909155600480546001600160a01b039485169083161790556005805492909316911617905560b0565b80516001600160a01b0381168114607d57600080fd5b919050565b60008060408385031215609457600080fd5b609b836067565b915060a7602084016067565b90509250929050565b611b13806100bf6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80639a2b5e63116100a2578063d746cdd011610071578063d746cdd0146102b2578063debffd9a146102c5578063f64b362a146102d8578063f65de599146102eb578063f9ae523d1461012557600080fd5b80639a2b5e631461024c578063b645fd631461026c578063c1f16b161461027f578063d36741481461029257600080fd5b806379ef33df116100de57806379ef33df146101f15780637f838211146102045780638a04d5e51461022457806392f4dfa21461023957600080fd5b806322dd2b1814610110578063537322e31461012557806358f5823c146101b35780636b884afb146101de575b600080fd5b61012361011e3660046112e2565b610301565b005b610184610133366004611329565b60408051808201909152600080825260208201525060009081526007602090815260409182902082518084019093525463ffffffff8116835264010000000090046001600160a01b03169082015290565b60408051825163ffffffff1681526020928301516001600160a01b031692810192909252015b60405180910390f35b6003546101c6906001600160a01b031681565b6040516001600160a01b0390911681526020016101aa565b6004546101c6906001600160a01b031681565b6005546101c6906001600160a01b031681565b610217610212366004611329565b6103b3565b6040516101aa9190611388565b61022c61055b565b6040516101aa919061141d565b6101236102473660046114c7565b610747565b61025f61025a36600461151a565b61082f565b6040516101aa9190611578565b61012361027a36600461158b565b61091e565b61012361028d36600461166c565b610bc3565b6102a56102a036600461151a565b610bef565b6040516101aa919061170b565b6101236102c0366004611744565b610d15565b6101236102d3366004611744565b610d7e565b6101236102e6366004611777565b610e31565b6102f3610eae565b6040516101aa9291906117a3565b60008263ffffffff16116103505760405162461bcd60e51b8152602060048201526011602482015270696e76616c69642073706c6974206e756d60781b60448201526064015b60405180910390fd5b6000838152600760209081526040808320805463ffffffff969096166001600160c01b0319909616959095176401000000006001600160a01b03959095169485021790945591815260068252918220600190810180549182018155835291200155565b606060096000838152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b82821015610550578382906000526020600020906003020160405180606001604052908160008201805461041b90611839565b80601f016020809104026020016040519081016040528092919081815260200182805461044790611839565b80156104945780601f1061046957610100808354040283529160200191610494565b820191906000526020600020905b81548152906001019060200180831161047757829003601f168201915b505050505081526020016001820180546104ad90611839565b80601f01602080910402602001604051908101604052809291908181526020018280546104d990611839565b80156105265780601f106104fb57610100808354040283529160200191610526565b820191906000526020600020905b81548152906001019060200180831161050957829003601f168201915b50505091835250506002919091015460ff16151560209182015290825260019290920191016103e8565b505050509050919050565b61057f60405180606001604052806060815260200160608152602001606081525090565b600060405180606001604052908160008201805461059c90611839565b80601f01602080910402602001604051908101604052809291908181526020018280546105c890611839565b80156106155780601f106105ea57610100808354040283529160200191610615565b820191906000526020600020905b8154815290600101906020018083116105f857829003601f168201915b5050505050815260200160018201805461062e90611839565b80601f016020809104026020016040519081016040528092919081815260200182805461065a90611839565b80156106a75780601f1061067c576101008083540402835291602001916106a7565b820191906000526020600020905b81548152906001019060200180831161068a57829003601f168201915b505050505081526020016002820180546106c090611839565b80601f01602080910402602001604051908101604052809291908181526020018280546106ec90611839565b80156107395780601f1061070e57610100808354040283529160200191610739565b820191906000526020600020905b81548152906001019060200180831161071c57829003601f168201915b505050505081525050905090565b6003546001600160a01b031633811461075f57600080fd5b61079e6040518060400160405280601e81526020017f456e726f6c6c205355206163636f756e742028416464726573733a25732900008152508561111f565b6040805160606020601f86018190040282018101835291810184815290918291908690869081908501838280828437600092018290525093855250506040805183815260208082018352948501526001600160a01b0389168352600690935250208151819061080d90826118c4565b5060208281015180516108269260018501920190611266565b50505050505050565b6001600160a01b0381166000908152600660205260408120600101546060918167ffffffffffffffff811115610867576108676115c0565b604051908082528060200260200182016040528015610890578160200160208202803683370190505b50905060005b6001600160a01b038516600090815260066020526040902060010154811015610916576001600160a01b03851660009081526006602052604090206001018054829081106108e6576108e6611984565b906000526020600020015482828151811061090357610903611984565b6020908102919091010152600101610896565b509392505050565b6000828152600860205260408120805461093790611839565b90501161094357600080fd5b6000828152600860205260408120600101805461095f90611839565b90501161096b57600080fd5b6040805160608101825260008481526008602052918220805482919061099090611839565b80601f01602080910402602001604051908101604052809291908181526020018280546109bc90611839565b8015610a095780601f106109de57610100808354040283529160200191610a09565b820191906000526020600020905b8154815290600101906020018083116109ec57829003601f168201915b50505050508152602001600860008681526020019081526020016000206001018054610a3490611839565b80601f0160208091040260200160405190810160405280929190818152602001828054610a6090611839565b8015610aad5780601f10610a8257610100808354040283529160200191610aad565b820191906000526020600020905b815481529060010190602001808311610a9057829003601f168201915b505050918352505083151560209182015260008581526009825260408120805460018101825590825291902082519293508392600390920201908190610af390826118c4565b5060208201516001820190610b0890826118c4565b50604091820151600291909101805460ff191691151591909117905580516060810182526000818301818152825282516020808201855282825280840191909152868252600890529190912081518190610b6290826118c4565b5060208201516001820190610b7790826118c4565b50905050610b8483611168565b6040805184815283151560208201527f961a406eff4d590fdb9e54bdba584518be2e6a7f91b34d5714330f83d571b031910160405180910390a1505050565b6000610bcf84826118c4565b506001610bdc82826118c4565b506002610be983826118c4565b50505050565b60408051808201909152606080825260208201526001600160a01b038216600090815260066020526040908190208151808301909252805482908290610c3490611839565b80601f0160208091040260200160405190810160405280929190818152602001828054610c6090611839565b8015610cad5780601f10610c8257610100808354040283529160200191610cad565b820191906000526020600020905b815481529060010190602001808311610c9057829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015610d0557602002820191906000526020600020905b815481526020019060010190808311610cf1575b5050505050815250509050919050565b60008381526008602052604081208054610d2e90611839565b905011610d3a57600080fd5b60008381526008602052604090206001018054610d5690611839565b159050610d6257600080fd5b6000838152600860205260409020600101610be982848361199a565b60008381526008602052604090208054610d9790611839565b159050610ddd5760405162461bcd60e51b815260206004820152601460248201527331b430b61034b99030b63932b0b23c9039b2ba1760611b6044820152606401610347565b6000838152600860205260409020610df682848361199a565b5050600a80546001810182556000919091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8019190915550565b60008281526007602052604090205463ffffffff16610e815760405162461bcd60e51b815260206004820152600c60248201526b696e76616c69642066696c6560a01b6044820152606401610347565b6001600160a01b031660009081526006602090815260408220600190810180549182018155835291200155565b600a54606090819060008167ffffffffffffffff811115610ed157610ed16115c0565b604051908082528060200260200182016040528015610efa578160200160208202803683370190505b50905060008267ffffffffffffffff811115610f1857610f186115c0565b604051908082528060200260200182016040528015610f5d57816020015b6040805180820190915260608082526020820152815260200190600190039081610f365790505b50905060005b83811015611114576000600a8281548110610f8057610f80611984565b9060005260206000200154905080848381518110610fa057610fa0611984565b60200260200101818152505060086000828152602001908152602001600020604051806040016040529081600082018054610fda90611839565b80601f016020809104026020016040519081016040528092919081815260200182805461100690611839565b80156110535780601f1061102857610100808354040283529160200191611053565b820191906000526020600020905b81548152906001019060200180831161103657829003601f168201915b5050505050815260200160018201805461106c90611839565b80601f016020809104026020016040519081016040528092919081815260200182805461109890611839565b80156110e55780601f106110ba576101008083540402835291602001916110e5565b820191906000526020600020905b8154815290600101906020018083116110c857829003601f168201915b50505050508152505083838151811061110057611100611984565b602090810291909101015250600101610f63565b509094909350915050565b6111648282604051602401611135929190611a5b565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b179052611245565b5050565b60008060005b600a548110156111b057600a818154811061118b5761118b611984565b906000526020600020015484036111a857600192508091506111b0565b60010161116e565b50805b600a546111c290600190611a9b565b81101561121857600a6111d6826001611ab4565b815481106111e6576111e6611984565b9060005260206000200154600a828154811061120457611204611984565b6000918252602090912001556001016111b3565b50600a80548061122a5761122a611ac7565b60019003818190600052602060002001600090559055505050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b8280548282559060005260206000209081019282156112a1579160200282015b828111156112a1578251825591602001919060010190611286565b506112ad9291506112b1565b5090565b5b808211156112ad57600081556001016112b2565b80356001600160a01b03811681146112dd57600080fd5b919050565b6000806000606084860312156112f757600080fd5b83359250602084013563ffffffff8116811461131257600080fd5b9150611320604085016112c6565b90509250925092565b60006020828403121561133b57600080fd5b5035919050565b6000815180845260005b818110156113685760208185018101518683018201520161134c565b506000602082860101526020601f19601f83011685010191505092915050565b600060208083018184528085518083526040925060408601915060408160051b87010184880160005b8381101561140f57603f198984030185528151606081518186526113d782870182611342565b915050888201518582038a8701526113ef8282611342565b9289015115159589019590955250948701949250908601906001016113b1565b509098975050505050505050565b6020815260008251606060208401526114396080840182611342565b90506020840151601f19808584030160408601526114578383611342565b92506040860151915080858403016060860152506114758282611342565b95945050505050565b60008083601f84011261149057600080fd5b50813567ffffffffffffffff8111156114a857600080fd5b6020830191508360208285010111156114c057600080fd5b9250929050565b6000806000604084860312156114dc57600080fd5b6114e5846112c6565b9250602084013567ffffffffffffffff81111561150157600080fd5b61150d8682870161147e565b9497909650939450505050565b60006020828403121561152c57600080fd5b611535826112c6565b9392505050565b60008151808452602080850194506020840160005b8381101561156d57815187529582019590820190600101611551565b509495945050505050565b602081526000611535602083018461153c565b6000806040838503121561159e57600080fd5b82359150602083013580151581146115b557600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff808411156115f1576115f16115c0565b604051601f8501601f19908116603f01168101908282118183101715611619576116196115c0565b8160405280935085815286868601111561163257600080fd5b858560208301376000602087830101525050509392505050565b600082601f83011261165d57600080fd5b611535838335602085016115d6565b60008060006060848603121561168157600080fd5b833567ffffffffffffffff8082111561169957600080fd5b818601915086601f8301126116ad57600080fd5b6116bc878335602085016115d6565b945060208601359150808211156116d257600080fd5b6116de8783880161164c565b935060408601359150808211156116f457600080fd5b506117018682870161164c565b9150509250925092565b6020815260008251604060208401526117276060840182611342565b90506020840151601f19848303016040850152611475828261153c565b60008060006040848603121561175957600080fd5b83359250602084013567ffffffffffffffff81111561150157600080fd5b6000806040838503121561178a57600080fd5b8235915061179a602084016112c6565b90509250929050565b600060408083526117b7604084018661153c565b6020848203818601528186518084528284019150828160051b85010183890160005b8381101561182957868303601f19018552815180518985526117fd8a860182611342565b91880151858303868a01529190506118158183611342565b9688019694505050908501906001016117d9565b50909a9950505050505050505050565b600181811c9082168061184d57607f821691505b60208210810361186d57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156118bf576000816000526020600020601f850160051c8101602086101561189c5750805b601f850160051c820191505b818110156118bb578281556001016118a8565b5050505b505050565b815167ffffffffffffffff8111156118de576118de6115c0565b6118f2816118ec8454611839565b84611873565b602080601f831160018114611927576000841561190f5750858301515b600019600386901b1c1916600185901b1785556118bb565b600085815260208120601f198616915b8281101561195657888601518255948401946001909101908401611937565b50858210156119745787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b67ffffffffffffffff8311156119b2576119b26115c0565b6119c6836119c08354611839565b83611873565b6000601f8411600181146119fa57600085156119e25750838201355b600019600387901b1c1916600186901b178355611a54565b600083815260209020601f19861690835b82811015611a2b5786850135825560209485019460019092019101611a0b565b5086821015611a485760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b604081526000611a6e6040830185611342565b905060018060a01b03831660208301529392505050565b634e487b7160e01b600052601160045260246000fd5b81810381811115611aae57611aae611a85565b92915050565b80820180821115611aae57611aae611a85565b634e487b7160e01b600052603160045260246000fdfea26469706673582212203c9103d1d6799b14319c7d0a3285a0f4e4708c6d978d839f996cfbaa51cdf68d64736f6c63430008190033",
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

// ReadFile is a free data retrieval call binding the contract method 0xf9ae523d.
//
// Solidity: function ReadFile(bytes32 _hash) view returns((uint32,address))
func (_XZ21 *XZ21Caller) ReadFile(opts *bind.CallOpts, _hash [32]byte) (XZ21FileProperty, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "ReadFile", _hash)

	if err != nil {
		return *new(XZ21FileProperty), err
	}

	out0 := *abi.ConvertType(out[0], new(XZ21FileProperty)).(*XZ21FileProperty)

	return out0, err

}

// ReadFile is a free data retrieval call binding the contract method 0xf9ae523d.
//
// Solidity: function ReadFile(bytes32 _hash) view returns((uint32,address))
func (_XZ21 *XZ21Session) ReadFile(_hash [32]byte) (XZ21FileProperty, error) {
	return _XZ21.Contract.ReadFile(&_XZ21.CallOpts, _hash)
}

// ReadFile is a free data retrieval call binding the contract method 0xf9ae523d.
//
// Solidity: function ReadFile(bytes32 _hash) view returns((uint32,address))
func (_XZ21 *XZ21CallerSession) ReadFile(_hash [32]byte) (XZ21FileProperty, error) {
	return _XZ21.Contract.ReadFile(&_XZ21.CallOpts, _hash)
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
