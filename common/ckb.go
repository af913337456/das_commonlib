package common

import (
	"errors"
	"fmt"
	"github.com/DA-Services/das_commonlib/ckb/celltype"
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
)

/**
 * Copyright (C), 2019-2021
 * FileName: ckb
 * Author:   LinGuanHong
 * Date:     2021/2/25 10:04 上午
 * Description:
 */

func LoadLiveCells(client rpc.Client, key *indexer.SearchKey, capLimit uint64, filter func(cell *indexer.LiveCell) bool) ([]indexer.LiveCell, uint64, error) {
	c := collector.NewLiveCellCollector(
		client, key, indexer.SearchOrderAsc, 100, "")
	iterator, err := c.Iterator()
	if err != nil {
		return nil, 0, fmt.Errorf("LoadLiveCells Collect failed: %s", err.Error())
	}
	liveCells := []indexer.LiveCell{}
	totalCap := uint64(0)
NextBatch:
	for iterator.HasNext() {
		liveCell, err := iterator.CurrentItem()
		if err != nil {
			return nil, 0, fmt.Errorf("LoadLiveCells, read iterator current err: %s", err.Error())
		}
		if filter != nil && !filter(liveCell) {
			continue
		}
		totalCap = totalCap + liveCell.Output.Capacity
		liveCells = append(liveCells, *liveCell)
		if err = iterator.Next(); err != nil {
			return nil, 0, fmt.Errorf("LoadLiveCells, read iterator next err: %s", err.Error())
		}
	}
	if totalCap < capLimit {
		iterator.Next()
		goto NextBatch
	}
	return liveCells, totalCap, nil
}

func GetScriptTypeFromLockScript(ckbSysScript *utils.SystemScripts, lockScript *types.Script) (celltype.LockScriptType, error) {
	lockCodeHash := lockScript.CodeHash
	switch lockCodeHash {
	case ckbSysScript.SecpSingleSigCell.CellHash:
		return celltype.ScriptType_User, nil
	case celltype.DasAnyOneCanSendCellInfo.CodeHash:
		return celltype.ScriptType_Any, nil
	case celltype.DasETHLockCellInfo.CodeHash:
		return celltype.ScriptType_ETH, nil
	case celltype.DasBTCLockCellInfo.CodeHash:
		return celltype.ScriptType_BTC, nil
	default:
		return -1, errors.New("invalid lockScript")
	}
}
