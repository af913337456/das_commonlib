package cellprovider

import (
	"errors"
	"fmt"
	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
	"github.com/DeAccountSystems/das_commonlib/ckb/gotype"
)

/**
 * Copyright (C), 2019-2021
 * FileName: account_cell
 * Author:   LinGuanHong
 * Date:     2021/7/16 5:00
 * Description:
 */

func (l *LiveCellPackObj) ToAccountCell() (*gotype.AccountCell, error) {
	if l.Obj == nil {
		return nil, errors.New("invalid accountCell, empty witness data")
	}
	accountCellData := l.Obj.(*celltype.AccountCellData)
	if accountCellData== nil {
		return nil, errors.New("invalid accountCell, null")
	}
	status, err := celltype.MoleculeU8ToGo(accountCellData.Status().RawData())
	if err != nil {
		return nil, fmt.Errorf("MoleculeU8ToGo err: %s", err.Error())
	}
	accountIdFromData, err := celltype.AccountIdFromOutputData(l.LiveCell.OutputData)
	if err != nil {
		return nil, fmt.Errorf("AccountIdFromOutputData err: %s", err.Error())
	}
	accountIdFromWitness := celltype.AccountCharsToAccount(*accountCellData.Account()).AccountId()
	if accountIdFromData != accountIdFromWitness {
		return nil, fmt.Errorf("accountId not equal")
	}
	return &gotype.AccountCell {
		CellCap:       l.CellCap,
		AccountId:     accountIdFromData,
		Status:        status,
		Point:         *l.LiveCell.OutPoint,
		WitnessStatus: celltype.AccountWitnessStatus_Exist,
		Data:          l.LiveCell.OutputData,
		WitnessData:   l.WitnessData,
		DasLockArgs:   l.LiveCell.Output.Lock.Args,
	}, nil
}