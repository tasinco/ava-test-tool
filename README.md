# ava-test-tool

Start a 5 node Ava testing environment.

## Howto run

```shell
go run ./cmd/app/
```

## Howto connect

You can interact with the C chain at http://localhost:9650

```go
package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

func main() {
	ctx := context.Background()

	rpcClnt, err := rpc.Dial("http://localhost:9650/ext/bc/C/rpc")
	if err != nil {
		log.Fatal(err)
	}
	ethClnt := ethclient.NewClient(rpcClnt)

	chainID, err := ethClnt.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(chainID)
}
```
