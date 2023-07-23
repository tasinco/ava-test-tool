package avalanche

import (
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/tasinco/ava-test-tool/avalanche/privatekey"
)

var (
	LocalNetPrivateKeysIDs = []string{
		"PrivateKey-ewoqjP7PxY4yr3iLTpLisriqt94hdyDFNgchSxGGztUrTXtNN",
		"PrivateKey-wHR4zmr9am94KVYnV2aRR4QXt78cuGebt1GpYNwJYEbfAGonj",
		"PrivateKey-AR874kuHtHpDk7ntffuEQ9cwiQLL2dz1DmJankW1EyXnz5fc7",
		"PrivateKey-Ntk8vV7zaWzAot2wuDXK4e9ZGFUnU49AYTDew5XUyYaNz2u9d",
		"PrivateKey-oLM8XbXxXmBHVbdKm2tRYQ1WdMj3b2NggftQpvDUXWSMtdY4i",
		"PrivateKey-2kjfDc9RVUQJnu3HQDGiVdxvhM9BmR3UTx7Aq8AJ82G2MspATy",
		"PrivateKey-2Rh5Gtu28ca7PS6rLfN6uou9ext8Y5xhoAJDdWPU7GESBLHtv6",
		"PrivateKey-2ZcbEPKkXjswsNRBGViGzruReAtTAxW9hsGeMc2GgppnJnDgne",
		"PrivateKey-22SYvqaRgFtPJfiZmswrCyE57UcssLVnNPDJ48PYAiCjKVAGy7",
		"PrivateKey-tYRsRPijLo6KD2azMLzkcB2ZUndU3a2dJ8kEqBtqesa85pWhB",
	}
	PrimaryLocalNetPrivateKey = LocalNetPrivateKeysIDs[0]

	LocalNetPrivateKeys = make([]*privatekey.PkInfo, len(LocalNetPrivateKeysIDs))
)

func init() {
	for pos, pkids := range LocalNetPrivateKeysIDs {
		pkInfo, err := privatekey.DecodeB32(pkids, constants.LocalID)
		if err != nil {
			panic(err)
		}
		LocalNetPrivateKeys[pos] = pkInfo
	}
}

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
