package builder

import (
	"encoding/hex"
	"fmt"
	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
)

/**
 * Copyright (C), 2019-2021
 * FileName: transaction_test
 * Author:   LinGuanHong
 * Date:     2021/3/24 5:16
 * Description:
 */

func Test_ActionTx(t *testing.T) {
	params := celltype.ActionParam_Owner
	actionBuilder := celltype.NewActionDataBuilder().Action(celltype.GoStrToMoleculeBytes(celltype.Action_WithdrawFromWallet))
	if params != nil {
		actionBuilder.Params(celltype.GoBytesToMoleculeBytes(params))
	}
	actionData := actionBuilder.Build()
	witnessBys := celltype.NewDasWitnessData(celltype.TableType_ACTION, actionData.AsSlice()).ToWitness()
	fmt.Println(hex.EncodeToString(witnessBys))
}

func Test_BuildTransaction(t *testing.T) {
	txBuilder := NewTransactionBuilder0("", nil, 0)
	i1 := &celltype.TypeInputCell{
		InputIndex: 0,
		Input:      types.CellInput{},
		LockType:   6,
		CellCap:    0,
	}
	i2 := &celltype.TypeInputCell{
		InputIndex: 0,
		Input:      types.CellInput{},
		LockType:   1,
		CellCap:    0,
	}
	i3 := &celltype.TypeInputCell{
		InputIndex: 0,
		Input:      types.CellInput{},
		LockType:   2,
		CellCap:    0,
	}
	i4 := &celltype.TypeInputCell{
		InputIndex: 0,
		Input:      types.CellInput{},
		LockType:   2,
		CellCap:    0,
	}
	i5 := &celltype.TypeInputCell{
		InputIndex: 0,
		Input:      types.CellInput{},
		LockType:   3,
		CellCap:    0,
	}
	txBuilder.inputList = append(txBuilder.inputList, i1)
	txBuilder.inputList = append(txBuilder.inputList, i2)
	txBuilder.inputList = append(txBuilder.inputList, i3)
	txBuilder.inputList = append(txBuilder.inputList, i4)
	txBuilder.inputList = append(txBuilder.inputList, i5)
	_ = txBuilder.BuildTransaction()
	t.Log(i1.InputIndex)
	t.Log(i2.InputIndex)
	t.Log(i3.InputIndex)
	t.Log(i4.InputIndex)
	t.Log(i5.InputIndex)
}
