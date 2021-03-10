package wallet

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/bech32"
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

func Test_InitWallet2(t *testing.T) {
	systemScript := &utils.SystemScripts{
		SecpSingleSigCell: &utils.SystemScriptCell{
			CellHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
			OutPoint: nil,
			HashType: "",
			DepType:  "",
		},
		SecpMultiSigCell: nil,
		DaoCell:          nil,
		ACPCell:          nil,
		SUDTCell:         nil,
		ChequeCell:       nil,
	}
	key, err := InitCkbWallet("1504c89d50057bcef660251abc4c75ca28f4ed9139cd32611a78f69559fb5168", systemScript)
	if err != nil {
		panic(err)
	}
	t.Log(hex.EncodeToString(key.LockScript.Args))
}

func Test_InitWallet(t *testing.T) {
	systemScript := &utils.SystemScripts{
		SecpSingleSigCell: &utils.SystemScriptCell{
			CellHash: types.HexToHash("3419a1c09eb2567f6552ee7a8ecffd64155cffe0f1796e6e61ec088d740c1356"),
			OutPoint: nil,
			HashType: "",
			DepType:  "",
		},
		SecpMultiSigCell: nil,
		DaoCell:          nil,
		ACPCell:          nil,
		SUDTCell:         nil,
		ChequeCell:       nil,
	}
	key, err := InitCkbWallet("1504c89d50057bcef660251abc4c75ca28f4ed9139cd32611a78f69559fb5168", systemScript)
	if err != nil {
		panic(err)
	}
	lockScript, err := key.Secp256k1Key.Script(systemScript)
	if err != nil {
		panic(err)
	}
	bys, err := lockScript.Serialize()
	if err != nil {
		panic(err)
	}
	address, err := bech32.Encode("ckb", bys)
	if err != nil {
		panic(err)
	}
	t.Log(address)
}

func Test_GetAddress(t *testing.T) {
	bs, err := hex.DecodeString("b39bbc0b3673c7d36450bc14cfcdad2d559c6c64")
	if err != nil {
		panic(err)
	}
	typebin, _ := hex.DecodeString("01")
	flag, _ := hex.DecodeString("00")

	payload := append(typebin, flag...)
	payload = append(payload, bs...)

	converted, err := bech32.ConvertBits(payload, 8, 5, true)
	if err != nil {
		panic(err)
	}
	address, err := bech32.Encode("ckb", converted)

	if err != nil {
		panic(err)
	}
	t.Log(address)
}

func Test_AddrToArgs(t *testing.T) {
	t.Log(GetLockScriptArgsFromShortAddress("ckb1qyqt8xaupvm8837nv3gtc9x0ekkj64vud3jqfwyw5v"))
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
	rawMsgHexBys := csprngEntropy(32)
	signMsg, err := key.Secp256k1Key.Sign(rawMsgHexBys)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(key.Secp256k1Key.PubKey()))
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
