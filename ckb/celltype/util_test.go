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

func Test_ParseProposeCellData(t *testing.T) {
	cellData := "0x646173000000001b0000000c000000170000000700000070726f706f736500000000"
	cellData = cellData[2:]
	cellDataBytes, err := hex.DecodeString(cellData)
	if err != nil {
		panic(err)
	}
	if das, err := ParseTxWitnessToDasWitnessObj(cellDataBytes); err != nil {
		panic(err)
	} else {
		index, err := MoleculeU32ToGo(das.MoleculeNewDataEntity.Index().RawData())
		if err != nil {
			panic(err)
		} else {
			t.Log(index)
		}
	}
}

func Test_CreateData(t *testing.T) {
	preAccountCellData :=
		NewPreAccountCellDataBuilder().
			Account(AccountCharsDefault()).
			CreatedAt(TimestampDefault()).
			OwnerLock(ScriptDefault()).
			RefundLock(ScriptDefault()).
			InviterWallet(BytesDefault()).
			ChannelWallet(GoBytesToMoleculeBytes([]byte("xx"))).
			Price(PriceConfigDefault()).
			Quote(Uint64Default()).
			Build()
	// new := NewDataEntityBuilder().
	// 	Index(GoUint32ToMoleculeU32(0)).
	// 	Version(GoUint32ToMoleculeU32(1)).
	// 	Entity(GoBytesToMoleculeBytes(preAccountCellData.AsSlice())).
	// 	Build()
	// d := NewDataBuilder().
	// 	Dep(DataEntityOptDefault()).
	// 	Old(DataEntityOptDefault()).
	// 	New(NewDataEntityOptBuilder().Set(new).Build()).
	// 	Build()
	// s := hex.EncodeToString(d.AsSlice())
	// t.Log(s)
	preAccountCell := NewPreAccountCell(TestNetPreAccountCell("", 0, 0, 0, nil, nil, &preAccountCellData))
	witnessBys := NewDasWitnessData(TableType_PRE_ACCOUNT_CELL, preAccountCell.TableData()).ToWitness()
	ret, err := ParseTxWitnessToDasWitnessObj(witnessBys)
	if err != nil {
		panic(err)
	}
	// rawData
	if preAccountCellData, err := PreAccountCellDataFromSlice(ret.MoleculeNewDataEntity.Entity().RawData(), false); err != nil {
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

func Test_RecoverData_From_BuildDasCommonMoleculeDataObj(t *testing.T) {
	createAt := NewTimestampBuilder().
		Set(GoTimeUnixToMoleculeBytes(time.Now().Unix())).Build()

	accountChars := NewAccountCharsBuilder()
	chars := []byte("iqyueq.bit")
	for _, item := range chars {
		accountChar :=
			NewAccountCharBuilder().
				CharSetName(GoUint32ToMoleculeU32(uint32(AccountChar_En))).
				Bytes(GoBytesToMoleculeBytes([]byte{item})).
				Build()
		accountChars.Push(accountChar)
	}

	inviterAccountId := GoBytesToMoleculeBytes(DasAccountFromStr("xxx.bit").AccountId().Bytes())
	args, _ := hex.DecodeString("b7526803f67ebe70aba6")
	preAccountCellData :=
		NewPreAccountCellDataBuilder().
			Account(accountChars.Build()).
			CreatedAt(createAt).
			OwnerLock(ScriptDefault()).
			RefundLock(GoCkbScriptToMoleculeScript(types.Script{
				CodeHash: types.HexToHash("123456aa"),
				HashType: types.HashTypeType,
				Args:     args,
			})).
			InviterWallet(inviterAccountId).
			ChannelWallet(GoBytesToMoleculeBytes([]byte("xx"))).
			Price(PriceConfigDefault()).
			Quote(GoUint64ToMoleculeU64(10086)).
			Build()
	preAccountCell := NewPreAccountCell(TestNetPreAccountCell("", 0, 0, 0, nil, nil, &preAccountCellData))
	witnessBys := NewDasWitnessData(preAccountCell.TableType(), preAccountCell.TableData()).ToWitness()
	ret, err := ParseTxWitnessToDasWitnessObj(witnessBys)
	if err != nil {
		panic(err)
	}
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
	witnessHex := "646173060000005f0200001000000010000000100000004f0200001000000014000000180000000000000001000000330200003302000024000000540100009d010000e6010000f401000002020000230200002b0200003001000034000000490000005e00000073000000880000009d000000b2000000c7000000dc000000f1000000060100001b010000150000000c00000010000000010000000100000031150000000c00000010000000010000000100000032150000000c00000010000000010000000100000033150000000c00000010000000010000000100000034150000000c00000010000000010000000100000035150000000c00000010000000010000000100000036150000000c00000010000000010000000100000037150000000c00000010000000010000000100000038150000000c0000001000000001000000010000002e150000000c00000010000000010000000100000062150000000c00000010000000010000000100000069150000000c00000010000000010000000100000074490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8011400000020af3b4ed1c7768a8b87d2fc27242c1c3a43d45f490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8011400000020af3b4ed1c7768a8b87d2fc27242c1c3a43d45f0a000000a8e756d8d2c5f06832240a000000a8e756d8d2c5f06832242100000010000000110000001900000008404b4c000000000020a1070000000000c735000000000000e1de4a6000000000"
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
