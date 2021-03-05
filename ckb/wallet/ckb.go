package wallet

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	ethSecp256k1 "github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/bech32"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
)

/**
 * Copyright (C), 2019-2020
 * FileName: contract_owner
 * Author:   LinGuanHong
 * Date:     2020/12/21 10:10 下午
 * Description:
 */

type CkbWalletObj struct {
	SystemScripts *utils.SystemScripts
	Secp256k1Key  *secp256k1.Secp256k1Key
	LockScript    *types.Script
}

func InitCkbWallet(privateKeyHex string, systemScript *utils.SystemScripts) (*CkbWalletObj, error) {
	key, err := secp256k1.HexToKey(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("InitCkbWallet HexToKey err: %s", err.Error())
	}
	lockScript, err := key.Script(systemScript)
	if err != nil {
		return nil, fmt.Errorf("InitCkbWallet LockScript err: %s", err.Error())
	}
	return &CkbWalletObj{
		Secp256k1Key:  key,
		LockScript:    lockScript,
		SystemScripts: systemScript,
	}, nil
}

func GetShortAddressFromLockScriptArgs(args string) {

}

func GetLockScriptArgsFromShortAddress(address string) (string, error) {
	_, bys, err := bech32.Decode(address)
	if err != nil {
		return "", fmt.Errorf("bech32.Decode err: %s", err.Error())
	}
	converted, err := bech32.ConvertBits(bys, 5, 8, false)
	if err != nil {
		return "", fmt.Errorf("bech32.ConvertBits err: %s", err.Error())
	}
	ret := hex.EncodeToString(converted)[4:]
	if len(ret) != 20 {
		return "", errors.New("invalid args len")
	}
	return ret, nil
}

func VerifySign(msg []byte, sign []byte, ckbPubkeyHex string) (bool, error) {
	recoveredPub, err := crypto.Ecrecover(msg, sign)
	if err != nil {
		return false, err
	}
	pubKey, err := crypto.UnmarshalPubkey(recoveredPub)
	if err != nil {
		return false, err
	}
	return hex.EncodeToString(ethSecp256k1.CompressPubkey(pubKey.X, pubKey.Y)) == ckbPubkeyHex, nil
}
