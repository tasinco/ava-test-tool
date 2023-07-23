package avalanche

import (
	"context"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/blocks"
	"github.com/ava-labs/coreth/plugin/evm"
)

func EvmUtxos(ctx context.Context, client evm.Client, sourceChainID ids.ID, addrs []string, fetchLimit uint32) (utxos map[ids.ID]*UTXOContainer, err error) {
	var (
		sourceChainIDStr = sourceChainID.String()
		index            api.Index
	)

	codec := blocks.Codec

	utxosmap := NewUtxoMap()
	for {
		utxosBytes, rindex, err := client.GetAtomicUTXOs(ctx, addrs, sourceChainIDStr, 1000, index.Address, index.UTXO)
		if err != nil {
			return utxos, err
		}

		if err := utxosmap.PraseUtxos(codec, utxosBytes); err != nil {
			return utxos, err
		}

		if uint32(len(utxosBytes)) < fetchLimit {
			break
		}

		index = rindex
	}

	return utxosmap.M, nil
}
