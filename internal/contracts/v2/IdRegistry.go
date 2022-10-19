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

// IdRegistryMetaData contains all meta data concerning the IdRegistry contract.
var IdRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Escrow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HasId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HasNoId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Invitable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoRecovery\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Registrable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint257\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CancelRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"ChangeHome\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"ChangeRecoveryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trustedCaller\",\"type\":\"address\"}],\"name\":\"ChangeTrustedCaller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableTrustedOnly\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"RequestRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"cancelRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"changeHome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"changeRecoveryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustedCaller\",\"type\":\"address\"}],\"name\":\"changeTrustedCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"completeRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"completeTransferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableTrustedOnly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"idOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"requestRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"requestTransferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"trustedRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IdRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IdRegistryMetaData.ABI instead.
var IdRegistryABI = IdRegistryMetaData.ABI

// IdRegistry is an auto generated Go binding around an Ethereum contract.
type IdRegistry struct {
	IdRegistryCaller     // Read-only binding to the contract
	IdRegistryTransactor // Write-only binding to the contract
	IdRegistryFilterer   // Log filterer for contract events
}

// IdRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdRegistrySession struct {
	Contract     *IdRegistry       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdRegistryCallerSession struct {
	Contract *IdRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IdRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdRegistryTransactorSession struct {
	Contract     *IdRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IdRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdRegistryRaw struct {
	Contract *IdRegistry // Generic contract binding to access the raw methods on
}

// IdRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdRegistryCallerRaw struct {
	Contract *IdRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IdRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdRegistryTransactorRaw struct {
	Contract *IdRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdRegistry creates a new instance of IdRegistry, bound to a specific deployed contract.
func NewIdRegistry(address common.Address, backend bind.ContractBackend) (*IdRegistry, error) {
	contract, err := bindIdRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdRegistry{IdRegistryCaller: IdRegistryCaller{contract: contract}, IdRegistryTransactor: IdRegistryTransactor{contract: contract}, IdRegistryFilterer: IdRegistryFilterer{contract: contract}}, nil
}

// NewIdRegistryCaller creates a new read-only instance of IdRegistry, bound to a specific deployed contract.
func NewIdRegistryCaller(address common.Address, caller bind.ContractCaller) (*IdRegistryCaller, error) {
	contract, err := bindIdRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdRegistryCaller{contract: contract}, nil
}

// NewIdRegistryTransactor creates a new write-only instance of IdRegistry, bound to a specific deployed contract.
func NewIdRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdRegistryTransactor, error) {
	contract, err := bindIdRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdRegistryTransactor{contract: contract}, nil
}

// NewIdRegistryFilterer creates a new log filterer instance of IdRegistry, bound to a specific deployed contract.
func NewIdRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdRegistryFilterer, error) {
	contract, err := bindIdRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdRegistryFilterer{contract: contract}, nil
}

// bindIdRegistry binds a generic wrapper to an already deployed contract.
func bindIdRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdRegistry *IdRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdRegistry.Contract.IdRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdRegistry *IdRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdRegistry.Contract.IdRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdRegistry *IdRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdRegistry.Contract.IdRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdRegistry *IdRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdRegistry *IdRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdRegistry *IdRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdRegistry.Contract.contract.Transact(opts, method, params...)
}

// IdOf is a free data retrieval call binding the contract method 0xd94fe832.
//
// Solidity: function idOf(address ) view returns(uint256)
func (_IdRegistry *IdRegistryCaller) IdOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "idOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdOf is a free data retrieval call binding the contract method 0xd94fe832.
//
// Solidity: function idOf(address ) view returns(uint256)
func (_IdRegistry *IdRegistrySession) IdOf(arg0 common.Address) (*big.Int, error) {
	return _IdRegistry.Contract.IdOf(&_IdRegistry.CallOpts, arg0)
}

