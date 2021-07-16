package cellprovider

import (
	"errors"
	"fmt"
	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
	"github.com/DeAccountSystems/das_commonlib/common"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2021
 * FileName: time_cell
 * Author:   LinGuanHong
 * Date:     2021/3/05 16:52
 * Description:
 */
func (l *LiveCellPackObj) ToHeightCellDep() *types.CellDep {
	return &types.CellDep{
		OutPoint: l.LiveCell.OutPoint,
		DepType:  types.DepTypeCode,
	}
}

func (l *LiveCellPackObj) LatestBlockHeight() (int64, error) {
	if l.LiveCell.OutputData == nil {
		return 0, errors.New("invalid blockHeightCell data")
	}
	return common.BytesToInt64(l.LiveCell.OutputData[2:]), nil
}

func LoadOneHeightCell(rpcClient rpc.Client) (*LiveCellPackObj, error) {
	searchKey := &indexer.SearchKey{
		Script:     celltype.DasHeightCellScript.Out.Script(),
		ScriptType: indexer.ScriptTypeType,
	}
	liveCells, _, err := common.LoadLiveCellsWithSize(rpcClient, searchKey, 20000*celltype.OneCkb, 10, true, false, nil)
	if err != nil {
		return nil, fmt.Errorf("LoadLiveCells err: %s", err.Error())
	}
	if size := len(liveCells); size == 0 {
		return nil, NewEmptyErr("timeCell")
	} else if size > 1 {
		return &LiveCellPackObj{
			LiveCell: &liveCells[1],
			CellCap:  liveCells[1].Output.Capacity,
		}, err
	}
	latestOne := &LiveCellPackObj{
		LiveCell: &liveCells[0],
		CellCap:  liveCells[0].Output.Capacity,
	}
	return latestOne, err
}
