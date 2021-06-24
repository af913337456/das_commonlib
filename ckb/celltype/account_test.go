package celltype

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
)

/**
 * Copyright (C), 2019-2020
 * FileName: data_test
 * Author:   LinGuanHong
 * Date:     2020/12/20 2:57
 * Description:
 */

func rpcClient() rpc.Client {
	rpcClient, err := rpc.DialWithIndexerContext(
		context.TODO(),
		"http://:8114",
		"http://:8116")
	if err != nil {
		panic(fmt.Errorf("init rpcClient failed: %s", err.Error()))
	}
	return rpcClient
}

func Test_DasLockCodeHashIndexType(t *testing.T) {
	t.Log(DasLockCodeHashIndexType(DasLockCodeHashIndexType_CKB_Normal).Bytes())
}

func Test_ExpiredAtFromOutputData(t *testing.T) {
	dataHex := "c4de24c38f1a22e65b9a1a24aaae7d4db37e7ae138e9d44651d76f1d179f95e8ee06f79afc0af40e7198faf1611a8fa5324263b3f2dd3b620000000062616161616161612e626974"
	dataBys, _ := hex.DecodeString(dataHex)
	expired, e := ExpiredAtFromOutputData(dataBys)
	if e != nil {
		panic(e)
	}
	t.Log(expired)
}

func Test_DefaultAccountCellDataBytes(t *testing.T) {
	id := DasAccountIdFromBytes([]byte("123"))
	nextId := DasAccountIdFromBytes([]byte("456"))
	bys := DefaultAccountCellDataBytes(id, nextId)
	t.Log(bys)
}

func Test_DasAccountIdFromBytes(t *testing.T) {
	id := DasAccountIdFromBytes([]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1})
	fmt.Println(id.HexStr())
}

func Test_EchoTypeId(t *testing.T) {
	t.Log(hexToArgsBytes("0x"))
	bys,_ := hex.DecodeString("d5eee5a3ac9d65658535b4bdad25e22a81c032f5bbdf5ace45605a33482eeb45")
	script := types.Script{
		CodeHash: types.HexToHash("0x00000000000000000000000000000000000000000000000000545950455f4944"),
		HashType: types.HashTypeType,
		Args:     bys,
	}
	serBys,_ := script.Serialize()
	bysRet, _ := blake2b.Blake256(serBys)
	t.Log(types.BytesToHash(bysRet))
	// t.Log(DasAccountCellScript.Out)
	// account_cell : 0x274775e475c1252b5333c20e1512b7b1296c4c5b52a25aa2ebd6e41f5894c41f
	// // 0x9878b226df9465c215fd3c94dc9f9bf6648d5bea48a24579cf83274fe13801d2
	// t.Log(DasWalletCellScript.Out)
	// t.Log(DasTimeCellScript.Out.TypeId())
}

func Test_InitSystemScript(t *testing.T) {
	fmt.Println(DasProposeCellScript.Dep.ToDepCell().OutPoint.TxHash.String())
	SetSystemCodeScriptOutPoint(DasProposeCellScript.Out.CodeHash, types.OutPoint{
		TxHash: types.HexToHash("111"),
	})
	obj,_ := SystemCodeScriptMap.Load(DasProposeCellScript.Out.CodeHash)
	item := obj.(*DASCellBaseInfo)
	fmt.Println(item.Dep.TxHash.String())
	fmt.Println(DasProposeCellScript.Dep.ToDepCell().OutPoint.TxHash.String())
	bys, err := SystemCodeScriptBytes()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(hex.EncodeToString(bys))
	}
}

func Test_EmojiAccountLen(t *testing.T) {
	t.Log(DasAccountFromStr("🏃‍♀️🏃.bit").AccountValidateLen())
	t.Log(DasAccountFromStr("🔥或许🏃‍♀️🏃linguan.bit").AccountValidateLen())
	t.Log(DasAccountFromStr("🏃‍♀️🏃linguan.bit").AccountValidateLen())
	t.Log(DasAccountFromStr("123.bit").AccountValidateLen())
	t.Log(DasAccountFromStr("11.bit").AccountValidateLen())
	t.Log(DasAccountFromStr("😄hj🌹.bit").AccountValidateLen())
	t.Log(DasAccountFromStr("😄🌹如此.bit").AccountValidateLen())
	// 🙈🙈🙈🙈🙈.bit
	t.Log(len(DasAccountFromStr("🙈🙈🙈🙈🙈🙈🙈.bit").Bytes()))
	t.Log(DasAccountFromStr("🙈🙈🙈🙈🙈.bit").AccountValidateLen())
}

func Test_AccountId(t *testing.T) {
	acc := DasAccountFromStr("12345666.bit")
	// 5bd281eef6f9d72d71a7
	t.Log(hex.EncodeToString(acc.AccountId().Bytes()))

	bys, _ := hex.DecodeString("00000000000000000000")
	t.Log(DasAccountIdFromBytes(bys).HexStr())
}

