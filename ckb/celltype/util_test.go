package celltype

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
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
	cellData := "0x64617302000000d100000010000000d1000000d1000000c10000001000000014000000180000000600000001000000a5000000a5000000200000002a0000005f0000009400000098000000a0000000a100000000000000000000000000350000001000000030000000310000000000000000000000000000000000000000000000000000000000000000000000000000000035000000100000003000000031000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000004000000"
	cellData = cellData[2:]
	cellDataBytes, err := hex.DecodeString(cellData)
	if err != nil {
		panic(err)
	}
	if das, err := ParseTxWitnessToDasWitnessObj(cellDataBytes); err != nil {
		panic(err)
	} else {
		t.Log(das.WitnessObj.TableType)
		if len(das.MoleculeDepDataEntity.AsSlice()) > 0 {
			panic("dep not empty")
		}
		if len(das.MoleculeNewDataEntity.AsSlice()) == 0 {
			panic("empty")
		}
		if das.MoleculeNewDataEntity.Entity().IsEmpty() {
			panic("empty")
		}
		accountCellData, err := AccountCellDataFromSlice(das.MoleculeNewDataEntity.Entity().RawData(), false)
		if err != nil {
			panic(err)
		}
		t.Log(MoleculeU32ToGo(accountCellData.Status().RawData()))
		_, err = MoleculeU32ToGo(das.MoleculeNewDataEntity.Index().RawData())
		if err != nil {
			panic(err)
		} else {
			t.Log("success")
			// newEntity := das.MoleculeNewDataEntity
			// depEntity := das.MoleculeDepDataEntity
			// if !newEntity.IsEmpty() && (depEntity == nil || depEntity.IsEmpty()) {
			// 	proposeCellData, err := ProposalCellDataFromSlice(newEntity.Entity().RawData(), false)
			// 	if err != nil {
			// 		panic(err)
			// 	}
			// 	lock, err := MoleculeScriptToGo(*proposeCellData.ProposerLock())
			// 	if err != nil {
			// 		panic(err)
			// 	}
			// 	t.Log(lock.CodeHash.String())
			// }
		}
	}
}

func Test_ParseAccountData(t *testing.T) {
	printf := func(hexStr string) {
		hexStr = hexStr[2:]
		bys, _ := hex.DecodeString(hexStr)
		id, err := AccountIdFromOutputData(bys)
		if err != nil {
			panic(err)
		}
		t.Log("id:", id.HexStr())
		nextId, err1 := NextAccountIdFromOutputData(bys)
		if err1 != nil {
			panic(err1)
		}
		t.Log("next:", nextId.HexStr())
	}
	printf("0xf44c70c8921227458a62fe683d84a90baa0386877d3f87b385d7cd3b872e3dfb717ce4f160d8ec367e3fffffffffffffffffffff40fb37620000000031313131313131312e626974")
}

type accountChar struct {
	CharSetName AccountCharType `json:"char_set_name"`
	Bytes       []byte          `json:"bytes"`
}
type accountChars []accountChar

