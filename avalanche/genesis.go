package avalanche

import (
	"errors"
	"fmt"
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/genesis"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/vms/avm/fxs"
	avmtxs "github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/nftfx"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/avalanchego/vms/platformvm/blocks"
	platformgenesis "github.com/ava-labs/avalanchego/vms/platformvm/genesis"
	platformtxs "github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/propertyfx"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

var (
	ErrIncorrectGenesisChainTxType = errors.New("incorrect genesis chain tx type")
)

type Genesis struct {
	Bytes           []byte
	AvaxAssetID     ids.ID
	Platform        *platformgenesis.Genesis
	Chains          map[ids.ID]*platformtxs.CreateChainTx
	ChainsByStr     map[string]*platformtxs.CreateChainTx
	PlatformVM      *platformvm.VM
	AvmCodec        codec.Manager
	PlatformvmCodec codec.Manager
	AvmParser       avmtxs.Parser
	AVMChain        *platformtxs.CreateChainTx
	EVMChain        *platformtxs.CreateChainTx
	AVMChainTx      *platformtxs.Tx
	EVMChainTX      *platformtxs.Tx
}

func NewGenesis(networkID uint32) (g *Genesis, err error) {
	g = &Genesis{}

	g.PlatformvmCodec = blocks.Codec

	platformGenesisBytes, assetID, err := genesis.FromConfig(genesis.GetConfig(networkID))
	if err != nil {
		return g, err
	}

	g.Bytes = platformGenesisBytes
	g.AvaxAssetID = assetID

	g.Platform, err = platformgenesis.Parse(platformGenesisBytes)
	if err != nil {
		return g, err
	}

	g.Chains = make(map[ids.ID]*platformtxs.CreateChainTx)
	g.ChainsByStr = make(map[string]*platformtxs.CreateChainTx)

	for _, chain := range g.Platform.Chains {
		createChainTx, ok := chain.Unsigned.(*platformtxs.CreateChainTx)
		if !ok {
			return g, ErrIncorrectGenesisChainTxType
		}

		g.ChainsByStr[chain.ID().String()] = createChainTx
		g.Chains[chain.ID()] = createChainTx

		if createChainTx.VMID == constants.AVMID {
			g.AVMChain = createChainTx
			g.AVMChainTx = chain
		}
		if createChainTx.VMID == constants.EVMID {
			g.EVMChain = createChainTx
			g.EVMChainTX = chain
		}
	}

	g.AvmParser, err = newAVMCodec(platformGenesisBytes)
	if err != nil {
		return g, err
	}

	return g, err
}

func newAVMCodec(genesisBytes []byte) (avmtxs.Parser, error) {
	g, err := genesis.VMGenesis(genesisBytes, constants.AVMID)
	if err != nil {
		return nil, err
	}

	createChainTx, ok := g.Unsigned.(*platformtxs.CreateChainTx)
	if !ok {
		return nil, ErrIncorrectGenesisChainTxType
	}

	var (
		fxIDs = createChainTx.FxIDs
		fxs   = make([]fxs.Fx, 0, len(fxIDs))
	)

	for _, fxID := range fxIDs {
		switch {
		case fxID == secp256k1fx.ID:
			fxs = append(fxs, &secp256k1fx.Fx{})
		case fxID == nftfx.ID:
			fxs = append(fxs, &nftfx.Fx{})
		case fxID == propertyfx.ID:
			fxs = append(fxs, &propertyfx.Fx{})
		default:
			panic(fmt.Sprint("unknown fx %v", fxID))
		}
	}

	return avmtxs.NewParser(fxs)
}
