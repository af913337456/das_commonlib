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

func Test_ParseTxWitnessToDasWitnessObj_ConfigCellType(t *testing.T) {
	wBytes, err := hex.DecodeString("6461730700000060010000100000001400000018000000008d27002c0100004801000028000000480000006800000088000000a8000000c8000000e80000000801000028010000a2c3a2b18da897bd24391a921956e45d245b46169d6acc9a0663316d15b51cb192d6a9525b9a054222982ab4740be6fe4281e65fff52ab252e7daf9306e12e3f4154b5f9114b8d2dd8323eead5d5e71d0959a2dc73f0672e829ae4dabffdb2d8e79953f024552e6130220a03d2497dc7c2f784f4297c69ba21d0c423915350e5274775e475c1252b5333c20e1512b7b1296c4c5b52a25aa2ebd6e41f5894c41f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009878b226df9465c215fd3c94dc9f9bf6648d5bea48a24579cf83274fe13801d2")
	if err != nil {
		panic(err)
	}
	obj, err := ParseTxWitnessToDasWitnessObj(wBytes)
	if err != nil {
		panic(err)
	}
	if configCellData, err := ConfigCellMainFromSlice(obj.MoleculeNewDataEntity.Entity().RawData(), false); err != nil {
		panic(err)
	} else {
		t.Log(MoleculeU32ToGo(configCellData.AccountExpirationGracePeriod().RawData()))
	}
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