func (chars accountChars) MoleculeAccountChars() AccountChars {
	accountChars := NewAccountCharsBuilder()
	for _, item := range chars {
		if string(item.Bytes) == "." {
			break
		}
		accountChar :=
			NewAccountCharBuilder().
				CharSetName(GoUint32ToMoleculeU32(uint32(item.CharSetName))).
				Bytes(GoBytesToMoleculeBytes(item.Bytes)).
				Build()
		accountChars.Push(accountChar)
		// fmt.Println(string(item.Bytes))
	}
	return accountChars.Build()
}
func Test_RecoverAccountIdFromChars(t *testing.T) {
	const testAccount = DasAccount("22222222.bit")
	t.Log(testAccount.AccountId().HexStr())
	accountChars := accountChars{}
	accountBytes := []byte(testAccount)
	for _, item := range accountBytes {
		accountChars = append(accountChars, accountChar{
			CharSetName: AccountChar_En,
			Bytes:       []byte{item},
		})
	}
	preAccountCellData :=
		NewPreAccountCellDataBuilder().
			Account(accountChars.MoleculeAccountChars()).
			CreatedAt(TimestampDefault()).
			OwnerLock(ScriptDefault()).
			RefundLock(ScriptDefault()).
			InviterWallet(BytesDefault()).
			ChannelWallet(GoBytesToMoleculeBytes([]byte("xx"))).
			Price(PriceConfigDefault()).
			Quote(Uint64Default()).
			Build()
	account := AccountCharsToAccount(*preAccountCellData.Account())
	t.Log(account)
	recover := AccountCharsToAccountId(*preAccountCellData.Account())
	t.Log(recover.HexStr())
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
	// preAccountCell := NewPreAccountCell(TestNetPreAccountCell("",&preAccountCellData))
	witnessBys := NewDasWitnessData(TableType_PRE_ACCOUNT_CELL, preAccountCellData.AsSlice()).ToWitness()
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
	preAccountCell := NewPreAccountCell(TestNetPreAccountCell("", &preAccountCellData))
	witnessBys := NewDasWitnessData(preAccountCell.TableType(), preAccountCellData.AsSlice()).ToWitness()
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

//
func Test_AccountDataFromBytes(t *testing.T) {
	witnessHex := "64617302000000a80300001000000010000000c1010000b101000010000000140000001800000001000000010000009501000095010000200000002a00000073000000bc000000880100009001000091010000d9806cd5bc7d52996b364900000010000000300000003100000058c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a630114000000c9f53b1d85356b60453f867610888d89a0b667ad4900000010000000300000003100000058c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a630114000000c9f53b1d85356b60453f867610888d89a0b667adcc00000024000000390000004e00000063000000780000008d000000a2000000b7000000150000000c00000010000000010000000100000031150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061cc825960000000000004000000e70100001000000014000000180000000100000001000000cb010000cb010000200000002a00000073000000bc000000880100009001000091010000d9806cd5bc7d52996b364900000010000000300000003100000058c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a630114000000c9f53b1d85356b60453f867610888d89a0b667ad4900000010000000300000003100000058c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a630114000000c9f53b1d85356b60453f867610888d89a0b667adcc00000024000000390000004e00000063000000780000008d000000a2000000b7000000150000000c00000010000000010000000100000031150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061cc82596000000000003a0000000800000032000000180000001e00000024000000280000002e000000020000006262020000006363000000000200000064640a000000"
	bys, err := hex.DecodeString(witnessHex)
	if err != nil {
		panic(err)
	}
	obj, err := ParseTxWitnessToDasWitnessObj(bys)
	if err != nil {
		panic(err)
	}
	accountCell, err := AccountCellDataFromSlice(obj.MoleculeNewDataEntity.Entity().RawData(), false)
	if err != nil {
		panic(err)
	}
	recordList := accountCell.Records()
	total := recordList.ItemCount()
	index := uint(0)
	for ; index < total; index++ {
		item := recordList.Get(index)
		fmt.Println("key:", string(item.RecordKey().RawData()))
		fmt.Println(string(item.RecordLabel().RawData()))
		fmt.Println(string(item.RecordType().RawData()))
		fmt.Println(string(item.RecordValue().RawData()))
		fmt.Println(MoleculeU32ToGo(item.RecordTtl().RawData()))
	}
}

func Test_GoUint32ToMoleculeU32(t *testing.T) {
	a := 1
	mu32 := GoUint32ToMoleculeU32(uint32(a))
	t.Log(MoleculeU32ToGo(mu32.RawData()))
}

func Test_ParseTxWitnessToDasWitnessObj_ConfigCellType(t *testing.T) {
	hexStr := "0x646173020000000c020000100000001000000010000000fc0100001000000014000000180000000100000001000000e0010000e0010000200000002a00000073000000bc000000d3010000db010000dc010000b0e9b753b2853a464029490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8011400000020af3b4ed1c7768a8b87d2fc27242c1c3a43d45f490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8011400000020af3b4ed1c7768a8b87d2fc27242c1c3a43d45f1701000030000000450000005a0000006f0000008400000099000000ae000000c3000000d8000000ed00000002010000150000000c0000001000000002000000010000006c150000000c00000010000000020000000100000069150000000c0000001000000002000000010000006e150000000c00000010000000020000000100000067150000000c00000010000000020000000100000075150000000c00000010000000020000000100000061150000000c0000001000000002000000010000006e150000000c00000010000000020000000100000068150000000c0000001000000002000000010000006f150000000c0000001000000002000000010000006e150000000c00000010000000020000000100000067d4ab5960000000000004000000"
	hexStr = hexStr[2:]
	wBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		panic(err)
	}
	obj, err := ParseTxWitnessToDasWitnessObj(wBytes)
	if err != nil {
		panic(err)
	}
	t.Log(obj.WitnessObj.TableType)
	if data, err := AccountCellDataFromSlice(obj.MoleculeNewDataEntity.Entity().RawData(), false); err != nil {
		panic(err)
	} else {
		account := AccountCharsToAccount(*data.Account())
		t.Log(account.Str())
		t.Log(account.AccountId().HexStr())
		t.Log(MoleculeU32ToGo(obj.MoleculeNewDataEntity.Index().RawData()))
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
