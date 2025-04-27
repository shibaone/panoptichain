// Package network defines objects that we'll use for understanding which network we are observing
package network

import (
	"fmt"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/log"
)

// Network is the interface that we need to identify a network.
type Network interface {
	GetName() string
}

const (
	ShibariumMainnetName  = "Shibarium"
	ShibariumPuppynetName = "Puppynet"

	HeimdallMainnetName  = "Heimdall Shibarium"
	HeimdallPuppynetName = "Heimdall Puppynet"

	EthereumName = "Ethereum"
	SepoliaName  = "Sepolia"
)

var ShibariumMainnet = EVMNetwork{Name: ShibariumMainnetName, ChainID: 109}
var ShibariumPuppynet = EVMNetwork{Name: ShibariumPuppynetName, ChainID: 157}

var HeimdallMainnet = TendermintNetwork{Name: HeimdallMainnetName, ChainID: "heimdall-109"}
var HeimdallPuppynet = TendermintNetwork{Name: HeimdallPuppynetName, ChainID: "heimdall-157"}

var Ethereum = EVMNetwork{Name: EthereumName, ChainID: 1}
var Sepolia = EVMNetwork{Name: SepoliaName, ChainID: 11155111}

var KnownNetworks = []Network{
	&ShibariumMainnet,
	&ShibariumPuppynet,

	&HeimdallMainnet,
	&HeimdallPuppynet,

	&Ethereum,
	&Sepolia,
}

// EVMNetwork is a specific type of network that is assumed to generally follow
// the norms of an Ethereum based network.
type EVMNetwork struct {
	Name    string
	ChainID uint64
}

// GetName returns the name that we've set for the network. It doesn't have to
// be canonical in anyway. Just something to match configs.
func (n *EVMNetwork) GetName() string {
	return n.Name
}

// GetChainID should return the configured chain id for the network.
func (n *EVMNetwork) GetChainID() uint64 {
	return n.ChainID
}

// TendermintNetwork is close to an EVMNetwork but generally the chain
// ID is a string and there won't be typical JSON RPC provers.
type TendermintNetwork struct {
	Name    string
	ChainID string
}

// GetName will return the configured name for this network.
func (n *TendermintNetwork) GetName() string {
	return n.Name
}

// GetChainID will return the chain id for the network.
func (n *TendermintNetwork) GetChainID() string {
	return n.ChainID
}

// GetNetworkByName is a convenience method to convert a name like
// "Ethereum" into a Network object.
func GetNetworkByName(name string) (Network, error) {
	for _, n := range config.Config().Networks {
		if n.GetName() == name {
			return n, nil
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
