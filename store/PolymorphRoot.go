// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// PolymorphRootParams is an auto generated low-level Go binding around an user-defined struct.
type PolymorphRootParams struct {
	Name                  string
	Symbol                string
	BaseURI               string
	DaoAddress            common.Address
	PremintedTokensCount  *big.Int
	BaseGenomeChangePrice *big.Int
	PolymorphPrice        *big.Int
	MaxSupply             *big.Int
	RandomizeGenomePrice  *big.Int
	BulkBuyLimit          *big.Int
	ArweaveAssetsJSON     string
	PolymorphV1Address    common.Address
}

// PolymorphRootMetaData contains all meta data concerning the PolymorphRoot contract.
var PolymorphRootMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"_daoAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"premintedTokensCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseGenomeChangePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_polymorphPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_randomizeGenomePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bulkBuyLimit\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_arweaveAssetsJSON\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_polymorphV1Address\",\"type\":\"address\"}],\"internalType\":\"structPolymorphRoot.Params\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"arweaveAssetsJSON\",\"type\":\"string\"}],\"name\":\"ArweaveAssetsJSONChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGenomeChange\",\"type\":\"uint256\"}],\"name\":\"BaseGenomeChangePriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"BaseURIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBulkBuyLimit\",\"type\":\"uint256\"}],\"name\":\"BulkBuyLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ConsumerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPolymorphPrice\",\"type\":\"uint256\"}],\"name\":\"PolymorphPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRandomizeGenomePriceChange\",\"type\":\"uint256\"}],\"name\":\"RandomizeGenomePriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gene\",\"type\":\"uint256\"}],\"name\":\"TokenBurnedAndMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGene\",\"type\":\"uint256\"}],\"name\":\"TokenMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldGene\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGene\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumPolymorph.PolymorphEventType\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"TokenMorphed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arweaveAssetsJSON\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseGenomeChangePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bulkBuy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bulkBuyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"burnAndMintNewPolymorph\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"burnCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGenomeChangePrice\",\"type\":\"uint256\"}],\"name\":\"changeBaseGenomeChangePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consumer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"changeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRandomizeGenomePrice\",\"type\":\"uint256\"}],\"name\":\"changeRandomizeGenomePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"consumerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"geneOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gene\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"genomeChanges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"genomeChnages\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isNotVirgin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genePosition\",\"type\":\"uint256\"}],\"name\":\"morphGene\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polymorphPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polymorphV1Contract\",\"outputs\":[{\"internalType\":\"contractPolymorph\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"priceForGenomeChange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"randomizeGenome\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomizeGenomePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_arweaveAssetsJSON\",\"type\":\"string\"}],\"name\":\"setArweaveAssetsJSON\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bulkBuyLimit\",\"type\":\"uint256\"}],\"name\":\"setBulkBuyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPolymorphPrice\",\"type\":\"uint256\"}],\"name\":\"setPolymorphPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"whitelistBridgeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistTunnelAddresses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gene\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVirgin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"genomeChangesCount\",\"type\":\"uint256\"}],\"name\":\"wormholeUpdateGene\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PolymorphRootABI is the input ABI used to generate the binding from.
// Deprecated: Use PolymorphRootMetaData.ABI instead.
var PolymorphRootABI = PolymorphRootMetaData.ABI

// PolymorphRoot is an auto generated Go binding around an Ethereum contract.
type PolymorphRoot struct {
	PolymorphRootCaller     // Read-only binding to the contract
	PolymorphRootTransactor // Write-only binding to the contract
	PolymorphRootFilterer   // Log filterer for contract events
}

// PolymorphRootCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolymorphRootCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolymorphRootTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolymorphRootTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolymorphRootFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolymorphRootFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolymorphRootSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolymorphRootSession struct {
	Contract     *PolymorphRoot    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolymorphRootCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolymorphRootCallerSession struct {
	Contract *PolymorphRootCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PolymorphRootTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolymorphRootTransactorSession struct {
	Contract     *PolymorphRootTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PolymorphRootRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolymorphRootRaw struct {
	Contract *PolymorphRoot // Generic contract binding to access the raw methods on
}

// PolymorphRootCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolymorphRootCallerRaw struct {
	Contract *PolymorphRootCaller // Generic read-only contract binding to access the raw methods on
}

// PolymorphRootTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolymorphRootTransactorRaw struct {
	Contract *PolymorphRootTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolymorphRoot creates a new instance of PolymorphRoot, bound to a specific deployed contract.
func NewPolymorphRoot(address common.Address, backend bind.ContractBackend) (*PolymorphRoot, error) {
	contract, err := bindPolymorphRoot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PolymorphRoot{PolymorphRootCaller: PolymorphRootCaller{contract: contract}, PolymorphRootTransactor: PolymorphRootTransactor{contract: contract}, PolymorphRootFilterer: PolymorphRootFilterer{contract: contract}}, nil
}

// NewPolymorphRootCaller creates a new read-only instance of PolymorphRoot, bound to a specific deployed contract.
func NewPolymorphRootCaller(address common.Address, caller bind.ContractCaller) (*PolymorphRootCaller, error) {
	contract, err := bindPolymorphRoot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolymorphRootCaller{contract: contract}, nil
}

// NewPolymorphRootTransactor creates a new write-only instance of PolymorphRoot, bound to a specific deployed contract.
func NewPolymorphRootTransactor(address common.Address, transactor bind.ContractTransactor) (*PolymorphRootTransactor, error) {
	contract, err := bindPolymorphRoot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolymorphRootTransactor{contract: contract}, nil
}

// NewPolymorphRootFilterer creates a new log filterer instance of PolymorphRoot, bound to a specific deployed contract.
func NewPolymorphRootFilterer(address common.Address, filterer bind.ContractFilterer) (*PolymorphRootFilterer, error) {
	contract, err := bindPolymorphRoot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolymorphRootFilterer{contract: contract}, nil
}

// bindPolymorphRoot binds a generic wrapper to an already deployed contract.
func bindPolymorphRoot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PolymorphRootABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolymorphRoot *PolymorphRootRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolymorphRoot.Contract.PolymorphRootCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolymorphRoot *PolymorphRootRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.PolymorphRootTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolymorphRoot *PolymorphRootRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.PolymorphRootTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolymorphRoot *PolymorphRootCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolymorphRoot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolymorphRoot *PolymorphRootTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolymorphRoot *PolymorphRootTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PolymorphRoot *PolymorphRootCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PolymorphRoot *PolymorphRootSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PolymorphRoot.Contract.DEFAULTADMINROLE(&_PolymorphRoot.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PolymorphRoot *PolymorphRootCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PolymorphRoot.Contract.DEFAULTADMINROLE(&_PolymorphRoot.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PolymorphRoot *PolymorphRootCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PolymorphRoot *PolymorphRootSession) MINTERROLE() ([32]byte, error) {
	return _PolymorphRoot.Contract.MINTERROLE(&_PolymorphRoot.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PolymorphRoot *PolymorphRootCallerSession) MINTERROLE() ([32]byte, error) {
	return _PolymorphRoot.Contract.MINTERROLE(&_PolymorphRoot.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_PolymorphRoot *PolymorphRootCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_PolymorphRoot *PolymorphRootSession) PAUSERROLE() ([32]byte, error) {
	return _PolymorphRoot.Contract.PAUSERROLE(&_PolymorphRoot.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_PolymorphRoot *PolymorphRootCallerSession) PAUSERROLE() ([32]byte, error) {
	return _PolymorphRoot.Contract.PAUSERROLE(&_PolymorphRoot.CallOpts)
}

// ArweaveAssetsJSON is a free data retrieval call binding the contract method 0xf528a627.
//
// Solidity: function arweaveAssetsJSON() view returns(string)
func (_PolymorphRoot *PolymorphRootCaller) ArweaveAssetsJSON(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "arweaveAssetsJSON")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ArweaveAssetsJSON is a free data retrieval call binding the contract method 0xf528a627.
//
// Solidity: function arweaveAssetsJSON() view returns(string)
func (_PolymorphRoot *PolymorphRootSession) ArweaveAssetsJSON() (string, error) {
	return _PolymorphRoot.Contract.ArweaveAssetsJSON(&_PolymorphRoot.CallOpts)
}

// ArweaveAssetsJSON is a free data retrieval call binding the contract method 0xf528a627.
//
// Solidity: function arweaveAssetsJSON() view returns(string)
func (_PolymorphRoot *PolymorphRootCallerSession) ArweaveAssetsJSON() (string, error) {
	return _PolymorphRoot.Contract.ArweaveAssetsJSON(&_PolymorphRoot.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PolymorphRoot *PolymorphRootCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PolymorphRoot *PolymorphRootSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PolymorphRoot.Contract.BalanceOf(&_PolymorphRoot.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PolymorphRoot *PolymorphRootCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PolymorphRoot.Contract.BalanceOf(&_PolymorphRoot.CallOpts, owner)
}

// BaseGenomeChangePrice is a free data retrieval call binding the contract method 0xce14617d.
//
// Solidity: function baseGenomeChangePrice() view returns(uint256)
func (_PolymorphRoot *PolymorphRootCaller) BaseGenomeChangePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "baseGenomeChangePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseGenomeChangePrice is a free data retrieval call binding the contract method 0xce14617d.
//
// Solidity: function baseGenomeChangePrice() view returns(uint256)
func (_PolymorphRoot *PolymorphRootSession) BaseGenomeChangePrice() (*big.Int, error) {
	return _PolymorphRoot.Contract.BaseGenomeChangePrice(&_PolymorphRoot.CallOpts)
}

// BaseGenomeChangePrice is a free data retrieval call binding the contract method 0xce14617d.
//
// Solidity: function baseGenomeChangePrice() view returns(uint256)
func (_PolymorphRoot *PolymorphRootCallerSession) BaseGenomeChangePrice() (*big.Int, error) {
	return _PolymorphRoot.Contract.BaseGenomeChangePrice(&_PolymorphRoot.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_PolymorphRoot *PolymorphRootCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_PolymorphRoot *PolymorphRootSession) BaseURI() (string, error) {
	return _PolymorphRoot.Contract.BaseURI(&_PolymorphRoot.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_PolymorphRoot *PolymorphRootCallerSession) BaseURI() (string, error) {
	return _PolymorphRoot.Contract.BaseURI(&_PolymorphRoot.CallOpts)
}

// BulkBuyLimit is a free data retrieval call binding the contract method 0xa49bccca.
//
// Solidity: function bulkBuyLimit() view returns(uint256)
func (_PolymorphRoot *PolymorphRootCaller) BulkBuyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "bulkBuyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BulkBuyLimit is a free data retrieval call binding the contract method 0xa49bccca.
//
// Solidity: function bulkBuyLimit() view returns(uint256)
func (_PolymorphRoot *PolymorphRootSession) BulkBuyLimit() (*big.Int, error) {
	return _PolymorphRoot.Contract.BulkBuyLimit(&_PolymorphRoot.CallOpts)
}

// BulkBuyLimit is a free data retrieval call binding the contract method 0xa49bccca.
//
// Solidity: function bulkBuyLimit() view returns(uint256)
func (_PolymorphRoot *PolymorphRootCallerSession) BulkBuyLimit() (*big.Int, error) {
	return _PolymorphRoot.Contract.BulkBuyLimit(&_PolymorphRoot.CallOpts)
}

// BurnCount is a free data retrieval call binding the contract method 0x15889e43.
//
// Solidity: function burnCount(address ) view returns(uint256)
func (_PolymorphRoot *PolymorphRootCaller) BurnCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "burnCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnCount is a free data retrieval call binding the contract method 0x15889e43.
//
// Solidity: function burnCount(address ) view returns(uint256)
func (_PolymorphRoot *PolymorphRootSession) BurnCount(arg0 common.Address) (*big.Int, error) {
	return _PolymorphRoot.Contract.BurnCount(&_PolymorphRoot.CallOpts, arg0)
}

// BurnCount is a free data retrieval call binding the contract method 0x15889e43.
//
// Solidity: function burnCount(address ) view returns(uint256)
func (_PolymorphRoot *PolymorphRootCallerSession) BurnCount(arg0 common.Address) (*big.Int, error) {
	return _PolymorphRoot.Contract.BurnCount(&_PolymorphRoot.CallOpts, arg0)
}

// ConsumerOf is a free data retrieval call binding the contract method 0xe5892331.
//
// Solidity: function consumerOf(uint256 _tokenId) view returns(address)
func (_PolymorphRoot *PolymorphRootCaller) ConsumerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "consumerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConsumerOf is a free data retrieval call binding the contract method 0xe5892331.
//
// Solidity: function consumerOf(uint256 _tokenId) view returns(address)
func (_PolymorphRoot *PolymorphRootSession) ConsumerOf(_tokenId *big.Int) (common.Address, error) {
	return _PolymorphRoot.Contract.ConsumerOf(&_PolymorphRoot.CallOpts, _tokenId)
}

// ConsumerOf is a free data retrieval call binding the contract method 0xe5892331.
//
// Solidity: function consumerOf(uint256 _tokenId) view returns(address)
func (_PolymorphRoot *PolymorphRootCallerSession) ConsumerOf(_tokenId *big.Int) (common.Address, error) {
	return _PolymorphRoot.Contract.ConsumerOf(&_PolymorphRoot.CallOpts, _tokenId)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_PolymorphRoot *PolymorphRootCaller) DaoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "daoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_PolymorphRoot *PolymorphRootSession) DaoAddress() (common.Address, error) {
	return _PolymorphRoot.Contract.DaoAddress(&_PolymorphRoot.CallOpts)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_PolymorphRoot *PolymorphRootCallerSession) DaoAddress() (common.Address, error) {
	return _PolymorphRoot.Contract.DaoAddress(&_PolymorphRoot.CallOpts)
}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_PolymorphRoot *PolymorphRootCaller) GeneOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "geneOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_PolymorphRoot *PolymorphRootSession) GeneOf(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphRoot.Contract.GeneOf(&_PolymorphRoot.CallOpts, tokenId)
}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_PolymorphRoot *PolymorphRootCallerSession) GeneOf(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphRoot.Contract.GeneOf(&_PolymorphRoot.CallOpts, tokenId)
}

// GenomeChanges is a free data retrieval call binding the contract method 0x23c8d07a.
//
// Solidity: function genomeChanges(uint256 tokenId) view returns(uint256 genomeChnages)
func (_PolymorphRoot *PolymorphRootCaller) GenomeChanges(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "genomeChanges", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenomeChanges is a free data retrieval call binding the contract method 0x23c8d07a.
//
// Solidity: function genomeChanges(uint256 tokenId) view returns(uint256 genomeChnages)
func (_PolymorphRoot *PolymorphRootSession) GenomeChanges(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphRoot.Contract.GenomeChanges(&_PolymorphRoot.CallOpts, tokenId)
}

// GenomeChanges is a free data retrieval call binding the contract method 0x23c8d07a.
//
// Solidity: function genomeChanges(uint256 tokenId) view returns(uint256 genomeChnages)
func (_PolymorphRoot *PolymorphRootCallerSession) GenomeChanges(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphRoot.Contract.GenomeChanges(&_PolymorphRoot.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_PolymorphRoot *PolymorphRootCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_PolymorphRoot *PolymorphRootSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _PolymorphRoot.Contract.GetApproved(&_PolymorphRoot.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_PolymorphRoot *PolymorphRootCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _PolymorphRoot.Contract.GetApproved(&_PolymorphRoot.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PolymorphRoot *PolymorphRootCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PolymorphRoot *PolymorphRootSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PolymorphRoot.Contract.GetRoleAdmin(&_PolymorphRoot.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PolymorphRoot *PolymorphRootCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PolymorphRoot.Contract.GetRoleAdmin(&_PolymorphRoot.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_PolymorphRoot *PolymorphRootCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_PolymorphRoot *PolymorphRootSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _PolymorphRoot.Contract.GetRoleMember(&_PolymorphRoot.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_PolymorphRoot *PolymorphRootCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _PolymorphRoot.Contract.GetRoleMember(&_PolymorphRoot.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_PolymorphRoot *PolymorphRootCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_PolymorphRoot *PolymorphRootSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _PolymorphRoot.Contract.GetRoleMemberCount(&_PolymorphRoot.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_PolymorphRoot *PolymorphRootCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _PolymorphRoot.Contract.GetRoleMemberCount(&_PolymorphRoot.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PolymorphRoot *PolymorphRootCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PolymorphRoot *PolymorphRootSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PolymorphRoot.Contract.HasRole(&_PolymorphRoot.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PolymorphRoot *PolymorphRootCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PolymorphRoot.Contract.HasRole(&_PolymorphRoot.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_PolymorphRoot *PolymorphRootCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_PolymorphRoot *PolymorphRootSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _PolymorphRoot.Contract.IsApprovedForAll(&_PolymorphRoot.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_PolymorphRoot *PolymorphRootCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _PolymorphRoot.Contract.IsApprovedForAll(&_PolymorphRoot.CallOpts, owner, operator)
}

// IsNotVirgin is a free data retrieval call binding the contract method 0x4df77416.
//
// Solidity: function isNotVirgin(uint256 ) view returns(bool)
func (_PolymorphRoot *PolymorphRootCaller) IsNotVirgin(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "isNotVirgin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNotVirgin is a free data retrieval call binding the contract method 0x4df77416.
//
// Solidity: function isNotVirgin(uint256 ) view returns(bool)
func (_PolymorphRoot *PolymorphRootSession) IsNotVirgin(arg0 *big.Int) (bool, error) {
	return _PolymorphRoot.Contract.IsNotVirgin(&_PolymorphRoot.CallOpts, arg0)
}

// IsNotVirgin is a free data retrieval call binding the contract method 0x4df77416.
//
// Solidity: function isNotVirgin(uint256 ) view returns(bool)
func (_PolymorphRoot *PolymorphRootCallerSession) IsNotVirgin(arg0 *big.Int) (bool, error) {
	return _PolymorphRoot.Contract.IsNotVirgin(&_PolymorphRoot.CallOpts, arg0)
}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_PolymorphRoot *PolymorphRootCaller) LastTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "lastTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_PolymorphRoot *PolymorphRootSession) LastTokenId() (*big.Int, error) {
	return _PolymorphRoot.Contract.LastTokenId(&_PolymorphRoot.CallOpts)
}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_PolymorphRoot *PolymorphRootCallerSession) LastTokenId() (*big.Int, error) {
	return _PolymorphRoot.Contract.LastTokenId(&_PolymorphRoot.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_PolymorphRoot *PolymorphRootCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_PolymorphRoot *PolymorphRootSession) MaxSupply() (*big.Int, error) {
	return _PolymorphRoot.Contract.MaxSupply(&_PolymorphRoot.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_PolymorphRoot *PolymorphRootCallerSession) MaxSupply() (*big.Int, error) {
	return _PolymorphRoot.Contract.MaxSupply(&_PolymorphRoot.CallOpts)
}

// Mint0 is a free data retrieval call binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) pure returns()
func (_PolymorphRoot *PolymorphRootCaller) Mint0(opts *bind.CallOpts, to common.Address) error {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "mint0", to)

	if err != nil {
		return err
	}

	return err

}

// Mint0 is a free data retrieval call binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) pure returns()
func (_PolymorphRoot *PolymorphRootSession) Mint0(to common.Address) error {
	return _PolymorphRoot.Contract.Mint0(&_PolymorphRoot.CallOpts, to)
}

// Mint0 is a free data retrieval call binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) pure returns()
func (_PolymorphRoot *PolymorphRootCallerSession) Mint0(to common.Address) error {
	return _PolymorphRoot.Contract.Mint0(&_PolymorphRoot.CallOpts, to)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PolymorphRoot *PolymorphRootCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PolymorphRoot *PolymorphRootSession) Name() (string, error) {
	return _PolymorphRoot.Contract.Name(&_PolymorphRoot.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PolymorphRoot *PolymorphRootCallerSession) Name() (string, error) {
	return _PolymorphRoot.Contract.Name(&_PolymorphRoot.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_PolymorphRoot *PolymorphRootCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_PolymorphRoot *PolymorphRootSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _PolymorphRoot.Contract.OwnerOf(&_PolymorphRoot.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_PolymorphRoot *PolymorphRootCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _PolymorphRoot.Contract.OwnerOf(&_PolymorphRoot.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PolymorphRoot *PolymorphRootCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PolymorphRoot *PolymorphRootSession) Paused() (bool, error) {
	return _PolymorphRoot.Contract.Paused(&_PolymorphRoot.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PolymorphRoot *PolymorphRootCallerSession) Paused() (bool, error) {
	return _PolymorphRoot.Contract.Paused(&_PolymorphRoot.CallOpts)
}

// PolymorphPrice is a free data retrieval call binding the contract method 0x704ec036.
//
// Solidity: function polymorphPrice() view returns(uint256)
func (_PolymorphRoot *PolymorphRootCaller) PolymorphPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "polymorphPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PolymorphPrice is a free data retrieval call binding the contract method 0x704ec036.
//
// Solidity: function polymorphPrice() view returns(uint256)
func (_PolymorphRoot *PolymorphRootSession) PolymorphPrice() (*big.Int, error) {
	return _PolymorphRoot.Contract.PolymorphPrice(&_PolymorphRoot.CallOpts)
}

// PolymorphPrice is a free data retrieval call binding the contract method 0x704ec036.
//
// Solidity: function polymorphPrice() view returns(uint256)
func (_PolymorphRoot *PolymorphRootCallerSession) PolymorphPrice() (*big.Int, error) {
	return _PolymorphRoot.Contract.PolymorphPrice(&_PolymorphRoot.CallOpts)
}

// PolymorphV1Contract is a free data retrieval call binding the contract method 0x25b081ff.
//
// Solidity: function polymorphV1Contract() view returns(address)
func (_PolymorphRoot *PolymorphRootCaller) PolymorphV1Contract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "polymorphV1Contract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolymorphV1Contract is a free data retrieval call binding the contract method 0x25b081ff.
//
// Solidity: function polymorphV1Contract() view returns(address)
func (_PolymorphRoot *PolymorphRootSession) PolymorphV1Contract() (common.Address, error) {
	return _PolymorphRoot.Contract.PolymorphV1Contract(&_PolymorphRoot.CallOpts)
}

// PolymorphV1Contract is a free data retrieval call binding the contract method 0x25b081ff.
//
// Solidity: function polymorphV1Contract() view returns(address)
func (_PolymorphRoot *PolymorphRootCallerSession) PolymorphV1Contract() (common.Address, error) {
	return _PolymorphRoot.Contract.PolymorphV1Contract(&_PolymorphRoot.CallOpts)
}

// PriceForGenomeChange is a free data retrieval call binding the contract method 0xd45351e5.
//
// Solidity: function priceForGenomeChange(uint256 tokenId) view returns(uint256 price)
func (_PolymorphRoot *PolymorphRootCaller) PriceForGenomeChange(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "priceForGenomeChange", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceForGenomeChange is a free data retrieval call binding the contract method 0xd45351e5.
//
// Solidity: function priceForGenomeChange(uint256 tokenId) view returns(uint256 price)
func (_PolymorphRoot *PolymorphRootSession) PriceForGenomeChange(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphRoot.Contract.PriceForGenomeChange(&_PolymorphRoot.CallOpts, tokenId)
}

// PriceForGenomeChange is a free data retrieval call binding the contract method 0xd45351e5.
//
// Solidity: function priceForGenomeChange(uint256 tokenId) view returns(uint256 price)
func (_PolymorphRoot *PolymorphRootCallerSession) PriceForGenomeChange(tokenId *big.Int) (*big.Int, error) {
	return _PolymorphRoot.Contract.PriceForGenomeChange(&_PolymorphRoot.CallOpts, tokenId)
}

// RandomizeGenomePrice is a free data retrieval call binding the contract method 0xec9c074c.
//
// Solidity: function randomizeGenomePrice() view returns(uint256)
func (_PolymorphRoot *PolymorphRootCaller) RandomizeGenomePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "randomizeGenomePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandomizeGenomePrice is a free data retrieval call binding the contract method 0xec9c074c.
//
// Solidity: function randomizeGenomePrice() view returns(uint256)
func (_PolymorphRoot *PolymorphRootSession) RandomizeGenomePrice() (*big.Int, error) {
	return _PolymorphRoot.Contract.RandomizeGenomePrice(&_PolymorphRoot.CallOpts)
}

// RandomizeGenomePrice is a free data retrieval call binding the contract method 0xec9c074c.
//
// Solidity: function randomizeGenomePrice() view returns(uint256)
func (_PolymorphRoot *PolymorphRootCallerSession) RandomizeGenomePrice() (*big.Int, error) {
	return _PolymorphRoot.Contract.RandomizeGenomePrice(&_PolymorphRoot.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PolymorphRoot *PolymorphRootCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PolymorphRoot *PolymorphRootSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PolymorphRoot.Contract.SupportsInterface(&_PolymorphRoot.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PolymorphRoot *PolymorphRootCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PolymorphRoot.Contract.SupportsInterface(&_PolymorphRoot.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PolymorphRoot *PolymorphRootCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PolymorphRoot *PolymorphRootSession) Symbol() (string, error) {
	return _PolymorphRoot.Contract.Symbol(&_PolymorphRoot.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PolymorphRoot *PolymorphRootCallerSession) Symbol() (string, error) {
	return _PolymorphRoot.Contract.Symbol(&_PolymorphRoot.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_PolymorphRoot *PolymorphRootCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_PolymorphRoot *PolymorphRootSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _PolymorphRoot.Contract.TokenByIndex(&_PolymorphRoot.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_PolymorphRoot *PolymorphRootCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _PolymorphRoot.Contract.TokenByIndex(&_PolymorphRoot.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_PolymorphRoot *PolymorphRootCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_PolymorphRoot *PolymorphRootSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _PolymorphRoot.Contract.TokenOfOwnerByIndex(&_PolymorphRoot.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_PolymorphRoot *PolymorphRootCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _PolymorphRoot.Contract.TokenOfOwnerByIndex(&_PolymorphRoot.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_PolymorphRoot *PolymorphRootCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_PolymorphRoot *PolymorphRootSession) TokenURI(tokenId *big.Int) (string, error) {
	return _PolymorphRoot.Contract.TokenURI(&_PolymorphRoot.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_PolymorphRoot *PolymorphRootCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _PolymorphRoot.Contract.TokenURI(&_PolymorphRoot.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PolymorphRoot *PolymorphRootCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PolymorphRoot *PolymorphRootSession) TotalSupply() (*big.Int, error) {
	return _PolymorphRoot.Contract.TotalSupply(&_PolymorphRoot.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PolymorphRoot *PolymorphRootCallerSession) TotalSupply() (*big.Int, error) {
	return _PolymorphRoot.Contract.TotalSupply(&_PolymorphRoot.CallOpts)
}

// WhitelistTunnelAddresses is a free data retrieval call binding the contract method 0xcccb6d0d.
//
// Solidity: function whitelistTunnelAddresses(address ) view returns(bool)
func (_PolymorphRoot *PolymorphRootCaller) WhitelistTunnelAddresses(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PolymorphRoot.contract.Call(opts, &out, "whitelistTunnelAddresses", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistTunnelAddresses is a free data retrieval call binding the contract method 0xcccb6d0d.
//
// Solidity: function whitelistTunnelAddresses(address ) view returns(bool)
func (_PolymorphRoot *PolymorphRootSession) WhitelistTunnelAddresses(arg0 common.Address) (bool, error) {
	return _PolymorphRoot.Contract.WhitelistTunnelAddresses(&_PolymorphRoot.CallOpts, arg0)
}

// WhitelistTunnelAddresses is a free data retrieval call binding the contract method 0xcccb6d0d.
//
// Solidity: function whitelistTunnelAddresses(address ) view returns(bool)
func (_PolymorphRoot *PolymorphRootCallerSession) WhitelistTunnelAddresses(arg0 common.Address) (bool, error) {
	return _PolymorphRoot.Contract.WhitelistTunnelAddresses(&_PolymorphRoot.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PolymorphRoot *PolymorphRootTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PolymorphRoot *PolymorphRootSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.Approve(&_PolymorphRoot.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.Approve(&_PolymorphRoot.TransactOpts, to, tokenId)
}

// BulkBuy is a paid mutator transaction binding the contract method 0xd5a83d3e.
//
// Solidity: function bulkBuy(uint256 amount) payable returns()
func (_PolymorphRoot *PolymorphRootTransactor) BulkBuy(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "bulkBuy", amount)
}

// BulkBuy is a paid mutator transaction binding the contract method 0xd5a83d3e.
//
// Solidity: function bulkBuy(uint256 amount) payable returns()
func (_PolymorphRoot *PolymorphRootSession) BulkBuy(amount *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.BulkBuy(&_PolymorphRoot.TransactOpts, amount)
}

// BulkBuy is a paid mutator transaction binding the contract method 0xd5a83d3e.
//
// Solidity: function bulkBuy(uint256 amount) payable returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) BulkBuy(amount *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.BulkBuy(&_PolymorphRoot.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_PolymorphRoot *PolymorphRootTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_PolymorphRoot *PolymorphRootSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.Burn(&_PolymorphRoot.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.Burn(&_PolymorphRoot.TransactOpts, tokenId)
}

// BurnAndMintNewPolymorph is a paid mutator transaction binding the contract method 0x572a1070.
//
// Solidity: function burnAndMintNewPolymorph(uint256[] tokenIds) returns()
func (_PolymorphRoot *PolymorphRootTransactor) BurnAndMintNewPolymorph(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "burnAndMintNewPolymorph", tokenIds)
}

// BurnAndMintNewPolymorph is a paid mutator transaction binding the contract method 0x572a1070.
//
// Solidity: function burnAndMintNewPolymorph(uint256[] tokenIds) returns()
func (_PolymorphRoot *PolymorphRootSession) BurnAndMintNewPolymorph(tokenIds []*big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.BurnAndMintNewPolymorph(&_PolymorphRoot.TransactOpts, tokenIds)
}

// BurnAndMintNewPolymorph is a paid mutator transaction binding the contract method 0x572a1070.
//
// Solidity: function burnAndMintNewPolymorph(uint256[] tokenIds) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) BurnAndMintNewPolymorph(tokenIds []*big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.BurnAndMintNewPolymorph(&_PolymorphRoot.TransactOpts, tokenIds)
}

// ChangeBaseGenomeChangePrice is a paid mutator transaction binding the contract method 0x289ea0a9.
//
// Solidity: function changeBaseGenomeChangePrice(uint256 newGenomeChangePrice) returns()
func (_PolymorphRoot *PolymorphRootTransactor) ChangeBaseGenomeChangePrice(opts *bind.TransactOpts, newGenomeChangePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "changeBaseGenomeChangePrice", newGenomeChangePrice)
}

// ChangeBaseGenomeChangePrice is a paid mutator transaction binding the contract method 0x289ea0a9.
//
// Solidity: function changeBaseGenomeChangePrice(uint256 newGenomeChangePrice) returns()
func (_PolymorphRoot *PolymorphRootSession) ChangeBaseGenomeChangePrice(newGenomeChangePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.ChangeBaseGenomeChangePrice(&_PolymorphRoot.TransactOpts, newGenomeChangePrice)
}

// ChangeBaseGenomeChangePrice is a paid mutator transaction binding the contract method 0x289ea0a9.
//
// Solidity: function changeBaseGenomeChangePrice(uint256 newGenomeChangePrice) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) ChangeBaseGenomeChangePrice(newGenomeChangePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.ChangeBaseGenomeChangePrice(&_PolymorphRoot.TransactOpts, newGenomeChangePrice)
}

// ChangeConsumer is a paid mutator transaction binding the contract method 0x70b5aecb.
//
// Solidity: function changeConsumer(address _consumer, uint256 _tokenId) returns()
func (_PolymorphRoot *PolymorphRootTransactor) ChangeConsumer(opts *bind.TransactOpts, _consumer common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "changeConsumer", _consumer, _tokenId)
}

// ChangeConsumer is a paid mutator transaction binding the contract method 0x70b5aecb.
//
// Solidity: function changeConsumer(address _consumer, uint256 _tokenId) returns()
func (_PolymorphRoot *PolymorphRootSession) ChangeConsumer(_consumer common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.ChangeConsumer(&_PolymorphRoot.TransactOpts, _consumer, _tokenId)
}

// ChangeConsumer is a paid mutator transaction binding the contract method 0x70b5aecb.
//
// Solidity: function changeConsumer(address _consumer, uint256 _tokenId) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) ChangeConsumer(_consumer common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.ChangeConsumer(&_PolymorphRoot.TransactOpts, _consumer, _tokenId)
}

// ChangeRandomizeGenomePrice is a paid mutator transaction binding the contract method 0x98c5c078.
//
// Solidity: function changeRandomizeGenomePrice(uint256 newRandomizeGenomePrice) returns()
func (_PolymorphRoot *PolymorphRootTransactor) ChangeRandomizeGenomePrice(opts *bind.TransactOpts, newRandomizeGenomePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "changeRandomizeGenomePrice", newRandomizeGenomePrice)
}

// ChangeRandomizeGenomePrice is a paid mutator transaction binding the contract method 0x98c5c078.
//
// Solidity: function changeRandomizeGenomePrice(uint256 newRandomizeGenomePrice) returns()
func (_PolymorphRoot *PolymorphRootSession) ChangeRandomizeGenomePrice(newRandomizeGenomePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.ChangeRandomizeGenomePrice(&_PolymorphRoot.TransactOpts, newRandomizeGenomePrice)
}

// ChangeRandomizeGenomePrice is a paid mutator transaction binding the contract method 0x98c5c078.
//
// Solidity: function changeRandomizeGenomePrice(uint256 newRandomizeGenomePrice) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) ChangeRandomizeGenomePrice(newRandomizeGenomePrice *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.ChangeRandomizeGenomePrice(&_PolymorphRoot.TransactOpts, newRandomizeGenomePrice)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PolymorphRoot *PolymorphRootTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PolymorphRoot *PolymorphRootSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.GrantRole(&_PolymorphRoot.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.GrantRole(&_PolymorphRoot.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_PolymorphRoot *PolymorphRootTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_PolymorphRoot *PolymorphRootSession) Mint() (*types.Transaction, error) {
	return _PolymorphRoot.Contract.Mint(&_PolymorphRoot.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) Mint() (*types.Transaction, error) {
	return _PolymorphRoot.Contract.Mint(&_PolymorphRoot.TransactOpts)
}

// MorphGene is a paid mutator transaction binding the contract method 0x56a5c926.
//
// Solidity: function morphGene(uint256 tokenId, uint256 genePosition) payable returns()
func (_PolymorphRoot *PolymorphRootTransactor) MorphGene(opts *bind.TransactOpts, tokenId *big.Int, genePosition *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "morphGene", tokenId, genePosition)
}

// MorphGene is a paid mutator transaction binding the contract method 0x56a5c926.
//
// Solidity: function morphGene(uint256 tokenId, uint256 genePosition) payable returns()
func (_PolymorphRoot *PolymorphRootSession) MorphGene(tokenId *big.Int, genePosition *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.MorphGene(&_PolymorphRoot.TransactOpts, tokenId, genePosition)
}

// MorphGene is a paid mutator transaction binding the contract method 0x56a5c926.
//
// Solidity: function morphGene(uint256 tokenId, uint256 genePosition) payable returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) MorphGene(tokenId *big.Int, genePosition *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.MorphGene(&_PolymorphRoot.TransactOpts, tokenId, genePosition)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolymorphRoot *PolymorphRootTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolymorphRoot *PolymorphRootSession) Pause() (*types.Transaction, error) {
	return _PolymorphRoot.Contract.Pause(&_PolymorphRoot.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) Pause() (*types.Transaction, error) {
	return _PolymorphRoot.Contract.Pause(&_PolymorphRoot.TransactOpts)
}

// RandomizeGenome is a paid mutator transaction binding the contract method 0x9e7bb467.
//
// Solidity: function randomizeGenome(uint256 tokenId) payable returns()
func (_PolymorphRoot *PolymorphRootTransactor) RandomizeGenome(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "randomizeGenome", tokenId)
}

// RandomizeGenome is a paid mutator transaction binding the contract method 0x9e7bb467.
//
// Solidity: function randomizeGenome(uint256 tokenId) payable returns()
func (_PolymorphRoot *PolymorphRootSession) RandomizeGenome(tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.RandomizeGenome(&_PolymorphRoot.TransactOpts, tokenId)
}

// RandomizeGenome is a paid mutator transaction binding the contract method 0x9e7bb467.
//
// Solidity: function randomizeGenome(uint256 tokenId) payable returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) RandomizeGenome(tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.RandomizeGenome(&_PolymorphRoot.TransactOpts, tokenId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PolymorphRoot *PolymorphRootTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PolymorphRoot *PolymorphRootSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.RenounceRole(&_PolymorphRoot.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.RenounceRole(&_PolymorphRoot.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PolymorphRoot *PolymorphRootTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PolymorphRoot *PolymorphRootSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.RevokeRole(&_PolymorphRoot.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.RevokeRole(&_PolymorphRoot.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphRoot *PolymorphRootTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphRoot *PolymorphRootSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SafeTransferFrom(&_PolymorphRoot.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SafeTransferFrom(&_PolymorphRoot.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_PolymorphRoot *PolymorphRootTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_PolymorphRoot *PolymorphRootSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SafeTransferFrom0(&_PolymorphRoot.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SafeTransferFrom0(&_PolymorphRoot.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PolymorphRoot *PolymorphRootTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PolymorphRoot *PolymorphRootSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SetApprovalForAll(&_PolymorphRoot.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SetApprovalForAll(&_PolymorphRoot.TransactOpts, operator, approved)
}

// SetArweaveAssetsJSON is a paid mutator transaction binding the contract method 0x56b1b300.
//
// Solidity: function setArweaveAssetsJSON(string _arweaveAssetsJSON) returns()
func (_PolymorphRoot *PolymorphRootTransactor) SetArweaveAssetsJSON(opts *bind.TransactOpts, _arweaveAssetsJSON string) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "setArweaveAssetsJSON", _arweaveAssetsJSON)
}

// SetArweaveAssetsJSON is a paid mutator transaction binding the contract method 0x56b1b300.
//
// Solidity: function setArweaveAssetsJSON(string _arweaveAssetsJSON) returns()
func (_PolymorphRoot *PolymorphRootSession) SetArweaveAssetsJSON(_arweaveAssetsJSON string) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SetArweaveAssetsJSON(&_PolymorphRoot.TransactOpts, _arweaveAssetsJSON)
}

// SetArweaveAssetsJSON is a paid mutator transaction binding the contract method 0x56b1b300.
//
// Solidity: function setArweaveAssetsJSON(string _arweaveAssetsJSON) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) SetArweaveAssetsJSON(_arweaveAssetsJSON string) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SetArweaveAssetsJSON(&_PolymorphRoot.TransactOpts, _arweaveAssetsJSON)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_PolymorphRoot *PolymorphRootTransactor) SetBaseURI(opts *bind.TransactOpts, _baseURI string) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "setBaseURI", _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_PolymorphRoot *PolymorphRootSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SetBaseURI(&_PolymorphRoot.TransactOpts, _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SetBaseURI(&_PolymorphRoot.TransactOpts, _baseURI)
}

// SetBulkBuyLimit is a paid mutator transaction binding the contract method 0x5e468dfd.
//
// Solidity: function setBulkBuyLimit(uint256 _bulkBuyLimit) returns()
func (_PolymorphRoot *PolymorphRootTransactor) SetBulkBuyLimit(opts *bind.TransactOpts, _bulkBuyLimit *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "setBulkBuyLimit", _bulkBuyLimit)
}

// SetBulkBuyLimit is a paid mutator transaction binding the contract method 0x5e468dfd.
//
// Solidity: function setBulkBuyLimit(uint256 _bulkBuyLimit) returns()
func (_PolymorphRoot *PolymorphRootSession) SetBulkBuyLimit(_bulkBuyLimit *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SetBulkBuyLimit(&_PolymorphRoot.TransactOpts, _bulkBuyLimit)
}

// SetBulkBuyLimit is a paid mutator transaction binding the contract method 0x5e468dfd.
//
// Solidity: function setBulkBuyLimit(uint256 _bulkBuyLimit) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) SetBulkBuyLimit(_bulkBuyLimit *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SetBulkBuyLimit(&_PolymorphRoot.TransactOpts, _bulkBuyLimit)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_PolymorphRoot *PolymorphRootTransactor) SetMaxSupply(opts *bind.TransactOpts, _maxSupply *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "setMaxSupply", _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_PolymorphRoot *PolymorphRootSession) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SetMaxSupply(&_PolymorphRoot.TransactOpts, _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SetMaxSupply(&_PolymorphRoot.TransactOpts, _maxSupply)
}

// SetPolymorphPrice is a paid mutator transaction binding the contract method 0x017f1e34.
//
// Solidity: function setPolymorphPrice(uint256 newPolymorphPrice) returns()
func (_PolymorphRoot *PolymorphRootTransactor) SetPolymorphPrice(opts *bind.TransactOpts, newPolymorphPrice *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "setPolymorphPrice", newPolymorphPrice)
}

// SetPolymorphPrice is a paid mutator transaction binding the contract method 0x017f1e34.
//
// Solidity: function setPolymorphPrice(uint256 newPolymorphPrice) returns()
func (_PolymorphRoot *PolymorphRootSession) SetPolymorphPrice(newPolymorphPrice *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SetPolymorphPrice(&_PolymorphRoot.TransactOpts, newPolymorphPrice)
}

// SetPolymorphPrice is a paid mutator transaction binding the contract method 0x017f1e34.
//
// Solidity: function setPolymorphPrice(uint256 newPolymorphPrice) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) SetPolymorphPrice(newPolymorphPrice *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.SetPolymorphPrice(&_PolymorphRoot.TransactOpts, newPolymorphPrice)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphRoot *PolymorphRootTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphRoot *PolymorphRootSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.TransferFrom(&_PolymorphRoot.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.TransferFrom(&_PolymorphRoot.TransactOpts, from, to, tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolymorphRoot *PolymorphRootTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolymorphRoot *PolymorphRootSession) Unpause() (*types.Transaction, error) {
	return _PolymorphRoot.Contract.Unpause(&_PolymorphRoot.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) Unpause() (*types.Transaction, error) {
	return _PolymorphRoot.Contract.Unpause(&_PolymorphRoot.TransactOpts)
}

// WhitelistBridgeAddress is a paid mutator transaction binding the contract method 0xab39a3c8.
//
// Solidity: function whitelistBridgeAddress(address bridgeAddress, bool status) returns()
func (_PolymorphRoot *PolymorphRootTransactor) WhitelistBridgeAddress(opts *bind.TransactOpts, bridgeAddress common.Address, status bool) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "whitelistBridgeAddress", bridgeAddress, status)
}

// WhitelistBridgeAddress is a paid mutator transaction binding the contract method 0xab39a3c8.
//
// Solidity: function whitelistBridgeAddress(address bridgeAddress, bool status) returns()
func (_PolymorphRoot *PolymorphRootSession) WhitelistBridgeAddress(bridgeAddress common.Address, status bool) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.WhitelistBridgeAddress(&_PolymorphRoot.TransactOpts, bridgeAddress, status)
}

// WhitelistBridgeAddress is a paid mutator transaction binding the contract method 0xab39a3c8.
//
// Solidity: function whitelistBridgeAddress(address bridgeAddress, bool status) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) WhitelistBridgeAddress(bridgeAddress common.Address, status bool) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.WhitelistBridgeAddress(&_PolymorphRoot.TransactOpts, bridgeAddress, status)
}

// WormholeUpdateGene is a paid mutator transaction binding the contract method 0x6a1c03dc.
//
// Solidity: function wormholeUpdateGene(uint256 tokenId, uint256 gene, bool isVirgin, uint256 genomeChangesCount) returns()
func (_PolymorphRoot *PolymorphRootTransactor) WormholeUpdateGene(opts *bind.TransactOpts, tokenId *big.Int, gene *big.Int, isVirgin bool, genomeChangesCount *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.contract.Transact(opts, "wormholeUpdateGene", tokenId, gene, isVirgin, genomeChangesCount)
}

// WormholeUpdateGene is a paid mutator transaction binding the contract method 0x6a1c03dc.
//
// Solidity: function wormholeUpdateGene(uint256 tokenId, uint256 gene, bool isVirgin, uint256 genomeChangesCount) returns()
func (_PolymorphRoot *PolymorphRootSession) WormholeUpdateGene(tokenId *big.Int, gene *big.Int, isVirgin bool, genomeChangesCount *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.WormholeUpdateGene(&_PolymorphRoot.TransactOpts, tokenId, gene, isVirgin, genomeChangesCount)
}

// WormholeUpdateGene is a paid mutator transaction binding the contract method 0x6a1c03dc.
//
// Solidity: function wormholeUpdateGene(uint256 tokenId, uint256 gene, bool isVirgin, uint256 genomeChangesCount) returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) WormholeUpdateGene(tokenId *big.Int, gene *big.Int, isVirgin bool, genomeChangesCount *big.Int) (*types.Transaction, error) {
	return _PolymorphRoot.Contract.WormholeUpdateGene(&_PolymorphRoot.TransactOpts, tokenId, gene, isVirgin, genomeChangesCount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PolymorphRoot *PolymorphRootTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolymorphRoot.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PolymorphRoot *PolymorphRootSession) Receive() (*types.Transaction, error) {
	return _PolymorphRoot.Contract.Receive(&_PolymorphRoot.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PolymorphRoot *PolymorphRootTransactorSession) Receive() (*types.Transaction, error) {
	return _PolymorphRoot.Contract.Receive(&_PolymorphRoot.TransactOpts)
}

// PolymorphRootApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PolymorphRoot contract.
type PolymorphRootApprovalIterator struct {
	Event *PolymorphRootApproval // Event containing the contract specifics and raw log

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
func (it *PolymorphRootApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootApproval)
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
		it.Event = new(PolymorphRootApproval)
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
func (it *PolymorphRootApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootApproval represents a Approval event raised by the PolymorphRoot contract.
type PolymorphRootApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PolymorphRoot *PolymorphRootFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*PolymorphRootApprovalIterator, error) {

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

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphRootApprovalIterator{contract: _PolymorphRoot.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PolymorphRoot *PolymorphRootFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PolymorphRootApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootApproval)
				if err := _PolymorphRoot.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_PolymorphRoot *PolymorphRootFilterer) ParseApproval(log types.Log) (*PolymorphRootApproval, error) {
	event := new(PolymorphRootApproval)
	if err := _PolymorphRoot.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the PolymorphRoot contract.
type PolymorphRootApprovalForAllIterator struct {
	Event *PolymorphRootApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PolymorphRootApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootApprovalForAll)
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
		it.Event = new(PolymorphRootApprovalForAll)
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
func (it *PolymorphRootApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootApprovalForAll represents a ApprovalForAll event raised by the PolymorphRoot contract.
type PolymorphRootApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PolymorphRoot *PolymorphRootFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PolymorphRootApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphRootApprovalForAllIterator{contract: _PolymorphRoot.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PolymorphRoot *PolymorphRootFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PolymorphRootApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootApprovalForAll)
				if err := _PolymorphRoot.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_PolymorphRoot *PolymorphRootFilterer) ParseApprovalForAll(log types.Log) (*PolymorphRootApprovalForAll, error) {
	event := new(PolymorphRootApprovalForAll)
	if err := _PolymorphRoot.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootArweaveAssetsJSONChangedIterator is returned from FilterArweaveAssetsJSONChanged and is used to iterate over the raw logs and unpacked data for ArweaveAssetsJSONChanged events raised by the PolymorphRoot contract.
type PolymorphRootArweaveAssetsJSONChangedIterator struct {
	Event *PolymorphRootArweaveAssetsJSONChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphRootArweaveAssetsJSONChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootArweaveAssetsJSONChanged)
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
		it.Event = new(PolymorphRootArweaveAssetsJSONChanged)
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
func (it *PolymorphRootArweaveAssetsJSONChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootArweaveAssetsJSONChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootArweaveAssetsJSONChanged represents a ArweaveAssetsJSONChanged event raised by the PolymorphRoot contract.
type PolymorphRootArweaveAssetsJSONChanged struct {
	ArweaveAssetsJSON string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterArweaveAssetsJSONChanged is a free log retrieval operation binding the contract event 0x4a826ca029d05af64e411551e15f7ee1e70af0b9bc43a31154ace86a863397b4.
//
// Solidity: event ArweaveAssetsJSONChanged(string arweaveAssetsJSON)
func (_PolymorphRoot *PolymorphRootFilterer) FilterArweaveAssetsJSONChanged(opts *bind.FilterOpts) (*PolymorphRootArweaveAssetsJSONChangedIterator, error) {

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "ArweaveAssetsJSONChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphRootArweaveAssetsJSONChangedIterator{contract: _PolymorphRoot.contract, event: "ArweaveAssetsJSONChanged", logs: logs, sub: sub}, nil
}

// WatchArweaveAssetsJSONChanged is a free log subscription operation binding the contract event 0x4a826ca029d05af64e411551e15f7ee1e70af0b9bc43a31154ace86a863397b4.
//
// Solidity: event ArweaveAssetsJSONChanged(string arweaveAssetsJSON)
func (_PolymorphRoot *PolymorphRootFilterer) WatchArweaveAssetsJSONChanged(opts *bind.WatchOpts, sink chan<- *PolymorphRootArweaveAssetsJSONChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "ArweaveAssetsJSONChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootArweaveAssetsJSONChanged)
				if err := _PolymorphRoot.contract.UnpackLog(event, "ArweaveAssetsJSONChanged", log); err != nil {
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

// ParseArweaveAssetsJSONChanged is a log parse operation binding the contract event 0x4a826ca029d05af64e411551e15f7ee1e70af0b9bc43a31154ace86a863397b4.
//
// Solidity: event ArweaveAssetsJSONChanged(string arweaveAssetsJSON)
func (_PolymorphRoot *PolymorphRootFilterer) ParseArweaveAssetsJSONChanged(log types.Log) (*PolymorphRootArweaveAssetsJSONChanged, error) {
	event := new(PolymorphRootArweaveAssetsJSONChanged)
	if err := _PolymorphRoot.contract.UnpackLog(event, "ArweaveAssetsJSONChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootBaseGenomeChangePriceChangedIterator is returned from FilterBaseGenomeChangePriceChanged and is used to iterate over the raw logs and unpacked data for BaseGenomeChangePriceChanged events raised by the PolymorphRoot contract.
type PolymorphRootBaseGenomeChangePriceChangedIterator struct {
	Event *PolymorphRootBaseGenomeChangePriceChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphRootBaseGenomeChangePriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootBaseGenomeChangePriceChanged)
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
		it.Event = new(PolymorphRootBaseGenomeChangePriceChanged)
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
func (it *PolymorphRootBaseGenomeChangePriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootBaseGenomeChangePriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootBaseGenomeChangePriceChanged represents a BaseGenomeChangePriceChanged event raised by the PolymorphRoot contract.
type PolymorphRootBaseGenomeChangePriceChanged struct {
	NewGenomeChange *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBaseGenomeChangePriceChanged is a free log retrieval operation binding the contract event 0xb1d78271daba9a366098d40b64d642a1399cabaa22c5234bacc87e92cef82ae6.
//
// Solidity: event BaseGenomeChangePriceChanged(uint256 newGenomeChange)
func (_PolymorphRoot *PolymorphRootFilterer) FilterBaseGenomeChangePriceChanged(opts *bind.FilterOpts) (*PolymorphRootBaseGenomeChangePriceChangedIterator, error) {

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "BaseGenomeChangePriceChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphRootBaseGenomeChangePriceChangedIterator{contract: _PolymorphRoot.contract, event: "BaseGenomeChangePriceChanged", logs: logs, sub: sub}, nil
}

// WatchBaseGenomeChangePriceChanged is a free log subscription operation binding the contract event 0xb1d78271daba9a366098d40b64d642a1399cabaa22c5234bacc87e92cef82ae6.
//
// Solidity: event BaseGenomeChangePriceChanged(uint256 newGenomeChange)
func (_PolymorphRoot *PolymorphRootFilterer) WatchBaseGenomeChangePriceChanged(opts *bind.WatchOpts, sink chan<- *PolymorphRootBaseGenomeChangePriceChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "BaseGenomeChangePriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootBaseGenomeChangePriceChanged)
				if err := _PolymorphRoot.contract.UnpackLog(event, "BaseGenomeChangePriceChanged", log); err != nil {
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

// ParseBaseGenomeChangePriceChanged is a log parse operation binding the contract event 0xb1d78271daba9a366098d40b64d642a1399cabaa22c5234bacc87e92cef82ae6.
//
// Solidity: event BaseGenomeChangePriceChanged(uint256 newGenomeChange)
func (_PolymorphRoot *PolymorphRootFilterer) ParseBaseGenomeChangePriceChanged(log types.Log) (*PolymorphRootBaseGenomeChangePriceChanged, error) {
	event := new(PolymorphRootBaseGenomeChangePriceChanged)
	if err := _PolymorphRoot.contract.UnpackLog(event, "BaseGenomeChangePriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootBaseURIChangedIterator is returned from FilterBaseURIChanged and is used to iterate over the raw logs and unpacked data for BaseURIChanged events raised by the PolymorphRoot contract.
type PolymorphRootBaseURIChangedIterator struct {
	Event *PolymorphRootBaseURIChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphRootBaseURIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootBaseURIChanged)
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
		it.Event = new(PolymorphRootBaseURIChanged)
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
func (it *PolymorphRootBaseURIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootBaseURIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootBaseURIChanged represents a BaseURIChanged event raised by the PolymorphRoot contract.
type PolymorphRootBaseURIChanged struct {
	BaseURI string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBaseURIChanged is a free log retrieval operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_PolymorphRoot *PolymorphRootFilterer) FilterBaseURIChanged(opts *bind.FilterOpts) (*PolymorphRootBaseURIChangedIterator, error) {

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "BaseURIChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphRootBaseURIChangedIterator{contract: _PolymorphRoot.contract, event: "BaseURIChanged", logs: logs, sub: sub}, nil
}

// WatchBaseURIChanged is a free log subscription operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_PolymorphRoot *PolymorphRootFilterer) WatchBaseURIChanged(opts *bind.WatchOpts, sink chan<- *PolymorphRootBaseURIChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "BaseURIChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootBaseURIChanged)
				if err := _PolymorphRoot.contract.UnpackLog(event, "BaseURIChanged", log); err != nil {
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

// ParseBaseURIChanged is a log parse operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_PolymorphRoot *PolymorphRootFilterer) ParseBaseURIChanged(log types.Log) (*PolymorphRootBaseURIChanged, error) {
	event := new(PolymorphRootBaseURIChanged)
	if err := _PolymorphRoot.contract.UnpackLog(event, "BaseURIChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootBulkBuyLimitChangedIterator is returned from FilterBulkBuyLimitChanged and is used to iterate over the raw logs and unpacked data for BulkBuyLimitChanged events raised by the PolymorphRoot contract.
type PolymorphRootBulkBuyLimitChangedIterator struct {
	Event *PolymorphRootBulkBuyLimitChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphRootBulkBuyLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootBulkBuyLimitChanged)
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
		it.Event = new(PolymorphRootBulkBuyLimitChanged)
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
func (it *PolymorphRootBulkBuyLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootBulkBuyLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootBulkBuyLimitChanged represents a BulkBuyLimitChanged event raised by the PolymorphRoot contract.
type PolymorphRootBulkBuyLimitChanged struct {
	NewBulkBuyLimit *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBulkBuyLimitChanged is a free log retrieval operation binding the contract event 0xa0e0113404674c6f545b966e8ec54db3066a6c720a0054f0bc4b0c900cfff243.
//
// Solidity: event BulkBuyLimitChanged(uint256 newBulkBuyLimit)
func (_PolymorphRoot *PolymorphRootFilterer) FilterBulkBuyLimitChanged(opts *bind.FilterOpts) (*PolymorphRootBulkBuyLimitChangedIterator, error) {

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "BulkBuyLimitChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphRootBulkBuyLimitChangedIterator{contract: _PolymorphRoot.contract, event: "BulkBuyLimitChanged", logs: logs, sub: sub}, nil
}

// WatchBulkBuyLimitChanged is a free log subscription operation binding the contract event 0xa0e0113404674c6f545b966e8ec54db3066a6c720a0054f0bc4b0c900cfff243.
//
// Solidity: event BulkBuyLimitChanged(uint256 newBulkBuyLimit)
func (_PolymorphRoot *PolymorphRootFilterer) WatchBulkBuyLimitChanged(opts *bind.WatchOpts, sink chan<- *PolymorphRootBulkBuyLimitChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "BulkBuyLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootBulkBuyLimitChanged)
				if err := _PolymorphRoot.contract.UnpackLog(event, "BulkBuyLimitChanged", log); err != nil {
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

// ParseBulkBuyLimitChanged is a log parse operation binding the contract event 0xa0e0113404674c6f545b966e8ec54db3066a6c720a0054f0bc4b0c900cfff243.
//
// Solidity: event BulkBuyLimitChanged(uint256 newBulkBuyLimit)
func (_PolymorphRoot *PolymorphRootFilterer) ParseBulkBuyLimitChanged(log types.Log) (*PolymorphRootBulkBuyLimitChanged, error) {
	event := new(PolymorphRootBulkBuyLimitChanged)
	if err := _PolymorphRoot.contract.UnpackLog(event, "BulkBuyLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootConsumerChangedIterator is returned from FilterConsumerChanged and is used to iterate over the raw logs and unpacked data for ConsumerChanged events raised by the PolymorphRoot contract.
type PolymorphRootConsumerChangedIterator struct {
	Event *PolymorphRootConsumerChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphRootConsumerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootConsumerChanged)
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
		it.Event = new(PolymorphRootConsumerChanged)
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
func (it *PolymorphRootConsumerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootConsumerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootConsumerChanged represents a ConsumerChanged event raised by the PolymorphRoot contract.
type PolymorphRootConsumerChanged struct {
	Owner    common.Address
	Consumer common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConsumerChanged is a free log retrieval operation binding the contract event 0x42ef856c2602f37ce625d252830bed486c5c8e9a4de8aa36cc3d15f304eb662b.
//
// Solidity: event ConsumerChanged(address indexed owner, address indexed consumer, uint256 indexed tokenId)
func (_PolymorphRoot *PolymorphRootFilterer) FilterConsumerChanged(opts *bind.FilterOpts, owner []common.Address, consumer []common.Address, tokenId []*big.Int) (*PolymorphRootConsumerChangedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "ConsumerChanged", ownerRule, consumerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphRootConsumerChangedIterator{contract: _PolymorphRoot.contract, event: "ConsumerChanged", logs: logs, sub: sub}, nil
}

// WatchConsumerChanged is a free log subscription operation binding the contract event 0x42ef856c2602f37ce625d252830bed486c5c8e9a4de8aa36cc3d15f304eb662b.
//
// Solidity: event ConsumerChanged(address indexed owner, address indexed consumer, uint256 indexed tokenId)
func (_PolymorphRoot *PolymorphRootFilterer) WatchConsumerChanged(opts *bind.WatchOpts, sink chan<- *PolymorphRootConsumerChanged, owner []common.Address, consumer []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "ConsumerChanged", ownerRule, consumerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootConsumerChanged)
				if err := _PolymorphRoot.contract.UnpackLog(event, "ConsumerChanged", log); err != nil {
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

// ParseConsumerChanged is a log parse operation binding the contract event 0x42ef856c2602f37ce625d252830bed486c5c8e9a4de8aa36cc3d15f304eb662b.
//
// Solidity: event ConsumerChanged(address indexed owner, address indexed consumer, uint256 indexed tokenId)
func (_PolymorphRoot *PolymorphRootFilterer) ParseConsumerChanged(log types.Log) (*PolymorphRootConsumerChanged, error) {
	event := new(PolymorphRootConsumerChanged)
	if err := _PolymorphRoot.contract.UnpackLog(event, "ConsumerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootMaxSupplyChangedIterator is returned from FilterMaxSupplyChanged and is used to iterate over the raw logs and unpacked data for MaxSupplyChanged events raised by the PolymorphRoot contract.
type PolymorphRootMaxSupplyChangedIterator struct {
	Event *PolymorphRootMaxSupplyChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphRootMaxSupplyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootMaxSupplyChanged)
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
		it.Event = new(PolymorphRootMaxSupplyChanged)
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
func (it *PolymorphRootMaxSupplyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootMaxSupplyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootMaxSupplyChanged represents a MaxSupplyChanged event raised by the PolymorphRoot contract.
type PolymorphRootMaxSupplyChanged struct {
	NewMaxSupply *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMaxSupplyChanged is a free log retrieval operation binding the contract event 0x28a10a2e0b5582da7164754cb994f6214b8af6aa7f7e003305fbc09e7106c513.
//
// Solidity: event MaxSupplyChanged(uint256 newMaxSupply)
func (_PolymorphRoot *PolymorphRootFilterer) FilterMaxSupplyChanged(opts *bind.FilterOpts) (*PolymorphRootMaxSupplyChangedIterator, error) {

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "MaxSupplyChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphRootMaxSupplyChangedIterator{contract: _PolymorphRoot.contract, event: "MaxSupplyChanged", logs: logs, sub: sub}, nil
}

// WatchMaxSupplyChanged is a free log subscription operation binding the contract event 0x28a10a2e0b5582da7164754cb994f6214b8af6aa7f7e003305fbc09e7106c513.
//
// Solidity: event MaxSupplyChanged(uint256 newMaxSupply)
func (_PolymorphRoot *PolymorphRootFilterer) WatchMaxSupplyChanged(opts *bind.WatchOpts, sink chan<- *PolymorphRootMaxSupplyChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "MaxSupplyChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootMaxSupplyChanged)
				if err := _PolymorphRoot.contract.UnpackLog(event, "MaxSupplyChanged", log); err != nil {
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

// ParseMaxSupplyChanged is a log parse operation binding the contract event 0x28a10a2e0b5582da7164754cb994f6214b8af6aa7f7e003305fbc09e7106c513.
//
// Solidity: event MaxSupplyChanged(uint256 newMaxSupply)
func (_PolymorphRoot *PolymorphRootFilterer) ParseMaxSupplyChanged(log types.Log) (*PolymorphRootMaxSupplyChanged, error) {
	event := new(PolymorphRootMaxSupplyChanged)
	if err := _PolymorphRoot.contract.UnpackLog(event, "MaxSupplyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PolymorphRoot contract.
type PolymorphRootPausedIterator struct {
	Event *PolymorphRootPaused // Event containing the contract specifics and raw log

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
func (it *PolymorphRootPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootPaused)
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
		it.Event = new(PolymorphRootPaused)
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
func (it *PolymorphRootPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootPaused represents a Paused event raised by the PolymorphRoot contract.
type PolymorphRootPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PolymorphRoot *PolymorphRootFilterer) FilterPaused(opts *bind.FilterOpts) (*PolymorphRootPausedIterator, error) {

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PolymorphRootPausedIterator{contract: _PolymorphRoot.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PolymorphRoot *PolymorphRootFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PolymorphRootPaused) (event.Subscription, error) {

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootPaused)
				if err := _PolymorphRoot.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PolymorphRoot *PolymorphRootFilterer) ParsePaused(log types.Log) (*PolymorphRootPaused, error) {
	event := new(PolymorphRootPaused)
	if err := _PolymorphRoot.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootPolymorphPriceChangedIterator is returned from FilterPolymorphPriceChanged and is used to iterate over the raw logs and unpacked data for PolymorphPriceChanged events raised by the PolymorphRoot contract.
type PolymorphRootPolymorphPriceChangedIterator struct {
	Event *PolymorphRootPolymorphPriceChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphRootPolymorphPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootPolymorphPriceChanged)
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
		it.Event = new(PolymorphRootPolymorphPriceChanged)
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
func (it *PolymorphRootPolymorphPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootPolymorphPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootPolymorphPriceChanged represents a PolymorphPriceChanged event raised by the PolymorphRoot contract.
type PolymorphRootPolymorphPriceChanged struct {
	NewPolymorphPrice *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPolymorphPriceChanged is a free log retrieval operation binding the contract event 0x6a08b3bba14e54ee218389c7c7444e619f3897465dc06757938cfd01a6957f6c.
//
// Solidity: event PolymorphPriceChanged(uint256 newPolymorphPrice)
func (_PolymorphRoot *PolymorphRootFilterer) FilterPolymorphPriceChanged(opts *bind.FilterOpts) (*PolymorphRootPolymorphPriceChangedIterator, error) {

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "PolymorphPriceChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphRootPolymorphPriceChangedIterator{contract: _PolymorphRoot.contract, event: "PolymorphPriceChanged", logs: logs, sub: sub}, nil
}

// WatchPolymorphPriceChanged is a free log subscription operation binding the contract event 0x6a08b3bba14e54ee218389c7c7444e619f3897465dc06757938cfd01a6957f6c.
//
// Solidity: event PolymorphPriceChanged(uint256 newPolymorphPrice)
func (_PolymorphRoot *PolymorphRootFilterer) WatchPolymorphPriceChanged(opts *bind.WatchOpts, sink chan<- *PolymorphRootPolymorphPriceChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "PolymorphPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootPolymorphPriceChanged)
				if err := _PolymorphRoot.contract.UnpackLog(event, "PolymorphPriceChanged", log); err != nil {
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

// ParsePolymorphPriceChanged is a log parse operation binding the contract event 0x6a08b3bba14e54ee218389c7c7444e619f3897465dc06757938cfd01a6957f6c.
//
// Solidity: event PolymorphPriceChanged(uint256 newPolymorphPrice)
func (_PolymorphRoot *PolymorphRootFilterer) ParsePolymorphPriceChanged(log types.Log) (*PolymorphRootPolymorphPriceChanged, error) {
	event := new(PolymorphRootPolymorphPriceChanged)
	if err := _PolymorphRoot.contract.UnpackLog(event, "PolymorphPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootRandomizeGenomePriceChangedIterator is returned from FilterRandomizeGenomePriceChanged and is used to iterate over the raw logs and unpacked data for RandomizeGenomePriceChanged events raised by the PolymorphRoot contract.
type PolymorphRootRandomizeGenomePriceChangedIterator struct {
	Event *PolymorphRootRandomizeGenomePriceChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphRootRandomizeGenomePriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootRandomizeGenomePriceChanged)
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
		it.Event = new(PolymorphRootRandomizeGenomePriceChanged)
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
func (it *PolymorphRootRandomizeGenomePriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootRandomizeGenomePriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootRandomizeGenomePriceChanged represents a RandomizeGenomePriceChanged event raised by the PolymorphRoot contract.
type PolymorphRootRandomizeGenomePriceChanged struct {
	NewRandomizeGenomePriceChange *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterRandomizeGenomePriceChanged is a free log retrieval operation binding the contract event 0xff4da8d01e7184cc8c9d6c57d64b336b1de6d676b6215408967bd071c8da7e3d.
//
// Solidity: event RandomizeGenomePriceChanged(uint256 newRandomizeGenomePriceChange)
func (_PolymorphRoot *PolymorphRootFilterer) FilterRandomizeGenomePriceChanged(opts *bind.FilterOpts) (*PolymorphRootRandomizeGenomePriceChangedIterator, error) {

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "RandomizeGenomePriceChanged")
	if err != nil {
		return nil, err
	}
	return &PolymorphRootRandomizeGenomePriceChangedIterator{contract: _PolymorphRoot.contract, event: "RandomizeGenomePriceChanged", logs: logs, sub: sub}, nil
}

// WatchRandomizeGenomePriceChanged is a free log subscription operation binding the contract event 0xff4da8d01e7184cc8c9d6c57d64b336b1de6d676b6215408967bd071c8da7e3d.
//
// Solidity: event RandomizeGenomePriceChanged(uint256 newRandomizeGenomePriceChange)
func (_PolymorphRoot *PolymorphRootFilterer) WatchRandomizeGenomePriceChanged(opts *bind.WatchOpts, sink chan<- *PolymorphRootRandomizeGenomePriceChanged) (event.Subscription, error) {

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "RandomizeGenomePriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootRandomizeGenomePriceChanged)
				if err := _PolymorphRoot.contract.UnpackLog(event, "RandomizeGenomePriceChanged", log); err != nil {
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

// ParseRandomizeGenomePriceChanged is a log parse operation binding the contract event 0xff4da8d01e7184cc8c9d6c57d64b336b1de6d676b6215408967bd071c8da7e3d.
//
// Solidity: event RandomizeGenomePriceChanged(uint256 newRandomizeGenomePriceChange)
func (_PolymorphRoot *PolymorphRootFilterer) ParseRandomizeGenomePriceChanged(log types.Log) (*PolymorphRootRandomizeGenomePriceChanged, error) {
	event := new(PolymorphRootRandomizeGenomePriceChanged)
	if err := _PolymorphRoot.contract.UnpackLog(event, "RandomizeGenomePriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PolymorphRoot contract.
type PolymorphRootRoleAdminChangedIterator struct {
	Event *PolymorphRootRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PolymorphRootRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootRoleAdminChanged)
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
		it.Event = new(PolymorphRootRoleAdminChanged)
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
func (it *PolymorphRootRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootRoleAdminChanged represents a RoleAdminChanged event raised by the PolymorphRoot contract.
type PolymorphRootRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PolymorphRoot *PolymorphRootFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PolymorphRootRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphRootRoleAdminChangedIterator{contract: _PolymorphRoot.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PolymorphRoot *PolymorphRootFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PolymorphRootRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootRoleAdminChanged)
				if err := _PolymorphRoot.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_PolymorphRoot *PolymorphRootFilterer) ParseRoleAdminChanged(log types.Log) (*PolymorphRootRoleAdminChanged, error) {
	event := new(PolymorphRootRoleAdminChanged)
	if err := _PolymorphRoot.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PolymorphRoot contract.
type PolymorphRootRoleGrantedIterator struct {
	Event *PolymorphRootRoleGranted // Event containing the contract specifics and raw log

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
func (it *PolymorphRootRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootRoleGranted)
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
		it.Event = new(PolymorphRootRoleGranted)
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
func (it *PolymorphRootRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootRoleGranted represents a RoleGranted event raised by the PolymorphRoot contract.
type PolymorphRootRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphRoot *PolymorphRootFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolymorphRootRoleGrantedIterator, error) {

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

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphRootRoleGrantedIterator{contract: _PolymorphRoot.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphRoot *PolymorphRootFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PolymorphRootRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootRoleGranted)
				if err := _PolymorphRoot.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_PolymorphRoot *PolymorphRootFilterer) ParseRoleGranted(log types.Log) (*PolymorphRootRoleGranted, error) {
	event := new(PolymorphRootRoleGranted)
	if err := _PolymorphRoot.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PolymorphRoot contract.
type PolymorphRootRoleRevokedIterator struct {
	Event *PolymorphRootRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PolymorphRootRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootRoleRevoked)
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
		it.Event = new(PolymorphRootRoleRevoked)
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
func (it *PolymorphRootRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootRoleRevoked represents a RoleRevoked event raised by the PolymorphRoot contract.
type PolymorphRootRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphRoot *PolymorphRootFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolymorphRootRoleRevokedIterator, error) {

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

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphRootRoleRevokedIterator{contract: _PolymorphRoot.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PolymorphRoot *PolymorphRootFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PolymorphRootRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootRoleRevoked)
				if err := _PolymorphRoot.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_PolymorphRoot *PolymorphRootFilterer) ParseRoleRevoked(log types.Log) (*PolymorphRootRoleRevoked, error) {
	event := new(PolymorphRootRoleRevoked)
	if err := _PolymorphRoot.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootTokenBurnedAndMintedIterator is returned from FilterTokenBurnedAndMinted and is used to iterate over the raw logs and unpacked data for TokenBurnedAndMinted events raised by the PolymorphRoot contract.
type PolymorphRootTokenBurnedAndMintedIterator struct {
	Event *PolymorphRootTokenBurnedAndMinted // Event containing the contract specifics and raw log

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
func (it *PolymorphRootTokenBurnedAndMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootTokenBurnedAndMinted)
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
		it.Event = new(PolymorphRootTokenBurnedAndMinted)
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
func (it *PolymorphRootTokenBurnedAndMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootTokenBurnedAndMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootTokenBurnedAndMinted represents a TokenBurnedAndMinted event raised by the PolymorphRoot contract.
type PolymorphRootTokenBurnedAndMinted struct {
	TokenId *big.Int
	Gene    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenBurnedAndMinted is a free log retrieval operation binding the contract event 0xb583e10e1b691f0fec5bb6e5bafd39038a17365ca548d3beb76c082261805583.
//
// Solidity: event TokenBurnedAndMinted(uint256 indexed tokenId, uint256 gene)
func (_PolymorphRoot *PolymorphRootFilterer) FilterTokenBurnedAndMinted(opts *bind.FilterOpts, tokenId []*big.Int) (*PolymorphRootTokenBurnedAndMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "TokenBurnedAndMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphRootTokenBurnedAndMintedIterator{contract: _PolymorphRoot.contract, event: "TokenBurnedAndMinted", logs: logs, sub: sub}, nil
}

// WatchTokenBurnedAndMinted is a free log subscription operation binding the contract event 0xb583e10e1b691f0fec5bb6e5bafd39038a17365ca548d3beb76c082261805583.
//
// Solidity: event TokenBurnedAndMinted(uint256 indexed tokenId, uint256 gene)
func (_PolymorphRoot *PolymorphRootFilterer) WatchTokenBurnedAndMinted(opts *bind.WatchOpts, sink chan<- *PolymorphRootTokenBurnedAndMinted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "TokenBurnedAndMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootTokenBurnedAndMinted)
				if err := _PolymorphRoot.contract.UnpackLog(event, "TokenBurnedAndMinted", log); err != nil {
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

// ParseTokenBurnedAndMinted is a log parse operation binding the contract event 0xb583e10e1b691f0fec5bb6e5bafd39038a17365ca548d3beb76c082261805583.
//
// Solidity: event TokenBurnedAndMinted(uint256 indexed tokenId, uint256 gene)
func (_PolymorphRoot *PolymorphRootFilterer) ParseTokenBurnedAndMinted(log types.Log) (*PolymorphRootTokenBurnedAndMinted, error) {
	event := new(PolymorphRootTokenBurnedAndMinted)
	if err := _PolymorphRoot.contract.UnpackLog(event, "TokenBurnedAndMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootTokenMintedIterator is returned from FilterTokenMinted and is used to iterate over the raw logs and unpacked data for TokenMinted events raised by the PolymorphRoot contract.
type PolymorphRootTokenMintedIterator struct {
	Event *PolymorphRootTokenMinted // Event containing the contract specifics and raw log

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
func (it *PolymorphRootTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootTokenMinted)
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
		it.Event = new(PolymorphRootTokenMinted)
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
func (it *PolymorphRootTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootTokenMinted represents a TokenMinted event raised by the PolymorphRoot contract.
type PolymorphRootTokenMinted struct {
	TokenId *big.Int
	NewGene *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenMinted is a free log retrieval operation binding the contract event 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da.
//
// Solidity: event TokenMinted(uint256 indexed tokenId, uint256 newGene)
func (_PolymorphRoot *PolymorphRootFilterer) FilterTokenMinted(opts *bind.FilterOpts, tokenId []*big.Int) (*PolymorphRootTokenMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "TokenMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphRootTokenMintedIterator{contract: _PolymorphRoot.contract, event: "TokenMinted", logs: logs, sub: sub}, nil
}

// WatchTokenMinted is a free log subscription operation binding the contract event 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da.
//
// Solidity: event TokenMinted(uint256 indexed tokenId, uint256 newGene)
func (_PolymorphRoot *PolymorphRootFilterer) WatchTokenMinted(opts *bind.WatchOpts, sink chan<- *PolymorphRootTokenMinted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "TokenMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootTokenMinted)
				if err := _PolymorphRoot.contract.UnpackLog(event, "TokenMinted", log); err != nil {
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

// ParseTokenMinted is a log parse operation binding the contract event 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da.
//
// Solidity: event TokenMinted(uint256 indexed tokenId, uint256 newGene)
func (_PolymorphRoot *PolymorphRootFilterer) ParseTokenMinted(log types.Log) (*PolymorphRootTokenMinted, error) {
	event := new(PolymorphRootTokenMinted)
	if err := _PolymorphRoot.contract.UnpackLog(event, "TokenMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootTokenMorphedIterator is returned from FilterTokenMorphed and is used to iterate over the raw logs and unpacked data for TokenMorphed events raised by the PolymorphRoot contract.
type PolymorphRootTokenMorphedIterator struct {
	Event *PolymorphRootTokenMorphed // Event containing the contract specifics and raw log

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
func (it *PolymorphRootTokenMorphedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootTokenMorphed)
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
		it.Event = new(PolymorphRootTokenMorphed)
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
func (it *PolymorphRootTokenMorphedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootTokenMorphedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootTokenMorphed represents a TokenMorphed event raised by the PolymorphRoot contract.
type PolymorphRootTokenMorphed struct {
	TokenId   *big.Int
	OldGene   *big.Int
	NewGene   *big.Int
	Price     *big.Int
	EventType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenMorphed is a free log retrieval operation binding the contract event 0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b.
//
// Solidity: event TokenMorphed(uint256 indexed tokenId, uint256 oldGene, uint256 newGene, uint256 price, uint8 eventType)
func (_PolymorphRoot *PolymorphRootFilterer) FilterTokenMorphed(opts *bind.FilterOpts, tokenId []*big.Int) (*PolymorphRootTokenMorphedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "TokenMorphed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphRootTokenMorphedIterator{contract: _PolymorphRoot.contract, event: "TokenMorphed", logs: logs, sub: sub}, nil
}

// WatchTokenMorphed is a free log subscription operation binding the contract event 0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b.
//
// Solidity: event TokenMorphed(uint256 indexed tokenId, uint256 oldGene, uint256 newGene, uint256 price, uint8 eventType)
func (_PolymorphRoot *PolymorphRootFilterer) WatchTokenMorphed(opts *bind.WatchOpts, sink chan<- *PolymorphRootTokenMorphed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "TokenMorphed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootTokenMorphed)
				if err := _PolymorphRoot.contract.UnpackLog(event, "TokenMorphed", log); err != nil {
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

// ParseTokenMorphed is a log parse operation binding the contract event 0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b.
//
// Solidity: event TokenMorphed(uint256 indexed tokenId, uint256 oldGene, uint256 newGene, uint256 price, uint8 eventType)
func (_PolymorphRoot *PolymorphRootFilterer) ParseTokenMorphed(log types.Log) (*PolymorphRootTokenMorphed, error) {
	event := new(PolymorphRootTokenMorphed)
	if err := _PolymorphRoot.contract.UnpackLog(event, "TokenMorphed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PolymorphRoot contract.
type PolymorphRootTransferIterator struct {
	Event *PolymorphRootTransfer // Event containing the contract specifics and raw log

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
func (it *PolymorphRootTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootTransfer)
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
		it.Event = new(PolymorphRootTransfer)
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
func (it *PolymorphRootTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootTransfer represents a Transfer event raised by the PolymorphRoot contract.
type PolymorphRootTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PolymorphRoot *PolymorphRootFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*PolymorphRootTransferIterator, error) {

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

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PolymorphRootTransferIterator{contract: _PolymorphRoot.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PolymorphRoot *PolymorphRootFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PolymorphRootTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootTransfer)
				if err := _PolymorphRoot.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_PolymorphRoot *PolymorphRootFilterer) ParseTransfer(log types.Log) (*PolymorphRootTransfer, error) {
	event := new(PolymorphRootTransfer)
	if err := _PolymorphRoot.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolymorphRootUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PolymorphRoot contract.
type PolymorphRootUnpausedIterator struct {
	Event *PolymorphRootUnpaused // Event containing the contract specifics and raw log

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
func (it *PolymorphRootUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolymorphRootUnpaused)
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
		it.Event = new(PolymorphRootUnpaused)
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
func (it *PolymorphRootUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolymorphRootUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolymorphRootUnpaused represents a Unpaused event raised by the PolymorphRoot contract.
type PolymorphRootUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PolymorphRoot *PolymorphRootFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PolymorphRootUnpausedIterator, error) {

	logs, sub, err := _PolymorphRoot.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PolymorphRootUnpausedIterator{contract: _PolymorphRoot.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PolymorphRoot *PolymorphRootFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PolymorphRootUnpaused) (event.Subscription, error) {

	logs, sub, err := _PolymorphRoot.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolymorphRootUnpaused)
				if err := _PolymorphRoot.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PolymorphRoot *PolymorphRootFilterer) ParseUnpaused(log types.Log) (*PolymorphRootUnpaused, error) {
	event := new(PolymorphRootUnpaused)
	if err := _PolymorphRoot.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
