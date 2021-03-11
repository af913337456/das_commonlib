package celltype

import (
	"encoding/binary"
	"encoding/hex"
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

func Test_PreAccountDataFromBytes(t *testing.T) {
	witnessHex := "64617309000000080000400000000000012004008200000001080000400208000400820100010000003000001200000000002000140a000008000400000800003202004090000200000000010110000000040000100400100080a14001a00110000000000400000000000000402000007042008000202000000080008050120000001000000800200100000002100000000000010001000600010800800000008a4000300020000000000000100081000000400600000080000020"
	bys, _ := hex.DecodeString(witnessHex)
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
