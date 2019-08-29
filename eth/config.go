// Copyright 2017 The go-severeum Authors
// This file is part of the go-severeum library.
//
// The go-severeum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-severeum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-severeum library. If not, see <http://www.gnu.org/licenses/>.

package sev

import (
	"math/big"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	"github.com/severeum/go-severeum/common"
	"github.com/severeum/go-severeum/common/hexutil"
	"github.com/severeum/go-severeum/consensus/sevash"
	"github.com/severeum/go-severeum/core"
	"github.com/severeum/go-severeum/sev/downloader"
	"github.com/severeum/go-severeum/sev/gasprice"
	"github.com/severeum/go-severeum/params"
)

// DefaultConfig contains default settings for use on the Severeum main net.
var DefaultConfig = Config{
	SyncMode: downloader.FastSync,
	Sevash: sevash.Config{
		CacheDir:       "sevash",
		CachesInMem:    2,
		CachesOnDisk:   3,
		DatasetsInMem:  1,
		DatasetsOnDisk: 2,
	},
	NetworkId:      1,
	LightPeers:     100,
	DatabaseCache:  512,
	TrieCleanCache: 256,
	TrieDirtyCache: 256,
	TrieTimeout:    60 * time.Minute,
	MinerGasFloor:  8000000,
	MinerGasCeil:   8000000,
	MinerGasPrice:  big.NewInt(params.GWei),
	MinerRecommit:  3 * time.Second,

	TxPool: core.DefaultTxPoolConfig,
	GPO: gasprice.Config{
		Blocks:     20,
		Percentile: 60,
	},
}

func init() {
	home := os.Getenv("HOME")
	if home == "" {
		if user, err := user.Current(); err == nil {
			home = user.HomeDir
		}
	}
	if runtime.GOOS == "windows" {
		DefaultConfig.Sevash.DatasetDir = filepath.Join(home, "AppData", "Sevash")
	} else {
		DefaultConfig.Sevash.DatasetDir = filepath.Join(home, ".sevash")
	}
}

//go:generate gencodec -type Config -field-override configMarshaling -formats toml -out gen_config.go

type Config struct {
	// The genesis block, which is inserted if the database is empty.
	// If nil, the Severeum main net block is used.
	Genesis *core.Genesis `toml:",omitempty"`

	// Protocol options
	NetworkId uint64 // Network ID to use for selecting peers to connect to
	SyncMode  downloader.SyncMode
	NoPruning bool

	// Whitelist of required block number -> hash values to accept
	Whitelist map[uint64]common.Hash `toml:"-"`

	// Light client options
	LightServ  int `toml:",omitempty"` // Maximum percentage of time allowed for serving LES requests
	LightPeers int `toml:",omitempty"` // Maximum number of LES client peers

	// Database options
	SkipBcVersionCheck bool `toml:"-"`
	DatabaseHandles    int  `toml:"-"`
	DatabaseCache      int
	TrieCleanCache     int
	TrieDirtyCache     int
	TrieTimeout        time.Duration

	// Mining-related options
	Severbase      common.Address `toml:",omitempty"`
	MinerNotify    []string       `toml:",omitempty"`
	MinerExtraData []byte         `toml:",omitempty"`
	MinerGasFloor  uint64
	MinerGasCeil   uint64
	MinerGasPrice  *big.Int
	MinerRecommit  time.Duration
	MinerNoverify  bool

	// Sevash options
	Sevash sevash.Config

	// Transaction pool options
	TxPool core.TxPoolConfig

	// Gas Price Oracle options
	GPO gasprice.Config

	// Enables tracking of SHA3 preimages in the VM
	EnablePreimageRecording bool

	// Miscellaneous options
	DocRoot string `toml:"-"`

	// Type of the EWASM interpreter ("" for default)
	EWASMInterpreter string

	// Type of the EVM interpreter ("" for default)
	EVMInterpreter string

	// Constantinople block override (TODO: remove after the fork)
	ConstantinopleOverride *big.Int
}

type configMarshaling struct {
	MinerExtraData hexutil.Bytes
}
