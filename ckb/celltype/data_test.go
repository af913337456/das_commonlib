package celltype

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/types"
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

func Test_InitSystemScript(t *testing.T) {
	InitSystemScript(SystemScript_ProposeCell, &DASCellBaseInfo{Out: DASCellBaseInfoOut{CodeHash: types.HexToHash("ab")}})
	fmt.Println(DasProposeCellScript.Out.CodeHash.String())
	fmt.Println(SystemCodeScriptMap[SystemScript_ProposeCell].Out.CodeHash.String())
	InitSystemScript(SystemScript_ProposeCell, &DASCellBaseInfo{Out: DASCellBaseInfoOut{CodeHash: types.HexToHash("abcc")}})
	fmt.Println(DasProposeCellScript.Out.CodeHash.String())
}

func Test_Ticker(t *testing.T) {
	go func() {
		ticker := time.NewTicker(time.Second * 2)
		for {
			select {
			case <-ticker.C:
				fmt.Println("1111")
			default:
				time.Sleep(time.Second)
				fmt.Println("2")
			}
		}
	}()
	select {}
}

func Test_AccountCharLen(t *testing.T) {
	// accountId 包含 bit
	// 取价格，不需要
	//
	fmt.Println(len([]rune("xx🌹你")))
	fmt.Println([]byte("🌹"))
	/**
	[
		{
			emoji
			[]byte("🌹")
		},
		{
			en
			[]byte("a")
		},
		{
			zh
			[]byte("你")
		}
	]
	*/
}

func Test_U64Bytes(t *testing.T) {
	d, _ := blake2b.Blake256([]byte("0"))
	t.Log(len(d))
	t.Log(len(GoUint64ToBytes(0)))
}

func Test_AccountChar(t *testing.T) {
	t.Log(len([]byte("account")))
}

func Test_CalAccountCellExpiredAt(t *testing.T) {
	// registerAt:=
	// 2021-01-28 18:02:50, 1611828171
	param := CalAccountCellExpiredAtParam{
		Quote:             1000, // 1000 ckb = 1 usd
		AccountCellCap:    178,
		PriceConfigNew:    1000000, // 10 usd
		PreAccountCellCap: 200,
	}
	timeSec, err := CalAccountCellExpiredAt(param, time.Now().Unix())
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