func Test_AccountCharLen(t *testing.T) {
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

func Test_PriceConfigs(t *testing.T) {
	getItem := func() *PriceConfig {
		p1 := NewPriceConfigBuilder().Length(GoUint8ToMoleculeU8(1)).Build()
		p2 := NewPriceConfigBuilder().Length(GoUint8ToMoleculeU8(2)).Build()
		p3 := NewPriceConfigBuilder().Length(GoUint8ToMoleculeU8(3)).Build()
		list := NewPriceConfigListBuilder().Push(p1).Push(p2).Push(p3).Build()
		fmt.Println(list.ItemCount())
		priceIndex := uint(0)
		accountBytesLen := uint8(len([]rune("123")))
		for ; priceIndex < list.ItemCount(); priceIndex++ {
			item := list.Get(priceIndex)
			accountLen, err := MoleculeU8ToGo(item.Length().AsSlice())
			if err != nil {
				panic(err)
			} else if accountLen == accountBytesLen {
				return item
			}
			fmt.Println(accountLen, accountBytesLen, priceIndex)
		}
		return nil
	}
	i := getItem()
	fmt.Println(i.Length().RawData())
}

func Test_U64Bytes(t *testing.T) {
	d, _ := blake2b.Blake256([]byte("0"))
	t.Log(len(d))
	t.Log(len(GoUint64ToBytes(0)))
}

func Test_AccountChar(t *testing.T) {
	t.Log(len([]byte("account")))
}

func Test_Blake2b_256(t *testing.T) {
	// 0xc9804583fc51c64512c0153264a707c254ae81ff
	bys, _ := blake2b.Blake160([]byte("das00007.bit"))
	t.Log(hex.EncodeToString(bys))
	t.Log(len(bys), bys)
}

func Test_ParseActionCell(t *testing.T) {
	// 0x64617305000000fd00000010000000fd000000fd000000ed0000001000000014000000180000000600000001000000d1000000d1000000140000005d0000006b00000073000000490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8011400000020af3b4ed1c7768a8b87d2fc27242c1c3a43d45f0a000000b7526803f67ebe70aba600000000000000005e00000008000000560000000c0000003100000025000000100000001a0000001b0000002f2e5c058a06dfc38cda00606f15d9fea831af648425000000100000001a0000001b000000606f15d9fea831af64840264e8380633ef848ec86a
	hexStr := "0x64617300000000210000000c0000001c0000000c000000656469745f7265636f7264730100000001"
	bys, err := hex.DecodeString(hexStr[2:])
	if err != nil {
		t.Fatal(err)
	}
	if witness, err := NewDasWitnessDataFromSlice(bys); err != nil {
		t.Fatal(err)
	} else {
		actionData, _ := ActionDataFromSlice(witness.TableBys, false)
		t.Log(witness.Tag,witness.TableType,hex.EncodeToString(witness.TableBys),string(actionData.Action().RawData()))
	}
}

func Test_ParsePreAccountCell(t *testing.T) {
	// 0x64617305000000fd00000010000000fd000000fd000000ed0000001000000014000000180000000600000001000000d1000000d1000000140000005d0000006b00000073000000490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8011400000020af3b4ed1c7768a8b87d2fc27242c1c3a43d45f0a000000b7526803f67ebe70aba600000000000000005e00000008000000560000000c0000003100000025000000100000001a0000001b0000002f2e5c058a06dfc38cda00606f15d9fea831af648425000000100000001a0000001b000000606f15d9fea831af64840264e8380633ef848ec86a
	hexStr := "0x646173000000001a0000000c00000016000000060000006465706c6f7900000000"
	bys, err := hex.DecodeString(hexStr[2:])
	if err != nil {
		t.Fatal(err)
	}
	if witness, err := NewDasWitnessDataFromSlice(bys); err != nil {
		t.Fatal(err)
	} else {
		actionData, _ := ActionDataFromSlice(witness.TableBys, false)
		t.Log(witness.Tag,witness.TableType,hex.EncodeToString(witness.TableBys),string(actionData.Action().RawData()))
	}
}

type tempCell struct {
	DasLockArgs []byte
}
func (a *tempCell) SetOwner(indexType DasLockCodeHashIndexType,args []byte)  {
	tempBytes := make([]byte,0,DasLockArgsMinBytesLen)
	tempBytes = append(tempBytes,indexType.Bytes()...)
	tempBytes = append(tempBytes,args...)
	tempBytes = append(tempBytes,a.DasLockArgs[DasLockArgsMinBytesLen/2:]...)
	a.DasLockArgs = tempBytes
}

func Test_SetOwner(t *testing.T) {
	owner := []byte{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}
	param := DasLockParam{
		OwnerCodeHashIndexByte: []byte{1},
		OwnerPubkeyHashByte:    owner,
		ManagerCodeHashIndex:   []byte{2},
		ManagerPubkeyHash:      owner,
	}
	tempCell := &tempCell{
		DasLockArgs: param.Bytes(),
	}
	repla := []byte{0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}
	tempCell.SetOwner(DasLockCodeHashIndexType_ETH_Normal,repla)
	fmt.Println(tempCell.DasLockArgs)
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
