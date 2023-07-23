package avalanche

import (
	"strings"

	"github.com/ava-labs/coreth/ethclient"
	"github.com/ava-labs/coreth/rpc"
)

type EvmClient struct {
	RpcClient *rpc.Client
	EthClient ethclient.Client
}

func NewEvmClient(addr string) (c *EvmClient, err error) {
	addr = strings.TrimPrefix(addr, "/")

	rc, err := rpc.Dial(addr + "/ext/bc/C/rpc")
	if err != nil {
		return c, err
	}
	ec := ethclient.NewClient(rc)

	return &EvmClient{
		RpcClient: rc,
		EthClient: ec,
	}, nil
}
