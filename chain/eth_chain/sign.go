package eth_chain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignByPrivateKey(hexKey, msg string) ([]byte, error) {
	key, _ := crypto.HexToECDSA(hexKey)
	data := crypto.Keccak256([]byte(msg))
	return crypto.Sign(data, key)
}

func VerifySign(hexAddr, msg string, sig []byte) (bool, error) {
	data := crypto.Keccak256([]byte(msg))
	addr := common.HexToAddress(hexAddr)
	pubKey, err := crypto.SigToPub(data, sig)
	if err != nil {
		return false, err
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return addr == recoveredAddr, nil
}
