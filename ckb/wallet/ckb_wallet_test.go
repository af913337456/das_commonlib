package wallet

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
	"io"
	"testing"
)

/**
 * Copyright (C), 2019-2020
 * FileName: ckb_wallet_test
 * Author:   LinGuanHong
 * Date:     2020/12/24 9:56 上午
 * Description:
 */

func Test_InitWallet(t *testing.T) {
	key, err := InitCkbWallet("", nil)
	if err != nil {
		panic(err)
	}
	t.Log(key.Secp256k1Key.PubKey())
}

func Test_VerifySign(t *testing.T) {
	key, err := InitCkbWallet(
		"1504c89d50057bcef660251abc4c75ca28f4ed9139cd32611a78f69559fb5168",
		&utils.SystemScripts{
			SecpSingleSigCell: &utils.SystemScriptCell{
				CellHash: types.HexToHash("123"),
				OutPoint: nil,
				HashType: "",
				DepType:  "",
			},
			SecpMultiSigCell: nil,
			DaoCell:          nil,
			ACPCell:          nil,
			SUDTCell:         nil,
			ChequeCell:       nil,
		})
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(key.LockScript.Args))
	// c75fd5f8add2a04db9ffcaf88b437d76f1812797
	rawMsgHexBys := csprngEntropy(32)
	signMsg, err := key.Secp256k1Key.Sign(rawMsgHexBys)
	if err != nil {
		panic(err)
	}
	fmt.Println("sign hex:", hex.EncodeToString(signMsg))
	pass, err := VerifySign(rawMsgHexBys, signMsg, hex.EncodeToString(key.Secp256k1Key.PubKey()))
	if err != nil {
		panic(err)
	}
	fmt.Println("verify sign:", pass)
}

func csprngEntropy(n int) []byte {
	buf := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}
	return buf
}
