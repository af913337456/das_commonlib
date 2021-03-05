package common

import (
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
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
		if totalCap >= capLimit { // enough
			break
		}
		if err = iterator.Next(); err != nil {
			return nil, 0, fmt.Errorf("LoadLiveCells, read iterator next err: %s", err.Error())
		}
	}
	if totalCap < capLimit {
		if err = iterator.Next(); err != nil {
			return nil, 0, fmt.Errorf("LoadLiveCells, read iterator next err: %s", err.Error())
		}
		goto NextBatch
	}
	return liveCells, totalCap, nil
}
