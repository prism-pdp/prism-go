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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_addrSP\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_addrTPA\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"AppendOwner\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"EnrollAccount\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"GetAccount\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.Account\",\"components\":[{\"name\":\"pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fileList\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetAuditingLogs\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structXZ21.AuditingLog[]\",\"components\":[{\"name\":\"chal\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetAuditingReqList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structXZ21.AuditingReq[]\",\"components\":[{\"name\":\"chal\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetChalList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetFileList\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GetParam\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.Param\",\"components\":[{\"name\":\"P\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"U\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"G\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ReadFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.FileProperty\",\"components\":[{\"name\":\"splitNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RegisterFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_splitNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"RegisterParam\",\"inputs\":[{\"name\":\"_p\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_g\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_u\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SearchFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXZ21.FileProperty\",\"components\":[{\"name\":\"splitNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SetAuditingResult\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SetChal\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_chal\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SetProof\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addrSM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrSP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addrTPA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Result\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_msg\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600f57600080fd5b50604051611f30380380611f30833981016040819052602c916082565b600380546001600160a01b03199081163317909155600480546001600160a01b039485169083161790556005805492909316911617905560b0565b80516001600160a01b0381168114607d57600080fd5b919050565b60008060408385031215609457600080fd5b609b836067565b915060a7602084016067565b90509250929050565b611e71806100bf6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063af595ae3116100a2578063d746cdd011610071578063d746cdd0146102d3578063debffd9a146102e6578063f64b362a146102f9578063f65de5991461030c578063f9ae523d1461013057600080fd5b8063af595ae314610277578063b645fd631461028d578063c1f16b16146102a0578063d3674148146102b357600080fd5b806379ef33df116100e957806379ef33df146101fc5780637f8382111461020f5780638a04d5e51461022f57806392f4dfa2146102445780639a2b5e631461025757600080fd5b806322dd2b181461011b578063537322e31461013057806358f5823c146101be5780636b884afb146101e9575b600080fd5b61012e6101293660046115ce565b610322565b005b61018f61013e366004611615565b60408051808201909152600080825260208201525060009081526007602090815260409182902082518084019093525463ffffffff8116835264010000000090046001600160a01b03169082015290565b60408051825163ffffffff1681526020928301516001600160a01b031692810192909252015b60405180910390f35b6003546101d1906001600160a01b031681565b6040516001600160a01b0390911681526020016101b5565b6004546101d1906001600160a01b031681565b6005546101d1906001600160a01b031681565b61022261021d366004611615565b6103d4565b6040516101b59190611674565b61023761057c565b6040516101b59190611709565b61012e6102523660046117b3565b610768565b61026a610265366004611806565b610850565b6040516101b59190611864565b61027f61093f565b6040516101b5929190611877565b61012e61029b3660046118e9565b610af6565b61012e6102ae3660046119ca565b610dd0565b6102c66102c1366004611806565b610dfc565b6040516101b59190611a69565b61012e6102e1366004611aa2565b610f22565b61012e6102f4366004611aa2565b611078565b61012e610307366004611ad5565b611205565b610314611282565b6040516101b5929190611b01565b60008263ffffffff16116103715760405162461bcd60e51b8152602060048201526011602482015270696e76616c69642073706c6974206e756d60781b60448201526064015b60405180910390fd5b6000838152600760209081526040808320805463ffffffff969096166001600160c01b0319909616959095176401000000006001600160a01b03959095169485021790945591815260068252918220600190810180549182018155835291200155565b606060096000838152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b82821015610571578382906000526020600020906003020160405180606001604052908160008201805461043c90611b97565b80601f016020809104026020016040519081016040528092919081815260200182805461046890611b97565b80156104b55780601f1061048a576101008083540402835291602001916104b5565b820191906000526020600020905b81548152906001019060200180831161049857829003601f168201915b505050505081526020016001820180546104ce90611b97565b80601f01602080910402602001604051908101604052809291908181526020018280546104fa90611b97565b80156105475780601f1061051c57610100808354040283529160200191610547565b820191906000526020600020905b81548152906001019060200180831161052a57829003601f168201915b50505091835250506002919091015460ff1615156020918201529082526001929092019101610409565b505050509050919050565b6105a060405180606001604052806060815260200160608152602001606081525090565b60006040518060600160405290816000820180546105bd90611b97565b80601f01602080910402602001604051908101604052809291908181526020018280546105e990611b97565b80156106365780601f1061060b57610100808354040283529160200191610636565b820191906000526020600020905b81548152906001019060200180831161061957829003601f168201915b5050505050815260200160018201805461064f90611b97565b80601f016020809104026020016040519081016040528092919081815260200182805461067b90611b97565b80156106c85780601f1061069d576101008083540402835291602001916106c8565b820191906000526020600020905b8154815290600101906020018083116106ab57829003601f168201915b505050505081526020016002820180546106e190611b97565b80601f016020809104026020016040519081016040528092919081815260200182805461070d90611b97565b801561075a5780601f1061072f5761010080835404028352916020019161075a565b820191906000526020600020905b81548152906001019060200180831161073d57829003601f168201915b505050505081525050905090565b6003546001600160a01b031633811461078057600080fd5b6107bf6040518060400160405280601e81526020017f456e726f6c6c205355206163636f756e742028416464726573733a2573290000815250856114e8565b6040805160606020601f86018190040282018101835291810184815290918291908690869081908501838280828437600092018290525093855250506040805183815260208082018352948501526001600160a01b0389168352600690935250208151819061082e9082611c22565b5060208281015180516108479260018501920190611552565b50505050505050565b6001600160a01b0381166000908152600660205260408120600101546060918167ffffffffffffffff8111156108885761088861191e565b6040519080825280602002602001820160405280156108b1578160200160208202803683370190505b50905060005b6001600160a01b038516600090815260066020526040902060010154811015610937576001600160a01b038516600090815260066020526040902060010180548290811061090757610907611ce2565b906000526020600020015482828151811061092457610924611ce2565b60209081029190910101526001016108b7565b509392505050565b600a54606090819060008167ffffffffffffffff8111156109625761096261191e565b60405190808252806020026020018201604052801561098b578160200160208202803683370190505b50905060008267ffffffffffffffff8111156109a9576109a961191e565b6040519080825280602002602001820160405280156109dc57816020015b60608152602001906001900390816109c75790505b50905060005b83811015610aeb576000600a82815481106109ff576109ff611ce2565b9060005260206000200154905080848381518110610a1f57610a1f611ce2565b60209081029190910181019190915260008281526008909152604090208054610a4790611b97565b80601f0160208091040260200160405190810160405280929190818152602001828054610a7390611b97565b8015610ac05780601f10610a9557610100808354040283529160200191610ac0565b820191906000526020600020905b815481529060010190602001808311610aa357829003601f168201915b5050505050838381518110610ad757610ad7611ce2565b6020908102919091010152506001016109e2565b509094909350915050565b60008060005b600b54811015610b3e57600b8181548110610b1957610b19611ce2565b90600052602060002001548503610b365760019250809150610b3e565b600101610afc565b5081610b815760405162461bcd60e51b8152602060048201526012602482015271556e6b6e6f776e20686173682076616c756560701b6044820152606401610368565b60408051606081018252600086815260086020529182208054829190610ba690611b97565b80601f0160208091040260200160405190810160405280929190818152602001828054610bd290611b97565b8015610c1f5780601f10610bf457610100808354040283529160200191610c1f565b820191906000526020600020905b815481529060010190602001808311610c0257829003601f168201915b50505050508152602001600860008881526020019081526020016000206001018054610c4a90611b97565b80601f0160208091040260200160405190810160405280929190818152602001828054610c7690611b97565b8015610cc35780601f10610c9857610100808354040283529160200191610cc3565b820191906000526020600020905b815481529060010190602001808311610ca657829003601f168201915b505050918352505085151560209182015260008781526009825260408120805460018101825590825291902082519293508392600390920201908190610d099082611c22565b5060208201516001820190610d1e9082611c22565b50604091909101516002909101805460ff1916911515919091179055815b600b54610d4b90600190611d0e565b811015610da157600b610d5f826001611d27565b81548110610d6f57610d6f611ce2565b9060005260206000200154600b8281548110610d8d57610d8d611ce2565b600091825260209091200155600101610d3c565b50600b805480610db357610db3611d3a565b600190038181906000526020600020016000905590555050505050565b6000610ddc8482611c22565b506001610de98282611c22565b506002610df68382611c22565b50505050565b60408051808201909152606080825260208201526001600160a01b038216600090815260066020526040908190208151808301909252805482908290610e4190611b97565b80601f0160208091040260200160405190810160405280929190818152602001828054610e6d90611b97565b8015610eba5780601f10610e8f57610100808354040283529160200191610eba565b820191906000526020600020905b815481529060010190602001808311610e9d57829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015610f1257602002820191906000526020600020905b815481526020019060010190808311610efe575b5050505050815250509050919050565b60008060005b600a54811015610f6a57600a8181548110610f4557610f45611ce2565b90600052602060002001548603610f625760019250809150610f6a565b600101610f28565b5081610fad5760405162461bcd60e51b8152602060048201526012602482015271556e6b6e6f776e20686173682076616c756560701b6044820152606401610368565b6000858152600860205260409020600101610fc9848683611d50565b50600b80546001810182556000919091527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db901859055805b600a5461101090600190611d0e565b81101561106657600a611024826001611d27565b8154811061103457611034611ce2565b9060005260206000200154600a828154811061105257611052611ce2565b600091825260209091200155600101611001565b50600a805480610db357610db3611d3a565b6000838152600860205260408120805461109190611b97565b905011156110f257336001600160a01b03167fee09d27cc57ad50ec0f8a69b47fed3feed38ac3215a611242b45c51ca2368a086040516110e9906020808252600490820152630536b69760e41b604082015260600190565b60405180910390a25b6040805160606020601f8501819004028201810183529181018381526000928291908690869081908501838280828437600092018290525093855250506040805160208181018352848252948501528883526008909352502081519192508291819061115e9082611c22565b50602082015160018201906111739082611c22565b5050600a80546001810182556000919091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8018590555060405133907fee09d27cc57ad50ec0f8a69b47fed3feed38ac3215a611242b45c51ca2368a08906111f7906020808252600790820152665375636365737360c81b604082015260600190565b60405180910390a250505050565b60008281526007602052604090205463ffffffff166112555760405162461bcd60e51b815260206004820152600c60248201526b696e76616c69642066696c6560a01b6044820152606401610368565b6001600160a01b031660009081526006602090815260408220600190810180549182018155835291200155565b600b54606090819060008167ffffffffffffffff8111156112a5576112a561191e565b6040519080825280602002602001820160405280156112ce578160200160208202803683370190505b50905060008267ffffffffffffffff8111156112ec576112ec61191e565b60405190808252806020026020018201604052801561133157816020015b604080518082019091526060808252602082015281526020019060019003908161130a5790505b50905060005b83811015610aeb576000600b828154811061135457611354611ce2565b906000526020600020015490508084838151811061137457611374611ce2565b602002602001018181525050600860008281526020019081526020016000206040518060400160405290816000820180546113ae90611b97565b80601f01602080910402602001604051908101604052809291908181526020018280546113da90611b97565b80156114275780601f106113fc57610100808354040283529160200191611427565b820191906000526020600020905b81548152906001019060200180831161140a57829003601f168201915b5050505050815260200160018201805461144090611b97565b80601f016020809104026020016040519081016040528092919081815260200182805461146c90611b97565b80156114b95780601f1061148e576101008083540402835291602001916114b9565b820191906000526020600020905b81548152906001019060200180831161149c57829003601f168201915b5050505050815250508383815181106114d4576114d4611ce2565b602090810291909101015250600101611337565b61152d82826040516024016114fe929190611e11565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b179052611531565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b82805482825590600052602060002090810192821561158d579160200282015b8281111561158d578251825591602001919060010190611572565b5061159992915061159d565b5090565b5b80821115611599576000815560010161159e565b80356001600160a01b03811681146115c957600080fd5b919050565b6000806000606084860312156115e357600080fd5b83359250602084013563ffffffff811681146115fe57600080fd5b915061160c604085016115b2565b90509250925092565b60006020828403121561162757600080fd5b5035919050565b6000815180845260005b8181101561165457602081850181015186830182015201611638565b506000602082860101526020601f19601f83011685010191505092915050565b600060208083018184528085518083526040925060408601915060408160051b87010184880160005b838110156116fb57603f198984030185528151606081518186526116c38287018261162e565b915050888201518582038a8701526116db828261162e565b92890151151595890195909552509487019492509086019060010161169d565b509098975050505050505050565b602081526000825160606020840152611725608084018261162e565b90506020840151601f1980858403016040860152611743838361162e565b9250604086015191508085840301606086015250611761828261162e565b95945050505050565b60008083601f84011261177c57600080fd5b50813567ffffffffffffffff81111561179457600080fd5b6020830191508360208285010111156117ac57600080fd5b9250929050565b6000806000604084860312156117c857600080fd5b6117d1846115b2565b9250602084013567ffffffffffffffff8111156117ed57600080fd5b6117f98682870161176a565b9497909650939450505050565b60006020828403121561181857600080fd5b611821826115b2565b9392505050565b60008151808452602080850194506020840160005b838110156118595781518752958201959082019060010161183d565b509495945050505050565b6020815260006118216020830184611828565b60408152600061188a6040830185611828565b6020838203818501528185518084528284019150828160051b85010183880160005b838110156118da57601f198784030185526118c883835161162e565b948601949250908501906001016118ac565b50909998505050505050505050565b600080604083850312156118fc57600080fd5b823591506020830135801515811461191357600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff8084111561194f5761194f61191e565b604051601f8501601f19908116603f011681019082821181831017156119775761197761191e565b8160405280935085815286868601111561199057600080fd5b858560208301376000602087830101525050509392505050565b600082601f8301126119bb57600080fd5b61182183833560208501611934565b6000806000606084860312156119df57600080fd5b833567ffffffffffffffff808211156119f757600080fd5b818601915086601f830112611a0b57600080fd5b611a1a87833560208501611934565b94506020860135915080821115611a3057600080fd5b611a3c878388016119aa565b93506040860135915080821115611a5257600080fd5b50611a5f868287016119aa565b9150509250925092565b602081526000825160406020840152611a85606084018261162e565b90506020840151601f198483030160408501526117618282611828565b600080600060408486031215611ab757600080fd5b83359250602084013567ffffffffffffffff8111156117ed57600080fd5b60008060408385031215611ae857600080fd5b82359150611af8602084016115b2565b90509250929050565b60006040808352611b156040840186611828565b6020848203818601528186518084528284019150828160051b85010183890160005b83811015611b8757868303601f1901855281518051898552611b5b8a86018261162e565b91880151858303868a0152919050611b73818361162e565b968801969450505090850190600101611b37565b50909a9950505050505050505050565b600181811c90821680611bab57607f821691505b602082108103611bcb57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115611c1d576000816000526020600020601f850160051c81016020861015611bfa5750805b601f850160051c820191505b81811015611c1957828155600101611c06565b5050505b505050565b815167ffffffffffffffff811115611c3c57611c3c61191e565b611c5081611c4a8454611b97565b84611bd1565b602080601f831160018114611c855760008415611c6d5750858301515b600019600386901b1c1916600185901b178555611c19565b600085815260208120601f198616915b82811015611cb457888601518255948401946001909101908401611c95565b5085821015611cd25787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b81810381811115611d2157611d21611cf8565b92915050565b80820180821115611d2157611d21611cf8565b634e487b7160e01b600052603160045260246000fd5b67ffffffffffffffff831115611d6857611d6861191e565b611d7c83611d768354611b97565b83611bd1565b6000601f841160018114611db05760008515611d985750838201355b600019600387901b1c1916600186901b178355611e0a565b600083815260209020601f19861690835b82811015611de15786850135825560209485019460019092019101611dc1565b5086821015611dfe5760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b604081526000611e24604083018561162e565b905060018060a01b0383166020830152939250505056fea26469706673582212209ff359ffc18482b26a54aaf3d77b62ce3d3c13aca6b36ffb8073223293e2879f64736f6c63430008190033",
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

