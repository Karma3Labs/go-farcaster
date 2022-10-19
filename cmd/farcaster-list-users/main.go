package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/eigentrust-io/go-farcaster/internal/contracts/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"sync"
)

const (
	FarcasterApiEndpoint = "https://api.farcaster.xyz"
	FarcasterV0Indexer   = FarcasterApiEndpoint + "/indexer"
	FarcasterV0Following = FarcasterV0Indexer + "/following"
	FarcasterV1Endpoint  = FarcasterApiEndpoint + "/v1"
	FarcasterV1Profiles  = FarcasterV1Endpoint + "/profiles"
	NodesCsv             = "nodes.csv"
	EdgesCsv             = "edges.csv"
	DefaultJobs          = 10
)

var (
	defaultJsonRpcEndpoint = os.Getenv("JSON_RPC")
	jsonRpcEndpoint        = flag.String(
		"json-rpc-endpoint", defaultJsonRpcEndpoint,
		"Goerli JSON RPC endpoint URL",
	)
	jobs = flag.Uint("jobs", DefaultJobs, "number of parallel jobs")
)

type ProfileResponse struct {
	Result struct {
		User User `json:"user"`
	} `json:"result"`
}

type User struct {
	AddrStr                  string  `json:"address"`
	Username                 string  `json:"username"`
	DisplayName              string  `json:"displayName"`
	Avatar                   Avatar  `json:"avatar"`
	FollowerCount            int     `json:"followerCount"`
	FollowingCount           int     `json:"followingCount"`
	IsViewerFollowing        bool    `json:"isViewerFollowing"`
	IsFollowingViewer        bool    `json:"isFollowingViewer"`
	Profile                  Profile `json:"profile"`
	ReferrerUsername         string  `json:"referrerUsername"`
	viewerCanSendDirectCasts bool    `json:"viewerCanSendDirectCasts"`
}

func (u User) Validate() error {
	return nil
}

func parseHexAddress(s string) (address common.Address, err error) {
	if common.IsHexAddress(s) {
		address = common.HexToAddress(s)
	} else {
		err = errors.Errorf("%s is not a valid address", s)
	}
	return
}

func (u User) Address() (address common.Address, err error) {
	return parseHexAddress(u.AddrStr)
}

type Avatar struct {
	UrlStr     string `json:"url"`
	IsVerified bool   `json:"isVerified"`
}

func (a Avatar) Url() (*url.URL, error) {
	return url.Parse(a.UrlStr)
}

type Profile struct {
	Bio                  Bio                  `json:"bio"`
	DirectMessageTargets DirectMessageTargets `json:"directMessageTargets"`
}

type Bio struct {
	Text     string    `json:"text"`
	Mentions []Mention `json:"mentions"`
}

type Mention struct {
	AddrStr      string `json:"address"`
	Username     string `json:"username"`
	DisplayName  string `json:"displayName"`
	Avatar       Avatar `json:"avatar"`
	RegisteredAt uint64 `json:"registeredAt"`
}

func (m Mention) Address() (common.Address, error) {
	return parseHexAddress(m.AddrStr)
}

type DirectMessageTargets struct {
	Telegram string `json:"telegram"`
}

type Following struct {
	AddrStr           string `json:"address"`
	Username          string `json:"username"`
	DisplayName       string `json:"displayName"`
	Avatar            Avatar `json:"avatar"`
	IsViewerFollowing bool   `json:"isViewerFollowing"`
	Verifications     []bool `json:"verifications"`
}

func (f Following) Address() (common.Address, error) {
	return parseHexAddress(f.AddrStr)
}

func GetUser(address common.Address) (*User, error) {
	url := fmt.Sprintf("%s/%s", FarcasterV1Profiles, address)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.WithMessagef(err, "GET %s", url)
	}
	if resp.StatusCode != 200 {
		return nil, errors.Errorf(
			"GET %s -> %d %s", url, resp.StatusCode, resp.Status,
		)
	}
	defer resp.Body.Close()
	var body ProfileResponse
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, errors.WithMessage(err, "cannot decode JSON response %#v")
	}
	return &body.Result.User, nil
}

func GetFollowing(address common.Address) ([]Following, error) {
	url := fmt.Sprintf("%s/%s", FarcasterV0Following, address)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.WithMessagef(err, "GET %s", url)
	}
	if resp.StatusCode != 200 {
		return nil, errors.Errorf(
			"GET %s -> %d %s", url, resp.StatusCode, resp.Status,
		)
	}
	defer resp.Body.Close()
	var followings []Following
	err = json.NewDecoder(resp.Body).Decode(&followings)
	if err != nil {
		return nil, errors.WithMessage(err, "cannot decode JSON response %#v")
	}
	return followings, nil
}

type Record struct {
	Fid        *uint256.Int
	User       *User
	Followings []Following
}

