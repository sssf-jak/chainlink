// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrfv2plus_wrapper_load_test_consumer

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

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

var VRFV2PlusWrapperLoadTestConsumerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vrfV2PlusWrapper\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"LINKAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyVRFWrapperCanFulfill\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"WrappedRequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"}],\"name\":\"WrapperRequestMade\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinkToken\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fulfilmentTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fulfilmentBlockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_vrfV2PlusWrapper\",\"outputs\":[{\"internalType\":\"contractIVRFV2PlusWrapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"_numWords\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_requestCount\",\"type\":\"uint16\"}],\"name\":\"makeRequests\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"_numWords\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_requestCount\",\"type\":\"uint16\"}],\"name\":\"makeRequestsNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_averageFulfillmentInMillions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_fastestFulfillment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_lastRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_requestCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_requests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fulfilmentTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fulfilmentBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"native\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_responseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_slowestFulfillment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"}],\"name\":\"setLinkToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052600060055560006006556103e76007553480156200002157600080fd5b5060405162001fe238038062001fe28339810160408190526200004491620001eb565b3380600084846001600160a01b038216156200007657600080546001600160a01b0319166001600160a01b0384161790555b60601b6001600160601b031916608052506001600160a01b038216620000e35760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600180546001600160a01b0319166001600160a01b03848116919091179091558116156200011657620001168162000121565b505050505062000223565b6001600160a01b0381163314156200017c5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000da565b600280546001600160a01b0319166001600160a01b03838116918217909255600154604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b80516001600160a01b0381168114620001e657600080fd5b919050565b60008060408385031215620001ff57600080fd5b6200020a83620001ce565b91506200021a60208401620001ce565b90509250929050565b60805160601c611d6f62000273600039600081816102e9015281816104b80152818161097501528181610da40152818161132a0152818161143f015281816115cf015261164d0152611d6f6000f3fe60806040526004361061016e5760003560e01c80639c24ea40116100cb578063d826f88f1161007f578063e76d516811610059578063e76d51681461044b578063f176596214610476578063f2fde38b1461049657600080fd5b8063d826f88f146103d6578063d8a4676f14610402578063dc1670db1461043557600080fd5b8063a168fa89116100b0578063a168fa891461030b578063afacbf9c146103a0578063b1e21749146103c057600080fd5b80639c24ea40146102b75780639ed0868d146102d757600080fd5b806374dba124116101225780637a8042bd116101075780637a8042bd1461022b57806384276d811461024b5780638da5cb5b1461026b57600080fd5b806374dba1241461020057806379ba50971461021657600080fd5b80631fe543e3116101535780631fe543e3146101b2578063557d2e92146101d4578063737144bc146101ea57600080fd5b806312065fe01461017a5780631757f11c1461019c57600080fd5b3661017557005b600080fd5b34801561018657600080fd5b50475b6040519081526020015b60405180910390f35b3480156101a857600080fd5b5061018960065481565b3480156101be57600080fd5b506101d26101cd366004611941565b6104b6565b005b3480156101e057600080fd5b5061018960045481565b3480156101f657600080fd5b5061018960055481565b34801561020c57600080fd5b5061018960075481565b34801561022257600080fd5b506101d2610558565b34801561023757600080fd5b506101d261024636600461190f565b610659565b34801561025757600080fd5b506101d261026636600461190f565b610747565b34801561027757600080fd5b5060015473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610193565b3480156102c357600080fd5b506101d26102d23660046118b0565b610837565b3480156102e357600080fd5b506102927f000000000000000000000000000000000000000000000000000000000000000081565b34801561031757600080fd5b5061036961032636600461190f565b600a602052600090815260409020805460018201546003830154600484015460058501546006860154600790960154949560ff9485169593949293919290911687565b604080519788529515156020880152948601939093526060850191909152608084015260a0830152151560c082015260e001610193565b3480156103ac57600080fd5b506101d26103bb366004611a30565b6108ce565b3480156103cc57600080fd5b5061018960085481565b3480156103e257600080fd5b506101d26000600581905560068190556103e76007556004819055600355565b34801561040e57600080fd5b5061042261041d36600461190f565b610b7a565b6040516101939796959493929190611b88565b34801561044157600080fd5b5061018960035481565b34801561045757600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff16610292565b34801561048257600080fd5b506101d2610491366004611a30565b610cfd565b3480156104a257600080fd5b506101d26104b13660046118b0565b610fa1565b7f00000000000000000000000000000000000000000000000000000000000000003373ffffffffffffffffffffffffffffffffffffffff821614610549576040517f8ba9316e00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff821660248201526044015b60405180910390fd5b6105538383610fb5565b505050565b60025473ffffffffffffffffffffffffffffffffffffffff1633146105d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610540565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560028054909116905560405173ffffffffffffffffffffffffffffffffffffffff909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b610661611194565b60005473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb61069e60015473ffffffffffffffffffffffffffffffffffffffff1690565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff909116600482015260248101849052604401602060405180830381600087803b15801561070b57600080fd5b505af115801561071f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061074391906118ed565b5050565b61074f611194565b600061077060015473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d80600081146107c7576040519150601f19603f3d011682016040523d82523d6000602084013e6107cc565b606091505b5050905080610743576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f77697468647261774e6174697665206661696c656400000000000000000000006044820152606401610540565b60005473ffffffffffffffffffffffffffffffffffffffff1615610887576040517f64f778ae00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6108d6611194565b60005b8161ffff168161ffff161015610b73576000610905604051806020016040528060001515815250611217565b90506000610915878787856112d3565b6008819055905060006109266114e4565b6040517f4306d35400000000000000000000000000000000000000000000000000000000815263ffffffff8a16600482015290915060009073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690634306d3549060240160206040518083038186803b1580156109b757600080fd5b505afa1580156109cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ef9190611928565b604080516101008101825282815260006020808301828152845183815280830186528486019081524260608601526080850184905260a0850189905260c0850184905260e08501849052898452600a8352949092208351815591516001830180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790559251805194955091939092610a97926002850192910190611825565b5060608201516003820155608082015160048083019190915560a0830151600583015560c0830151600683015560e090920151600790910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790558054906000610b0a83611ccb565b9091555050600083815260096020526040908190208390555183907f5f56b4c20db9f5b294cbf6f681368de4a992a27e2de2ee702dcf2cbbfa791ec490610b549084815260200190565b60405180910390a2505050508080610b6b90611ca9565b9150506108d9565b5050505050565b6000818152600a602052604081205481906060908290819081908190610bfc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f72657175657374206e6f7420666f756e640000000000000000000000000000006044820152606401610540565b6000888152600a6020908152604080832081516101008101835281548152600182015460ff16151581850152600282018054845181870281018701865281815292959394860193830182828015610c7257602002820191906000526020600020905b815481526020019060010190808311610c5e575b50505050508152602001600382015481526020016004820154815260200160058201548152602001600682015481526020016007820160009054906101000a900460ff1615151515815250509050806000015181602001518260400151836060015184608001518560a001518660c00151975097509750975097509750975050919395979092949650565b610d05611194565b60005b8161ffff168161ffff161015610b73576000610d34604051806020016040528060011515815250611217565b90506000610d4487878785611581565b600881905590506000610d556114e4565b6040517f4b16093500000000000000000000000000000000000000000000000000000000815263ffffffff8a16600482015290915060009073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690634b1609359060240160206040518083038186803b158015610de657600080fd5b505afa158015610dfa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1e9190611928565b604080516101008101825282815260006020808301828152845183815280830186528486019081524260608601526080850184905260a0850189905260c08501849052600160e086018190528a8552600a84529590932084518155905194810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001695151595909517909455905180519495509193610ec59260028501920190611825565b5060608201516003820155608082015160048083019190915560a0830151600583015560c0830151600683015560e090920151600790910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790558054906000610f3883611ccb565b9091555050600083815260096020526040908190208390555183907f5f56b4c20db9f5b294cbf6f681368de4a992a27e2de2ee702dcf2cbbfa791ec490610f829084815260200190565b60405180910390a2505050508080610f9990611ca9565b915050610d08565b610fa9611194565b610fb281611707565b50565b6000828152600a602052604090205461102a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f72657175657374206e6f7420666f756e640000000000000000000000000000006044820152606401610540565b60006110346114e4565b600084815260096020526040812054919250906110519083611c92565b9050600061106282620f4240611c55565b90506006548211156110745760068290555b600754821061108557600754611087565b815b60075560035461109757806110ca565b6003546110a5906001611c02565b816003546005546110b69190611c55565b6110c09190611c02565b6110ca9190611c1a565b600555600380549060006110dd83611ccb565b90915550506000858152600a60209081526040909120600181810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169091179055855161113592600290920191870190611825565b506000858152600a602052604090819020426004820155600681018590555490517f6c84e12b4c188e61f1b4727024a5cf05c025fa58467e5eedf763c0744c89da7b916111859188918891611b5f565b60405180910390a15050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314611215576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610540565b565b60607f92fd13387c7fe7befbc38d303d6468778fb9731bc4583f17d92989c6fcfdeaaa8260405160240161125091511515815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915292915050565b600080546040517f4306d35400000000000000000000000000000000000000000000000000000000815263ffffffff8716600482015273ffffffffffffffffffffffffffffffffffffffff91821691634000aea0917f000000000000000000000000000000000000000000000000000000000000000091821690634306d3549060240160206040518083038186803b15801561136e57600080fd5b505afa158015611382573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113a69190611928565b888888886040516020016113bd9493929190611bcf565b6040516020818303038152906040526040518463ffffffff1660e01b81526004016113ea93929190611b2a565b602060405180830381600087803b15801561140457600080fd5b505af1158015611418573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061143c91906118ed565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663fc2a88c36040518163ffffffff1660e01b815260040160206040518083038186803b1580156114a357600080fd5b505afa1580156114b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114db9190611928565b95945050505050565b6000466114f0816117fe565b1561157a57606473ffffffffffffffffffffffffffffffffffffffff1663a3b1b31d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561153c57600080fd5b505afa158015611550573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115749190611928565b91505090565b4391505090565b6040517f4b16093500000000000000000000000000000000000000000000000000000000815263ffffffff85166004820152600090819073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690634b1609359060240160206040518083038186803b15801561161157600080fd5b505afa158015611625573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116499190611928565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639cfc058e82888888886040518663ffffffff1660e01b81526004016116ab9493929190611bcf565b6020604051808303818588803b1580156116c457600080fd5b505af11580156116d8573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906116fd9190611928565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116331415611787576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610540565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600154604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b600061a4b1821480611812575062066eed82145b8061181f575062066eee82145b92915050565b828054828255906000526020600020908101928215611860579160200282015b82811115611860578251825591602001919060010190611845565b5061186c929150611870565b5090565b5b8082111561186c5760008155600101611871565b803561ffff8116811461189757600080fd5b919050565b803563ffffffff8116811461189757600080fd5b6000602082840312156118c257600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146118e657600080fd5b9392505050565b6000602082840312156118ff57600080fd5b815180151581146118e657600080fd5b60006020828403121561192157600080fd5b5035919050565b60006020828403121561193a57600080fd5b5051919050565b6000806040838503121561195457600080fd5b8235915060208084013567ffffffffffffffff8082111561197457600080fd5b818601915086601f83011261198857600080fd5b81358181111561199a5761199a611d33565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f830116810181811085821117156119dd576119dd611d33565b604052828152858101935084860182860187018b10156119fc57600080fd5b600095505b83861015611a1f578035855260019590950194938601938601611a01565b508096505050505050509250929050565b60008060008060808587031215611a4657600080fd5b611a4f8561189c565b9350611a5d60208601611885565b9250611a6b6040860161189c565b9150611a7960608601611885565b905092959194509250565b600081518084526020808501945080840160005b83811015611ab457815187529582019590820190600101611a98565b509495945050505050565b6000815180845260005b81811015611ae557602081850181015186830182015201611ac9565b81811115611af7576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff841681528260208201526060604082015260006114db6060830184611abf565b838152606060208201526000611b786060830185611a84565b9050826040830152949350505050565b878152861515602082015260e060408201526000611ba960e0830188611a84565b90508560608301528460808301528360a08301528260c083015298975050505050505050565b600063ffffffff808716835261ffff86166020840152808516604084015250608060608301526116fd6080830184611abf565b60008219821115611c1557611c15611d04565b500190565b600082611c50577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611c8d57611c8d611d04565b500290565b600082821015611ca457611ca4611d04565b500390565b600061ffff80831681811415611cc157611cc1611d04565b6001019392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611cfd57611cfd611d04565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c6343000806000a",
}

