package celltype

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
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

func Test_BuildDasCommonMoleculeDataObj(t *testing.T) {
	BuildDasCommonMoleculeDataObj(0,0,0,nil,nil,&AccountCellData{})
}

func Test_ParseCellData(t *testing.T) {
	cellData := "0x64617306000000781100001000000010000000100000006811000010000000140000001800000008000000010000004c1100004c1100000c0000004100000035000000100000003000000031000000000000000000000000000000000000000000000000000000000000000000000000000000000b110000980000000b0100007e010000f101000064020000d70200004a030000bd03000030040000a30400001605000089050000fc0500006f060000e206000055070000b20700002508000098080000f508000068090000db0900004e0a0000c10a0000340b0000a70b00001a0c00008d0c0000000d0000730d0000e60d0000590e0000cc0e00003f0f0000b20f00002510000098100000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a000000047873fe90306ef58279f97b8dfeb1d43c3318ec97047873fe90306ef58279f97b8dfeb1d43c3318ec9780eb3d2d01000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003111937d72b14fc0b998038f4ab68737cceeb15bf03111937d72b14fc0b998038f4ab68737cceeb15bf20f6ac2c01000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003e427f4202c3d43cf2a538e1a3ed5a34b63d0715003e427f4202c3d43cf2a538e1a3ed5a34b63d0715000c1812f01000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a000000039041bd678e37f1d19f916551968ee94f20772e5f039041bd678e37f1d19f916551968ee94f20772e5f80414d3601000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003381472ba3b0c1c492245baa190ad878adddbf44a03381472ba3b0c1c492245baa190ad878adddbf44a20a1343101000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003eb1c4cb922d77c05b8ed4708ffabb1923159456703eb1c4cb922d77c05b8ed4708ffabb1923159456720a1343101000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a0000000326b2099f0ce962134fc86b715135641b0f7a93420326b2099f0ce962134fc86b715135641b0f7a9342a076783301000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003fc9d2a3ad8a0e0275f967eb9f51dbcf2441fb1c703fc9d2a3ad8a0e0275f967eb9f51dbcf2441fb1c780414d3601000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003dcc6fec3f5393c453b38e9af061fa5c4285cc36303dcc6fec3f5393c453b38e9af061fa5c4285cc36380414d3601000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003c715d96c3541191ff31da6c6b7ad508971a43af103c715d96c3541191ff31da6c6b7ad508971a43af1402c6f3701000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003f2e305b759bad7e0818adf207a81585d184e733103f2e305b759bad7e0818adf207a81585d184e733140d65f2e01000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a0000000373ce4c69c1e4b26b324f62da4a2439bae8957d380373ce4c69c1e4b26b324f62da4a2439bae8957d38a0cbf02e01000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003374d410f0edf510562aed617e8de0a47baf5cd6003374d410f0edf510562aed617e8de0a47baf5cd60e0e0ce2d01000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a000000038969efdea33e87f5333b7dbf6ce1d282d7458208038969efdea33e87f5333b7dbf6ce1d282d745820800c1812f01000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003c554d9b045b9a06a4c51cccf654b3fd3d36c99c403c554d9b045b9a06a4c51cccf654b3fd3d36c99c440d7f63b010000005d0000000c00000055000000490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8011400000075bebe0707641658cb8020b9233de32c20c3e172a076783301000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a0000000364619f4f65dcb33d15eb5687112718e1ace16f340364619f4f65dcb33d15eb5687112718e1ace16f340017913801000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003060bbae03ef52f1b47db247215da0fb87ff4b2eb03060bbae03ef52f1b47db247215da0fb87ff4b2eb80414d36010000005d0000000c00000055000000490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce80114000000d6465f6faf694011c770bb4e5aeed5b32865183b0076dd4101000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a000000033ad47c7fb550b4a9c46e779639e9984b12feb480033ad47c7fb550b4a9c46e779639e9984b12feb4804081e73201000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a0000000374f60592118f89ad0a0de93a2fd150da9f82d6940374f60592118f89ad0a0de93a2fd150da9f82d694e0e0ce2d01000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a000000038e30102bc827530643966de21202db6254c19c74038e30102bc827530643966de21202db6254c19c7480414d3601000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003a1562071beebf0d778312263d0936c707e12964a03a1562071beebf0d778312263d0936c707e12964a80414d3601000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a000000043d12f8d6f0ea36d2f01553beb4810d12d3658d2a043d12f8d6f0ea36d2f01553beb4810d12d3658d2a8096c53101000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a000000032ce62318bc5bf2eeb34b8f2f12880fdef20e356f032ce62318bc5bf2eeb34b8f2f12880fdef20e356fc001b33901000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a0000000360b01504cb8fc73d0eff81e203e46e3e8d6704880360b01504cb8fc73d0eff81e203e46e3e8d670488a0cbf02e01000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003255749ce8ab1dd14c09787e7405a15ff17c2445903255749ce8ab1dd14c09787e7405a15ff17c2445980414d3601000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a000000032ba1473cb3973c288312a92fb8930bb0af2cae02032ba1473cb3973c288312a92fb8930bb0af2cae02c001b33901000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a000000036be89802619bc5063b09e7280034d0189bfc763f036be89802619bc5063b09e7280034d0189bfc763fa076783301000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003d8083816f42c8b47ed1151b8275ec2874cba6c3c03d8083816f42c8b47ed1151b8275ec2874cba6c3c80414d3601000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a000000032932d13c93a64ba81ea4356ee0e58ce4bd6b0680032932d13c93a64ba81ea4356ee0e58ce4bd6b0680402c6f3701000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a0000000388a72a5b05cc9150a5e96d01d2eb794b98bdbfb90388a72a5b05cc9150a5e96d01d2eb794b98bdbfb940d65f2e01000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a000000037120ee7a79962549f19349da887607ba6dae6230037120ee7a79962549f19349da887607ba6dae6230a0cbf02e01000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003b5868663c79b206d52896b5a805de4c616a3a03a03b5868663c79b206d52896b5a805de4c616a3a03a4081e73201000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003540cb04ebab67e05a620b97bb367ac5e4ed68f0903540cb04ebab67e05a620b97bb367ac5e4ed68f0900c1812f01000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003471862fc58e7cf2346c2a2929899b0184d40e27903471862fc58e7cf2346c2a2929899b0184d40e2794081e73201000000730000000c0000006b0000005f0000001000000030000000310000009376c3b5811942960a846691e16e477cf43d7c7fa654067c9948dfcd09a32137012a00000003773bcce3b8b41a37ce59fd95f7cbccbff2cfd2d003773bcce3b8b41a37ce59fd95f7cbccbff2cfd2d0e036de3601000000"
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
		cellData, err := IncomeCellDataFromSlice(das.MoleculeNewDataEntity.Entity().RawData(), false)
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
	witnessHex := "64617301000000ce02000010000000100000006f0100005f010000100000001400000018000000010000000100000043010000430100001c000000300000002e010000360100003e0100003f01000006980d3ab90a0c42285eb056016037738bb6a86ffe0000002c00000041000000560000006b0000008000000095000000aa000000bf000000d4000000e9000000150000000c00000010000000020000000100000074150000000c00000010000000020000000100000072150000000c00000010000000020000000100000079150000000c0000001000000002000000010000006e150000000c00000010000000020000000100000064150000000c00000010000000020000000100000061150000000c0000001000000002000000010000006d150000000c00000010000000020000000100000065150000000c00000010000000020000000100000072150000000c000000100000000200000001000000655418c46000000000000000000000000000040000005f010000100000001400000018000000000000000100000043010000430100001c000000300000002e010000360100003e0100003f01000006980d3ab90a0c42285eb056016037738bb6a86ffe0000002c00000041000000560000006b0000008000000095000000aa000000bf000000d4000000e9000000150000000c00000010000000020000000100000074150000000c00000010000000020000000100000072150000000c00000010000000020000000100000079150000000c0000001000000002000000010000006e150000000c00000010000000020000000100000064150000000c00000010000000020000000100000061150000000c0000001000000002000000010000006d150000000c00000010000000020000000100000065150000000c00000010000000020000000100000072150000000c000000100000000200000001000000655418c4600000000000000000000000000004000000"
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
	fmt.Println("account cell field count:",accountCell.Len())
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

func Test_EchoAction(t *testing.T) {
	s1  := ScriptDefault()
	s12 := ScriptDefault()
	action := NewActionDataBuilder().
		Action(GoStrToMoleculeBytes(Action_EditManager)).
		Params(GoBytesToMoleculeBytes(append(s1.AsSlice(),s12.AsSlice()...))).
		Build()
	// 646173000000001f0000000c0000001b0000000b0000006275795f6163636f756e7400000000
	// 64617300000000210000000c0000001b0000000b0000006275795f6163636f756e74020000003131
	// 64617300000000230000000c0000001b0000000b0000006275795f6163636f756e740400000034333435
	// 64617300000000890000000c0000001b0000000b0000006275795f6163636f756e746a00000035000000100000003000000031000000000000000000000000000000000000000000000000000000000000000000000000000000003500000010000000300000003100000000000000000000000000000000000000000000000000000000000000000000000000000000
	// 64617300000000b10000000c0000001b0000000b0000006275795f6163636f756e7492000000490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8011400000019b04faf5b6e76e6d6640344b23dc16ffd9010ec490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8011400000019b04faf5b6e76e6d6640344b23dc16ffd9010ec
	witnessBys := NewDasWitnessData(TableType_Action, action.AsSlice()).ToWitness()
	fmt.Println(hex.EncodeToString(witnessBys))
}

func Test_NewAccountCell(t *testing.T) {
	accountCellData := NewAccountCellDataBuilder().
		Id(AccountIdDefault()).
		Status(Uint8Default()).
		Account(AccountCharsDefault()).
		RegisteredAt(Uint64Default()).
		Records(RecordsDefault()).Build()
	fmt.Println(accountCellData.Len())
}

func Test_ParseTxWitnessToDasWitnessObj_ConfigCellType(t *testing.T) {
	hexStr := "0x646173050000003e0200001000000010000000100000002e020000100000001400000018000000000000000100000012020000120200002c000000f80000002d0100005b01000073010000a8010000dd010000fe010000060200000a020000cc00000024000000390000004e00000063000000780000008d000000a2000000b7000000150000000c0000001000000002000000010000007a150000000c0000001000000002000000010000006f150000000c0000001000000002000000010000006e150000000c00000010000000020000000100000061150000000c00000010000000020000000100000074150000000c00000010000000020000000100000069150000000c0000001000000002000000010000006f150000000c0000001000000002000000010000006e35000000100000003000000031000000000000000000000000000000000000000000000000000000000000000000000000000000002a000000053a6cab3323833f53754db4202f5741756c436ede053a6cab3323833f53754db4202f5741756c436ede140000000000000000000000000000000000000000000000350000001000000030000000310000000000000000000000000000000000000000000000000000000000000000000000000000000035000000100000003000000031000000000000000000000000000000000000000000000000000000000000000000000000000000002100000010000000110000001900000008404b4c0000000000404b4c00000000008e32000000000000f401000074a74a6100000000"
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
		inviterLock,err := data.InviterLock().IntoScript()
		if err != nil {
			panic(err)
		}
		fmt.Println("CodeHash:",hex.EncodeToString(inviterLock.CodeHash().RawData()))
		fmt.Println("HashType:",hex.EncodeToString(inviterLock.HashType().AsSlice()))
		fmt.Println("Args Len:",inviterLock.Args().Len())
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