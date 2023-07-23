package avalanche

import (
	"strings"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/vms/avm"
	"github.com/ava-labs/coreth/plugin/evm"
)

type Client struct {
	Info info.Client
	Avm  avm.Client
	Evm  evm.Client
}

func NewClient(addr string) *Client {
	addr = strings.TrimPrefix(addr, "/")

	return &Client{
		Info: info.NewClient(addr),
		Avm:  avm.NewClient(addr, "X"),
		Evm:  evm.NewClient(addr, "C"),
	}
}
