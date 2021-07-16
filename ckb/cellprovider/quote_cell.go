package cellprovider

import (
	"fmt"
	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
	"github.com/DeAccountSystems/das_commonlib/ckb/gotype"
	"github.com/DeAccountSystems/das_commonlib/common"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2021
 * FileName: time_cell
 * Author:   LinGuanHong
 * Date:     2021/2/25 10:52
 * Description:
 */

func (l *LiveCellPackObj) ToQuoteCell() *gotype.QuoteCell {
	return &gotype.QuoteCell{
		Data: l.LiveCell.OutputData,
		CellDep: types.CellDep{
			OutPoint: l.LiveCell.OutPoint,
			DepType:  types.DepTypeCode,
		},
	}
}

func LoadOneQuoteCell(rpcClient rpc.Client) (*LiveCellPackObj, error) {
	searchKey := &indexer.SearchKey{
		Script:     celltype.DasQuoteCellScript.Out.Script(),
		ScriptType: indexer.ScriptTypeType,
	}
	liveCells, _, err := common.LoadLiveCells(rpcClient, searchKey, 400*celltype.OneCkb, true, false, nil)
	if err != nil {
		return nil, fmt.Errorf("LoadLiveCells err: %s", err.Error())
	}
	if size := len(liveCells); size == 0 {
		return nil, NewEmptyErr("quoteCell")
	} else if size > 1 {
		return &LiveCellPackObj{
			LiveCell: &liveCells[1],
			CellCap:  liveCells[1].Output.Capacity,
		}, err
	}
	return &LiveCellPackObj{
		LiveCell: &liveCells[0],
		CellCap:  liveCells[0].Output.Capacity,
	}, err
}
