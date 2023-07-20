package privatekey

import (
	"crypto/ecdsa"
	"encoding/hex"
	"strings"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/cb58"
	"github.com/ava-labs/avalanchego/utils/constants"
	avacrypto "github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type PkInfo struct {
	PrivKeyBytes string                `json:"privKeyBytes"`
	PrivKey      *secp256k1.PrivateKey `json:"-"`
	XPrivKey     *avacrypto.PrivateKey `json:"-"`
	PrivKeyECDSA *ecdsa.PrivateKey     `json:"-"`
	Caddr        common.Address        `json:"caddr"`
	CaddrShort   ids.ShortID           `json:"-"`
	Xaddr        string                `json:"xaddr"`
	XaddrShort   ids.ShortID           `json:"-"`
}

// DecodeB32 format PrivateKey-xxxxx
func DecodeB32(key string, network uint32) (pkInfo *PkInfo, err error) {
	hrp := constants.GetHRP(network)

	avacryptofactor := avacrypto.Factory{}

	pkInfo = &PkInfo{}

	trimmedPrivateKey := strings.TrimPrefix(key, avacrypto.PrivateKeyPrefix)
	privKeyBytes, err := cb58.Decode(trimmedPrivateKey)
	if err != nil {
		return pkInfo, err
	}
	pkInfo.PrivKeyBytes = hex.EncodeToString(privKeyBytes)

	pkInfo.PrivKey = secp256k1.PrivKeyFromBytes(privKeyBytes)

	pkInfo.PrivKeyECDSA = pkInfo.PrivKey.ToECDSA()
	pkInfo.Caddr = crypto.PubkeyToAddress(pkInfo.PrivKeyECDSA.PublicKey)

	copy(pkInfo.CaddrShort[:], pkInfo.Caddr[:])

	pkInfo.XPrivKey, err = avacryptofactor.ToPrivateKey(privKeyBytes)
	if err != nil {
		return pkInfo, err
	}

	pkInfo.XaddrShort = pkInfo.XPrivKey.PublicKey().Address()

	pkInfo.Xaddr, err = address.FormatBech32(hrp, pkInfo.XaddrShort[:])
	if err != nil {
		return pkInfo, err
	}

	return pkInfo, nil
}
