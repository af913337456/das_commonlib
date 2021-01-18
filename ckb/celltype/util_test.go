package celltype

import (
	"encoding/binary"
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
