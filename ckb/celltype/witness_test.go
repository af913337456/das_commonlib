package celltype

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"math/big"
	"strings"
	"testing"
	"time"
)

/**
 * Copyright (C), 2019-2020
 * FileName: util_test
 * Author:   LinGuanHong
 * Date:     2020/12/27 12:09
 * Description:
 */

func Test_FindTargetTypeScriptByInputList(t *testing.T) {
	inputList := []*types.CellInput{
		{
			Since: 0,
			PreviousOutput: &types.OutPoint{
				TxHash: types.HexToHash("0xd6590562d4b6ac365399575611e83c8ab86e09429d6fb36846ee15d8febcc8c4"),
				Index:  0,
			},
		},
	}
	ret,err := FindTargetTypeScriptByInputList(&ReqFindTargetTypeScriptParam{
		Ctx:       context.TODO(),
		RpcClient: rpcClient(),
		InputList: inputList,
		IsLock:    false,
		CodeHash:  DasProposeCellScript.Out.CodeHash,
	})
	if err != nil {
		panic(err)
	}
	err = GetTargetCellFromWitness(ret.Tx, func(rawWitnessData []byte, witnessParseObj *ParseDasWitnessBysDataObj) (bool, error) {
		witnessDataObj := witnessParseObj.WitnessObj
		switch witnessDataObj.TableType {
		case TableType_PROPOSE_CELL:
			t.Log("found!",hex.EncodeToString(rawWitnessData))
			return true, nil
		}
		return false, nil
	}, func(err error) {
		t.Log(err.Error())
	})
	if err != nil {
		panic(err)
	}
}

func Test_ParseCellData(t *testing.T) {
	cellData := "0x64617304000000a10400001000000010000000100000009104000010000000140000001800000000000000010000007504000075040000100000005900000061000000490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce801140000000daf676ba7d268446c6e79da5c58dda393f0be2500000000000000001404000024000000a2000000200100009e0100001c0200009a02000018030000960300007e0000000c00000045000000390000001000000024000000250000001b8cc4589ac521e9012aa5897b02e66ba451fee8001ba67e59c90bb93d192ccdde189b05e460adc0be390000001000000024000000250000001ba67e59c90bb93d192ccdde189b05e460adc0be021ba73bd4911823f9fbaeb01a5edf06142d60848d7e0000000c00000045000000390000001000000024000000250000008bc7aca87fd6683a6dd1938a29824893850cd6f7008bcb51b36347ddc7286c165f94c440a75d06de3d390000001000000024000000250000008bcb51b36347ddc7286c165f94c440a75d06de3d028bdd0f376d9ea1a1f9d65b4be1614d3bbbd637d77e0000000c00000045000000390000001000000024000000250000008bfad4b1328ad88a3f37e96848c690518eb60ea5008c00359fb689f81e35ee98a4d159e6a975984667390000001000000024000000250000008c00359fb689f81e35ee98a4d159e6a975984667028c01bfa41d36f76b52a1851e3af62b29316ebaec7e0000000c00000045000000390000001000000024000000250000008de0ea9a6ca11543b2ce6a0be078fca284fa143c008decee1865de003f0c84315c8533b524ec6d9f66390000001000000024000000250000008decee1865de003f0c84315c8533b524ec6d9f66028e1a5abb4e87095ada530482d15bbb06e0101cea7e0000000c00000045000000390000001000000024000000250000009fe757403c2302af8082faedad9152786e90d1f9009febcae703695b2f0f157c5e5311ca89773243ce390000001000000024000000250000009febcae703695b2f0f157c5e5311ca89773243ce029fefbe339dcd8916240bb29c74991ea0bbf127c37e0000000c0000004500000039000000100000002400000025000000cd0b1a1efb26770098e27a8e25e4b576c4abade500cd17ff8dec8216fc62bc5c75e769533710e2cd2439000000100000002400000025000000cd17ff8dec8216fc62bc5c75e769533710e2cd2402cd1a97f68c3f35c40b5db70a4332240a88fd0c2b7e0000000c0000004500000039000000100000002400000025000000ecb623d8418feb37086b959eb1b95a9400647b2c00ecb62404a5b5f572cca7c25bd19b19181a38528039000000100000002400000025000000ecb62404a5b5f572cca7c25bd19b19181a38528002ecbf880a753dca8b75cb89ddd1214a64e1021bd57e0000000c0000004500000039000000100000002400000025000000f18c985f090f80f2bf2b9b32fff79a3a77f5f04200f19429d4aed43a822d13e5cb2b70a22a2c2fc2e039000000100000002400000025000000f19429d4aed43a822d13e5cb2b70a22a2c2fc2e002f19bf464a527548675960e489831b140365912ec"
	if strings.HasPrefix(cellData,"0x") {
		cellData = cellData[2:]
	}
	cellDataBytes, err := hex.DecodeString(cellData)
	if err != nil {
		panic(err)
	}
	if das, err := ParseTxWitnessToDasWitnessObj(cellDataBytes); err != nil {
		panic(err)
	} else {
		t.Log(das.WitnessObj.TableType,string(das.WitnessObj.TableBys))
		if len(das.MoleculeNewDataEntity.AsSlice()) == 0 {
			panic("empty 1")
		}
		if das.MoleculeNewDataEntity.Entity().IsEmpty() {
			panic("empty 2")
		}
		cellData, err := ProposalCellDataFromSlice(das.MoleculeNewDataEntity.Entity().RawData(), false)
		if err != nil {
			panic(err)
		}
		fmt.Println("entity hex:",hex.EncodeToString(cellData.AsSlice()))
		bys,err := blake2b.Blake256(cellData.AsSlice())
		if err != nil {
			panic(err)
		}
		fmt.Println("hex:",hex.EncodeToString(bys))
		// fmt.Println("itemCount:",cellData.Records().ItemCount())
		// // t.Log(MoleculeU32ToGo(accountCellData.Status().RawData()))
		// _, err = MoleculeU32ToGo(das.MoleculeNewDataEntity.Index().RawData())
		// if err != nil {
		// 	panic(err)
		// } else {
		// 	t.Log("success")
		// 	// newEntity := das.MoleculeNewDataEntity
		// 	// depEntity := das.MoleculeDepDataEntity
		// 	// if !newEntity.IsEmpty() && (depEntity == nil || depEntity.IsEmpty()) {
		// 	// 	proposeCellData, err := ProposalCellDataFromSlice(newEntity.Entity().RawData(), false)
		// 	// 	if err != nil {
		// 	// 		panic(err)
		// 	// 	}
		// 	// 	lock, err := MoleculeScriptToGo(*proposeCellData.ProposerLock())
		// 	// 	if err != nil {
		// 	// 		panic(err)
		// 	// 	}
		// 	// 	t.Log(lock.CodeHash.String())
		// 	// }
		// }
	}
}

