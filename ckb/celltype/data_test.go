package celltype

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"testing"
	"time"
)

/**
 * Copyright (C), 2019-2020
 * FileName: data_test
 * Author:   LinGuanHong
 * Date:     2020/12/20 2:57 下午
 * Description:
 */

func Test_CalAccountCellExpiredAt(t *testing.T) {
	// registerAt:=
	// 2021-01-28 18:02:50, 1611828171
	param := CalAccountCellExpiredAtParam{
		Quote:              1,
		AccountCellCap:     178,
		AccountConfigPrice: 10,
		AccountBytesLen:    uint32(len([]byte("nice.bit"))),
		PreAccountCellCap:  300,
	}
	timeSec, err := CalAccountCellExpiredAt(param, 1611828171)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(timeSec)
		fmt.Println(time.Unix(int64(timeSec), 0).String())
		bys, _ := json.Marshal(param)
		t.Log(string(bys))
	}
}

func Test_Blake2b_256(t *testing.T) {
	bys, _ := blake2b.Blake256([]byte("lafdqalqvbappo"))
	t.Log(len(bys), bys)
}

func Test_ParseActionCell(t *testing.T) {
	hexStr := "1e000000646173001a0000000c0000001600000006000000636f6e66696700000000"
	bys, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatal(err)
	}
	if witness, err := NewDasWitnessDataFromSlice(bys); err != nil {
		t.Fatal(err)
	} else {
		t.Log(witness.Tag, witness.TableType)
	}
}

// func Test_StateCellData(t *testing.T) {
// 	stateCell := NewStateCellDataBuilder()
// 	rootHash := HashFromSliceUnchecked([]byte("hello world!h"))
// 	stateCell.ReservedAccountRoot(*rootHash)
// 	// dataBytes := stateCell.Build()
// 	raw := string(stateCell.reserved_account_root.AsSlice())
// 	t.Log("raw ===> ", raw)
// 	t.Log("rawHex ===> ", hex.EncodeToString(stateCell.reserved_account_root.RawData()))
//
// }
