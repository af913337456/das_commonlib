package wallet

import (
	"fmt"
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
	Secp256k1Key *secp256k1.Secp256k1Key
	LockScript   *types.Script
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
		Secp256k1Key: key,
		LockScript:   lockScript,
	}, nil
}
