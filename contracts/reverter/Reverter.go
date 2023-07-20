// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reverter

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

// ReverterMetaData contains all meta data concerning the Reverter contract.
var ReverterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"getEnableReceive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"setEnableReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060015f819055506103a0806100245f395ff3fe60806040526004361061002c575f3560e01c80633ec60ee3146100dd57806384e2329c1461010557610087565b366100875761003961012f565b5f6001540361007d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610074906101f2565b60405180910390fd5b61008561017c565b005b61008f61012f565b5f600154036100d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100ca9061025a565b60405180910390fd5b6100db61017c565b005b3480156100e8575f80fd5b5061010360048036038101906100fe91906102af565b610185565b005b348015610110575f80fd5b5061011961018f565b60405161012691906102e9565b60405180910390f35b60025f5403610173576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016a9061034c565b60405180910390fd5b60025f81905550565b60015f81905550565b8060018190555050565b5f600154905090565b5f82825260208201905092915050565b7f62656361757365207265636569766500000000000000000000000000000000005f82015250565b5f6101dc600f83610198565b91506101e7826101a8565b602082019050919050565b5f6020820190508181035f830152610209816101d0565b9050919050565b7f626563617573652066616c6c6261636b000000000000000000000000000000005f82015250565b5f610244601083610198565b915061024f82610210565b602082019050919050565b5f6020820190508181035f83015261027181610238565b9050919050565b5f80fd5b5f819050919050565b61028e8161027c565b8114610298575f80fd5b50565b5f813590506102a981610285565b92915050565b5f602082840312156102c4576102c3610278565b5b5f6102d18482850161029b565b91505092915050565b6102e38161027c565b82525050565b5f6020820190506102fc5f8301846102da565b92915050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c005f82015250565b5f610336601f83610198565b915061034182610302565b602082019050919050565b5f6020820190508181035f8301526103638161032a565b905091905056fea26469706673582212202095edeee4a9c8260284f5e52634ec27b6bfee887d248072faf62528e4eb6f5f64736f6c63430008150033",
}

// ReverterABI is the input ABI used to generate the binding from.
// Deprecated: Use ReverterMetaData.ABI instead.
var ReverterABI = ReverterMetaData.ABI

// ReverterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReverterMetaData.Bin instead.
var ReverterBin = ReverterMetaData.Bin

// DeployReverter deploys a new Ethereum contract, binding an instance of Reverter to it.
func DeployReverter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Reverter, error) {
	parsed, err := ReverterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReverterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Reverter{ReverterCaller: ReverterCaller{contract: contract}, ReverterTransactor: ReverterTransactor{contract: contract}, ReverterFilterer: ReverterFilterer{contract: contract}}, nil
}

// Reverter is an auto generated Go binding around an Ethereum contract.
type Reverter struct {
	ReverterCaller     // Read-only binding to the contract
	ReverterTransactor // Write-only binding to the contract
	ReverterFilterer   // Log filterer for contract events
}

// ReverterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReverterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReverterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReverterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReverterSession struct {
	Contract     *Reverter         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReverterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReverterCallerSession struct {
	Contract *ReverterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ReverterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReverterTransactorSession struct {
	Contract     *ReverterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ReverterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReverterRaw struct {
	Contract *Reverter // Generic contract binding to access the raw methods on
}

// ReverterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReverterCallerRaw struct {
	Contract *ReverterCaller // Generic read-only contract binding to access the raw methods on
}

// ReverterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReverterTransactorRaw struct {
	Contract *ReverterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReverter creates a new instance of Reverter, bound to a specific deployed contract.
func NewReverter(address common.Address, backend bind.ContractBackend) (*Reverter, error) {
	contract, err := bindReverter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reverter{ReverterCaller: ReverterCaller{contract: contract}, ReverterTransactor: ReverterTransactor{contract: contract}, ReverterFilterer: ReverterFilterer{contract: contract}}, nil
}

// NewReverterCaller creates a new read-only instance of Reverter, bound to a specific deployed contract.
func NewReverterCaller(address common.Address, caller bind.ContractCaller) (*ReverterCaller, error) {
	contract, err := bindReverter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReverterCaller{contract: contract}, nil
}

// NewReverterTransactor creates a new write-only instance of Reverter, bound to a specific deployed contract.
func NewReverterTransactor(address common.Address, transactor bind.ContractTransactor) (*ReverterTransactor, error) {
	contract, err := bindReverter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReverterTransactor{contract: contract}, nil
}

// NewReverterFilterer creates a new log filterer instance of Reverter, bound to a specific deployed contract.
func NewReverterFilterer(address common.Address, filterer bind.ContractFilterer) (*ReverterFilterer, error) {
	contract, err := bindReverter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReverterFilterer{contract: contract}, nil
}

// bindReverter binds a generic wrapper to an already deployed contract.
func bindReverter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReverterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reverter *ReverterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reverter.Contract.ReverterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reverter *ReverterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reverter.Contract.ReverterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reverter *ReverterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reverter.Contract.ReverterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reverter *ReverterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reverter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reverter *ReverterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reverter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reverter *ReverterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reverter.Contract.contract.Transact(opts, method, params...)
}

// GetEnableReceive is a free data retrieval call binding the contract method 0x84e2329c.
//
// Solidity: function getEnableReceive() view returns(uint256)
func (_Reverter *ReverterCaller) GetEnableReceive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reverter.contract.Call(opts, &out, "getEnableReceive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEnableReceive is a free data retrieval call binding the contract method 0x84e2329c.
//
// Solidity: function getEnableReceive() view returns(uint256)
func (_Reverter *ReverterSession) GetEnableReceive() (*big.Int, error) {
	return _Reverter.Contract.GetEnableReceive(&_Reverter.CallOpts)
}

// GetEnableReceive is a free data retrieval call binding the contract method 0x84e2329c.
//
// Solidity: function getEnableReceive() view returns(uint256)
func (_Reverter *ReverterCallerSession) GetEnableReceive() (*big.Int, error) {
	return _Reverter.Contract.GetEnableReceive(&_Reverter.CallOpts)
}

// SetEnableReceive is a paid mutator transaction binding the contract method 0x3ec60ee3.
//
// Solidity: function setEnableReceive(uint256 v) returns()
func (_Reverter *ReverterTransactor) SetEnableReceive(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _Reverter.contract.Transact(opts, "setEnableReceive", v)
}

// SetEnableReceive is a paid mutator transaction binding the contract method 0x3ec60ee3.
//
// Solidity: function setEnableReceive(uint256 v) returns()
func (_Reverter *ReverterSession) SetEnableReceive(v *big.Int) (*types.Transaction, error) {
	return _Reverter.Contract.SetEnableReceive(&_Reverter.TransactOpts, v)
}

// SetEnableReceive is a paid mutator transaction binding the contract method 0x3ec60ee3.
//
// Solidity: function setEnableReceive(uint256 v) returns()
func (_Reverter *ReverterTransactorSession) SetEnableReceive(v *big.Int) (*types.Transaction, error) {
	return _Reverter.Contract.SetEnableReceive(&_Reverter.TransactOpts, v)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Reverter *ReverterTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Reverter.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Reverter *ReverterSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Reverter.Contract.Fallback(&_Reverter.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Reverter *ReverterTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Reverter.Contract.Fallback(&_Reverter.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Reverter *ReverterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reverter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Reverter *ReverterSession) Receive() (*types.Transaction, error) {
	return _Reverter.Contract.Receive(&_Reverter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Reverter *ReverterTransactorSession) Receive() (*types.Transaction, error) {
	return _Reverter.Contract.Receive(&_Reverter.TransactOpts)
}
