package contracts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	GoerliIdRegistryAddrStr     = "0xda107a1caf36d198b12c16c7b6a1d1c795978c42"
	GoerliIdRegistryGenesis     = 7648795
	GoerliNameRegistryAddrStr   = "0xe3be01d99baa8db9905b33a3ca391238234b79d1"
	GoerliNameRegistryGenesis   = 7648795
	GoerliBundleRegistryAddrStr = "0xdb647193df79ce69b5d34549aae98d519223f682"
	GoerliBundleRegistryGenesis = 7648795
)

var (
	GoerliIdRegistryAddress     = common.HexToAddress(GoerliIdRegistryAddrStr)
	GoerliNameRegistryAddress   = common.HexToAddress(GoerliNameRegistryAddrStr)
	GoerliBundleRegistryAddress = common.HexToAddress(GoerliBundleRegistryAddrStr)
)