var VRFV2PlusWrapperLoadTestConsumerABI = VRFV2PlusWrapperLoadTestConsumerMetaData.ABI

var VRFV2PlusWrapperLoadTestConsumerBin = VRFV2PlusWrapperLoadTestConsumerMetaData.Bin

func DeployVRFV2PlusWrapperLoadTestConsumer(auth *bind.TransactOpts, backend bind.ContractBackend, _link common.Address, _vrfV2PlusWrapper common.Address) (common.Address, *types.Transaction, *VRFV2PlusWrapperLoadTestConsumer, error) {
	parsed, err := VRFV2PlusWrapperLoadTestConsumerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFV2PlusWrapperLoadTestConsumerBin), backend, _link, _vrfV2PlusWrapper)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFV2PlusWrapperLoadTestConsumer{VRFV2PlusWrapperLoadTestConsumerCaller: VRFV2PlusWrapperLoadTestConsumerCaller{contract: contract}, VRFV2PlusWrapperLoadTestConsumerTransactor: VRFV2PlusWrapperLoadTestConsumerTransactor{contract: contract}, VRFV2PlusWrapperLoadTestConsumerFilterer: VRFV2PlusWrapperLoadTestConsumerFilterer{contract: contract}}, nil
}

type VRFV2PlusWrapperLoadTestConsumer struct {
	address common.Address
	abi     abi.ABI
	VRFV2PlusWrapperLoadTestConsumerCaller
	VRFV2PlusWrapperLoadTestConsumerTransactor
	VRFV2PlusWrapperLoadTestConsumerFilterer
}

