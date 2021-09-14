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
			Capacity: 1212121212121,
			Lock:     nil,
			Type:     nil,
			Data:     []byte("123424232"),
		},
	}
	m := &MMJson{
		action:          "{\"action\": \"edit_records\",\"params\": \"0x01\"}",
		fee:             10000201021,
		inputsCapacity:  45621163888,
		outputsCapacity: 2978378266,
		dasMessage:      "123",
		digest:          "123",
		inputs:          "[{\"capacity\":\"221. CKB\",\"lock\":\"das-lock,0x01,0x0019b04faf5b6e76e6d6640344b23dc16ffd9010...\",\"type\":\"account-cell-type,0x01,0x\",\"data\":\"{ account: linguaniii.bit, expired_at: 1661877499 }\",\"extraData\":\"{ status: 0, records_hash: 0x55478d76900611eb079b22088081124ed6c8bae21a05dd1a0d197efcc7c114ce }\"}]",
		outputs:         "[{\"capacity\":\"220.9999 CKB\",\"lock\":\"das-lock,0x01,0x0019b04faf5b6e76e6d6640344b23dc16ffd9010...\",\"type\":\"account-cell-type,0x01,0x\",\"data\":\"{ account: linguaniii.bit, expired_at: 1661877499 }\",\"extraData\":\"{ status: 0, records_hash: 0xa34bb356af1a5260ff86dfbba27d74bf697a453ee6d44a6517cfe62cf8f0e94a }\"}]",
	}
	_ = m.FillInputs(inputs, nil)
	_ = m.FillOutputs(inputs, nil)
	obj, err := m.Build()
	fmt.Println("build err:", err)
	bys, _ := json.MarshalIndent(obj, " ", " ")
	fmt.Println(string(bys))
}

func Test_CreateWithdrawPlainText(t *testing.T) {
	bys, _ := hex.DecodeString("00dc36477cf2434288a5502120ef0fd919ae37c15500dc36477cf2434288a5502120ef0fd919ae37c155")
	bys3, _ := hex.DecodeString("001c36477cf2434288a5502120ef0fd919ae37c15500dc36477cf2434288a5502120ef0fd919ae37c155")
	inputs := []gotype.WithdrawDasLockCell{
		{
			OutPoint:       nil,
			LockScriptArgs: bys,
			CellCap:        12345678,
		},
		{
			OutPoint:       nil,
			LockScriptArgs: bys3,
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
	bys2, _ := hex.DecodeString("dc36477cf2434288a5502120ef0fd919ae37c155")
	output := WithdrawPlainTextOutputParam{
		ReceiverCkbScript: types.Script{
			CodeHash: types.Hash{},
			HashType: "",
			Args:     bys2,
		},
		Amount: 9863781321,
	}
	mmjson := new(MMJson)
	mmjson.FillWithdrawDasMessage(true, inputs, output)
	fmt.Println(mmjson)
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
	fmt.Println(removeSuffixZeroChar(new(big.Rat).
		Quo(
			new(big.Rat).SetInt(new(big.Int).SetUint64(10*celltype.OneCkb)),
			new(big.Rat).SetInt(new(big.Int).SetUint64(celltype.OneCkb))).
		FloatString(8)))
}

func TestJson(t *testing.T) {
	//{
	//	"capacity":"225 CKB",
	//	"data":"{ account: tangzhihong005.bit, expired_at: 1662629612 }",
	//	"extraData":"{ status: 0, records_hash: 0x55478d76900611eb079b22088081124ed6c8bae21a05dd1a0d197efcc7c114ce }",
	//	"lock":"das-lock,0x01,0x0515a33588908cf8edb27d1abe3852bf287abd38...",
	//	"type":"account-cell-type,0x01,0x"
	//}
	var retList []inputOutputParam712
	var ijson inputOutputParam712
	ijson.Capacity = "225 CKB"
	ijson.Data = "{ account: tangzhihong005.bit, expired_at: 1662629612 }"
	ijson.ExtraData = "{ status: 0, records_hash: 0x55478d76900611eb079b22088081124ed6c8bae21a05dd1a0d197efcc7c114ce }"
	ijson.LockStr = "das-lock,0x01,0x0515a33588908cf8edb27d1abe3852bf287abd38..."
	ijson.TypeStr = "account-cell-type,0x01,0x"
	retList = append(retList, ijson, ijson)
	data, _ := json.Marshal(retList)
	fmt.Println(string(data))
}

func TestCkbValueStr(t *testing.T) {
	capStr := ckbValueStr(15450790000)
	fmt.Println(capStr)
	fmt.Println(removeSuffixZeroChar(capStr))
}
