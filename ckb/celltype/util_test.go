package celltype

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"math/big"
	"testing"
	"time"
)

/**
 * Copyright (C), 2019-2020
 * FileName: util_test
 * Author:   LinGuanHong
 * Date:     2020/12/27 12:09 下午
 * Description:
 */

func Test_RecoverData_From_BuildDasCommonMoleculeDataObj(t *testing.T) {
	createAt := NewTimestampBuilder().
		Set(GoTimeUnixToMoleculeBytes(time.Now().Unix())).Build()
	inviterAccountId := GoBytesToMoleculeBytes(DasAccountFromStr("xxx.bit").AccountId().Bytes())
	preAccountCellData :=
		NewPreAccountCellDataBuilder().
			Account(AccountCharsDefault()).
			CreatedAt(createAt).
			OwnerLock(ScriptDefault()).
			RefundLock(GoCkbScriptToMoleculeScript(types.Script{
				CodeHash: types.HexToHash("123456aa"),
				HashType: types.HashTypeType,
				Args:     nil,
			})).
			InviterWallet(inviterAccountId).
			ChannelWallet(GoBytesToMoleculeBytes([]byte("xx"))).
			Price(PriceConfigDefault()).
			Quote(GoUint64ToMoleculeU64(10086)).
			Build()
	preAccountCell := NewPreAccountCell(TestNetPreAccountCell(0, 0, 0, nil, nil, &preAccountCellData))
	tableBys := preAccountCell.TableData()
	data := DataFromSliceUnchecked(tableBys)
	ret := &ParseDasWitnessBysDataObj{}
	ret.MoleculeData = data
	if data.Dep().IsNone() {
		ret.MoleculeDepDataEntity = nil
	} else {
		ret.MoleculeDepDataEntity = DataEntityFromSliceUnchecked(data.Dep().AsSlice())
	}
	if data.Old().IsNone() {
		ret.MoleculeOldDataEntity = nil
	} else {
		ret.MoleculeOldDataEntity = DataEntityFromSliceUnchecked(data.Old().AsSlice())
	}
	ret.MoleculeNewDataEntity = DataEntityFromSliceUnchecked(data.New().AsSlice())
	if preAccountCellData, err := PreAccountCellDataFromSlice(ret.MoleculeNewDataEntity.Entity().AsSlice(), false); err != nil {
		panic(err)
	} else {
		t.Log(string(preAccountCellData.ChannelWallet().RawData()))
		script, err := MoleculeScriptToGo(*preAccountCellData.RefundLock())
		if err != nil {
			panic(err)
		}
		t.Log(script.CodeHash.String())
		t.Log(MoleculeU64ToGo(preAccountCellData.Quote().RawData()))
	}
}

func Test_PreAccountDataFromBytes(t *testing.T) {
	witnessHex := "64617306000000e9010000100000001000000010000000d90100001000000014000000180000000000000001000000c101000024000000220100004b010000740100008201000090010000b1010000b9010000fe0000002c00000041000000560000006b0000008000000095000000aa000000bf000000d4000000e9000000150000000c0000001000000001000000010000006c150000000c00000010000000010000000100000067150000000c00000010000000010000000100000068150000000c0000001000000001000000010000005f150000000c00000010000000010000000100000031150000000c00000010000000010000000100000032150000000c0000001000000001000000010000002e150000000c00000010000000010000000100000062150000000c00000010000000010000000100000069150000000c0000001000000001000000010000007429000000100000001000000011000000011400000020af3b4ed1c7768a8b87d2fc27242c1c3a43d45f29000000100000001000000011000000011400000020af3b4ed1c7768a8b87d2fc27242c1c3a43d45f0a000000ddd1866d906c5d39824b0a000000ddd1866d906c5d39824b2100000010000000110000001900000006c0cf6a000000000060ae0a0000000000c73500000000000058f7496000000000"
	bys, err := hex.DecodeString(witnessHex)
	if err != nil {
		panic(err)
	}
	obj, err := ParseTxWitnessToDasWitnessObj(bys)
	if err != nil {
		panic(err)
	}
	preAccountCell, err := PreAccountCellDataFromSlice(obj.MoleculeNewDataEntity.Entity().RawData(), false)
	if err != nil {
		panic(err)
	}
	script, err := MoleculeScriptToGo(*preAccountCell.OwnerLock())
	if err != nil {
		panic(err)
	}
	t.Log(hex.EncodeToString(script.Args))
}

func Test_ParseTxWitnessToDasWitnessObj_ConfigCellType(t *testing.T) {
	wBytes, err := hex.DecodeString("64617307000000540000000c000000300000002400000014000000180000001c00000020000000000000000000000080510100e80300002400000014000000180000001c00000020000000008d2700008d270080510100e8030000")
	if err != nil {
		panic(err)
	}
	obj, err := ParseTxWitnessToDasWitnessObj(wBytes)
	if err != nil {
		panic(err)
	}
	t.Log(obj.WitnessObj.TableType)
	// if configCellData, err := ConfigCellRegisterFromSlice(obj.MoleculeNewDataEntity.Entity().RawData(), false); err != nil {
	// 	panic(err)
	// } else {
	// 	t.Log(MoleculeU32ToGo(configCellData.ApplyMinWaitingBlockNumber().RawData()))
	// }
}

func Test_GoTimestampToMoleculeBytes(t *testing.T) {
	timeNowSec := time.Now().Unix()
	t.Log(timeNowSec)
	ret := GoTimeUnixToMoleculeBytes(timeNowSec)
	_mt := NewTimestampBuilder().Set(ret).Build()
	_rd := _mt.RawData()
	t.Log(byteToInt64(_rd))
	t.Log(new(big.Int).SetBytes(_rd).String())
}

func byteToInt64(bys []byte) int64 {
	return int64(binary.LittleEndian.Uint64(bys))
}
