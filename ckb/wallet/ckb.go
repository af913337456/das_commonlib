package wallet

import (
	"encoding/hex"
	"fmt"
	ethSecp256k1 "github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
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

func VerifySign(msg []byte, sign []byte, ckbPubkeyHex string) (bool, error) {
	pubKey, err := ethSecp256k1.RecoverPubkey(msg, sign)
	if err != nil {
		return false, err
	}
	bys, err := blake2b.Blake160(pubKey)
	if err != nil {
		return false, err
	}
	return hex.EncodeToString(bys) == ckbPubkeyHex, nil
}
