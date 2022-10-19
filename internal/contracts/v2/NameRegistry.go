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

// NameRegistryMetaData contains all meta data concerning the NameRegistry contract.
var NameRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitReplay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Escrow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidName\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecovery\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Invitable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoRecovery\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBiddable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInvitable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotModerator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRenewable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTreasurer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Registered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Registrable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"CancelRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ChangeFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"ChangePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"ChangeRecoveryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trustedCaller\",\"type\":\"address\"}],\"name\":\"ChangeTrustedCaller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"ChangeVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableTrustedOnly\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inviterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inviteeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"fname\",\"type\":\"bytes16\"}],\"name\":\"Invite\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"Renew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"RequestRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"changePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"changeRecoveryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustedCaller\",\"type\":\"address\"}],\"name\":\"changeTrustedCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"changeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"completeRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableTrustedOnly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"expiryOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"fname\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"generateCommit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"commit\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commit\",\"type\":\"bytes32\"}],\"name\":\"makeCommit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"reclaim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recoveryClockOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recoveryDestinationOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recoveryOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"fname\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"requestRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"timestampOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedOnly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"fname\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recovery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inviter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invitee\",\"type\":\"uint256\"}],\"name\":\"trustedRegister\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NameRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use NameRegistryMetaData.ABI instead.
var NameRegistryABI = NameRegistryMetaData.ABI

// NameRegistry is an auto generated Go binding around an Ethereum contract.
type NameRegistry struct {
	NameRegistryCaller     // Read-only binding to the contract
	NameRegistryTransactor // Write-only binding to the contract
	NameRegistryFilterer   // Log filterer for contract events
}

// NameRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type NameRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NameRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NameRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NameRegistrySession struct {
	Contract     *NameRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NameRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NameRegistryCallerSession struct {
	Contract *NameRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NameRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NameRegistryTransactorSession struct {
	Contract     *NameRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NameRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type NameRegistryRaw struct {
	Contract *NameRegistry // Generic contract binding to access the raw methods on
}

// NameRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NameRegistryCallerRaw struct {
	Contract *NameRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// NameRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NameRegistryTransactorRaw struct {
	Contract *NameRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNameRegistry creates a new instance of NameRegistry, bound to a specific deployed contract.
func NewNameRegistry(address common.Address, backend bind.ContractBackend) (*NameRegistry, error) {
	contract, err := bindNameRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NameRegistry{NameRegistryCaller: NameRegistryCaller{contract: contract}, NameRegistryTransactor: NameRegistryTransactor{contract: contract}, NameRegistryFilterer: NameRegistryFilterer{contract: contract}}, nil
}

// NewNameRegistryCaller creates a new read-only instance of NameRegistry, bound to a specific deployed contract.
func NewNameRegistryCaller(address common.Address, caller bind.ContractCaller) (*NameRegistryCaller, error) {
	contract, err := bindNameRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NameRegistryCaller{contract: contract}, nil
}

// NewNameRegistryTransactor creates a new write-only instance of NameRegistry, bound to a specific deployed contract.
func NewNameRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*NameRegistryTransactor, error) {
	contract, err := bindNameRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NameRegistryTransactor{contract: contract}, nil
}

// NewNameRegistryFilterer creates a new log filterer instance of NameRegistry, bound to a specific deployed contract.
func NewNameRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*NameRegistryFilterer, error) {
	contract, err := bindNameRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NameRegistryFilterer{contract: contract}, nil
}