// IdOf is a free data retrieval call binding the contract method 0xd94fe832.
//
// Solidity: function idOf(address ) view returns(uint256)
func (_IdRegistry *IdRegistryCallerSession) IdOf(arg0 common.Address) (*big.Int, error) {
	return _IdRegistry.Contract.IdOf(&_IdRegistry.CallOpts, arg0)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_IdRegistry *IdRegistryCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_IdRegistry *IdRegistrySession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _IdRegistry.Contract.IsTrustedForwarder(&_IdRegistry.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_IdRegistry *IdRegistryCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _IdRegistry.Contract.IsTrustedForwarder(&_IdRegistry.CallOpts, forwarder)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdRegistry *IdRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdRegistry *IdRegistrySession) Owner() (common.Address, error) {
	return _IdRegistry.Contract.Owner(&_IdRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdRegistry *IdRegistryCallerSession) Owner() (common.Address, error) {
	return _IdRegistry.Contract.Owner(&_IdRegistry.CallOpts)
}

// TransferOwnership is a free data retrieval call binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address ) view returns()
func (_IdRegistry *IdRegistryCaller) TransferOwnership(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _IdRegistry.contract.Call(opts, &out, "transferOwnership", arg0)

	if err != nil {
		return err
	}

	return err

}

// TransferOwnership is a free data retrieval call binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address ) view returns()
func (_IdRegistry *IdRegistrySession) TransferOwnership(arg0 common.Address) error {
	return _IdRegistry.Contract.TransferOwnership(&_IdRegistry.CallOpts, arg0)
}

// TransferOwnership is a free data retrieval call binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address ) view returns()
func (_IdRegistry *IdRegistryCallerSession) TransferOwnership(arg0 common.Address) error {
	return _IdRegistry.Contract.TransferOwnership(&_IdRegistry.CallOpts, arg0)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0xc90db447.
//
// Solidity: function cancelRecovery(address from) returns()
func (_IdRegistry *IdRegistryTransactor) CancelRecovery(opts *bind.TransactOpts, from common.Address) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "cancelRecovery", from)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0xc90db447.
//
// Solidity: function cancelRecovery(address from) returns()
func (_IdRegistry *IdRegistrySession) CancelRecovery(from common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.CancelRecovery(&_IdRegistry.TransactOpts, from)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0xc90db447.
//
// Solidity: function cancelRecovery(address from) returns()
func (_IdRegistry *IdRegistryTransactorSession) CancelRecovery(from common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.CancelRecovery(&_IdRegistry.TransactOpts, from)
}

// ChangeHome is a paid mutator transaction binding the contract method 0xaa217f25.
//
// Solidity: function changeHome(string url) returns()
func (_IdRegistry *IdRegistryTransactor) ChangeHome(opts *bind.TransactOpts, url string) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "changeHome", url)
}

// ChangeHome is a paid mutator transaction binding the contract method 0xaa217f25.
//
// Solidity: function changeHome(string url) returns()
func (_IdRegistry *IdRegistrySession) ChangeHome(url string) (*types.Transaction, error) {
	return _IdRegistry.Contract.ChangeHome(&_IdRegistry.TransactOpts, url)
}

// ChangeHome is a paid mutator transaction binding the contract method 0xaa217f25.
//
// Solidity: function changeHome(string url) returns()
func (_IdRegistry *IdRegistryTransactorSession) ChangeHome(url string) (*types.Transaction, error) {
	return _IdRegistry.Contract.ChangeHome(&_IdRegistry.TransactOpts, url)
}

// ChangeRecoveryAddress is a paid mutator transaction binding the contract method 0xf1f0b224.
//
// Solidity: function changeRecoveryAddress(address recovery) returns()
func (_IdRegistry *IdRegistryTransactor) ChangeRecoveryAddress(opts *bind.TransactOpts, recovery common.Address) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "changeRecoveryAddress", recovery)
}

// ChangeRecoveryAddress is a paid mutator transaction binding the contract method 0xf1f0b224.
//
// Solidity: function changeRecoveryAddress(address recovery) returns()
func (_IdRegistry *IdRegistrySession) ChangeRecoveryAddress(recovery common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.ChangeRecoveryAddress(&_IdRegistry.TransactOpts, recovery)
}

// ChangeRecoveryAddress is a paid mutator transaction binding the contract method 0xf1f0b224.
//
// Solidity: function changeRecoveryAddress(address recovery) returns()
func (_IdRegistry *IdRegistryTransactorSession) ChangeRecoveryAddress(recovery common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.ChangeRecoveryAddress(&_IdRegistry.TransactOpts, recovery)
}

// ChangeTrustedCaller is a paid mutator transaction binding the contract method 0x881b1956.
//
// Solidity: function changeTrustedCaller(address _trustedCaller) returns()
func (_IdRegistry *IdRegistryTransactor) ChangeTrustedCaller(opts *bind.TransactOpts, _trustedCaller common.Address) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "changeTrustedCaller", _trustedCaller)
}

// ChangeTrustedCaller is a paid mutator transaction binding the contract method 0x881b1956.
//
// Solidity: function changeTrustedCaller(address _trustedCaller) returns()
func (_IdRegistry *IdRegistrySession) ChangeTrustedCaller(_trustedCaller common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.ChangeTrustedCaller(&_IdRegistry.TransactOpts, _trustedCaller)
}

