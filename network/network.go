// Package network defines objects that we'll use for understanding which network we are observing
package network

import (
	"fmt"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/log"
)

// Network is the interface needed to identify a network.
type Network interface {
	GetName() string
	GetChainID() uint64
	IsPolygonPoS() bool
	IsPolygonZkEVM() bool
}

const (
	PolygonMainnetName = "Polygon Mainnet"
	PolygonMumbaiName  = "Polygon Mumbai"
	PolygonAmoyName    = "Polygon Amoy"

	ZkEVMMainnetName = "zkEVM Mainnet"
	ZkEVMTestnetName = "zkEVM Testnet"
	ZkEVMCardonaName = "zkEVM Cardona"
	ZkEVMBaliName    = "zkEVM Bali"

	EthereumName = "Ethereum"
	SepoliaName  = "Sepolia"
	GoerliName   = "Goerli"
)

var PolygonMainnet = config.Network{Name: PolygonMainnetName, ChainID: 137, PolygonPoS: true}
var PolygonMumbai = config.Network{Name: PolygonMumbaiName, ChainID: 80001, PolygonPoS: true}
var PolygonAmoy = config.Network{Name: PolygonAmoyName, ChainID: 80002, PolygonPoS: true}

var ZkEVMMainnet = config.Network{Name: ZkEVMMainnetName, ChainID: 1101, PolygonZkEVM: true}
var ZkEVMTestnet = config.Network{Name: ZkEVMTestnetName, ChainID: 1442, PolygonZkEVM: true}
var ZkEVMCardona = config.Network{Name: ZkEVMCardonaName, ChainID: 2442, PolygonZkEVM: true}
var ZkEVMBali = config.Network{Name: ZkEVMBaliName, ChainID: 2440, PolygonZkEVM: true}

var Ethereum = config.Network{Name: EthereumName, ChainID: 1}
var Sepolia = config.Network{Name: SepoliaName, ChainID: 11155111}
var Goerli = config.Network{Name: GoerliName, ChainID: 5}

var KnownNetworks = []Network{
	&PolygonMainnet,
	&PolygonMumbai,
	&PolygonAmoy,

	&ZkEVMMainnet,
	&ZkEVMTestnet,
	&ZkEVMCardona,
	&ZkEVMBali,

	&Ethereum,
	&Goerli,
	&Sepolia,
}

// GetNetworkByName converts a name like "Ethereum" into a Network object.
func GetNetworkByName(name string) (Network, error) {
	for _, n := range config.Config().Networks {
		if n.GetName() == name {
			return &n, nil
		}
	}

	for _, n := range KnownNetworks {
		if n.GetName() == name {
			return n, nil
		}
	}

	return nil, fmt.Errorf("the network name %s is not recognized", name)
}

func init() {
	names := make(map[string]struct{}, 0)
	for _, k := range KnownNetworks {
		if _, ok := names[k.GetName()]; ok {
			log.Fatal().Str("name", k.GetName()).Msg("Two networks with the same name detected")
		}

		names[k.GetName()] = struct{}{}
	}
}
