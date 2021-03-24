package builder

import (
	"github.com/DA-Services/das_commonlib/ckb/celltype"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
)

/**
 * Copyright (C), 2019-2021
 * FileName: transaction_test
 * Author:   LinGuanHong
 * Date:     2021/3/24 5:16 下午
 * Description:
 */

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
	txBuilder.BuildTransaction()
	t.Log(i1.InputIndex)
	t.Log(i2.InputIndex)
	t.Log(i3.InputIndex)
	t.Log(i4.InputIndex)
	t.Log(i5.InputIndex)

}
