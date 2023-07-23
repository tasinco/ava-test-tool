package avalanche

import (
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/components/avax"
)

type UTXOContainer struct {
	Utxos []*avax.UTXO
	Amt   uint64
}

type UTXOMap struct {
	M map[ids.ID]*UTXOContainer
}

func NewUtxoMap() *UTXOMap {
	return &UTXOMap{M: make(map[ids.ID]*UTXOContainer)}
}

func (m *UTXOMap) PraseUtxos(codec codec.Manager, utxosBytes [][]byte) (err error) {
	for _, utxoBytes := range utxosBytes {
		var utxo avax.UTXO
		_, err := codec.Unmarshal(utxoBytes, &utxo)
		if err != nil {
			return err
		}

		if _, ok := m.M[utxo.AssetID()]; !ok {
			m.M[utxo.AssetID()] = &UTXOContainer{}
		}

		m.M[utxo.AssetID()].Utxos = append(m.M[utxo.AssetID()].Utxos, &utxo)
		if amount, ok := utxo.Out.(avax.Amounter); ok {
			m.M[utxo.AssetID()].Amt += amount.Amount()
		}
	}
	return nil
}