type VRFV2PlusWrapperLoadTestConsumerCaller struct {
	contract *bind.BoundContract
}

type VRFV2PlusWrapperLoadTestConsumerTransactor struct {
	contract *bind.BoundContract
}

type VRFV2PlusWrapperLoadTestConsumerFilterer struct {
	contract *bind.BoundContract
}

type VRFV2PlusWrapperLoadTestConsumerSession struct {
	Contract     *VRFV2PlusWrapperLoadTestConsumer
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFV2PlusWrapperLoadTestConsumerCallerSession struct {
	Contract *VRFV2PlusWrapperLoadTestConsumerCaller
	CallOpts bind.CallOpts
}

type VRFV2PlusWrapperLoadTestConsumerTransactorSession struct {
	Contract     *VRFV2PlusWrapperLoadTestConsumerTransactor
	TransactOpts bind.TransactOpts
}

type VRFV2PlusWrapperLoadTestConsumerRaw struct {
	Contract *VRFV2PlusWrapperLoadTestConsumer
}

type VRFV2PlusWrapperLoadTestConsumerCallerRaw struct {
	Contract *VRFV2PlusWrapperLoadTestConsumerCaller
}

type VRFV2PlusWrapperLoadTestConsumerTransactorRaw struct {
	Contract *VRFV2PlusWrapperLoadTestConsumerTransactor
}

func NewVRFV2PlusWrapperLoadTestConsumer(address common.Address, backend bind.ContractBackend) (*VRFV2PlusWrapperLoadTestConsumer, error) {
	abi, err := abi.JSON(strings.NewReader(VRFV2PlusWrapperLoadTestConsumerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFV2PlusWrapperLoadTestConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusWrapperLoadTestConsumer{address: address, abi: abi, VRFV2PlusWrapperLoadTestConsumerCaller: VRFV2PlusWrapperLoadTestConsumerCaller{contract: contract}, VRFV2PlusWrapperLoadTestConsumerTransactor: VRFV2PlusWrapperLoadTestConsumerTransactor{contract: contract}, VRFV2PlusWrapperLoadTestConsumerFilterer: VRFV2PlusWrapperLoadTestConsumerFilterer{contract: contract}}, nil
}

func NewVRFV2PlusWrapperLoadTestConsumerCaller(address common.Address, caller bind.ContractCaller) (*VRFV2PlusWrapperLoadTestConsumerCaller, error) {
	contract, err := bindVRFV2PlusWrapperLoadTestConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusWrapperLoadTestConsumerCaller{contract: contract}, nil
}

func NewVRFV2PlusWrapperLoadTestConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFV2PlusWrapperLoadTestConsumerTransactor, error) {
	contract, err := bindVRFV2PlusWrapperLoadTestConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusWrapperLoadTestConsumerTransactor{contract: contract}, nil
}

func NewVRFV2PlusWrapperLoadTestConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFV2PlusWrapperLoadTestConsumerFilterer, error) {
	contract, err := bindVRFV2PlusWrapperLoadTestConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusWrapperLoadTestConsumerFilterer{contract: contract}, nil
}

func bindVRFV2PlusWrapperLoadTestConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFV2PlusWrapperLoadTestConsumerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.VRFV2PlusWrapperLoadTestConsumerCaller.contract.Call(opts, result, method, params...)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.VRFV2PlusWrapperLoadTestConsumerTransactor.contract.Transfer(opts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.VRFV2PlusWrapperLoadTestConsumerTransactor.contract.Transact(opts, method, params...)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.contract.Transfer(opts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.contract.Transact(opts, method, params...)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusWrapperLoadTestConsumer.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) GetBalance() (*big.Int, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.GetBalance(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCallerSession) GetBalance() (*big.Int, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.GetBalance(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCaller) GetLinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFV2PlusWrapperLoadTestConsumer.contract.Call(opts, &out, "getLinkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) GetLinkToken() (common.Address, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.GetLinkToken(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCallerSession) GetLinkToken() (common.Address, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.GetLinkToken(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCaller) GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (GetRequestStatus,

	error) {
	var out []interface{}
	err := _VRFV2PlusWrapperLoadTestConsumer.contract.Call(opts, &out, "getRequestStatus", _requestId)

	outstruct := new(GetRequestStatus)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paid = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fulfilled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.RandomWords = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.RequestTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FulfilmentTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RequestBlockNumber = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.FulfilmentBlockNumber = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) GetRequestStatus(_requestId *big.Int) (GetRequestStatus,

	error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.GetRequestStatus(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts, _requestId)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCallerSession) GetRequestStatus(_requestId *big.Int) (GetRequestStatus,

	error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.GetRequestStatus(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts, _requestId)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCaller) IVrfV2PlusWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFV2PlusWrapperLoadTestConsumer.contract.Call(opts, &out, "i_vrfV2PlusWrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) IVrfV2PlusWrapper() (common.Address, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.IVrfV2PlusWrapper(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCallerSession) IVrfV2PlusWrapper() (common.Address, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.IVrfV2PlusWrapper(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFV2PlusWrapperLoadTestConsumer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) Owner() (common.Address, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.Owner(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCallerSession) Owner() (common.Address, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.Owner(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCaller) SAverageFulfillmentInMillions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusWrapperLoadTestConsumer.contract.Call(opts, &out, "s_averageFulfillmentInMillions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) SAverageFulfillmentInMillions() (*big.Int, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SAverageFulfillmentInMillions(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCallerSession) SAverageFulfillmentInMillions() (*big.Int, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SAverageFulfillmentInMillions(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCaller) SFastestFulfillment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusWrapperLoadTestConsumer.contract.Call(opts, &out, "s_fastestFulfillment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) SFastestFulfillment() (*big.Int, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SFastestFulfillment(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCallerSession) SFastestFulfillment() (*big.Int, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SFastestFulfillment(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCaller) SLastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusWrapperLoadTestConsumer.contract.Call(opts, &out, "s_lastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) SLastRequestId() (*big.Int, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SLastRequestId(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCallerSession) SLastRequestId() (*big.Int, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SLastRequestId(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCaller) SRequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusWrapperLoadTestConsumer.contract.Call(opts, &out, "s_requestCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) SRequestCount() (*big.Int, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SRequestCount(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCallerSession) SRequestCount() (*big.Int, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SRequestCount(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCaller) SRequests(opts *bind.CallOpts, arg0 *big.Int) (SRequests,

	error) {
	var out []interface{}
	err := _VRFV2PlusWrapperLoadTestConsumer.contract.Call(opts, &out, "s_requests", arg0)

	outstruct := new(SRequests)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paid = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fulfilled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.RequestTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FulfilmentTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RequestBlockNumber = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.FulfilmentBlockNumber = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Native = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) SRequests(arg0 *big.Int) (SRequests,

	error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SRequests(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts, arg0)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCallerSession) SRequests(arg0 *big.Int) (SRequests,

	error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SRequests(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts, arg0)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCaller) SResponseCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusWrapperLoadTestConsumer.contract.Call(opts, &out, "s_responseCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) SResponseCount() (*big.Int, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SResponseCount(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCallerSession) SResponseCount() (*big.Int, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SResponseCount(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCaller) SSlowestFulfillment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusWrapperLoadTestConsumer.contract.Call(opts, &out, "s_slowestFulfillment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) SSlowestFulfillment() (*big.Int, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SSlowestFulfillment(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerCallerSession) SSlowestFulfillment() (*big.Int, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SSlowestFulfillment(&_VRFV2PlusWrapperLoadTestConsumer.CallOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.contract.Transact(opts, "acceptOwnership")
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.AcceptOwnership(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.AcceptOwnership(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactor) MakeRequests(opts *bind.TransactOpts, _callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32, _requestCount uint16) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.contract.Transact(opts, "makeRequests", _callbackGasLimit, _requestConfirmations, _numWords, _requestCount)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) MakeRequests(_callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32, _requestCount uint16) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.MakeRequests(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts, _callbackGasLimit, _requestConfirmations, _numWords, _requestCount)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactorSession) MakeRequests(_callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32, _requestCount uint16) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.MakeRequests(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts, _callbackGasLimit, _requestConfirmations, _numWords, _requestCount)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactor) MakeRequestsNative(opts *bind.TransactOpts, _callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32, _requestCount uint16) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.contract.Transact(opts, "makeRequestsNative", _callbackGasLimit, _requestConfirmations, _numWords, _requestCount)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) MakeRequestsNative(_callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32, _requestCount uint16) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.MakeRequestsNative(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts, _callbackGasLimit, _requestConfirmations, _numWords, _requestCount)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactorSession) MakeRequestsNative(_callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32, _requestCount uint16) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.MakeRequestsNative(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts, _callbackGasLimit, _requestConfirmations, _numWords, _requestCount)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.contract.Transact(opts, "rawFulfillRandomWords", _requestId, _randomWords)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) RawFulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.RawFulfillRandomWords(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts, _requestId, _randomWords)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactorSession) RawFulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.RawFulfillRandomWords(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts, _requestId, _randomWords)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.contract.Transact(opts, "reset")
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) Reset() (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.Reset(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactorSession) Reset() (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.Reset(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactor) SetLinkToken(opts *bind.TransactOpts, _link common.Address) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.contract.Transact(opts, "setLinkToken", _link)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) SetLinkToken(_link common.Address) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SetLinkToken(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts, _link)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactorSession) SetLinkToken(_link common.Address) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.SetLinkToken(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts, _link)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.TransferOwnership(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts, to)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.TransferOwnership(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts, to)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactor) WithdrawLink(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.contract.Transact(opts, "withdrawLink", amount)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) WithdrawLink(amount *big.Int) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.WithdrawLink(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts, amount)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactorSession) WithdrawLink(amount *big.Int) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.WithdrawLink(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts, amount)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactor) WithdrawNative(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.contract.Transact(opts, "withdrawNative", amount)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) WithdrawNative(amount *big.Int) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.WithdrawNative(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts, amount)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactorSession) WithdrawNative(amount *big.Int) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.WithdrawNative(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts, amount)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.contract.RawTransact(opts, nil)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerSession) Receive() (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.Receive(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts)
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerTransactorSession) Receive() (*types.Transaction, error) {
	return _VRFV2PlusWrapperLoadTestConsumer.Contract.Receive(&_VRFV2PlusWrapperLoadTestConsumer.TransactOpts)
}

type VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequestedIterator struct {
	Event *VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequested)
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

func (it *VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2PlusWrapperLoadTestConsumer.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequestedIterator{contract: _VRFV2PlusWrapperLoadTestConsumer.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2PlusWrapperLoadTestConsumer.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequested)
				if err := _VRFV2PlusWrapperLoadTestConsumer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequested, error) {
	event := new(VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequested)
	if err := _VRFV2PlusWrapperLoadTestConsumer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFV2PlusWrapperLoadTestConsumerOwnershipTransferredIterator struct {
	Event *VRFV2PlusWrapperLoadTestConsumerOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV2PlusWrapperLoadTestConsumerOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV2PlusWrapperLoadTestConsumerOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFV2PlusWrapperLoadTestConsumerOwnershipTransferred)
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

func (it *VRFV2PlusWrapperLoadTestConsumerOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFV2PlusWrapperLoadTestConsumerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV2PlusWrapperLoadTestConsumerOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2PlusWrapperLoadTestConsumerOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2PlusWrapperLoadTestConsumer.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusWrapperLoadTestConsumerOwnershipTransferredIterator{contract: _VRFV2PlusWrapperLoadTestConsumer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFV2PlusWrapperLoadTestConsumerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV2PlusWrapperLoadTestConsumer.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV2PlusWrapperLoadTestConsumerOwnershipTransferred)
				if err := _VRFV2PlusWrapperLoadTestConsumer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerFilterer) ParseOwnershipTransferred(log types.Log) (*VRFV2PlusWrapperLoadTestConsumerOwnershipTransferred, error) {
	event := new(VRFV2PlusWrapperLoadTestConsumerOwnershipTransferred)
	if err := _VRFV2PlusWrapperLoadTestConsumer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilledIterator struct {
	Event *VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilled)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilled)
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

func (it *VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilled struct {
	RequestId   *big.Int
	RandomWords []*big.Int
	Payment     *big.Int
	Raw         types.Log
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerFilterer) FilterWrappedRequestFulfilled(opts *bind.FilterOpts) (*VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilledIterator, error) {

	logs, sub, err := _VRFV2PlusWrapperLoadTestConsumer.contract.FilterLogs(opts, "WrappedRequestFulfilled")
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilledIterator{contract: _VRFV2PlusWrapperLoadTestConsumer.contract, event: "WrappedRequestFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerFilterer) WatchWrappedRequestFulfilled(opts *bind.WatchOpts, sink chan<- *VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilled) (event.Subscription, error) {

	logs, sub, err := _VRFV2PlusWrapperLoadTestConsumer.contract.WatchLogs(opts, "WrappedRequestFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilled)
				if err := _VRFV2PlusWrapperLoadTestConsumer.contract.UnpackLog(event, "WrappedRequestFulfilled", log); err != nil {
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

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerFilterer) ParseWrappedRequestFulfilled(log types.Log) (*VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilled, error) {
	event := new(VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilled)
	if err := _VRFV2PlusWrapperLoadTestConsumer.contract.UnpackLog(event, "WrappedRequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFV2PlusWrapperLoadTestConsumerWrapperRequestMadeIterator struct {
	Event *VRFV2PlusWrapperLoadTestConsumerWrapperRequestMade

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV2PlusWrapperLoadTestConsumerWrapperRequestMadeIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV2PlusWrapperLoadTestConsumerWrapperRequestMade)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFV2PlusWrapperLoadTestConsumerWrapperRequestMade)
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

func (it *VRFV2PlusWrapperLoadTestConsumerWrapperRequestMadeIterator) Error() error {
	return it.fail
}

func (it *VRFV2PlusWrapperLoadTestConsumerWrapperRequestMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV2PlusWrapperLoadTestConsumerWrapperRequestMade struct {
	RequestId *big.Int
	Paid      *big.Int
	Raw       types.Log
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerFilterer) FilterWrapperRequestMade(opts *bind.FilterOpts, requestId []*big.Int) (*VRFV2PlusWrapperLoadTestConsumerWrapperRequestMadeIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFV2PlusWrapperLoadTestConsumer.contract.FilterLogs(opts, "WrapperRequestMade", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusWrapperLoadTestConsumerWrapperRequestMadeIterator{contract: _VRFV2PlusWrapperLoadTestConsumer.contract, event: "WrapperRequestMade", logs: logs, sub: sub}, nil
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerFilterer) WatchWrapperRequestMade(opts *bind.WatchOpts, sink chan<- *VRFV2PlusWrapperLoadTestConsumerWrapperRequestMade, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFV2PlusWrapperLoadTestConsumer.contract.WatchLogs(opts, "WrapperRequestMade", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV2PlusWrapperLoadTestConsumerWrapperRequestMade)
				if err := _VRFV2PlusWrapperLoadTestConsumer.contract.UnpackLog(event, "WrapperRequestMade", log); err != nil {
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

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumerFilterer) ParseWrapperRequestMade(log types.Log) (*VRFV2PlusWrapperLoadTestConsumerWrapperRequestMade, error) {
	event := new(VRFV2PlusWrapperLoadTestConsumerWrapperRequestMade)
	if err := _VRFV2PlusWrapperLoadTestConsumer.contract.UnpackLog(event, "WrapperRequestMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetRequestStatus struct {
	Paid                  *big.Int
	Fulfilled             bool
	RandomWords           []*big.Int
	RequestTimestamp      *big.Int
	FulfilmentTimestamp   *big.Int
	RequestBlockNumber    *big.Int
	FulfilmentBlockNumber *big.Int
}
type SRequests struct {
	Paid                  *big.Int
	Fulfilled             bool
	RequestTimestamp      *big.Int
	FulfilmentTimestamp   *big.Int
	RequestBlockNumber    *big.Int
	FulfilmentBlockNumber *big.Int
	Native                bool
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumer) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFV2PlusWrapperLoadTestConsumer.abi.Events["OwnershipTransferRequested"].ID:
		return _VRFV2PlusWrapperLoadTestConsumer.ParseOwnershipTransferRequested(log)
	case _VRFV2PlusWrapperLoadTestConsumer.abi.Events["OwnershipTransferred"].ID:
		return _VRFV2PlusWrapperLoadTestConsumer.ParseOwnershipTransferred(log)
	case _VRFV2PlusWrapperLoadTestConsumer.abi.Events["WrappedRequestFulfilled"].ID:
		return _VRFV2PlusWrapperLoadTestConsumer.ParseWrappedRequestFulfilled(log)
	case _VRFV2PlusWrapperLoadTestConsumer.abi.Events["WrapperRequestMade"].ID:
		return _VRFV2PlusWrapperLoadTestConsumer.ParseWrapperRequestMade(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VRFV2PlusWrapperLoadTestConsumerOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilled) Topic() common.Hash {
	return common.HexToHash("0x6c84e12b4c188e61f1b4727024a5cf05c025fa58467e5eedf763c0744c89da7b")
}

func (VRFV2PlusWrapperLoadTestConsumerWrapperRequestMade) Topic() common.Hash {
	return common.HexToHash("0x5f56b4c20db9f5b294cbf6f681368de4a992a27e2de2ee702dcf2cbbfa791ec4")
}

func (_VRFV2PlusWrapperLoadTestConsumer *VRFV2PlusWrapperLoadTestConsumer) Address() common.Address {
	return _VRFV2PlusWrapperLoadTestConsumer.address
}

type VRFV2PlusWrapperLoadTestConsumerInterface interface {
	GetBalance(opts *bind.CallOpts) (*big.Int, error)

	GetLinkToken(opts *bind.CallOpts) (common.Address, error)

	GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (GetRequestStatus,

		error)

	IVrfV2PlusWrapper(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SAverageFulfillmentInMillions(opts *bind.CallOpts) (*big.Int, error)

	SFastestFulfillment(opts *bind.CallOpts) (*big.Int, error)

	SLastRequestId(opts *bind.CallOpts) (*big.Int, error)

	SRequestCount(opts *bind.CallOpts) (*big.Int, error)

	SRequests(opts *bind.CallOpts, arg0 *big.Int) (SRequests,

		error)

	SResponseCount(opts *bind.CallOpts) (*big.Int, error)

	SSlowestFulfillment(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	MakeRequests(opts *bind.TransactOpts, _callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32, _requestCount uint16) (*types.Transaction, error)

	MakeRequestsNative(opts *bind.TransactOpts, _callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32, _requestCount uint16) (*types.Transaction, error)

	RawFulfillRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error)

	Reset(opts *bind.TransactOpts) (*types.Transaction, error)

	SetLinkToken(opts *bind.TransactOpts, _link common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawLink(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	WithdrawNative(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VRFV2PlusWrapperLoadTestConsumerOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV2PlusWrapperLoadTestConsumerOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFV2PlusWrapperLoadTestConsumerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VRFV2PlusWrapperLoadTestConsumerOwnershipTransferred, error)

	FilterWrappedRequestFulfilled(opts *bind.FilterOpts) (*VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilledIterator, error)

	WatchWrappedRequestFulfilled(opts *bind.WatchOpts, sink chan<- *VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilled) (event.Subscription, error)

	ParseWrappedRequestFulfilled(log types.Log) (*VRFV2PlusWrapperLoadTestConsumerWrappedRequestFulfilled, error)

	FilterWrapperRequestMade(opts *bind.FilterOpts, requestId []*big.Int) (*VRFV2PlusWrapperLoadTestConsumerWrapperRequestMadeIterator, error)

	WatchWrapperRequestMade(opts *bind.WatchOpts, sink chan<- *VRFV2PlusWrapperLoadTestConsumerWrapperRequestMade, requestId []*big.Int) (event.Subscription, error)

	ParseWrapperRequestMade(log types.Log) (*VRFV2PlusWrapperLoadTestConsumerWrapperRequestMade, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
