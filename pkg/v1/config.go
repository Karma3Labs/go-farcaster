package farcaster_v1

import "github.com/ethereum/go-ethereum/ethclient"

// Config contains configuration items needed for interfacing with
// the Farcaster API smart contracts.
type Config struct {
	// EthClient is used for interfacing with the Ethereum smart contracts.
	EthClient *ethclient.Client
}
