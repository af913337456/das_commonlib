package eth_chain

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core"
	"strings"
)

func VerifySignEIP712(address, signHex, jsonStr string) error {
	var typedData core.TypedData
	if err := json.Unmarshal([]byte(jsonStr), &typedData); err != nil {
		return err
	}
	typedDataHash, _ := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	domainSeparator, _ := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash := crypto.Keccak256Hash(rawData)
	userAddress := common.HexToAddress(address)
	if strings.HasPrefix(signHex, "0x") {
		signHex = signHex[2:]
	}
	signature, _ := hex.DecodeString(signHex)
	if len(signature) != 65 {
		return fmt.Errorf("invalid signature length: %d", len(signature))
	}
	if signature[64] != 27 && signature[64] != 28 {
		return fmt.Errorf("invalid recovery id: %d", signature[64])
	}
	signature[64] -= 27
	pubKeyRaw, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		return fmt.Errorf("invalid signature: %s", err.Error())
	}
	pubKey, err := crypto.UnmarshalPubkey(pubKeyRaw)
	if err != nil {
		return err
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	if !bytes.Equal(userAddress.Bytes(), recoveredAddr.Bytes()) {
		return fmt.Errorf("addresses do not match")
	}
	return nil
}