// ChangeTrustedCaller is a paid mutator transaction binding the contract method 0x881b1956.
//
// Solidity: function changeTrustedCaller(address _trustedCaller) returns()
func (_IdRegistry *IdRegistryTransactorSession) ChangeTrustedCaller(_trustedCaller common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.ChangeTrustedCaller(&_IdRegistry.TransactOpts, _trustedCaller)
}

// CompleteRecovery is a paid mutator transaction binding the contract method 0xf480dd7e.
//
// Solidity: function completeRecovery(address from) returns()
func (_IdRegistry *IdRegistryTransactor) CompleteRecovery(opts *bind.TransactOpts, from common.Address) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "completeRecovery", from)
}

// CompleteRecovery is a paid mutator transaction binding the contract method 0xf480dd7e.
//
// Solidity: function completeRecovery(address from) returns()
func (_IdRegistry *IdRegistrySession) CompleteRecovery(from common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.CompleteRecovery(&_IdRegistry.TransactOpts, from)
}

// CompleteRecovery is a paid mutator transaction binding the contract method 0xf480dd7e.
//
// Solidity: function completeRecovery(address from) returns()
func (_IdRegistry *IdRegistryTransactorSession) CompleteRecovery(from common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.CompleteRecovery(&_IdRegistry.TransactOpts, from)
}

// CompleteTransferOwnership is a paid mutator transaction binding the contract method 0x53f0447e.
//
// Solidity: function completeTransferOwnership() returns()
func (_IdRegistry *IdRegistryTransactor) CompleteTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "completeTransferOwnership")
}

// CompleteTransferOwnership is a paid mutator transaction binding the contract method 0x53f0447e.
//
// Solidity: function completeTransferOwnership() returns()
func (_IdRegistry *IdRegistrySession) CompleteTransferOwnership() (*types.Transaction, error) {
	return _IdRegistry.Contract.CompleteTransferOwnership(&_IdRegistry.TransactOpts)
}

// CompleteTransferOwnership is a paid mutator transaction binding the contract method 0x53f0447e.
//
// Solidity: function completeTransferOwnership() returns()
func (_IdRegistry *IdRegistryTransactorSession) CompleteTransferOwnership() (*types.Transaction, error) {
	return _IdRegistry.Contract.CompleteTransferOwnership(&_IdRegistry.TransactOpts)
}

// DisableTrustedOnly is a paid mutator transaction binding the contract method 0x6e9bde49.
//
// Solidity: function disableTrustedOnly() returns()
func (_IdRegistry *IdRegistryTransactor) DisableTrustedOnly(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "disableTrustedOnly")
}

// DisableTrustedOnly is a paid mutator transaction binding the contract method 0x6e9bde49.
//
// Solidity: function disableTrustedOnly() returns()
func (_IdRegistry *IdRegistrySession) DisableTrustedOnly() (*types.Transaction, error) {
	return _IdRegistry.Contract.DisableTrustedOnly(&_IdRegistry.TransactOpts)
}

// DisableTrustedOnly is a paid mutator transaction binding the contract method 0x6e9bde49.
//
// Solidity: function disableTrustedOnly() returns()
func (_IdRegistry *IdRegistryTransactorSession) DisableTrustedOnly() (*types.Transaction, error) {
	return _IdRegistry.Contract.DisableTrustedOnly(&_IdRegistry.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0xede49739.
//
// Solidity: function register(address to, address recovery, string url) returns()
func (_IdRegistry *IdRegistryTransactor) Register(opts *bind.TransactOpts, to common.Address, recovery common.Address, url string) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "register", to, recovery, url)
}

// Register is a paid mutator transaction binding the contract method 0xede49739.
//
// Solidity: function register(address to, address recovery, string url) returns()
func (_IdRegistry *IdRegistrySession) Register(to common.Address, recovery common.Address, url string) (*types.Transaction, error) {
	return _IdRegistry.Contract.Register(&_IdRegistry.TransactOpts, to, recovery, url)
}

// Register is a paid mutator transaction binding the contract method 0xede49739.
//
// Solidity: function register(address to, address recovery, string url) returns()
func (_IdRegistry *IdRegistryTransactorSession) Register(to common.Address, recovery common.Address, url string) (*types.Transaction, error) {
	return _IdRegistry.Contract.Register(&_IdRegistry.TransactOpts, to, recovery, url)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdRegistry *IdRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdRegistry *IdRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _IdRegistry.Contract.RenounceOwnership(&_IdRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdRegistry *IdRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IdRegistry.Contract.RenounceOwnership(&_IdRegistry.TransactOpts)
}

// RequestRecovery is a paid mutator transaction binding the contract method 0x36bacd73.
//
// Solidity: function requestRecovery(address from, address to) returns()
func (_IdRegistry *IdRegistryTransactor) RequestRecovery(opts *bind.TransactOpts, from common.Address, to common.Address) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "requestRecovery", from, to)
}

