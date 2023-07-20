# Sol contract

Simple contract to accept a balance or reject a balance.

## Howto fetch solc

```shell
cd ~/{src}
git clone git@github.com:ethereum/solc-bin.git
```

## Howto build abigen

```shell
cd ~/{src}
git clone git@github.com:ethereum/go-ethereum.git
cd ~/{src}/go-ethereum
go build -o ~/go/bin/abigen ./cmd/abigen
```

## build

```shell
PATH=~/{src}/solc-bin/linux-amd64:~/go/bin:$PATH} make
```

output
```text
solc-linux-amd64-latest --abi Reverter.sol -o build --overwrite
Compiler run successful. Artifact(s) can be found in directory "build".
solc-linux-amd64-latest --bin Reverter.sol -o build --overwrite
Compiler run successful. Artifact(s) can be found in directory "build".
abigen --bin=build/Reverter.bin --abi=build/Reverter.abi --pkg=reverter --type=Reverter --out=Reverter.go
```