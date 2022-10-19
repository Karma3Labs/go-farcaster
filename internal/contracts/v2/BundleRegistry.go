// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// BundleRegistryBatchUser is an auto generated low-level Go binding around an user-defined struct.
type BundleRegistryBatchUser struct {
	To       common.Address
	Username [16]byte
}

// BundleRegistryMetaData contains all meta data concerning the BundleRegistry contract.
var BundleRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_idRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nameRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedCaller\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trustedCaller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ChangeTrustedCaller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedCaller\",\"type\":\"address\"}],\"name\":\"changeTrustedCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes16\",\"name\":\"username\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"inviter\",\"type\":\"uint256\"}],\"name\":\"partialTrustedRegister\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes16\",\"name\":\"username\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes16\",\"name\":\"username\",\"type\":\"bytes16\"}],\"internalType\":\"structBundleRegistry.BatchUser[]\",\"name\":\"users\",\"type\":\"tuple[]\"}],\"name\":\"trustedBatchRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes16\",\"name\":\"username\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"inviter\",\"type\":\"uint256\"}],\"name\":\"trustedRegister\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BundleRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use BundleRegistryMetaData.ABI instead.
var BundleRegistryABI = BundleRegistryMetaData.ABI

// BundleRegistry is an auto generated Go binding around an Ethereum contract.
type BundleRegistry struct {
	BundleRegistryCaller     // Read-only binding to the contract
	BundleRegistryTransactor // Write-only binding to the contract
	BundleRegistryFilterer   // Log filterer for contract events
}

// BundleRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BundleRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundleRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BundleRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundleRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BundleRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundleRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BundleRegistrySession struct {
	Contract     *BundleRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BundleRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BundleRegistryCallerSession struct {
	Contract *BundleRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BundleRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BundleRegistryTransactorSession struct {
	Contract     *BundleRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BundleRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BundleRegistryRaw struct {
	Contract *BundleRegistry // Generic contract binding to access the raw methods on
}

// BundleRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BundleRegistryCallerRaw struct {
	Contract *BundleRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// BundleRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BundleRegistryTransactorRaw struct {
	Contract *BundleRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBundleRegistry creates a new instance of BundleRegistry, bound to a specific deployed contract.
func NewBundleRegistry(address common.Address, backend bind.ContractBackend) (*BundleRegistry, error) {
	contract, err := bindBundleRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BundleRegistry{BundleRegistryCaller: BundleRegistryCaller{contract: contract}, BundleRegistryTransactor: BundleRegistryTransactor{contract: contract}, BundleRegistryFilterer: BundleRegistryFilterer{contract: contract}}, nil
}

// NewBundleRegistryCaller creates a new read-only instance of BundleRegistry, bound to a specific deployed contract.
func NewBundleRegistryCaller(address common.Address, caller bind.ContractCaller) (*BundleRegistryCaller, error) {
	contract, err := bindBundleRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BundleRegistryCaller{contract: contract}, nil
}

// NewBundleRegistryTransactor creates a new write-only instance of BundleRegistry, bound to a specific deployed contract.
func NewBundleRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*BundleRegistryTransactor, error) {
	contract, err := bindBundleRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BundleRegistryTransactor{contract: contract}, nil
}

// NewBundleRegistryFilterer creates a new log filterer instance of BundleRegistry, bound to a specific deployed contract.
func NewBundleRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*BundleRegistryFilterer, error) {
	contract, err := bindBundleRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BundleRegistryFilterer{contract: contract}, nil
}

