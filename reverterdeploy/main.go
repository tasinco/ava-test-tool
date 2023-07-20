package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/tasinco/ava-test-tool/contracts/reverter"
	"github.com/tasinco/ava-test-tool/ethcalls"
)

var Bech32HRP = constants.GetHRP(constants.LocalID)

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

	// private key from certs for ava local network
	privateKey, err := crypto.HexToECDSA("56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027")
	if err != nil {
		log.Fatal(err)
	}
	addr := crypto.PubkeyToAddress(privateKey.PublicKey)
	avaxAddr, err := address.FormatBech32(Bech32HRP, addr.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	// private key from certs for ava local network
	privateKey2, err := crypto.HexToECDSA("56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccd558d8027")
	if err != nil {
		log.Fatal(err)
	}
	addr2 := crypto.PubkeyToAddress(privateKey2.PublicKey)
	avaxAddr2, err := address.FormatBech32(Bech32HRP, addr2.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	// check balances
	addr1bal, err := ethClnt.BalanceAt(ctx, addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	addr2bal, err := ethClnt.BalanceAt(ctx, addr2, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("addr", avaxAddr, addr1bal, addr.String())
	log.Println("addr2", avaxAddr2, addr2bal, addr2.String())

	// deploy the contract
	reverterContractAddr, rtx, reverter, err := reverter.DeployReverter(getAuth(ctx, privateKey, chainID), ethClnt)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = waitForTx(ctx, "reverter", ethClnt, rtx.Hash())

	log.Println("reverterContractAddr", reverterContractAddr.String())

	// disable receive
	tx, err := reverter.SetEnableReceive(getAuth(ctx, privateKey, chainID), new(big.Int))
	if err != nil {
		log.Fatal(err)
	}
	_, _ = waitForTx(ctx, "revert disable", ethClnt, tx.Hash())

	// send some money
	gasPrice, signedTx, err := doSend(ethClnt, ctx, addr, new(big.Int).SetUint64((1)*params.Ether), reverterContractAddr, chainID, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("receive", tx.Hash(), gasPrice)
	_, _ = waitForTx(ctx, "receive", ethClnt, signedTx.Hash())

	// check balances
	addr1bal, err = ethClnt.BalanceAt(ctx, addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	addr2bal, err = ethClnt.BalanceAt(ctx, addr2, nil)
	if err != nil {
		log.Fatal(err)
	}
	rContractBal, err := ethClnt.BalanceAt(ctx, reverterContractAddr, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("addr", avaxAddr, addr1bal, addr.String())
	log.Println("addr2", avaxAddr2, addr2bal, addr2.String())
	log.Println("contract balance", rContractBal)

	// debug the trace (it will revert)
	c, err := ethcalls.EthDebugTrace(ctx, rpcClnt, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(toJson(&c))

	// re-enable the contract
	tx, err = reverter.SetEnableReceive(getAuth(ctx, privateKey, chainID), new(big.Int).SetInt64(1))
	if err != nil {
		log.Fatal(err)
	}
	_, _ = waitForTx(ctx, "revert disable", ethClnt, tx.Hash())

	// send more money
	gasPrice, signedTx, err = doSend(ethClnt, ctx, addr, new(big.Int).SetUint64((1*params.Ether)+12345), reverterContractAddr, chainID, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("receive", tx.Hash(), gasPrice)
	_, _ = waitForTx(ctx, "receive", ethClnt, signedTx.Hash())

	// re-check balances
	addr1bal, err = ethClnt.BalanceAt(ctx, addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	addr2bal, err = ethClnt.BalanceAt(ctx, addr2, nil)
	if err != nil {
		log.Fatal(err)
	}
	rContractBal, err = ethClnt.BalanceAt(ctx, reverterContractAddr, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("addr", avaxAddr, addr1bal, addr.String())
	log.Println("addr2", avaxAddr2, addr2bal, addr2.String())
	log.Println("contract balance", rContractBal)

	// debug the call it won't revert and the contract balance has increased.
	c, err = ethcalls.EthDebugTrace(ctx, rpcClnt, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(toJson(&c))
}

func getAuth(ctx context.Context, privateKey *ecdsa.PrivateKey, chainID *big.Int) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Context = ctx
	return auth
}

func doSend(ethClnt *ethclient.Client, ctx context.Context, addr common.Address, value *big.Int, spliiterContractAddr common.Address, chainID *big.Int, privateKey *ecdsa.PrivateKey) (*big.Int, *types.Transaction, error) {
	var (
		gasPrice *big.Int
		signedTx *types.Transaction
	)

	nonce, err := ethClnt.PendingNonceAt(ctx, addr)
	if err != nil {
		return gasPrice, signedTx, err
	}

	gasLimit := uint64(3000000) // in units

	head, err := ethClnt.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return gasPrice, signedTx, err
	}

	gasPrice = new(big.Int).Add(head.BaseFee, big.NewInt(1))

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &spliiterContractAddr,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     nil,
	})

	signedTx, err = types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		return gasPrice, signedTx, err
	}

	err = ethClnt.SendTransaction(ctx, signedTx)
	if err != nil {
		return gasPrice, signedTx, err
	}

	return gasPrice, signedTx, nil
}

func waitForTx(ctx context.Context, tag string, ec *ethclient.Client, txhash common.Hash) (tr *types.Receipt, err error) {
	for {
		tr, err = ec.TransactionReceipt(ctx, txhash)
		if err != nil {
			if !strings.Contains(err.Error(), "not found") {
				log.Println("transaction fail", tag, err)
			}
			time.Sleep(500 * time.Millisecond)
			continue
		}
		break
	}

	return tr, nil
}

func toJson(tx interface{}) string {
	jd, _ := json.Marshal(tx)
	return string(jd)
}