// bindNameRegistry binds a generic wrapper to an already deployed contract.
func bindNameRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NameRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameRegistry *NameRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NameRegistry.Contract.NameRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameRegistry *NameRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameRegistry.Contract.NameRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameRegistry *NameRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameRegistry.Contract.NameRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameRegistry *NameRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NameRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameRegistry *NameRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameRegistry *NameRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NameRegistry *NameRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NameRegistry *NameRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NameRegistry.Contract.DEFAULTADMINROLE(&_NameRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NameRegistry *NameRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NameRegistry.Contract.DEFAULTADMINROLE(&_NameRegistry.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NameRegistry *NameRegistryCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NameRegistry *NameRegistrySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NameRegistry.Contract.BalanceOf(&_NameRegistry.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NameRegistry *NameRegistryCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NameRegistry.Contract.BalanceOf(&_NameRegistry.CallOpts, owner)
}

// ExpiryOf is a free data retrieval call binding the contract method 0xbaef73e9.
//
// Solidity: function expiryOf(uint256 ) view returns(uint256)
func (_NameRegistry *NameRegistryCaller) ExpiryOf(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "expiryOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpiryOf is a free data retrieval call binding the contract method 0xbaef73e9.
//
// Solidity: function expiryOf(uint256 ) view returns(uint256)
func (_NameRegistry *NameRegistrySession) ExpiryOf(arg0 *big.Int) (*big.Int, error) {
	return _NameRegistry.Contract.ExpiryOf(&_NameRegistry.CallOpts, arg0)
}

// ExpiryOf is a free data retrieval call binding the contract method 0xbaef73e9.
//
// Solidity: function expiryOf(uint256 ) view returns(uint256)
func (_NameRegistry *NameRegistryCallerSession) ExpiryOf(arg0 *big.Int) (*big.Int, error) {
	return _NameRegistry.Contract.ExpiryOf(&_NameRegistry.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_NameRegistry *NameRegistryCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_NameRegistry *NameRegistrySession) Fee() (*big.Int, error) {
	return _NameRegistry.Contract.Fee(&_NameRegistry.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_NameRegistry *NameRegistryCallerSession) Fee() (*big.Int, error) {
	return _NameRegistry.Contract.Fee(&_NameRegistry.CallOpts)
}

// GenerateCommit is a free data retrieval call binding the contract method 0x9fc999e4.
//
// Solidity: function generateCommit(bytes16 fname, address to, bytes32 secret, address recovery) pure returns(bytes32 commit)
func (_NameRegistry *NameRegistryCaller) GenerateCommit(opts *bind.CallOpts, fname [16]byte, to common.Address, secret [32]byte, recovery common.Address) ([32]byte, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "generateCommit", fname, to, secret, recovery)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GenerateCommit is a free data retrieval call binding the contract method 0x9fc999e4.
//
// Solidity: function generateCommit(bytes16 fname, address to, bytes32 secret, address recovery) pure returns(bytes32 commit)
func (_NameRegistry *NameRegistrySession) GenerateCommit(fname [16]byte, to common.Address, secret [32]byte, recovery common.Address) ([32]byte, error) {
	return _NameRegistry.Contract.GenerateCommit(&_NameRegistry.CallOpts, fname, to, secret, recovery)
}

// GenerateCommit is a free data retrieval call binding the contract method 0x9fc999e4.
//
// Solidity: function generateCommit(bytes16 fname, address to, bytes32 secret, address recovery) pure returns(bytes32 commit)
func (_NameRegistry *NameRegistryCallerSession) GenerateCommit(fname [16]byte, to common.Address, secret [32]byte, recovery common.Address) ([32]byte, error) {
	return _NameRegistry.Contract.GenerateCommit(&_NameRegistry.CallOpts, fname, to, secret, recovery)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NameRegistry *NameRegistryCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NameRegistry *NameRegistrySession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NameRegistry.Contract.GetApproved(&_NameRegistry.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NameRegistry *NameRegistryCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NameRegistry.Contract.GetApproved(&_NameRegistry.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NameRegistry *NameRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NameRegistry *NameRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NameRegistry.Contract.GetRoleAdmin(&_NameRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NameRegistry *NameRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NameRegistry.Contract.GetRoleAdmin(&_NameRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NameRegistry *NameRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NameRegistry *NameRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NameRegistry.Contract.HasRole(&_NameRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NameRegistry *NameRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NameRegistry.Contract.HasRole(&_NameRegistry.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NameRegistry *NameRegistryCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NameRegistry *NameRegistrySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NameRegistry.Contract.IsApprovedForAll(&_NameRegistry.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NameRegistry *NameRegistryCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NameRegistry.Contract.IsApprovedForAll(&_NameRegistry.CallOpts, owner, operator)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_NameRegistry *NameRegistryCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_NameRegistry *NameRegistrySession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _NameRegistry.Contract.IsTrustedForwarder(&_NameRegistry.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_NameRegistry *NameRegistryCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _NameRegistry.Contract.IsTrustedForwarder(&_NameRegistry.CallOpts, forwarder)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NameRegistry *NameRegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NameRegistry *NameRegistrySession) Name() (string, error) {
	return _NameRegistry.Contract.Name(&_NameRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NameRegistry *NameRegistryCallerSession) Name() (string, error) {
	return _NameRegistry.Contract.Name(&_NameRegistry.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NameRegistry *NameRegistryCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NameRegistry *NameRegistrySession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NameRegistry.Contract.OwnerOf(&_NameRegistry.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NameRegistry *NameRegistryCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NameRegistry.Contract.OwnerOf(&_NameRegistry.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NameRegistry *NameRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NameRegistry *NameRegistrySession) Paused() (bool, error) {
	return _NameRegistry.Contract.Paused(&_NameRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NameRegistry *NameRegistryCallerSession) Paused() (bool, error) {
	return _NameRegistry.Contract.Paused(&_NameRegistry.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_NameRegistry *NameRegistryCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_NameRegistry *NameRegistrySession) Pool() (common.Address, error) {
	return _NameRegistry.Contract.Pool(&_NameRegistry.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_NameRegistry *NameRegistryCallerSession) Pool() (common.Address, error) {
	return _NameRegistry.Contract.Pool(&_NameRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NameRegistry *NameRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NameRegistry *NameRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _NameRegistry.Contract.ProxiableUUID(&_NameRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NameRegistry *NameRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _NameRegistry.Contract.ProxiableUUID(&_NameRegistry.CallOpts)
}

// RecoveryClockOf is a free data retrieval call binding the contract method 0xac817ccc.
//
// Solidity: function recoveryClockOf(uint256 ) view returns(uint256)
func (_NameRegistry *NameRegistryCaller) RecoveryClockOf(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "recoveryClockOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RecoveryClockOf is a free data retrieval call binding the contract method 0xac817ccc.
//
// Solidity: function recoveryClockOf(uint256 ) view returns(uint256)
func (_NameRegistry *NameRegistrySession) RecoveryClockOf(arg0 *big.Int) (*big.Int, error) {
	return _NameRegistry.Contract.RecoveryClockOf(&_NameRegistry.CallOpts, arg0)
}

// RecoveryClockOf is a free data retrieval call binding the contract method 0xac817ccc.
//
// Solidity: function recoveryClockOf(uint256 ) view returns(uint256)
func (_NameRegistry *NameRegistryCallerSession) RecoveryClockOf(arg0 *big.Int) (*big.Int, error) {
	return _NameRegistry.Contract.RecoveryClockOf(&_NameRegistry.CallOpts, arg0)
}

// RecoveryDestinationOf is a free data retrieval call binding the contract method 0xcb17e04a.
//
// Solidity: function recoveryDestinationOf(uint256 ) view returns(address)
func (_NameRegistry *NameRegistryCaller) RecoveryDestinationOf(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "recoveryDestinationOf", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoveryDestinationOf is a free data retrieval call binding the contract method 0xcb17e04a.
//
// Solidity: function recoveryDestinationOf(uint256 ) view returns(address)
func (_NameRegistry *NameRegistrySession) RecoveryDestinationOf(arg0 *big.Int) (common.Address, error) {
	return _NameRegistry.Contract.RecoveryDestinationOf(&_NameRegistry.CallOpts, arg0)
}

// RecoveryDestinationOf is a free data retrieval call binding the contract method 0xcb17e04a.
//
// Solidity: function recoveryDestinationOf(uint256 ) view returns(address)
func (_NameRegistry *NameRegistryCallerSession) RecoveryDestinationOf(arg0 *big.Int) (common.Address, error) {
	return _NameRegistry.Contract.RecoveryDestinationOf(&_NameRegistry.CallOpts, arg0)
}

// RecoveryOf is a free data retrieval call binding the contract method 0xfa1a1b25.
//
// Solidity: function recoveryOf(uint256 ) view returns(address)
func (_NameRegistry *NameRegistryCaller) RecoveryOf(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "recoveryOf", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoveryOf is a free data retrieval call binding the contract method 0xfa1a1b25.
//
// Solidity: function recoveryOf(uint256 ) view returns(address)
func (_NameRegistry *NameRegistrySession) RecoveryOf(arg0 *big.Int) (common.Address, error) {
	return _NameRegistry.Contract.RecoveryOf(&_NameRegistry.CallOpts, arg0)
}

// RecoveryOf is a free data retrieval call binding the contract method 0xfa1a1b25.
//
// Solidity: function recoveryOf(uint256 ) view returns(address)
func (_NameRegistry *NameRegistryCallerSession) RecoveryOf(arg0 *big.Int) (common.Address, error) {
	return _NameRegistry.Contract.RecoveryOf(&_NameRegistry.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NameRegistry *NameRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NameRegistry *NameRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NameRegistry.Contract.SupportsInterface(&_NameRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NameRegistry *NameRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NameRegistry.Contract.SupportsInterface(&_NameRegistry.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NameRegistry *NameRegistryCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NameRegistry *NameRegistrySession) Symbol() (string, error) {
	return _NameRegistry.Contract.Symbol(&_NameRegistry.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NameRegistry *NameRegistryCallerSession) Symbol() (string, error) {
	return _NameRegistry.Contract.Symbol(&_NameRegistry.CallOpts)
}

// TimestampOf is a free data retrieval call binding the contract method 0x76fa0b8a.
//
// Solidity: function timestampOf(bytes32 ) view returns(uint256)
func (_NameRegistry *NameRegistryCaller) TimestampOf(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "timestampOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimestampOf is a free data retrieval call binding the contract method 0x76fa0b8a.
//
// Solidity: function timestampOf(bytes32 ) view returns(uint256)
func (_NameRegistry *NameRegistrySession) TimestampOf(arg0 [32]byte) (*big.Int, error) {
	return _NameRegistry.Contract.TimestampOf(&_NameRegistry.CallOpts, arg0)
}

// TimestampOf is a free data retrieval call binding the contract method 0x76fa0b8a.
//
// Solidity: function timestampOf(bytes32 ) view returns(uint256)
func (_NameRegistry *NameRegistryCallerSession) TimestampOf(arg0 [32]byte) (*big.Int, error) {
	return _NameRegistry.Contract.TimestampOf(&_NameRegistry.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) pure returns(string)
func (_NameRegistry *NameRegistryCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) pure returns(string)
func (_NameRegistry *NameRegistrySession) TokenURI(tokenId *big.Int) (string, error) {
	return _NameRegistry.Contract.TokenURI(&_NameRegistry.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) pure returns(string)
func (_NameRegistry *NameRegistryCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NameRegistry.Contract.TokenURI(&_NameRegistry.CallOpts, tokenId)
}

// TrustedCaller is a free data retrieval call binding the contract method 0x268f0760.
//
// Solidity: function trustedCaller() view returns(address)
func (_NameRegistry *NameRegistryCaller) TrustedCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "trustedCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedCaller is a free data retrieval call binding the contract method 0x268f0760.
//
// Solidity: function trustedCaller() view returns(address)
func (_NameRegistry *NameRegistrySession) TrustedCaller() (common.Address, error) {
	return _NameRegistry.Contract.TrustedCaller(&_NameRegistry.CallOpts)
}

// TrustedCaller is a free data retrieval call binding the contract method 0x268f0760.
//
// Solidity: function trustedCaller() view returns(address)
func (_NameRegistry *NameRegistryCallerSession) TrustedCaller() (common.Address, error) {
	return _NameRegistry.Contract.TrustedCaller(&_NameRegistry.CallOpts)
}

// TrustedOnly is a free data retrieval call binding the contract method 0x6b2ddd4e.
//
// Solidity: function trustedOnly() view returns(uint256)
func (_NameRegistry *NameRegistryCaller) TrustedOnly(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "trustedOnly")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TrustedOnly is a free data retrieval call binding the contract method 0x6b2ddd4e.
//
// Solidity: function trustedOnly() view returns(uint256)
func (_NameRegistry *NameRegistrySession) TrustedOnly() (*big.Int, error) {
	return _NameRegistry.Contract.TrustedOnly(&_NameRegistry.CallOpts)
}

// TrustedOnly is a free data retrieval call binding the contract method 0x6b2ddd4e.
//
// Solidity: function trustedOnly() view returns(uint256)
func (_NameRegistry *NameRegistryCallerSession) TrustedOnly() (*big.Int, error) {
	return _NameRegistry.Contract.TrustedOnly(&_NameRegistry.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_NameRegistry *NameRegistryCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_NameRegistry *NameRegistrySession) Vault() (common.Address, error) {
	return _NameRegistry.Contract.Vault(&_NameRegistry.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_NameRegistry *NameRegistryCallerSession) Vault() (common.Address, error) {
	return _NameRegistry.Contract.Vault(&_NameRegistry.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NameRegistry *NameRegistryTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NameRegistry *NameRegistrySession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.Approve(&_NameRegistry.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NameRegistry *NameRegistryTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.Approve(&_NameRegistry.TransactOpts, to, tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x5e751000.
//
// Solidity: function bid(address to, uint256 tokenId, address recovery) payable returns()
func (_NameRegistry *NameRegistryTransactor) Bid(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, recovery common.Address) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "bid", to, tokenId, recovery)
}

// Bid is a paid mutator transaction binding the contract method 0x5e751000.
//
// Solidity: function bid(address to, uint256 tokenId, address recovery) payable returns()
func (_NameRegistry *NameRegistrySession) Bid(to common.Address, tokenId *big.Int, recovery common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.Bid(&_NameRegistry.TransactOpts, to, tokenId, recovery)
}

// Bid is a paid mutator transaction binding the contract method 0x5e751000.
//
// Solidity: function bid(address to, uint256 tokenId, address recovery) payable returns()
func (_NameRegistry *NameRegistryTransactorSession) Bid(to common.Address, tokenId *big.Int, recovery common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.Bid(&_NameRegistry.TransactOpts, to, tokenId, recovery)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0x10de2676.
//
// Solidity: function cancelRecovery(uint256 tokenId) returns()
func (_NameRegistry *NameRegistryTransactor) CancelRecovery(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "cancelRecovery", tokenId)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0x10de2676.
//
// Solidity: function cancelRecovery(uint256 tokenId) returns()
func (_NameRegistry *NameRegistrySession) CancelRecovery(tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.CancelRecovery(&_NameRegistry.TransactOpts, tokenId)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0x10de2676.
//
// Solidity: function cancelRecovery(uint256 tokenId) returns()
func (_NameRegistry *NameRegistryTransactorSession) CancelRecovery(tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.CancelRecovery(&_NameRegistry.TransactOpts, tokenId)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 _fee) returns()
func (_NameRegistry *NameRegistryTransactor) ChangeFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "changeFee", _fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 _fee) returns()
func (_NameRegistry *NameRegistrySession) ChangeFee(_fee *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.ChangeFee(&_NameRegistry.TransactOpts, _fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 _fee) returns()
func (_NameRegistry *NameRegistryTransactorSession) ChangeFee(_fee *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.ChangeFee(&_NameRegistry.TransactOpts, _fee)
}

// ChangePool is a paid mutator transaction binding the contract method 0x4339bc30.
//
// Solidity: function changePool(address _pool) returns()
func (_NameRegistry *NameRegistryTransactor) ChangePool(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "changePool", _pool)
}

// ChangePool is a paid mutator transaction binding the contract method 0x4339bc30.
//
// Solidity: function changePool(address _pool) returns()
func (_NameRegistry *NameRegistrySession) ChangePool(_pool common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.ChangePool(&_NameRegistry.TransactOpts, _pool)
}

// ChangePool is a paid mutator transaction binding the contract method 0x4339bc30.
//
// Solidity: function changePool(address _pool) returns()
func (_NameRegistry *NameRegistryTransactorSession) ChangePool(_pool common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.ChangePool(&_NameRegistry.TransactOpts, _pool)
}

// ChangeRecoveryAddress is a paid mutator transaction binding the contract method 0x9950e7ee.
//
// Solidity: function changeRecoveryAddress(uint256 tokenId, address recovery) returns()
func (_NameRegistry *NameRegistryTransactor) ChangeRecoveryAddress(opts *bind.TransactOpts, tokenId *big.Int, recovery common.Address) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "changeRecoveryAddress", tokenId, recovery)
}

// ChangeRecoveryAddress is a paid mutator transaction binding the contract method 0x9950e7ee.
//
// Solidity: function changeRecoveryAddress(uint256 tokenId, address recovery) returns()
func (_NameRegistry *NameRegistrySession) ChangeRecoveryAddress(tokenId *big.Int, recovery common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.ChangeRecoveryAddress(&_NameRegistry.TransactOpts, tokenId, recovery)
}

// ChangeRecoveryAddress is a paid mutator transaction binding the contract method 0x9950e7ee.
//
// Solidity: function changeRecoveryAddress(uint256 tokenId, address recovery) returns()
func (_NameRegistry *NameRegistryTransactorSession) ChangeRecoveryAddress(tokenId *big.Int, recovery common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.ChangeRecoveryAddress(&_NameRegistry.TransactOpts, tokenId, recovery)
}

// ChangeTrustedCaller is a paid mutator transaction binding the contract method 0x881b1956.
//
// Solidity: function changeTrustedCaller(address _trustedCaller) returns()
func (_NameRegistry *NameRegistryTransactor) ChangeTrustedCaller(opts *bind.TransactOpts, _trustedCaller common.Address) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "changeTrustedCaller", _trustedCaller)
}

// ChangeTrustedCaller is a paid mutator transaction binding the contract method 0x881b1956.
//
// Solidity: function changeTrustedCaller(address _trustedCaller) returns()
func (_NameRegistry *NameRegistrySession) ChangeTrustedCaller(_trustedCaller common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.ChangeTrustedCaller(&_NameRegistry.TransactOpts, _trustedCaller)
}

// ChangeTrustedCaller is a paid mutator transaction binding the contract method 0x881b1956.
//
// Solidity: function changeTrustedCaller(address _trustedCaller) returns()
func (_NameRegistry *NameRegistryTransactorSession) ChangeTrustedCaller(_trustedCaller common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.ChangeTrustedCaller(&_NameRegistry.TransactOpts, _trustedCaller)
}

// ChangeVault is a paid mutator transaction binding the contract method 0x60e232a9.
//
// Solidity: function changeVault(address _vault) returns()
func (_NameRegistry *NameRegistryTransactor) ChangeVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "changeVault", _vault)
}

// ChangeVault is a paid mutator transaction binding the contract method 0x60e232a9.
//
// Solidity: function changeVault(address _vault) returns()
func (_NameRegistry *NameRegistrySession) ChangeVault(_vault common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.ChangeVault(&_NameRegistry.TransactOpts, _vault)
}

// ChangeVault is a paid mutator transaction binding the contract method 0x60e232a9.
//
// Solidity: function changeVault(address _vault) returns()
func (_NameRegistry *NameRegistryTransactorSession) ChangeVault(_vault common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.ChangeVault(&_NameRegistry.TransactOpts, _vault)
}

// CompleteRecovery is a paid mutator transaction binding the contract method 0x85fb8449.
//
// Solidity: function completeRecovery(uint256 tokenId) returns()
func (_NameRegistry *NameRegistryTransactor) CompleteRecovery(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "completeRecovery", tokenId)
}

// CompleteRecovery is a paid mutator transaction binding the contract method 0x85fb8449.
//
// Solidity: function completeRecovery(uint256 tokenId) returns()
func (_NameRegistry *NameRegistrySession) CompleteRecovery(tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.CompleteRecovery(&_NameRegistry.TransactOpts, tokenId)
}

// CompleteRecovery is a paid mutator transaction binding the contract method 0x85fb8449.
//
// Solidity: function completeRecovery(uint256 tokenId) returns()
func (_NameRegistry *NameRegistryTransactorSession) CompleteRecovery(tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.CompleteRecovery(&_NameRegistry.TransactOpts, tokenId)
}

// DisableTrustedOnly is a paid mutator transaction binding the contract method 0x6e9bde49.
//
// Solidity: function disableTrustedOnly() returns()
func (_NameRegistry *NameRegistryTransactor) DisableTrustedOnly(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "disableTrustedOnly")
}

// DisableTrustedOnly is a paid mutator transaction binding the contract method 0x6e9bde49.
//
// Solidity: function disableTrustedOnly() returns()
func (_NameRegistry *NameRegistrySession) DisableTrustedOnly() (*types.Transaction, error) {
	return _NameRegistry.Contract.DisableTrustedOnly(&_NameRegistry.TransactOpts)
}

// DisableTrustedOnly is a paid mutator transaction binding the contract method 0x6e9bde49.
//
// Solidity: function disableTrustedOnly() returns()
func (_NameRegistry *NameRegistryTransactorSession) DisableTrustedOnly() (*types.Transaction, error) {
	return _NameRegistry.Contract.DisableTrustedOnly(&_NameRegistry.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NameRegistry *NameRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NameRegistry *NameRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.GrantRole(&_NameRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NameRegistry *NameRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.GrantRole(&_NameRegistry.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f15b414.
//
// Solidity: function initialize(string _tokenName, string _tokenSymbol, address _vault, address _pool) returns()
func (_NameRegistry *NameRegistryTransactor) Initialize(opts *bind.TransactOpts, _tokenName string, _tokenSymbol string, _vault common.Address, _pool common.Address) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "initialize", _tokenName, _tokenSymbol, _vault, _pool)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f15b414.
//
// Solidity: function initialize(string _tokenName, string _tokenSymbol, address _vault, address _pool) returns()
func (_NameRegistry *NameRegistrySession) Initialize(_tokenName string, _tokenSymbol string, _vault common.Address, _pool common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.Initialize(&_NameRegistry.TransactOpts, _tokenName, _tokenSymbol, _vault, _pool)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f15b414.
//
// Solidity: function initialize(string _tokenName, string _tokenSymbol, address _vault, address _pool) returns()
func (_NameRegistry *NameRegistryTransactorSession) Initialize(_tokenName string, _tokenSymbol string, _vault common.Address, _pool common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.Initialize(&_NameRegistry.TransactOpts, _tokenName, _tokenSymbol, _vault, _pool)
}

// MakeCommit is a paid mutator transaction binding the contract method 0x0c97fc36.
//
// Solidity: function makeCommit(bytes32 commit) returns()
func (_NameRegistry *NameRegistryTransactor) MakeCommit(opts *bind.TransactOpts, commit [32]byte) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "makeCommit", commit)
}

// MakeCommit is a paid mutator transaction binding the contract method 0x0c97fc36.
//
// Solidity: function makeCommit(bytes32 commit) returns()
func (_NameRegistry *NameRegistrySession) MakeCommit(commit [32]byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.MakeCommit(&_NameRegistry.TransactOpts, commit)
}

// MakeCommit is a paid mutator transaction binding the contract method 0x0c97fc36.
//
// Solidity: function makeCommit(bytes32 commit) returns()
func (_NameRegistry *NameRegistryTransactorSession) MakeCommit(commit [32]byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.MakeCommit(&_NameRegistry.TransactOpts, commit)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NameRegistry *NameRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NameRegistry *NameRegistrySession) Pause() (*types.Transaction, error) {
	return _NameRegistry.Contract.Pause(&_NameRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NameRegistry *NameRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _NameRegistry.Contract.Pause(&_NameRegistry.TransactOpts)
}

// Reclaim is a paid mutator transaction binding the contract method 0x2dabbeed.
//
// Solidity: function reclaim(uint256 tokenId) payable returns()
func (_NameRegistry *NameRegistryTransactor) Reclaim(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "reclaim", tokenId)
}

// Reclaim is a paid mutator transaction binding the contract method 0x2dabbeed.
//
// Solidity: function reclaim(uint256 tokenId) payable returns()
func (_NameRegistry *NameRegistrySession) Reclaim(tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.Reclaim(&_NameRegistry.TransactOpts, tokenId)
}

// Reclaim is a paid mutator transaction binding the contract method 0x2dabbeed.
//
// Solidity: function reclaim(uint256 tokenId) payable returns()
func (_NameRegistry *NameRegistryTransactorSession) Reclaim(tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.Reclaim(&_NameRegistry.TransactOpts, tokenId)
}

// Register is a paid mutator transaction binding the contract method 0x464ac22c.
//
// Solidity: function register(bytes16 fname, address to, bytes32 secret, address recovery) payable returns()
func (_NameRegistry *NameRegistryTransactor) Register(opts *bind.TransactOpts, fname [16]byte, to common.Address, secret [32]byte, recovery common.Address) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "register", fname, to, secret, recovery)
}

// Register is a paid mutator transaction binding the contract method 0x464ac22c.
//
// Solidity: function register(bytes16 fname, address to, bytes32 secret, address recovery) payable returns()
func (_NameRegistry *NameRegistrySession) Register(fname [16]byte, to common.Address, secret [32]byte, recovery common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.Register(&_NameRegistry.TransactOpts, fname, to, secret, recovery)
}

// Register is a paid mutator transaction binding the contract method 0x464ac22c.
//
// Solidity: function register(bytes16 fname, address to, bytes32 secret, address recovery) payable returns()
func (_NameRegistry *NameRegistryTransactorSession) Register(fname [16]byte, to common.Address, secret [32]byte, recovery common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.Register(&_NameRegistry.TransactOpts, fname, to, secret, recovery)
}

// Renew is a paid mutator transaction binding the contract method 0x5baa7509.
//
// Solidity: function renew(uint256 tokenId) payable returns()
func (_NameRegistry *NameRegistryTransactor) Renew(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "renew", tokenId)
}

// Renew is a paid mutator transaction binding the contract method 0x5baa7509.
//
// Solidity: function renew(uint256 tokenId) payable returns()
func (_NameRegistry *NameRegistrySession) Renew(tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.Renew(&_NameRegistry.TransactOpts, tokenId)
}

// Renew is a paid mutator transaction binding the contract method 0x5baa7509.
//
// Solidity: function renew(uint256 tokenId) payable returns()
func (_NameRegistry *NameRegistryTransactorSession) Renew(tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.Renew(&_NameRegistry.TransactOpts, tokenId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NameRegistry *NameRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NameRegistry *NameRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.RenounceRole(&_NameRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NameRegistry *NameRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.RenounceRole(&_NameRegistry.TransactOpts, role, account)
}

// RequestRecovery is a paid mutator transaction binding the contract method 0x6473b35e.
//
// Solidity: function requestRecovery(uint256 tokenId, address to) returns()
func (_NameRegistry *NameRegistryTransactor) RequestRecovery(opts *bind.TransactOpts, tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "requestRecovery", tokenId, to)
}

// RequestRecovery is a paid mutator transaction binding the contract method 0x6473b35e.
//
// Solidity: function requestRecovery(uint256 tokenId, address to) returns()
func (_NameRegistry *NameRegistrySession) RequestRecovery(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.RequestRecovery(&_NameRegistry.TransactOpts, tokenId, to)
}

// RequestRecovery is a paid mutator transaction binding the contract method 0x6473b35e.
//
// Solidity: function requestRecovery(uint256 tokenId, address to) returns()
func (_NameRegistry *NameRegistryTransactorSession) RequestRecovery(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.RequestRecovery(&_NameRegistry.TransactOpts, tokenId, to)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NameRegistry *NameRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NameRegistry *NameRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.RevokeRole(&_NameRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NameRegistry *NameRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.RevokeRole(&_NameRegistry.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NameRegistry *NameRegistryTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NameRegistry *NameRegistrySession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.SafeTransferFrom(&_NameRegistry.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NameRegistry *NameRegistryTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.SafeTransferFrom(&_NameRegistry.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_NameRegistry *NameRegistryTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_NameRegistry *NameRegistrySession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.SafeTransferFrom0(&_NameRegistry.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_NameRegistry *NameRegistryTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.SafeTransferFrom0(&_NameRegistry.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NameRegistry *NameRegistryTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NameRegistry *NameRegistrySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NameRegistry.Contract.SetApprovalForAll(&_NameRegistry.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NameRegistry *NameRegistryTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NameRegistry.Contract.SetApprovalForAll(&_NameRegistry.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NameRegistry *NameRegistryTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NameRegistry *NameRegistrySession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.TransferFrom(&_NameRegistry.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NameRegistry *NameRegistryTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.TransferFrom(&_NameRegistry.TransactOpts, from, to, tokenId)
}

// TrustedRegister is a paid mutator transaction binding the contract method 0x3f09a95a.
//
// Solidity: function trustedRegister(bytes16 fname, address to, address recovery, uint256 inviter, uint256 invitee) payable returns()
func (_NameRegistry *NameRegistryTransactor) TrustedRegister(opts *bind.TransactOpts, fname [16]byte, to common.Address, recovery common.Address, inviter *big.Int, invitee *big.Int) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "trustedRegister", fname, to, recovery, inviter, invitee)
}

// TrustedRegister is a paid mutator transaction binding the contract method 0x3f09a95a.
//
// Solidity: function trustedRegister(bytes16 fname, address to, address recovery, uint256 inviter, uint256 invitee) payable returns()
func (_NameRegistry *NameRegistrySession) TrustedRegister(fname [16]byte, to common.Address, recovery common.Address, inviter *big.Int, invitee *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.TrustedRegister(&_NameRegistry.TransactOpts, fname, to, recovery, inviter, invitee)
}

// TrustedRegister is a paid mutator transaction binding the contract method 0x3f09a95a.
//
// Solidity: function trustedRegister(bytes16 fname, address to, address recovery, uint256 inviter, uint256 invitee) payable returns()
func (_NameRegistry *NameRegistryTransactorSession) TrustedRegister(fname [16]byte, to common.Address, recovery common.Address, inviter *big.Int, invitee *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.TrustedRegister(&_NameRegistry.TransactOpts, fname, to, recovery, inviter, invitee)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NameRegistry *NameRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NameRegistry *NameRegistrySession) Unpause() (*types.Transaction, error) {
	return _NameRegistry.Contract.Unpause(&_NameRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NameRegistry *NameRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _NameRegistry.Contract.Unpause(&_NameRegistry.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NameRegistry *NameRegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NameRegistry *NameRegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpgradeTo(&_NameRegistry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NameRegistry *NameRegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpgradeTo(&_NameRegistry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NameRegistry *NameRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NameRegistry *NameRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpgradeToAndCall(&_NameRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NameRegistry *NameRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpgradeToAndCall(&_NameRegistry.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NameRegistry *NameRegistryTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NameRegistry *NameRegistrySession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.Withdraw(&_NameRegistry.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NameRegistry *NameRegistryTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.Withdraw(&_NameRegistry.TransactOpts, amount)
}

// NameRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the NameRegistry contract.
type NameRegistryAdminChangedIterator struct {
	Event *NameRegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *NameRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryAdminChanged)
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
		it.Event = new(NameRegistryAdminChanged)
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
func (it *NameRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryAdminChanged represents a AdminChanged event raised by the NameRegistry contract.
type NameRegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NameRegistry *NameRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*NameRegistryAdminChangedIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &NameRegistryAdminChangedIterator{contract: _NameRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NameRegistry *NameRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *NameRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryAdminChanged)
				if err := _NameRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NameRegistry *NameRegistryFilterer) ParseAdminChanged(log types.Log) (*NameRegistryAdminChanged, error) {
	event := new(NameRegistryAdminChanged)
	if err := _NameRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NameRegistry contract.
type NameRegistryApprovalIterator struct {
	Event *NameRegistryApproval // Event containing the contract specifics and raw log

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
func (it *NameRegistryApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryApproval)
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
		it.Event = new(NameRegistryApproval)
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
func (it *NameRegistryApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryApproval represents a Approval event raised by the NameRegistry contract.
type NameRegistryApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NameRegistry *NameRegistryFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NameRegistryApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryApprovalIterator{contract: _NameRegistry.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NameRegistry *NameRegistryFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NameRegistryApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryApproval)
				if err := _NameRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NameRegistry *NameRegistryFilterer) ParseApproval(log types.Log) (*NameRegistryApproval, error) {
	event := new(NameRegistryApproval)
	if err := _NameRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the NameRegistry contract.
type NameRegistryApprovalForAllIterator struct {
	Event *NameRegistryApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NameRegistryApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryApprovalForAll)
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
		it.Event = new(NameRegistryApprovalForAll)
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
func (it *NameRegistryApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryApprovalForAll represents a ApprovalForAll event raised by the NameRegistry contract.
type NameRegistryApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NameRegistry *NameRegistryFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NameRegistryApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryApprovalForAllIterator{contract: _NameRegistry.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NameRegistry *NameRegistryFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NameRegistryApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryApprovalForAll)
				if err := _NameRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NameRegistry *NameRegistryFilterer) ParseApprovalForAll(log types.Log) (*NameRegistryApprovalForAll, error) {
	event := new(NameRegistryApprovalForAll)
	if err := _NameRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the NameRegistry contract.
type NameRegistryBeaconUpgradedIterator struct {
	Event *NameRegistryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *NameRegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryBeaconUpgraded)
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
		it.Event = new(NameRegistryBeaconUpgraded)
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
func (it *NameRegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryBeaconUpgraded represents a BeaconUpgraded event raised by the NameRegistry contract.
type NameRegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NameRegistry *NameRegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*NameRegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryBeaconUpgradedIterator{contract: _NameRegistry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NameRegistry *NameRegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *NameRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryBeaconUpgraded)
				if err := _NameRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NameRegistry *NameRegistryFilterer) ParseBeaconUpgraded(log types.Log) (*NameRegistryBeaconUpgraded, error) {
	event := new(NameRegistryBeaconUpgraded)
	if err := _NameRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryCancelRecoveryIterator is returned from FilterCancelRecovery and is used to iterate over the raw logs and unpacked data for CancelRecovery events raised by the NameRegistry contract.
type NameRegistryCancelRecoveryIterator struct {
	Event *NameRegistryCancelRecovery // Event containing the contract specifics and raw log

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
func (it *NameRegistryCancelRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryCancelRecovery)
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
		it.Event = new(NameRegistryCancelRecovery)
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
func (it *NameRegistryCancelRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryCancelRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryCancelRecovery represents a CancelRecovery event raised by the NameRegistry contract.
type NameRegistryCancelRecovery struct {
	By      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCancelRecovery is a free log retrieval operation binding the contract event 0x6181d4215ebc71e962cc193554c17f05a825da06230fdf9ece45081f09cb206f.
//
// Solidity: event CancelRecovery(address indexed by, uint256 indexed tokenId)
func (_NameRegistry *NameRegistryFilterer) FilterCancelRecovery(opts *bind.FilterOpts, by []common.Address, tokenId []*big.Int) (*NameRegistryCancelRecoveryIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "CancelRecovery", byRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryCancelRecoveryIterator{contract: _NameRegistry.contract, event: "CancelRecovery", logs: logs, sub: sub}, nil
}

// WatchCancelRecovery is a free log subscription operation binding the contract event 0x6181d4215ebc71e962cc193554c17f05a825da06230fdf9ece45081f09cb206f.
//
// Solidity: event CancelRecovery(address indexed by, uint256 indexed tokenId)
func (_NameRegistry *NameRegistryFilterer) WatchCancelRecovery(opts *bind.WatchOpts, sink chan<- *NameRegistryCancelRecovery, by []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "CancelRecovery", byRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryCancelRecovery)
				if err := _NameRegistry.contract.UnpackLog(event, "CancelRecovery", log); err != nil {
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
// Solidity: event CancelRecovery(address indexed by, uint256 indexed tokenId)
func (_NameRegistry *NameRegistryFilterer) ParseCancelRecovery(log types.Log) (*NameRegistryCancelRecovery, error) {
	event := new(NameRegistryCancelRecovery)
	if err := _NameRegistry.contract.UnpackLog(event, "CancelRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryChangeFeeIterator is returned from FilterChangeFee and is used to iterate over the raw logs and unpacked data for ChangeFee events raised by the NameRegistry contract.
type NameRegistryChangeFeeIterator struct {
	Event *NameRegistryChangeFee // Event containing the contract specifics and raw log

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
func (it *NameRegistryChangeFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryChangeFee)
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
		it.Event = new(NameRegistryChangeFee)
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
func (it *NameRegistryChangeFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryChangeFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryChangeFee represents a ChangeFee event raised by the NameRegistry contract.
type NameRegistryChangeFee struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChangeFee is a free log retrieval operation binding the contract event 0xfba1d84f2e30311f1380f2355f294fbedd53264c2504378e2c4b2133dea16670.
//
// Solidity: event ChangeFee(uint256 fee)
func (_NameRegistry *NameRegistryFilterer) FilterChangeFee(opts *bind.FilterOpts) (*NameRegistryChangeFeeIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "ChangeFee")
	if err != nil {
		return nil, err
	}
	return &NameRegistryChangeFeeIterator{contract: _NameRegistry.contract, event: "ChangeFee", logs: logs, sub: sub}, nil
}

// WatchChangeFee is a free log subscription operation binding the contract event 0xfba1d84f2e30311f1380f2355f294fbedd53264c2504378e2c4b2133dea16670.
//
// Solidity: event ChangeFee(uint256 fee)
func (_NameRegistry *NameRegistryFilterer) WatchChangeFee(opts *bind.WatchOpts, sink chan<- *NameRegistryChangeFee) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "ChangeFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryChangeFee)
				if err := _NameRegistry.contract.UnpackLog(event, "ChangeFee", log); err != nil {
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

// ParseChangeFee is a log parse operation binding the contract event 0xfba1d84f2e30311f1380f2355f294fbedd53264c2504378e2c4b2133dea16670.
//
// Solidity: event ChangeFee(uint256 fee)
func (_NameRegistry *NameRegistryFilterer) ParseChangeFee(log types.Log) (*NameRegistryChangeFee, error) {
	event := new(NameRegistryChangeFee)
	if err := _NameRegistry.contract.UnpackLog(event, "ChangeFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryChangePoolIterator is returned from FilterChangePool and is used to iterate over the raw logs and unpacked data for ChangePool events raised by the NameRegistry contract.
type NameRegistryChangePoolIterator struct {
	Event *NameRegistryChangePool // Event containing the contract specifics and raw log

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
func (it *NameRegistryChangePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryChangePool)
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
		it.Event = new(NameRegistryChangePool)
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
func (it *NameRegistryChangePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryChangePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryChangePool represents a ChangePool event raised by the NameRegistry contract.
type NameRegistryChangePool struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterChangePool is a free log retrieval operation binding the contract event 0x79025dab199855650264c602de305dcf5c292b8e5b4470ef271724a79d0343f4.
//
// Solidity: event ChangePool(address indexed pool)
func (_NameRegistry *NameRegistryFilterer) FilterChangePool(opts *bind.FilterOpts, pool []common.Address) (*NameRegistryChangePoolIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "ChangePool", poolRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryChangePoolIterator{contract: _NameRegistry.contract, event: "ChangePool", logs: logs, sub: sub}, nil
}

// WatchChangePool is a free log subscription operation binding the contract event 0x79025dab199855650264c602de305dcf5c292b8e5b4470ef271724a79d0343f4.
//
// Solidity: event ChangePool(address indexed pool)
func (_NameRegistry *NameRegistryFilterer) WatchChangePool(opts *bind.WatchOpts, sink chan<- *NameRegistryChangePool, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "ChangePool", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryChangePool)
				if err := _NameRegistry.contract.UnpackLog(event, "ChangePool", log); err != nil {
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

// ParseChangePool is a log parse operation binding the contract event 0x79025dab199855650264c602de305dcf5c292b8e5b4470ef271724a79d0343f4.
//
// Solidity: event ChangePool(address indexed pool)
func (_NameRegistry *NameRegistryFilterer) ParseChangePool(log types.Log) (*NameRegistryChangePool, error) {
	event := new(NameRegistryChangePool)
	if err := _NameRegistry.contract.UnpackLog(event, "ChangePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryChangeRecoveryAddressIterator is returned from FilterChangeRecoveryAddress and is used to iterate over the raw logs and unpacked data for ChangeRecoveryAddress events raised by the NameRegistry contract.
type NameRegistryChangeRecoveryAddressIterator struct {
	Event *NameRegistryChangeRecoveryAddress // Event containing the contract specifics and raw log

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
func (it *NameRegistryChangeRecoveryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryChangeRecoveryAddress)
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
		it.Event = new(NameRegistryChangeRecoveryAddress)
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
func (it *NameRegistryChangeRecoveryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryChangeRecoveryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryChangeRecoveryAddress represents a ChangeRecoveryAddress event raised by the NameRegistry contract.
type NameRegistryChangeRecoveryAddress struct {
	TokenId  *big.Int
	Recovery common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChangeRecoveryAddress is a free log retrieval operation binding the contract event 0x8e700b803af43e14651431cd73c9fe7d11b131ad797576a70b893ce5766f65c3.
//
// Solidity: event ChangeRecoveryAddress(uint256 indexed tokenId, address indexed recovery)
func (_NameRegistry *NameRegistryFilterer) FilterChangeRecoveryAddress(opts *bind.FilterOpts, tokenId []*big.Int, recovery []common.Address) (*NameRegistryChangeRecoveryAddressIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var recoveryRule []interface{}
	for _, recoveryItem := range recovery {
		recoveryRule = append(recoveryRule, recoveryItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "ChangeRecoveryAddress", tokenIdRule, recoveryRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryChangeRecoveryAddressIterator{contract: _NameRegistry.contract, event: "ChangeRecoveryAddress", logs: logs, sub: sub}, nil
}

// WatchChangeRecoveryAddress is a free log subscription operation binding the contract event 0x8e700b803af43e14651431cd73c9fe7d11b131ad797576a70b893ce5766f65c3.
//
// Solidity: event ChangeRecoveryAddress(uint256 indexed tokenId, address indexed recovery)
func (_NameRegistry *NameRegistryFilterer) WatchChangeRecoveryAddress(opts *bind.WatchOpts, sink chan<- *NameRegistryChangeRecoveryAddress, tokenId []*big.Int, recovery []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var recoveryRule []interface{}
	for _, recoveryItem := range recovery {
		recoveryRule = append(recoveryRule, recoveryItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "ChangeRecoveryAddress", tokenIdRule, recoveryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryChangeRecoveryAddress)
				if err := _NameRegistry.contract.UnpackLog(event, "ChangeRecoveryAddress", log); err != nil {
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
// Solidity: event ChangeRecoveryAddress(uint256 indexed tokenId, address indexed recovery)
func (_NameRegistry *NameRegistryFilterer) ParseChangeRecoveryAddress(log types.Log) (*NameRegistryChangeRecoveryAddress, error) {
	event := new(NameRegistryChangeRecoveryAddress)
	if err := _NameRegistry.contract.UnpackLog(event, "ChangeRecoveryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryChangeTrustedCallerIterator is returned from FilterChangeTrustedCaller and is used to iterate over the raw logs and unpacked data for ChangeTrustedCaller events raised by the NameRegistry contract.
type NameRegistryChangeTrustedCallerIterator struct {
	Event *NameRegistryChangeTrustedCaller // Event containing the contract specifics and raw log

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
func (it *NameRegistryChangeTrustedCallerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryChangeTrustedCaller)
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
		it.Event = new(NameRegistryChangeTrustedCaller)
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
func (it *NameRegistryChangeTrustedCallerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryChangeTrustedCallerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryChangeTrustedCaller represents a ChangeTrustedCaller event raised by the NameRegistry contract.
type NameRegistryChangeTrustedCaller struct {
	TrustedCaller common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChangeTrustedCaller is a free log retrieval operation binding the contract event 0x255ba3357fefe42b361c216b6e0bc5541f1e6ea4c6178d4a45ad8dd7ec28139d.
//
// Solidity: event ChangeTrustedCaller(address indexed trustedCaller)
func (_NameRegistry *NameRegistryFilterer) FilterChangeTrustedCaller(opts *bind.FilterOpts, trustedCaller []common.Address) (*NameRegistryChangeTrustedCallerIterator, error) {

	var trustedCallerRule []interface{}
	for _, trustedCallerItem := range trustedCaller {
		trustedCallerRule = append(trustedCallerRule, trustedCallerItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "ChangeTrustedCaller", trustedCallerRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryChangeTrustedCallerIterator{contract: _NameRegistry.contract, event: "ChangeTrustedCaller", logs: logs, sub: sub}, nil
}

// WatchChangeTrustedCaller is a free log subscription operation binding the contract event 0x255ba3357fefe42b361c216b6e0bc5541f1e6ea4c6178d4a45ad8dd7ec28139d.
//
// Solidity: event ChangeTrustedCaller(address indexed trustedCaller)
func (_NameRegistry *NameRegistryFilterer) WatchChangeTrustedCaller(opts *bind.WatchOpts, sink chan<- *NameRegistryChangeTrustedCaller, trustedCaller []common.Address) (event.Subscription, error) {

	var trustedCallerRule []interface{}
	for _, trustedCallerItem := range trustedCaller {
		trustedCallerRule = append(trustedCallerRule, trustedCallerItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "ChangeTrustedCaller", trustedCallerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryChangeTrustedCaller)
				if err := _NameRegistry.contract.UnpackLog(event, "ChangeTrustedCaller", log); err != nil {
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
func (_NameRegistry *NameRegistryFilterer) ParseChangeTrustedCaller(log types.Log) (*NameRegistryChangeTrustedCaller, error) {
	event := new(NameRegistryChangeTrustedCaller)
	if err := _NameRegistry.contract.UnpackLog(event, "ChangeTrustedCaller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryChangeVaultIterator is returned from FilterChangeVault and is used to iterate over the raw logs and unpacked data for ChangeVault events raised by the NameRegistry contract.
type NameRegistryChangeVaultIterator struct {
	Event *NameRegistryChangeVault // Event containing the contract specifics and raw log

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
func (it *NameRegistryChangeVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryChangeVault)
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
		it.Event = new(NameRegistryChangeVault)
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
func (it *NameRegistryChangeVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryChangeVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryChangeVault represents a ChangeVault event raised by the NameRegistry contract.
type NameRegistryChangeVault struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChangeVault is a free log retrieval operation binding the contract event 0x646d70535c6b451b92021874a72abd441f122ba1c0b8f24d074352bd169fad3f.
//
// Solidity: event ChangeVault(address indexed vault)
func (_NameRegistry *NameRegistryFilterer) FilterChangeVault(opts *bind.FilterOpts, vault []common.Address) (*NameRegistryChangeVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "ChangeVault", vaultRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryChangeVaultIterator{contract: _NameRegistry.contract, event: "ChangeVault", logs: logs, sub: sub}, nil
}

// WatchChangeVault is a free log subscription operation binding the contract event 0x646d70535c6b451b92021874a72abd441f122ba1c0b8f24d074352bd169fad3f.
//
// Solidity: event ChangeVault(address indexed vault)
func (_NameRegistry *NameRegistryFilterer) WatchChangeVault(opts *bind.WatchOpts, sink chan<- *NameRegistryChangeVault, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "ChangeVault", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryChangeVault)
				if err := _NameRegistry.contract.UnpackLog(event, "ChangeVault", log); err != nil {
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

// ParseChangeVault is a log parse operation binding the contract event 0x646d70535c6b451b92021874a72abd441f122ba1c0b8f24d074352bd169fad3f.
//
// Solidity: event ChangeVault(address indexed vault)
func (_NameRegistry *NameRegistryFilterer) ParseChangeVault(log types.Log) (*NameRegistryChangeVault, error) {
	event := new(NameRegistryChangeVault)
	if err := _NameRegistry.contract.UnpackLog(event, "ChangeVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryDisableTrustedOnlyIterator is returned from FilterDisableTrustedOnly and is used to iterate over the raw logs and unpacked data for DisableTrustedOnly events raised by the NameRegistry contract.
type NameRegistryDisableTrustedOnlyIterator struct {
	Event *NameRegistryDisableTrustedOnly // Event containing the contract specifics and raw log

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
func (it *NameRegistryDisableTrustedOnlyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryDisableTrustedOnly)
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
		it.Event = new(NameRegistryDisableTrustedOnly)
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
func (it *NameRegistryDisableTrustedOnlyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryDisableTrustedOnlyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryDisableTrustedOnly represents a DisableTrustedOnly event raised by the NameRegistry contract.
type NameRegistryDisableTrustedOnly struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableTrustedOnly is a free log retrieval operation binding the contract event 0x03732e5295a5bd18e6ef95b03b41aa8bcadae292a7ef40468144c7a727dfa8b5.
//
// Solidity: event DisableTrustedOnly()
func (_NameRegistry *NameRegistryFilterer) FilterDisableTrustedOnly(opts *bind.FilterOpts) (*NameRegistryDisableTrustedOnlyIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "DisableTrustedOnly")
	if err != nil {
		return nil, err
	}
	return &NameRegistryDisableTrustedOnlyIterator{contract: _NameRegistry.contract, event: "DisableTrustedOnly", logs: logs, sub: sub}, nil
}

// WatchDisableTrustedOnly is a free log subscription operation binding the contract event 0x03732e5295a5bd18e6ef95b03b41aa8bcadae292a7ef40468144c7a727dfa8b5.
//
// Solidity: event DisableTrustedOnly()
func (_NameRegistry *NameRegistryFilterer) WatchDisableTrustedOnly(opts *bind.WatchOpts, sink chan<- *NameRegistryDisableTrustedOnly) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "DisableTrustedOnly")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryDisableTrustedOnly)
				if err := _NameRegistry.contract.UnpackLog(event, "DisableTrustedOnly", log); err != nil {
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
func (_NameRegistry *NameRegistryFilterer) ParseDisableTrustedOnly(log types.Log) (*NameRegistryDisableTrustedOnly, error) {
	event := new(NameRegistryDisableTrustedOnly)
	if err := _NameRegistry.contract.UnpackLog(event, "DisableTrustedOnly", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NameRegistry contract.
type NameRegistryInitializedIterator struct {
	Event *NameRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *NameRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryInitialized)
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
		it.Event = new(NameRegistryInitialized)
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
func (it *NameRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryInitialized represents a Initialized event raised by the NameRegistry contract.
type NameRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NameRegistry *NameRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*NameRegistryInitializedIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NameRegistryInitializedIterator{contract: _NameRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NameRegistry *NameRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NameRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryInitialized)
				if err := _NameRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NameRegistry *NameRegistryFilterer) ParseInitialized(log types.Log) (*NameRegistryInitialized, error) {
	event := new(NameRegistryInitialized)
	if err := _NameRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryInviteIterator is returned from FilterInvite and is used to iterate over the raw logs and unpacked data for Invite events raised by the NameRegistry contract.
type NameRegistryInviteIterator struct {
	Event *NameRegistryInvite // Event containing the contract specifics and raw log

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
func (it *NameRegistryInviteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryInvite)
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
		it.Event = new(NameRegistryInvite)
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
func (it *NameRegistryInviteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryInviteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryInvite represents a Invite event raised by the NameRegistry contract.
type NameRegistryInvite struct {
	InviterId *big.Int
	InviteeId *big.Int
	Fname     [16]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInvite is a free log retrieval operation binding the contract event 0xed994b8dfbd359de8b535931832fe533e23820fbb73ce69d8dde9bd677989f39.
//
// Solidity: event Invite(uint256 indexed inviterId, uint256 indexed inviteeId, bytes16 indexed fname)
func (_NameRegistry *NameRegistryFilterer) FilterInvite(opts *bind.FilterOpts, inviterId []*big.Int, inviteeId []*big.Int, fname [][16]byte) (*NameRegistryInviteIterator, error) {

	var inviterIdRule []interface{}
	for _, inviterIdItem := range inviterId {
		inviterIdRule = append(inviterIdRule, inviterIdItem)
	}
	var inviteeIdRule []interface{}
	for _, inviteeIdItem := range inviteeId {
		inviteeIdRule = append(inviteeIdRule, inviteeIdItem)
	}
	var fnameRule []interface{}
	for _, fnameItem := range fname {
		fnameRule = append(fnameRule, fnameItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "Invite", inviterIdRule, inviteeIdRule, fnameRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryInviteIterator{contract: _NameRegistry.contract, event: "Invite", logs: logs, sub: sub}, nil
}

// WatchInvite is a free log subscription operation binding the contract event 0xed994b8dfbd359de8b535931832fe533e23820fbb73ce69d8dde9bd677989f39.
//
// Solidity: event Invite(uint256 indexed inviterId, uint256 indexed inviteeId, bytes16 indexed fname)
func (_NameRegistry *NameRegistryFilterer) WatchInvite(opts *bind.WatchOpts, sink chan<- *NameRegistryInvite, inviterId []*big.Int, inviteeId []*big.Int, fname [][16]byte) (event.Subscription, error) {

	var inviterIdRule []interface{}
	for _, inviterIdItem := range inviterId {
		inviterIdRule = append(inviterIdRule, inviterIdItem)
	}
	var inviteeIdRule []interface{}
	for _, inviteeIdItem := range inviteeId {
		inviteeIdRule = append(inviteeIdRule, inviteeIdItem)
	}
	var fnameRule []interface{}
	for _, fnameItem := range fname {
		fnameRule = append(fnameRule, fnameItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "Invite", inviterIdRule, inviteeIdRule, fnameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryInvite)
				if err := _NameRegistry.contract.UnpackLog(event, "Invite", log); err != nil {
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

// ParseInvite is a log parse operation binding the contract event 0xed994b8dfbd359de8b535931832fe533e23820fbb73ce69d8dde9bd677989f39.
//
// Solidity: event Invite(uint256 indexed inviterId, uint256 indexed inviteeId, bytes16 indexed fname)
func (_NameRegistry *NameRegistryFilterer) ParseInvite(log types.Log) (*NameRegistryInvite, error) {
	event := new(NameRegistryInvite)
	if err := _NameRegistry.contract.UnpackLog(event, "Invite", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the NameRegistry contract.
type NameRegistryPausedIterator struct {
	Event *NameRegistryPaused // Event containing the contract specifics and raw log

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
func (it *NameRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryPaused)
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
		it.Event = new(NameRegistryPaused)
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
func (it *NameRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryPaused represents a Paused event raised by the NameRegistry contract.
type NameRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NameRegistry *NameRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*NameRegistryPausedIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NameRegistryPausedIterator{contract: _NameRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NameRegistry *NameRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NameRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryPaused)
				if err := _NameRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NameRegistry *NameRegistryFilterer) ParsePaused(log types.Log) (*NameRegistryPaused, error) {
	event := new(NameRegistryPaused)
	if err := _NameRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryRenewIterator is returned from FilterRenew and is used to iterate over the raw logs and unpacked data for Renew events raised by the NameRegistry contract.
type NameRegistryRenewIterator struct {
	Event *NameRegistryRenew // Event containing the contract specifics and raw log

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
func (it *NameRegistryRenewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryRenew)
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
		it.Event = new(NameRegistryRenew)
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
func (it *NameRegistryRenewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryRenewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryRenew represents a Renew event raised by the NameRegistry contract.
type NameRegistryRenew struct {
	TokenId *big.Int
	Expiry  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRenew is a free log retrieval operation binding the contract event 0xbf5b84fa6df1868d005e90d05ee12a6c49025be6f38d2807f183743676744c16.
//
// Solidity: event Renew(uint256 indexed tokenId, uint256 expiry)
func (_NameRegistry *NameRegistryFilterer) FilterRenew(opts *bind.FilterOpts, tokenId []*big.Int) (*NameRegistryRenewIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "Renew", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryRenewIterator{contract: _NameRegistry.contract, event: "Renew", logs: logs, sub: sub}, nil
}

// WatchRenew is a free log subscription operation binding the contract event 0xbf5b84fa6df1868d005e90d05ee12a6c49025be6f38d2807f183743676744c16.
//
// Solidity: event Renew(uint256 indexed tokenId, uint256 expiry)
func (_NameRegistry *NameRegistryFilterer) WatchRenew(opts *bind.WatchOpts, sink chan<- *NameRegistryRenew, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "Renew", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryRenew)
				if err := _NameRegistry.contract.UnpackLog(event, "Renew", log); err != nil {
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

// ParseRenew is a log parse operation binding the contract event 0xbf5b84fa6df1868d005e90d05ee12a6c49025be6f38d2807f183743676744c16.
//
// Solidity: event Renew(uint256 indexed tokenId, uint256 expiry)
func (_NameRegistry *NameRegistryFilterer) ParseRenew(log types.Log) (*NameRegistryRenew, error) {
	event := new(NameRegistryRenew)
	if err := _NameRegistry.contract.UnpackLog(event, "Renew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryRequestRecoveryIterator is returned from FilterRequestRecovery and is used to iterate over the raw logs and unpacked data for RequestRecovery events raised by the NameRegistry contract.
type NameRegistryRequestRecoveryIterator struct {
	Event *NameRegistryRequestRecovery // Event containing the contract specifics and raw log

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
func (it *NameRegistryRequestRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryRequestRecovery)
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
		it.Event = new(NameRegistryRequestRecovery)
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
func (it *NameRegistryRequestRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryRequestRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryRequestRecovery represents a RequestRecovery event raised by the NameRegistry contract.
type NameRegistryRequestRecovery struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRequestRecovery is a free log retrieval operation binding the contract event 0xfab80e8bf242ed27bf595552dfdddbdd794f201d6dfcd8df7347f82f8e1f1f9b.
//
// Solidity: event RequestRecovery(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NameRegistry *NameRegistryFilterer) FilterRequestRecovery(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NameRegistryRequestRecoveryIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "RequestRecovery", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryRequestRecoveryIterator{contract: _NameRegistry.contract, event: "RequestRecovery", logs: logs, sub: sub}, nil
}

// WatchRequestRecovery is a free log subscription operation binding the contract event 0xfab80e8bf242ed27bf595552dfdddbdd794f201d6dfcd8df7347f82f8e1f1f9b.
//
// Solidity: event RequestRecovery(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NameRegistry *NameRegistryFilterer) WatchRequestRecovery(opts *bind.WatchOpts, sink chan<- *NameRegistryRequestRecovery, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "RequestRecovery", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryRequestRecovery)
				if err := _NameRegistry.contract.UnpackLog(event, "RequestRecovery", log); err != nil {
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
// Solidity: event RequestRecovery(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NameRegistry *NameRegistryFilterer) ParseRequestRecovery(log types.Log) (*NameRegistryRequestRecovery, error) {
	event := new(NameRegistryRequestRecovery)
	if err := _NameRegistry.contract.UnpackLog(event, "RequestRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the NameRegistry contract.
type NameRegistryRoleAdminChangedIterator struct {
	Event *NameRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *NameRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryRoleAdminChanged)
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
		it.Event = new(NameRegistryRoleAdminChanged)
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
func (it *NameRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the NameRegistry contract.
type NameRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NameRegistry *NameRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NameRegistryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryRoleAdminChangedIterator{contract: _NameRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NameRegistry *NameRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NameRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryRoleAdminChanged)
				if err := _NameRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NameRegistry *NameRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*NameRegistryRoleAdminChanged, error) {
	event := new(NameRegistryRoleAdminChanged)
	if err := _NameRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the NameRegistry contract.
type NameRegistryRoleGrantedIterator struct {
	Event *NameRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *NameRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryRoleGranted)
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
		it.Event = new(NameRegistryRoleGranted)
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
func (it *NameRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryRoleGranted represents a RoleGranted event raised by the NameRegistry contract.
type NameRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NameRegistry *NameRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NameRegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryRoleGrantedIterator{contract: _NameRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NameRegistry *NameRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NameRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryRoleGranted)
				if err := _NameRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NameRegistry *NameRegistryFilterer) ParseRoleGranted(log types.Log) (*NameRegistryRoleGranted, error) {
	event := new(NameRegistryRoleGranted)
	if err := _NameRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the NameRegistry contract.
type NameRegistryRoleRevokedIterator struct {
	Event *NameRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *NameRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryRoleRevoked)
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
		it.Event = new(NameRegistryRoleRevoked)
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
func (it *NameRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryRoleRevoked represents a RoleRevoked event raised by the NameRegistry contract.
type NameRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NameRegistry *NameRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NameRegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryRoleRevokedIterator{contract: _NameRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NameRegistry *NameRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NameRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryRoleRevoked)
				if err := _NameRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NameRegistry *NameRegistryFilterer) ParseRoleRevoked(log types.Log) (*NameRegistryRoleRevoked, error) {
	event := new(NameRegistryRoleRevoked)
	if err := _NameRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NameRegistry contract.
type NameRegistryTransferIterator struct {
	Event *NameRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *NameRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryTransfer)
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
		it.Event = new(NameRegistryTransfer)
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
func (it *NameRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryTransfer represents a Transfer event raised by the NameRegistry contract.
type NameRegistryTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NameRegistry *NameRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NameRegistryTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryTransferIterator{contract: _NameRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NameRegistry *NameRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NameRegistryTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryTransfer)
				if err := _NameRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NameRegistry *NameRegistryFilterer) ParseTransfer(log types.Log) (*NameRegistryTransfer, error) {
	event := new(NameRegistryTransfer)
	if err := _NameRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the NameRegistry contract.
type NameRegistryUnpausedIterator struct {
	Event *NameRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *NameRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryUnpaused)
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
		it.Event = new(NameRegistryUnpaused)
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
func (it *NameRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryUnpaused represents a Unpaused event raised by the NameRegistry contract.
type NameRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NameRegistry *NameRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NameRegistryUnpausedIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NameRegistryUnpausedIterator{contract: _NameRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NameRegistry *NameRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NameRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryUnpaused)
				if err := _NameRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NameRegistry *NameRegistryFilterer) ParseUnpaused(log types.Log) (*NameRegistryUnpaused, error) {
	event := new(NameRegistryUnpaused)
	if err := _NameRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NameRegistry contract.
type NameRegistryUpgradedIterator struct {
	Event *NameRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *NameRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryUpgraded)
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
		it.Event = new(NameRegistryUpgraded)
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
func (it *NameRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryUpgraded represents a Upgraded event raised by the NameRegistry contract.
type NameRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NameRegistry *NameRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NameRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NameRegistryUpgradedIterator{contract: _NameRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NameRegistry *NameRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NameRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryUpgraded)
				if err := _NameRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NameRegistry *NameRegistryFilterer) ParseUpgraded(log types.Log) (*NameRegistryUpgraded, error) {
	event := new(NameRegistryUpgraded)
	if err := _NameRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
