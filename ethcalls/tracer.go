package ethcalls

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/rpc"
)

var tracer = "callTracer"

// tracerTimeout = "180s"

type Call struct {
	Type    string         `json:"type"`
	From    common.Address `json:"from"`
	To      common.Address `json:"to"`
	Input   string         `json:"input"`
	Value   *hexutil.Big   `json:"value"`
	Gas     *hexutil.Big   `json:"gas"`
	GasUsed *hexutil.Big   `json:"gasUsed"`
	Revert  bool           `json:"revert"`
	Error   string         `json:"error,omitempty"`
	Calls   []*Call        `json:"calls,omitempty"`
}

func EthDebugTrace(ctx context.Context, clnt *rpc.Client, txHash common.Hash) (Call, error) {
	var results Call
	args := &tracers.TraceConfig{
		// Timeout: &tracerTimeout,
		Tracer: &tracer,
	}

	txHashVal := txHash.String()

	if err := clnt.CallContext(ctx, &results, "debug_traceTransaction", txHashVal, args); err != nil {
		return results, err
	}

	return results, nil
}
