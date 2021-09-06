package rule712

import (
	"encoding/hex"
	"encoding/json"
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

func Test_AppendAccountCellInput(t *testing.T) {
	list := InputOutputParam712List{}
	list.AppendAccountCellInput(&gotype.AccountCell{
		CellCap:       123,
		AccountId:     celltype.DasAccountId{},
		Status:        0,
		Point:         types.OutPoint{},
		WitnessStatus: 0,
		Data:          nil,
		WitnessData:   nil,
		DasLockArgs:   nil,
	})
	fmt.Println(list[0].Capacity)
}

func Test_MMJsonBuild(t *testing.T) {
	inputs := InputOutputParam712List{
		{
			Capacity:  1212121212121,
			Lock:      nil,
			Type:      nil,
			Data:      []byte("123424232"),
		},
	}
	m := &MMJson{
		action:          "{\"action\": \"transfer_account\", \"params\": \"0x1111111111,0x222222222222,0x3333333333333\"}",
		fee:             10000201021,
		inputsCapacity:  45621163888,
		outputsCapacity: 2978378266,
		plainText:       "123",
		digest:          "123",
		inputs:          "",
		outputs:         "",
	}
	_ = m.FillInputs(inputs,nil)
	_ = m.FillOutputs(inputs,nil)
	obj,err := m.Build()
	fmt.Println(err)
	bys,_ := json.MarshalIndent(obj," "," ")
	fmt.Println(string(bys))
}

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
	new(MMJson).FillWithdrawPlainText(true,inputs,output)
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