func Test_PrintProposeCellLink(t *testing.T) {
	cellData := "0x64617305000000b1010000100000001000000010000000a101000010000000140000001800000000000000010000008501000085010000140000005d0000006b00000073000000490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8011400000020af3b4ed1c7768a8b87d2fc27242c1c3a43d45f0a000000b7526803f67ebe70aba60000000000000000120100001000000066000000bc000000560000000c0000003100000025000000100000001a0000001b0000001ceba7416f1392fc15d6001da30727b69b2db9126025000000100000001a0000001b0000001da30727b69b2db912600224f7e86151c0b593c23e560000000c0000003100000025000000100000001a0000001b00000064e8380633ef848ec86a00717ce4f160d8ec367e3f25000000100000001a0000001b000000717ce4f160d8ec367e3f02717ce4f160d8ec367e3f560000000c0000003100000025000000100000001a0000001b000000717ce4f160d8ec367e3f0080b35d34622e0d49a9b825000000100000001a0000001b00000080b35d34622e0d49a9b8028413d752ccbfeb88e0e6"
	cellData = cellData[2:]
	cellDataBytes, err := hex.DecodeString(cellData)
	if err != nil {
		panic(err)
	}
	if das, err := ParseTxWitnessToDasWitnessObj(cellDataBytes); err != nil {
		panic(err)
	} else {
		t.Log(das.WitnessObj.TableType)
		if len(das.MoleculeNewDataEntity.AsSlice()) == 0 {
			panic("empty")
		}
		if das.MoleculeNewDataEntity.Entity().IsEmpty() {
			panic("empty")
		}
		fmt.Println(MoleculeU32ToGo(das.MoleculeNewDataEntity.Index().RawData()))
		proposeData, err := ProposalCellDataFromSlice(das.MoleculeNewDataEntity.Entity().RawData(), false)
		if err != nil {
			panic(err)
		}
		list,err := ProposeWitnessSliceDataObjectListFromBytes(proposeData.AsSlice())
		bys, err := json.MarshalIndent(list, " ", " ")
		if err != nil {
			fmt.Println("ProposeTxWitnessDataList json err: ", err.Error())
		} else {
			fmt.Println(string(bys))
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
		account, err1 := AccountFromOutputData(bys)
		if err1 != nil {
			panic(err1)
		}
		t.Log("account:", account)

	}
	printf("0x46d316da038978b44ed0fa1f5590553f4f15ff879f2f04cfdeae72cd1c7cb7a700000000000000000000ffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000000")
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
// func Test_RecoverAccountIdFromChars(t *testing.T) {
// 	const testAccount = DasAccount("22222222.bit")
// 	t.Log(testAccount.AccountId().HexStr())
// 	accountChars := accountChars{}
// 	accountBytes := []byte(testAccount)
// 	for _, item := range accountBytes {
// 		accountChars = append(accountChars, accountChar{
// 			CharSetName: AccountChar_En,
// 			Bytes:       []byte{item},
// 		})
// 	}
// 	preAccountCellData :=
// 		NewPreAccountCellDataBuilder().
// 			Account(accountChars.MoleculeAccountChars()).
// 			CreatedAt(TimestampDefault()).
// 			InviterWallet(BytesDefault()).
// 			ChannelWallet(GoBytesToMoleculeBytes([]byte("xx"))).
// 			Price(PriceConfigDefault()).
// 			Quote(Uint64Default()).
// 			Build()
// 	account := AccountCharsToAccount(*preAccountCellData.Account())
// 	t.Log(account)
// 	recover := AccountCharsToAccountId(*preAccountCellData.Account())
// 	t.Log(recover.HexStr())
// }
//
// func Test_CreateData(t *testing.T) {
// 	preAccountCellData :=
// 		NewPreAccountCellDataBuilder().
// 			Account(AccountCharsDefault()).
// 			CreatedAt(TimestampDefault()).
// 			RefundLock(ScriptDefault()).
// 			InviterWallet(BytesDefault()).
// 			ChannelWallet(GoBytesToMoleculeBytes([]byte("xx"))).
// 			Price(PriceConfigDefault()).
// 			Quote(Uint64Default()).
// 			Build()
// 	// new := NewDataEntityBuilder().
// 	// 	Index(GoUint32ToMoleculeU32(0)).
// 	// 	Version(GoUint32ToMoleculeU32(1)).
// 	// 	Entity(GoBytesToMoleculeBytes(preAccountCellData.AsSlice())).
// 	// 	Build()
// 	// d := NewDataBuilder().
// 	// 	Dep(DataEntityOptDefault()).
// 	// 	Old(DataEntityOptDefault()).
// 	// 	New(NewDataEntityOptBuilder().Set(new).Build()).
// 	// 	Build()
// 	// s := hex.EncodeToString(d.AsSlice())
// 	// t.Log(s)
// 	// preAccountCell := NewPreAccountCell(TestNetPreAccountCell("",&preAccountCellData))
// 	witnessBys := NewDasWitnessData(TableType_PRE_ACCOUNT_CELL, preAccountCellData.AsSlice()).ToWitness()
// 	ret, err := ParseTxWitnessToDasWitnessObj(witnessBys)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// rawData
// 	if preAccountCellData, err := PreAccountCellDataFromSlice(ret.MoleculeNewDataEntity.Entity().RawData(), false); err != nil {
// 		panic(err)
// 	} else {
// 		t.Log(string(preAccountCellData.ChannelWallet().RawData()))
// 		script, err := MoleculeScriptToGo(*preAccountCellData.RefundLock())
// 		if err != nil {
// 			panic(err)
// 		}
// 		t.Log(script.CodeHash.String())
// 		t.Log(MoleculeU64ToGo(preAccountCellData.Quote().RawData()))
// 	}
// }

// func Test_RecoverData_From_BuildDasCommonMoleculeDataObj(t *testing.T) {
// 	createAt := NewTimestampBuilder().
// 		Set(GoTimeUnixToMoleculeBytes(time.Now().Unix())).Build()
//
// 	accountChars := NewAccountCharsBuilder()
// 	chars := []byte("iqyueq.bit")
// 	for _, item := range chars {
// 		accountChar :=
// 			NewAccountCharBuilder().
// 				CharSetName(GoUint32ToMoleculeU32(uint32(AccountChar_En))).
// 				Bytes(GoBytesToMoleculeBytes([]byte{item})).
// 				Build()
// 		accountChars.Push(accountChar)
// 	}
//
// 	inviterAccountId := GoBytesToMoleculeBytes(DasAccountFromStr("xxx.bit").AccountId().Bytes())
// 	args, _ := hex.DecodeString("b7526803f67ebe70aba6")
// 	preAccountCellData :=
// 		NewPreAccountCellDataBuilder().
// 			Account(accountChars.Build()).
// 			CreatedAt(createAt).
// 			RefundLock(GoCkbScriptToMoleculeScript(types.Script{
// 				CodeHash: types.HexToHash("123456aa"),
// 				HashType: types.HashTypeType,
// 				Args:     args,
// 			})).
// 			InviterWallet(inviterAccountId).
// 			ChannelWallet(GoBytesToMoleculeBytes([]byte("xx"))).
// 			Price(PriceConfigDefault()).
// 			Quote(GoUint64ToMoleculeU64(10086)).
// 			Build()
// 	preAccountCell := NewPreAccountCell(TestNetPreAccountCell("", &preAccountCellData))
// 	witnessBys := NewDasWitnessData(preAccountCell.TableType(), preAccountCellData.AsSlice()).ToWitness()
// 	ret, err := ParseTxWitnessToDasWitnessObj(witnessBys)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if preAccountCellData, err := PreAccountCellDataFromSlice(ret.MoleculeNewDataEntity.Entity().AsSlice(), false); err != nil {
// 		panic(err)
// 	} else {
// 		t.Log(string(preAccountCellData.ChannelWallet().RawData()))
// 		script, err := MoleculeScriptToGo(*preAccountCellData.RefundLock())
// 		if err != nil {
// 			panic(err)
// 		}
// 		t.Log(script.CodeHash.String())
// 		t.Log(MoleculeU64ToGo(preAccountCellData.Quote().RawData()))
// 	}
// }

func Test_PreAccountDataFromBytes(t *testing.T) {
	witnessHex := "0x64617306000000ca010000100000001000000010000000ba01000010000000140000001800000000000000010000009e0100009e01000028000000f40000003d010000570100005b010000690100008a0100009201000096010000cc00000024000000390000004e00000063000000780000008d000000a2000000b7000000150000000c00000010000000010000000100000031150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061150000000c00000010000000020000000100000061150000000c000000100000000200000001000000614900000010000000300000003100000058c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a630114000000c9f53b1d85356b60453f867610888d89a0b667ad160000000003c9f53b1d85356b60453f867610888d89a0b667ad000000000a000000b7526803f67ebe70aba62100000010000000110000001900000008404b4c000000000020a1070000000000e663000000000000000000009ff9986000000000"
	if strings.HasPrefix(witnessHex,"0x") {
		witnessHex = witnessHex[2:]
	}
	bys, err := hex.DecodeString(witnessHex)
	if err != nil {
		panic(err)
	}
	obj, err := ParseTxWitnessToDasWitnessObj(bys)
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(obj.MoleculeNewDataEntity.Entity().RawData()))
	preAccountCell, err := PreAccountCellDataFromSlice(obj.MoleculeNewDataEntity.Entity().RawData(), false)
	if err != nil {
		panic(err)
	}
	// script, err := mo(*preAccountCell.OwnerLockArgs())
	// if err != nil {
	// 	panic(err)
	// }
	// t.Log(hex.EncodeToString(script.Args))
	t.Log(preAccountCell)
	t.Log(hex.EncodeToString(preAccountCell.OwnerLockArgs().RawData()))
	if hex.EncodeToString(preAccountCell.OwnerLockArgs().RawData())[2] == '0' {
		fmt.Println("bad")
	}
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
	hexStr := "0x64617305000000d6010000100000001000000010000000c60100001000000014000000180000000000000001000000aa010000aa0100002c000000c60000000f010000280100002c0100002c01000075010000960100009e010000a20100009a0000001c00000031000000460000005b0000007000000085000000150000000c00000010000000020000000100000072150000000c00000010000000020000000100000073150000000c0000001000000002000000010000006c150000000c00000010000000020000000100000069150000000c0000001000000002000000010000006e150000000c0000001000000002000000010000006b4900000010000000300000003100000058c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a6301140000000400f20fd1f498647c5239d3e17e0e6275e4fdfa15000000030400f20fd1f498647c5239d3e17e0e6275e4fdfa00000000490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8011400000019b04faf5b6e76e6d6640344b23dc16ffd9010ec2100000010000000110000001900000006c0cf6a0000000000c0cf6a0000000000064100000000000000000000a77cbf6000000000"
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
	if data, err := PreAccountCellDataFromSlice(obj.MoleculeNewDataEntity.Entity().RawData(), false); err != nil {
		panic(err)
	} else {
		fmt.Println(hex.EncodeToString(data.OwnerLockArgs().RawData()))
		fmt.Println("invalid:",isInvalidOwnerLock(data.OwnerLockArgs().RawData()))
	}
}

func isInvalidOwnerLock(bys []byte) bool {
	return hex.EncodeToString(bys)[2] == '0'
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

func Test_NewDasWitnessData(t *testing.T) {
	obj := NewDasWitnessData(1, []byte("usa"))
	t.Log(hex.EncodeToString(obj.ToWitness()))
	das, err := NewDasWitnessDataFromSlice(obj.ToWitness())
	if err != nil {
		panic(err.Error())
	} else {
		t.Log(hex.EncodeToString(das.ToWitness()))
		t.Log(das.Tag, das.TableType)
	}
}