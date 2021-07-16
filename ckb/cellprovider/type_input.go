package cellprovider

import (
	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2021
 * FileName: type_input
 * Author:   LinGuanHong
 * Date:     2021/7/16 5:05
 * Description:
 */

func (l *LiveCellPackObj) TypeInputCell(lockScriptType celltype.LockScriptType) *celltype.TypeInputCell {
	return &celltype.TypeInputCell{
		Input: types.CellInput{
			PreviousOutput: l.LiveCell.OutPoint,
		},
		LockType: lockScriptType,
		CellCap:  l.CellCap,
	}
}