// RequestRecovery is a paid mutator transaction binding the contract method 0x36bacd73.
//
// Solidity: function requestRecovery(address from, address to) returns()
func (_IdRegistry *IdRegistrySession) RequestRecovery(from common.Address, to common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.RequestRecovery(&_IdRegistry.TransactOpts, from, to)
}

// RequestRecovery is a paid mutator transaction binding the contract method 0x36bacd73.
//
// Solidity: function requestRecovery(address from, address to) returns()
func (_IdRegistry *IdRegistryTransactorSession) RequestRecovery(from common.Address, to common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.RequestRecovery(&_IdRegistry.TransactOpts, from, to)
}

// RequestTransferOwnership is a paid mutator transaction binding the contract method 0x9d6fa618.
//
// Solidity: function requestTransferOwnership(address newOwner) returns()
func (_IdRegistry *IdRegistryTransactor) RequestTransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "requestTransferOwnership", newOwner)
}

// RequestTransferOwnership is a paid mutator transaction binding the contract method 0x9d6fa618.
//
// Solidity: function requestTransferOwnership(address newOwner) returns()
func (_IdRegistry *IdRegistrySession) RequestTransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.RequestTransferOwnership(&_IdRegistry.TransactOpts, newOwner)
}

// RequestTransferOwnership is a paid mutator transaction binding the contract method 0x9d6fa618.
//
// Solidity: function requestTransferOwnership(address newOwner) returns()
func (_IdRegistry *IdRegistryTransactorSession) RequestTransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.RequestTransferOwnership(&_IdRegistry.TransactOpts, newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address to) returns()
func (_IdRegistry *IdRegistryTransactor) Transfer(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "transfer", to)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address to) returns()
func (_IdRegistry *IdRegistrySession) Transfer(to common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.Transfer(&_IdRegistry.TransactOpts, to)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address to) returns()
func (_IdRegistry *IdRegistryTransactorSession) Transfer(to common.Address) (*types.Transaction, error) {
	return _IdRegistry.Contract.Transfer(&_IdRegistry.TransactOpts, to)
}

// TrustedRegister is a paid mutator transaction binding the contract method 0x052a30a3.
//
// Solidity: function trustedRegister(address to, address recovery, string url) returns()
func (_IdRegistry *IdRegistryTransactor) TrustedRegister(opts *bind.TransactOpts, to common.Address, recovery common.Address, url string) (*types.Transaction, error) {
	return _IdRegistry.contract.Transact(opts, "trustedRegister", to, recovery, url)
}

// TrustedRegister is a paid mutator transaction binding the contract method 0x052a30a3.
//
// Solidity: function trustedRegister(address to, address recovery, string url) returns()
func (_IdRegistry *IdRegistrySession) TrustedRegister(to common.Address, recovery common.Address, url string) (*types.Transaction, error) {
	return _IdRegistry.Contract.TrustedRegister(&_IdRegistry.TransactOpts, to, recovery, url)
}

// TrustedRegister is a paid mutator transaction binding the contract method 0x052a30a3.
//
// Solidity: function trustedRegister(address to, address recovery, string url) returns()
func (_IdRegistry *IdRegistryTransactorSession) TrustedRegister(to common.Address, recovery common.Address, url string) (*types.Transaction, error) {
	return _IdRegistry.Contract.TrustedRegister(&_IdRegistry.TransactOpts, to, recovery, url)
}