// bindBundleRegistry binds a generic wrapper to an already deployed contract.
func bindBundleRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BundleRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BundleRegistry *BundleRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BundleRegistry.Contract.BundleRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BundleRegistry *BundleRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleRegistry.Contract.BundleRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BundleRegistry *BundleRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BundleRegistry.Contract.BundleRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BundleRegistry *BundleRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BundleRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BundleRegistry *BundleRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BundleRegistry *BundleRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BundleRegistry.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BundleRegistry *BundleRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BundleRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BundleRegistry *BundleRegistrySession) Owner() (common.Address, error) {
	return _BundleRegistry.Contract.Owner(&_BundleRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BundleRegistry *BundleRegistryCallerSession) Owner() (common.Address, error) {
	return _BundleRegistry.Contract.Owner(&_BundleRegistry.CallOpts)
}

// ChangeTrustedCaller is a paid mutator transaction binding the contract method 0x881b1956.
//
// Solidity: function changeTrustedCaller(address newTrustedCaller) returns()
func (_BundleRegistry *BundleRegistryTransactor) ChangeTrustedCaller(opts *bind.TransactOpts, newTrustedCaller common.Address) (*types.Transaction, error) {
	return _BundleRegistry.contract.Transact(opts, "changeTrustedCaller", newTrustedCaller)
}

// ChangeTrustedCaller is a paid mutator transaction binding the contract method 0x881b1956.
//
// Solidity: function changeTrustedCaller(address newTrustedCaller) returns()
func (_BundleRegistry *BundleRegistrySession) ChangeTrustedCaller(newTrustedCaller common.Address) (*types.Transaction, error) {
	return _BundleRegistry.Contract.ChangeTrustedCaller(&_BundleRegistry.TransactOpts, newTrustedCaller)
}

// ChangeTrustedCaller is a paid mutator transaction binding the contract method 0x881b1956.
//
// Solidity: function changeTrustedCaller(address newTrustedCaller) returns()
func (_BundleRegistry *BundleRegistryTransactorSession) ChangeTrustedCaller(newTrustedCaller common.Address) (*types.Transaction, error) {
	return _BundleRegistry.Contract.ChangeTrustedCaller(&_BundleRegistry.TransactOpts, newTrustedCaller)
}

// PartialTrustedRegister is a paid mutator transaction binding the contract method 0xb5d860c0.
//
// Solidity: function partialTrustedRegister(address to, address recovery, string url, bytes16 username, uint256 inviter) payable returns()
func (_BundleRegistry *BundleRegistryTransactor) PartialTrustedRegister(opts *bind.TransactOpts, to common.Address, recovery common.Address, url string, username [16]byte, inviter *big.Int) (*types.Transaction, error) {
	return _BundleRegistry.contract.Transact(opts, "partialTrustedRegister", to, recovery, url, username, inviter)
}

// PartialTrustedRegister is a paid mutator transaction binding the contract method 0xb5d860c0.
//
// Solidity: function partialTrustedRegister(address to, address recovery, string url, bytes16 username, uint256 inviter) payable returns()
func (_BundleRegistry *BundleRegistrySession) PartialTrustedRegister(to common.Address, recovery common.Address, url string, username [16]byte, inviter *big.Int) (*types.Transaction, error) {
	return _BundleRegistry.Contract.PartialTrustedRegister(&_BundleRegistry.TransactOpts, to, recovery, url, username, inviter)
}

// PartialTrustedRegister is a paid mutator transaction binding the contract method 0xb5d860c0.
//
// Solidity: function partialTrustedRegister(address to, address recovery, string url, bytes16 username, uint256 inviter) payable returns()
func (_BundleRegistry *BundleRegistryTransactorSession) PartialTrustedRegister(to common.Address, recovery common.Address, url string, username [16]byte, inviter *big.Int) (*types.Transaction, error) {
	return _BundleRegistry.Contract.PartialTrustedRegister(&_BundleRegistry.TransactOpts, to, recovery, url, username, inviter)
}

// Register is a paid mutator transaction binding the contract method 0xcb737fef.
//
// Solidity: function register(address to, address recovery, string url, bytes16 username, bytes32 secret) payable returns()
func (_BundleRegistry *BundleRegistryTransactor) Register(opts *bind.TransactOpts, to common.Address, recovery common.Address, url string, username [16]byte, secret [32]byte) (*types.Transaction, error) {
	return _BundleRegistry.contract.Transact(opts, "register", to, recovery, url, username, secret)
}

// Register is a paid mutator transaction binding the contract method 0xcb737fef.
//
// Solidity: function register(address to, address recovery, string url, bytes16 username, bytes32 secret) payable returns()
func (_BundleRegistry *BundleRegistrySession) Register(to common.Address, recovery common.Address, url string, username [16]byte, secret [32]byte) (*types.Transaction, error) {
	return _BundleRegistry.Contract.Register(&_BundleRegistry.TransactOpts, to, recovery, url, username, secret)
}

// Register is a paid mutator transaction binding the contract method 0xcb737fef.
//
// Solidity: function register(address to, address recovery, string url, bytes16 username, bytes32 secret) payable returns()
func (_BundleRegistry *BundleRegistryTransactorSession) Register(to common.Address, recovery common.Address, url string, username [16]byte, secret [32]byte) (*types.Transaction, error) {
	return _BundleRegistry.Contract.Register(&_BundleRegistry.TransactOpts, to, recovery, url, username, secret)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BundleRegistry *BundleRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BundleRegistry *BundleRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _BundleRegistry.Contract.RenounceOwnership(&_BundleRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BundleRegistry *BundleRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BundleRegistry.Contract.RenounceOwnership(&_BundleRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BundleRegistry *BundleRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BundleRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BundleRegistry *BundleRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BundleRegistry.Contract.TransferOwnership(&_BundleRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BundleRegistry *BundleRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BundleRegistry.Contract.TransferOwnership(&_BundleRegistry.TransactOpts, newOwner)
}

// TrustedBatchRegister is a paid mutator transaction binding the contract method 0xac566c93.
//
// Solidity: function trustedBatchRegister((address,bytes16)[] users) returns()
func (_BundleRegistry *BundleRegistryTransactor) TrustedBatchRegister(opts *bind.TransactOpts, users []BundleRegistryBatchUser) (*types.Transaction, error) {
	return _BundleRegistry.contract.Transact(opts, "trustedBatchRegister", users)
}

// TrustedBatchRegister is a paid mutator transaction binding the contract method 0xac566c93.
//
// Solidity: function trustedBatchRegister((address,bytes16)[] users) returns()
func (_BundleRegistry *BundleRegistrySession) TrustedBatchRegister(users []BundleRegistryBatchUser) (*types.Transaction, error) {
	return _BundleRegistry.Contract.TrustedBatchRegister(&_BundleRegistry.TransactOpts, users)
}

// TrustedBatchRegister is a paid mutator transaction binding the contract method 0xac566c93.
//
// Solidity: function trustedBatchRegister((address,bytes16)[] users) returns()
func (_BundleRegistry *BundleRegistryTransactorSession) TrustedBatchRegister(users []BundleRegistryBatchUser) (*types.Transaction, error) {
	return _BundleRegistry.Contract.TrustedBatchRegister(&_BundleRegistry.TransactOpts, users)
}

// TrustedRegister is a paid mutator transaction binding the contract method 0x3c4d2003.
//
// Solidity: function trustedRegister(address to, address recovery, string url, bytes16 username, uint256 inviter) payable returns()
func (_BundleRegistry *BundleRegistryTransactor) TrustedRegister(opts *bind.TransactOpts, to common.Address, recovery common.Address, url string, username [16]byte, inviter *big.Int) (*types.Transaction, error) {
	return _BundleRegistry.contract.Transact(opts, "trustedRegister", to, recovery, url, username, inviter)
}

// TrustedRegister is a paid mutator transaction binding the contract method 0x3c4d2003.
//
// Solidity: function trustedRegister(address to, address recovery, string url, bytes16 username, uint256 inviter) payable returns()
func (_BundleRegistry *BundleRegistrySession) TrustedRegister(to common.Address, recovery common.Address, url string, username [16]byte, inviter *big.Int) (*types.Transaction, error) {
	return _BundleRegistry.Contract.TrustedRegister(&_BundleRegistry.TransactOpts, to, recovery, url, username, inviter)
}

// TrustedRegister is a paid mutator transaction binding the contract method 0x3c4d2003.
//
// Solidity: function trustedRegister(address to, address recovery, string url, bytes16 username, uint256 inviter) payable returns()
func (_BundleRegistry *BundleRegistryTransactorSession) TrustedRegister(to common.Address, recovery common.Address, url string, username [16]byte, inviter *big.Int) (*types.Transaction, error) {
	return _BundleRegistry.Contract.TrustedRegister(&_BundleRegistry.TransactOpts, to, recovery, url, username, inviter)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BundleRegistry *BundleRegistryTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleRegistry.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BundleRegistry *BundleRegistrySession) Receive() (*types.Transaction, error) {
	return _BundleRegistry.Contract.Receive(&_BundleRegistry.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BundleRegistry *BundleRegistryTransactorSession) Receive() (*types.Transaction, error) {
	return _BundleRegistry.Contract.Receive(&_BundleRegistry.TransactOpts)
}

// BundleRegistryChangeTrustedCallerIterator is returned from FilterChangeTrustedCaller and is used to iterate over the raw logs and unpacked data for ChangeTrustedCaller events raised by the BundleRegistry contract.
type BundleRegistryChangeTrustedCallerIterator struct {
	Event *BundleRegistryChangeTrustedCaller // Event containing the contract specifics and raw log

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
func (it *BundleRegistryChangeTrustedCallerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BundleRegistryChangeTrustedCaller)
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
		it.Event = new(BundleRegistryChangeTrustedCaller)
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
func (it *BundleRegistryChangeTrustedCallerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BundleRegistryChangeTrustedCallerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BundleRegistryChangeTrustedCaller represents a ChangeTrustedCaller event raised by the BundleRegistry contract.
type BundleRegistryChangeTrustedCaller struct {
	TrustedCaller common.Address
	Owner         common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChangeTrustedCaller is a free log retrieval operation binding the contract event 0x21d1bc879be7b8899ffb08788ea98d9f917258db26f2ed08daf62a3a4c3f643b.
//
// Solidity: event ChangeTrustedCaller(address indexed trustedCaller, address indexed owner)
func (_BundleRegistry *BundleRegistryFilterer) FilterChangeTrustedCaller(opts *bind.FilterOpts, trustedCaller []common.Address, owner []common.Address) (*BundleRegistryChangeTrustedCallerIterator, error) {

	var trustedCallerRule []interface{}
	for _, trustedCallerItem := range trustedCaller {
		trustedCallerRule = append(trustedCallerRule, trustedCallerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BundleRegistry.contract.FilterLogs(opts, "ChangeTrustedCaller", trustedCallerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BundleRegistryChangeTrustedCallerIterator{contract: _BundleRegistry.contract, event: "ChangeTrustedCaller", logs: logs, sub: sub}, nil
}

// WatchChangeTrustedCaller is a free log subscription operation binding the contract event 0x21d1bc879be7b8899ffb08788ea98d9f917258db26f2ed08daf62a3a4c3f643b.
//
// Solidity: event ChangeTrustedCaller(address indexed trustedCaller, address indexed owner)
func (_BundleRegistry *BundleRegistryFilterer) WatchChangeTrustedCaller(opts *bind.WatchOpts, sink chan<- *BundleRegistryChangeTrustedCaller, trustedCaller []common.Address, owner []common.Address) (event.Subscription, error) {

	var trustedCallerRule []interface{}
	for _, trustedCallerItem := range trustedCaller {
		trustedCallerRule = append(trustedCallerRule, trustedCallerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BundleRegistry.contract.WatchLogs(opts, "ChangeTrustedCaller", trustedCallerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BundleRegistryChangeTrustedCaller)
				if err := _BundleRegistry.contract.UnpackLog(event, "ChangeTrustedCaller", log); err != nil {
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

// ParseChangeTrustedCaller is a log parse operation binding the contract event 0x21d1bc879be7b8899ffb08788ea98d9f917258db26f2ed08daf62a3a4c3f643b.
//
// Solidity: event ChangeTrustedCaller(address indexed trustedCaller, address indexed owner)
func (_BundleRegistry *BundleRegistryFilterer) ParseChangeTrustedCaller(log types.Log) (*BundleRegistryChangeTrustedCaller, error) {
	event := new(BundleRegistryChangeTrustedCaller)
	if err := _BundleRegistry.contract.UnpackLog(event, "ChangeTrustedCaller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BundleRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BundleRegistry contract.
type BundleRegistryOwnershipTransferredIterator struct {
	Event *BundleRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BundleRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BundleRegistryOwnershipTransferred)
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
		it.Event = new(BundleRegistryOwnershipTransferred)
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
func (it *BundleRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BundleRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BundleRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the BundleRegistry contract.
type BundleRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BundleRegistry *BundleRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BundleRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BundleRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BundleRegistryOwnershipTransferredIterator{contract: _BundleRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BundleRegistry *BundleRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BundleRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BundleRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BundleRegistryOwnershipTransferred)
				if err := _BundleRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BundleRegistry *BundleRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*BundleRegistryOwnershipTransferred, error) {
	event := new(BundleRegistryOwnershipTransferred)
	if err := _BundleRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
