all: build

.PHONY: build
build:
	solc-linux-amd64-latest --abi Reverter.sol -o build --overwrite
	solc-linux-amd64-latest --bin Reverter.sol -o build --overwrite
	abigen --bin=build/Reverter.bin --abi=build/Reverter.abi --pkg=reverter --type=Reverter --out=Reverter.go