// GetChalList is a free data retrieval call binding the contract method 0xaf595ae3.
//
// Solidity: function GetChalList() view returns(bytes32[], bytes[])
func (_XZ21 *XZ21Caller) GetChalList(opts *bind.CallOpts) ([][32]byte, [][]byte, error) {
	var out []interface{}
	err := _XZ21.contract.Call(opts, &out, "GetChalList")

	if err != nil {
		return *new([][32]byte), *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	out1 := *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)

	return out0, out1, err

}

// GetChalList is a free data retrieval call binding the contract method 0xaf595ae3.
//
// Solidity: function GetChalList() view returns(bytes32[], bytes[])
func (_XZ21 *XZ21Session) GetChalList() ([][32]byte, [][]byte, error) {
	return _XZ21.Contract.GetChalList(&_XZ21.CallOpts)
}

// GetChalList is a free data retrieval call binding the contract method 0xaf595ae3.
//
// Solidity: function GetChalList() view returns(bytes32[], bytes[])
func (_XZ21 *XZ21CallerSession) GetChalList() ([][32]byte, [][]byte, error) {
	return _XZ21.Contract.GetChalList(&_XZ21.CallOpts)
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

// XZ21ResultIterator is returned from FilterResult and is used to iterate over the raw logs and unpacked data for Result events raised by the XZ21 contract.
type XZ21ResultIterator struct {
	Event *XZ21Result // Event containing the contract specifics and raw log

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
func (it *XZ21ResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XZ21Result)
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
		it.Event = new(XZ21Result)
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
func (it *XZ21ResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XZ21ResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XZ21Result represents a Result event raised by the XZ21 contract.
type XZ21Result struct {
	From common.Address
	Msg  string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterResult is a free log retrieval operation binding the contract event 0xee09d27cc57ad50ec0f8a69b47fed3feed38ac3215a611242b45c51ca2368a08.
//
// Solidity: event Result(address indexed _from, string _msg)
func (_XZ21 *XZ21Filterer) FilterResult(opts *bind.FilterOpts, _from []common.Address) (*XZ21ResultIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _XZ21.contract.FilterLogs(opts, "Result", _fromRule)
	if err != nil {
		return nil, err
	}
	return &XZ21ResultIterator{contract: _XZ21.contract, event: "Result", logs: logs, sub: sub}, nil
}

// WatchResult is a free log subscription operation binding the contract event 0xee09d27cc57ad50ec0f8a69b47fed3feed38ac3215a611242b45c51ca2368a08.
//
// Solidity: event Result(address indexed _from, string _msg)
func (_XZ21 *XZ21Filterer) WatchResult(opts *bind.WatchOpts, sink chan<- *XZ21Result, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _XZ21.contract.WatchLogs(opts, "Result", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XZ21Result)
				if err := _XZ21.contract.UnpackLog(event, "Result", log); err != nil {
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

// ParseResult is a log parse operation binding the contract event 0xee09d27cc57ad50ec0f8a69b47fed3feed38ac3215a611242b45c51ca2368a08.
//
// Solidity: event Result(address indexed _from, string _msg)
func (_XZ21 *XZ21Filterer) ParseResult(log types.Log) (*XZ21Result, error) {
	event := new(XZ21Result)
	if err := _XZ21.contract.UnpackLog(event, "Result", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