func main() {
	flag.Parse()

	if *jsonRpcEndpoint == "" {
		fmt.Fprintln(os.Stderr, "--json-rpc-endpoint is not specified")
		os.Exit(64)
	}
	logger := log.Logger
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, *jsonRpcEndpoint)
	if err != nil {
		logger.Fatal().Err(err).Msg("cannot connect to JSON RPC endpoint")
	}
	address := contracts.GoerliIdRegistryAddress
	idRegistry, err := contracts.NewIdRegistry(address, client)
	if err != nil {
		logger.Fatal().Err(err).Stringer(
			"address", address,
		).Msg("cannot create IdRegistry")
	}
	bn, err := client.BlockNumber(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("cannot get the head block number")
	}
	logger.Info().Uint64("blockNumber", bn).Msg("chain height")
	filterOpts := bind.FilterOpts{
		Start:   contracts.GoerliIdRegistryGenesis,
		End:     &bn,
		Context: ctx,
	}

	records := make(chan Record)

	go func() {
		workerSlots := make(chan bool, *jobs)
		var outstanding sync.WaitGroup
		more := true
		for fid := uint256.NewInt(1); more; fid = fid.AddUint64(fid, 1) {
			workerSlots <- true
			outstanding.Add(1)
			go func(fid *uint256.Int) {
				defer func() {
					outstanding.Done()
					<-workerSlots
				}()
				fids := []*big.Int{fid.ToBig()}
				registerEvents, err := idRegistry.FilterRegister(
					&filterOpts, nil, fids,
				)
				logger := logger.With().Stringer("fid", fid).Logger()
				if err != nil {
					logger.Err(err).Msg("cannot fetch register event")
					return
				}
				if !registerEvents.Next() {
					more = false
					return
				}
				ev := registerEvents.Event
				bn := ev.Raw.BlockNumber
				idx := ev.Raw.Index
				addr := ev.To
				if registerEvents.Next() {
					logger.Error().Msg("duplicate registration")
					return
				}
				transferEvents, err := idRegistry.FilterTransfer(
					&filterOpts, nil, nil, fids,
				)
				if err != nil {
					logger.Err(err).Msg("cannot fetch transfer events")
					return
				}
				for transferEvents.Next() {
					ev := transferEvents.Event
					if ev.Raw.BlockNumber < bn || ev.Raw.BlockNumber == bn && ev.Raw.Index < idx {
						logger.Warn().Msg("event went backward")
						continue
					}
					logger.Info().Stringer("from", ev.From).Stringer(
						"to", ev.To,
					).Msg("user transferred")
					bn = ev.Raw.BlockNumber
					idx = ev.Raw.Index
					addr = ev.To
				}
				logger = logger.With().Stringer("address", addr).Logger()
				user, err := GetUser(addr)
				if err != nil {
					logger.Err(err).Msg("cannot fetch user profile")
					return
				}
				username := user.Username
				logger = logger.With().Str("username", username).Logger()
				followings, err := GetFollowing(addr)
				if err != nil {
					logger.Err(err).Msg("cannot fetch followings")
					return
				}
				if len(followings) != user.FollowingCount {
					logger.Warn().Int(
						"lenFollowings", len(followings),
					).Int(
						"followingCount", user.FollowingCount,
					).Msg("following count mismatch")
				}
				records <- Record{Fid: fid, User: user, Followings: followings}
			}(fid.Clone())
		}
		outstanding.Wait()
		close(records)
	}()

	nodesCsv, err := os.Create(NodesCsv)
	if err != nil {
		logger.Fatal().Str("filename", NodesCsv).Msg("cannot open node file")
	}

	edgesCsv, err := os.Create(EdgesCsv)
	if err != nil {
		logger.Fatal().Str("filename", EdgesCsv).Msg("cannot open edge file")
	}

	nodesWriter := csv.NewWriter(nodesCsv)
	edgesWriter := csv.NewWriter(edgesCsv)
	defer func() {
		nodesWriter.Flush()
		if err := nodesWriter.Error(); err != nil {
			logger.Err(err).Str(
				"filename", NodesCsv,
			).Msg("cannot flush nodes file")
		}
		edgesWriter.Flush()
		if err := edgesWriter.Error(); err != nil {
			logger.Err(err).Str(
				"filename", EdgesCsv,
			).Msg("cannot flush edges file")
		}
		edgesWriter.Flush()
		if err := nodesCsv.Close(); err != nil {
			logger.Err(err).Str(
				"filename", NodesCsv,
			).Msg("cannot close nodes file")
		}
		if err := edgesCsv.Close(); err != nil {
			logger.Err(err).Str(
				"filename", EdgesCsv,
			).Msg("cannot close edges file")
		}
	}()
	for record := range records {
		if err := nodesWriter.Write(
			[]string{
				record.Fid.String(),
				record.User.AddrStr,
				record.User.Username,
			},
		); err != nil {
			logger.Err(err).Str(
				"filename", NodesCsv,
			).Msg("cannot write node record")
			return
		}
		for _, following := range record.Followings {
			if err := edgesWriter.Write(
				[]string{
					record.User.AddrStr,
					record.User.Username,
					following.AddrStr,
					following.Username,
				},
			); err != nil {
				logger.Err(err).Str(
					"filename", EdgesCsv,
				).Msg("cannot write edge record")
				return
			}
		}
	}
}