// IdRegistryCancelRecoveryIterator is returned from FilterCancelRecovery and is used to iterate over the raw logs and unpacked data for CancelRecovery events raised by the IdRegistry contract.
type IdRegistryCancelRecoveryIterator struct {
	Event *IdRegistryCancelRecovery // Event containing the contract specifics and raw log

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
func (it *IdRegistryCancelRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryCancelRecovery)
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
		it.Event = new(IdRegistryCancelRecovery)
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
func (it *IdRegistryCancelRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryCancelRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryCancelRecovery represents a CancelRecovery event raised by the IdRegistry contract.
type IdRegistryCancelRecovery struct {
	By  common.Address
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancelRecovery is a free log retrieval operation binding the contract event 0x6181d4215ebc71e962cc193554c17f05a825da06230fdf9ece45081f09cb206f.
//
// Solidity: event CancelRecovery(address indexed by, uint256 indexed id)
func (_IdRegistry *IdRegistryFilterer) FilterCancelRecovery(opts *bind.FilterOpts, by []common.Address, id []*big.Int) (*IdRegistryCancelRecoveryIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "CancelRecovery", byRule, idRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryCancelRecoveryIterator{contract: _IdRegistry.contract, event: "CancelRecovery", logs: logs, sub: sub}, nil
}

// WatchCancelRecovery is a free log subscription operation binding the contract event 0x6181d4215ebc71e962cc193554c17f05a825da06230fdf9ece45081f09cb206f.
//
// Solidity: event CancelRecovery(address indexed by, uint256 indexed id)
func (_IdRegistry *IdRegistryFilterer) WatchCancelRecovery(opts *bind.WatchOpts, sink chan<- *IdRegistryCancelRecovery, by []common.Address, id []*big.Int) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "CancelRecovery", byRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryCancelRecovery)
				if err := _IdRegistry.contract.UnpackLog(event, "CancelRecovery", log); err != nil {
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

// ParseCancelRecovery is a log parse operation binding the contract event 0x6181d4215ebc71e962cc193554c17f05a825da06230fdf9ece45081f09cb206f.
//
// Solidity: event CancelRecovery(address indexed by, uint256 indexed id)
func (_IdRegistry *IdRegistryFilterer) ParseCancelRecovery(log types.Log) (*IdRegistryCancelRecovery, error) {
	event := new(IdRegistryCancelRecovery)
	if err := _IdRegistry.contract.UnpackLog(event, "CancelRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryChangeHomeIterator is returned from FilterChangeHome and is used to iterate over the raw logs and unpacked data for ChangeHome events raised by the IdRegistry contract.
type IdRegistryChangeHomeIterator struct {
	Event *IdRegistryChangeHome // Event containing the contract specifics and raw log

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
func (it *IdRegistryChangeHomeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryChangeHome)
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
		it.Event = new(IdRegistryChangeHome)
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
func (it *IdRegistryChangeHomeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryChangeHomeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryChangeHome represents a ChangeHome event raised by the IdRegistry contract.
type IdRegistryChangeHome struct {
	Id  *big.Int
	Url string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChangeHome is a free log retrieval operation binding the contract event 0x5d926244a310a9e23d7caac05945165ecf7ef6f4a47fae87eb5e8a005629fdb0.
//
// Solidity: event ChangeHome(uint256 indexed id, string url)
func (_IdRegistry *IdRegistryFilterer) FilterChangeHome(opts *bind.FilterOpts, id []*big.Int) (*IdRegistryChangeHomeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "ChangeHome", idRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryChangeHomeIterator{contract: _IdRegistry.contract, event: "ChangeHome", logs: logs, sub: sub}, nil
}

// WatchChangeHome is a free log subscription operation binding the contract event 0x5d926244a310a9e23d7caac05945165ecf7ef6f4a47fae87eb5e8a005629fdb0.
//
// Solidity: event ChangeHome(uint256 indexed id, string url)
func (_IdRegistry *IdRegistryFilterer) WatchChangeHome(opts *bind.WatchOpts, sink chan<- *IdRegistryChangeHome, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "ChangeHome", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryChangeHome)
				if err := _IdRegistry.contract.UnpackLog(event, "ChangeHome", log); err != nil {
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

// ParseChangeHome is a log parse operation binding the contract event 0x5d926244a310a9e23d7caac05945165ecf7ef6f4a47fae87eb5e8a005629fdb0.
//
// Solidity: event ChangeHome(uint256 indexed id, string url)
func (_IdRegistry *IdRegistryFilterer) ParseChangeHome(log types.Log) (*IdRegistryChangeHome, error) {
	event := new(IdRegistryChangeHome)
	if err := _IdRegistry.contract.UnpackLog(event, "ChangeHome", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryChangeRecoveryAddressIterator is returned from FilterChangeRecoveryAddress and is used to iterate over the raw logs and unpacked data for ChangeRecoveryAddress events raised by the IdRegistry contract.
type IdRegistryChangeRecoveryAddressIterator struct {
	Event *IdRegistryChangeRecoveryAddress // Event containing the contract specifics and raw log

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
func (it *IdRegistryChangeRecoveryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryChangeRecoveryAddress)
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
		it.Event = new(IdRegistryChangeRecoveryAddress)
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
func (it *IdRegistryChangeRecoveryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryChangeRecoveryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryChangeRecoveryAddress represents a ChangeRecoveryAddress event raised by the IdRegistry contract.
type IdRegistryChangeRecoveryAddress struct {
	Id       *big.Int
	Recovery common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChangeRecoveryAddress is a free log retrieval operation binding the contract event 0x8e700b803af43e14651431cd73c9fe7d11b131ad797576a70b893ce5766f65c3.
//
// Solidity: event ChangeRecoveryAddress(uint256 indexed id, address indexed recovery)
func (_IdRegistry *IdRegistryFilterer) FilterChangeRecoveryAddress(opts *bind.FilterOpts, id []*big.Int, recovery []common.Address) (*IdRegistryChangeRecoveryAddressIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var recoveryRule []interface{}
	for _, recoveryItem := range recovery {
		recoveryRule = append(recoveryRule, recoveryItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "ChangeRecoveryAddress", idRule, recoveryRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryChangeRecoveryAddressIterator{contract: _IdRegistry.contract, event: "ChangeRecoveryAddress", logs: logs, sub: sub}, nil
}

// WatchChangeRecoveryAddress is a free log subscription operation binding the contract event 0x8e700b803af43e14651431cd73c9fe7d11b131ad797576a70b893ce5766f65c3.
//
// Solidity: event ChangeRecoveryAddress(uint256 indexed id, address indexed recovery)
func (_IdRegistry *IdRegistryFilterer) WatchChangeRecoveryAddress(opts *bind.WatchOpts, sink chan<- *IdRegistryChangeRecoveryAddress, id []*big.Int, recovery []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var recoveryRule []interface{}
	for _, recoveryItem := range recovery {
		recoveryRule = append(recoveryRule, recoveryItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "ChangeRecoveryAddress", idRule, recoveryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryChangeRecoveryAddress)
				if err := _IdRegistry.contract.UnpackLog(event, "ChangeRecoveryAddress", log); err != nil {
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

// ParseChangeRecoveryAddress is a log parse operation binding the contract event 0x8e700b803af43e14651431cd73c9fe7d11b131ad797576a70b893ce5766f65c3.
//
// Solidity: event ChangeRecoveryAddress(uint256 indexed id, address indexed recovery)
func (_IdRegistry *IdRegistryFilterer) ParseChangeRecoveryAddress(log types.Log) (*IdRegistryChangeRecoveryAddress, error) {
	event := new(IdRegistryChangeRecoveryAddress)
	if err := _IdRegistry.contract.UnpackLog(event, "ChangeRecoveryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryChangeTrustedCallerIterator is returned from FilterChangeTrustedCaller and is used to iterate over the raw logs and unpacked data for ChangeTrustedCaller events raised by the IdRegistry contract.
type IdRegistryChangeTrustedCallerIterator struct {
	Event *IdRegistryChangeTrustedCaller // Event containing the contract specifics and raw log

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
func (it *IdRegistryChangeTrustedCallerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryChangeTrustedCaller)
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
		it.Event = new(IdRegistryChangeTrustedCaller)
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
func (it *IdRegistryChangeTrustedCallerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryChangeTrustedCallerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryChangeTrustedCaller represents a ChangeTrustedCaller event raised by the IdRegistry contract.
type IdRegistryChangeTrustedCaller struct {
	TrustedCaller common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChangeTrustedCaller is a free log retrieval operation binding the contract event 0x255ba3357fefe42b361c216b6e0bc5541f1e6ea4c6178d4a45ad8dd7ec28139d.
//
// Solidity: event ChangeTrustedCaller(address indexed trustedCaller)
func (_IdRegistry *IdRegistryFilterer) FilterChangeTrustedCaller(opts *bind.FilterOpts, trustedCaller []common.Address) (*IdRegistryChangeTrustedCallerIterator, error) {

	var trustedCallerRule []interface{}
	for _, trustedCallerItem := range trustedCaller {
		trustedCallerRule = append(trustedCallerRule, trustedCallerItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "ChangeTrustedCaller", trustedCallerRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryChangeTrustedCallerIterator{contract: _IdRegistry.contract, event: "ChangeTrustedCaller", logs: logs, sub: sub}, nil
}

// WatchChangeTrustedCaller is a free log subscription operation binding the contract event 0x255ba3357fefe42b361c216b6e0bc5541f1e6ea4c6178d4a45ad8dd7ec28139d.
//
// Solidity: event ChangeTrustedCaller(address indexed trustedCaller)
func (_IdRegistry *IdRegistryFilterer) WatchChangeTrustedCaller(opts *bind.WatchOpts, sink chan<- *IdRegistryChangeTrustedCaller, trustedCaller []common.Address) (event.Subscription, error) {

	var trustedCallerRule []interface{}
	for _, trustedCallerItem := range trustedCaller {
		trustedCallerRule = append(trustedCallerRule, trustedCallerItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "ChangeTrustedCaller", trustedCallerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryChangeTrustedCaller)
				if err := _IdRegistry.contract.UnpackLog(event, "ChangeTrustedCaller", log); err != nil {
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

// ParseChangeTrustedCaller is a log parse operation binding the contract event 0x255ba3357fefe42b361c216b6e0bc5541f1e6ea4c6178d4a45ad8dd7ec28139d.
//
// Solidity: event ChangeTrustedCaller(address indexed trustedCaller)
func (_IdRegistry *IdRegistryFilterer) ParseChangeTrustedCaller(log types.Log) (*IdRegistryChangeTrustedCaller, error) {
	event := new(IdRegistryChangeTrustedCaller)
	if err := _IdRegistry.contract.UnpackLog(event, "ChangeTrustedCaller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryDisableTrustedOnlyIterator is returned from FilterDisableTrustedOnly and is used to iterate over the raw logs and unpacked data for DisableTrustedOnly events raised by the IdRegistry contract.
type IdRegistryDisableTrustedOnlyIterator struct {
	Event *IdRegistryDisableTrustedOnly // Event containing the contract specifics and raw log

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
func (it *IdRegistryDisableTrustedOnlyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryDisableTrustedOnly)
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
		it.Event = new(IdRegistryDisableTrustedOnly)
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
func (it *IdRegistryDisableTrustedOnlyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryDisableTrustedOnlyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryDisableTrustedOnly represents a DisableTrustedOnly event raised by the IdRegistry contract.
type IdRegistryDisableTrustedOnly struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableTrustedOnly is a free log retrieval operation binding the contract event 0x03732e5295a5bd18e6ef95b03b41aa8bcadae292a7ef40468144c7a727dfa8b5.
//
// Solidity: event DisableTrustedOnly()
func (_IdRegistry *IdRegistryFilterer) FilterDisableTrustedOnly(opts *bind.FilterOpts) (*IdRegistryDisableTrustedOnlyIterator, error) {

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "DisableTrustedOnly")
	if err != nil {
		return nil, err
	}
	return &IdRegistryDisableTrustedOnlyIterator{contract: _IdRegistry.contract, event: "DisableTrustedOnly", logs: logs, sub: sub}, nil
}

// WatchDisableTrustedOnly is a free log subscription operation binding the contract event 0x03732e5295a5bd18e6ef95b03b41aa8bcadae292a7ef40468144c7a727dfa8b5.
//
// Solidity: event DisableTrustedOnly()
func (_IdRegistry *IdRegistryFilterer) WatchDisableTrustedOnly(opts *bind.WatchOpts, sink chan<- *IdRegistryDisableTrustedOnly) (event.Subscription, error) {

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "DisableTrustedOnly")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryDisableTrustedOnly)
				if err := _IdRegistry.contract.UnpackLog(event, "DisableTrustedOnly", log); err != nil {
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

// ParseDisableTrustedOnly is a log parse operation binding the contract event 0x03732e5295a5bd18e6ef95b03b41aa8bcadae292a7ef40468144c7a727dfa8b5.
//
// Solidity: event DisableTrustedOnly()
func (_IdRegistry *IdRegistryFilterer) ParseDisableTrustedOnly(log types.Log) (*IdRegistryDisableTrustedOnly, error) {
	event := new(IdRegistryDisableTrustedOnly)
	if err := _IdRegistry.contract.UnpackLog(event, "DisableTrustedOnly", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IdRegistry contract.
type IdRegistryOwnershipTransferredIterator struct {
	Event *IdRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryOwnershipTransferred)
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
		it.Event = new(IdRegistryOwnershipTransferred)
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
func (it *IdRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the IdRegistry contract.
type IdRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdRegistry *IdRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryOwnershipTransferredIterator{contract: _IdRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdRegistry *IdRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryOwnershipTransferred)
				if err := _IdRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IdRegistry *IdRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*IdRegistryOwnershipTransferred, error) {
	event := new(IdRegistryOwnershipTransferred)
	if err := _IdRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the IdRegistry contract.
type IdRegistryRegisterIterator struct {
	Event *IdRegistryRegister // Event containing the contract specifics and raw log

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
func (it *IdRegistryRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryRegister)
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
		it.Event = new(IdRegistryRegister)
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
func (it *IdRegistryRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryRegister represents a Register event raised by the IdRegistry contract.
type IdRegistryRegister struct {
	To       common.Address
	Id       *big.Int
	Recovery common.Address
	Url      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0x3cd6a0ffcc37406d9958e09bba79ff19d8237819eb2e1911f9edbce656499c87.
//
// Solidity: event Register(address indexed to, uint256 indexed id, address recovery, string url)
func (_IdRegistry *IdRegistryFilterer) FilterRegister(opts *bind.FilterOpts, to []common.Address, id []*big.Int) (*IdRegistryRegisterIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "Register", toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryRegisterIterator{contract: _IdRegistry.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0x3cd6a0ffcc37406d9958e09bba79ff19d8237819eb2e1911f9edbce656499c87.
//
// Solidity: event Register(address indexed to, uint256 indexed id, address recovery, string url)
func (_IdRegistry *IdRegistryFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *IdRegistryRegister, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "Register", toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryRegister)
				if err := _IdRegistry.contract.UnpackLog(event, "Register", log); err != nil {
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

// ParseRegister is a log parse operation binding the contract event 0x3cd6a0ffcc37406d9958e09bba79ff19d8237819eb2e1911f9edbce656499c87.
//
// Solidity: event Register(address indexed to, uint256 indexed id, address recovery, string url)
func (_IdRegistry *IdRegistryFilterer) ParseRegister(log types.Log) (*IdRegistryRegister, error) {
	event := new(IdRegistryRegister)
	if err := _IdRegistry.contract.UnpackLog(event, "Register", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryRequestRecoveryIterator is returned from FilterRequestRecovery and is used to iterate over the raw logs and unpacked data for RequestRecovery events raised by the IdRegistry contract.
type IdRegistryRequestRecoveryIterator struct {
	Event *IdRegistryRequestRecovery // Event containing the contract specifics and raw log

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
func (it *IdRegistryRequestRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryRequestRecovery)
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
		it.Event = new(IdRegistryRequestRecovery)
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
func (it *IdRegistryRequestRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryRequestRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryRequestRecovery represents a RequestRecovery event raised by the IdRegistry contract.
type IdRegistryRequestRecovery struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRequestRecovery is a free log retrieval operation binding the contract event 0xfab80e8bf242ed27bf595552dfdddbdd794f201d6dfcd8df7347f82f8e1f1f9b.
//
// Solidity: event RequestRecovery(address indexed from, address indexed to, uint256 indexed id)
func (_IdRegistry *IdRegistryFilterer) FilterRequestRecovery(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*IdRegistryRequestRecoveryIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "RequestRecovery", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryRequestRecoveryIterator{contract: _IdRegistry.contract, event: "RequestRecovery", logs: logs, sub: sub}, nil
}

// WatchRequestRecovery is a free log subscription operation binding the contract event 0xfab80e8bf242ed27bf595552dfdddbdd794f201d6dfcd8df7347f82f8e1f1f9b.
//
// Solidity: event RequestRecovery(address indexed from, address indexed to, uint256 indexed id)
func (_IdRegistry *IdRegistryFilterer) WatchRequestRecovery(opts *bind.WatchOpts, sink chan<- *IdRegistryRequestRecovery, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "RequestRecovery", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryRequestRecovery)
				if err := _IdRegistry.contract.UnpackLog(event, "RequestRecovery", log); err != nil {
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

// ParseRequestRecovery is a log parse operation binding the contract event 0xfab80e8bf242ed27bf595552dfdddbdd794f201d6dfcd8df7347f82f8e1f1f9b.
//
// Solidity: event RequestRecovery(address indexed from, address indexed to, uint256 indexed id)
func (_IdRegistry *IdRegistryFilterer) ParseRequestRecovery(log types.Log) (*IdRegistryRequestRecovery, error) {
	event := new(IdRegistryRequestRecovery)
	if err := _IdRegistry.contract.UnpackLog(event, "RequestRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IdRegistry contract.
type IdRegistryTransferIterator struct {
	Event *IdRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *IdRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdRegistryTransfer)
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
		it.Event = new(IdRegistryTransfer)
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
func (it *IdRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdRegistryTransfer represents a Transfer event raised by the IdRegistry contract.
type IdRegistryTransfer struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_IdRegistry *IdRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*IdRegistryTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.FilterLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &IdRegistryTransferIterator{contract: _IdRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_IdRegistry *IdRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IdRegistryTransfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IdRegistry.contract.WatchLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdRegistryTransfer)
				if err := _IdRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_IdRegistry *IdRegistryFilterer) ParseTransfer(log types.Log) (*IdRegistryTransfer, error) {
	event := new(IdRegistryTransfer)
	if err := _IdRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
