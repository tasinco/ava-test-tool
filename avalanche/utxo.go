package avalanche

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/rpc"
	"github.com/ava-labs/avalanchego/vms/avm"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/avalanchego/vms/platformvm/blocks"
)

var (
	_ UTXOClient = platformvm.Client(nil)
	_ UTXOClient = avm.Client(nil)
)

type UTXOClient interface {
	GetAtomicUTXOs(
		ctx context.Context,
		addrs []ids.ShortID,
		sourceChain string,
		limit uint32,
		startAddress ids.ShortID,
		startUTXOID ids.ID,
		options ...rpc.Option,
	) ([][]byte, ids.ShortID, ids.ID, error)
}

// Utxos fetch all spendable utxos
func Utxos(ctx context.Context, client UTXOClient, sourceChainID string, addrs []ids.ShortID, fetchLimit uint32) (utxos map[ids.ID]*UTXOContainer, err error) {
	var (
		startAddr ids.ShortID
		startUTXO ids.ID
	)

	codec := blocks.Codec

	utxosmap := NewUtxoMap()
	for {
		utxosBytes, endAddr, endUTXO, err := client.GetAtomicUTXOs(ctx, addrs, sourceChainID, fetchLimit, startAddr, startUTXO)
		if err != nil {
			return utxos, err
		}

		if err := utxosmap.PraseUtxos(codec, utxosBytes); err != nil {
			return utxos, err
		}

		if uint32(len(utxosBytes)) < fetchLimit {
			break
		}

		startAddr = endAddr
		startUTXO = endUTXO
	}

	return utxosmap.M, nil
}
