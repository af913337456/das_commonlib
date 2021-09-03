package rule712

import (
	"encoding/hex"
	"fmt"
	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
	"github.com/DeAccountSystems/das_commonlib/ckb/gotype"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"math/big"
	"testing"
)

/**
 * Copyright (C), 2019-2021
 * FileName: 712_test
 * Author:   LinGuanHong
 * Date:     2021/9/3 3:02
 * Description:
 */

func Test_CreateWithdrawPlainText(t *testing.T) {
	bys,_ := hex.DecodeString("00dc36477cf2434288a5502120ef0fd919ae37c15500dc36477cf2434288a5502120ef0fd919ae37c155")
	inputs := []gotype.WithdrawDasLockCell{
		{
			OutPoint:       nil,
			LockScriptArgs: bys,
			CellCap:        12345678,
		},
		{
			OutPoint:       nil,
			LockScriptArgs: bys,
			CellCap:        12345678,
		},
		{
			OutPoint:       nil,
			LockScriptArgs: bys,
			CellCap:        12345678,
		},
		{
			OutPoint:       nil,
			LockScriptArgs: bys,
			CellCap:        12345678,
		},
	}
	bys2,_ := hex.DecodeString("dc36477cf2434288a5502120ef0fd919ae37c155")
	output := WithdrawPlainTextOutputParam{
		ReceiverCkbScript: types.Script{
			CodeHash: types.Hash{},
			HashType: "",
			Args:     bys2,
		},
		Amount: 9863781321,
	}
	str := CreateWithdrawPlainText(true,inputs,output)
	fmt.Println(str)
}

func Test_QuoCkbValue(t *testing.T) {
	fmt.Println(removeSuffixZeroChar(new(big.Rat).
		Quo(
			new(big.Rat).SetInt(new(big.Int).SetUint64(1234560001102000)),
			new(big.Rat).SetInt(new(big.Int).SetUint64(celltype.OneCkb))).
		FloatString(8)))
	fmt.Println(removeSuffixZeroChar(new(big.Rat).
		Quo(
			new(big.Rat).SetInt(new(big.Int).SetUint64(1234560001102)),
			new(big.Rat).SetInt(new(big.Int).SetUint64(celltype.OneCkb))).
		FloatString(8)))
	fmt.Println(removeSuffixZeroChar(new(big.Rat).
		Quo(
			new(big.Rat).SetInt(new(big.Int).SetUint64(1234560001102010)),
			new(big.Rat).SetInt(new(big.Int).SetUint64(celltype.OneCkb))).
		FloatString(8)))
	fmt.Println(removeSuffixZeroChar(new(big.Rat).
		Quo(
			new(big.Rat).SetInt(new(big.Int).SetUint64(1102000)),
			new(big.Rat).SetInt(new(big.Int).SetUint64(celltype.OneCkb))).
		FloatString(8)))
	fmt.Println(removeSuffixZeroChar(new(big.Rat).
		Quo(
			new(big.Rat).SetInt(new(big.Int).SetUint64(11020001)),
			new(big.Rat).SetInt(new(big.Int).SetUint64(celltype.OneCkb))).
		FloatString(8)))
}

