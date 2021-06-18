package repeatchecker

import (
	"encoding/hex"
	"sync"
	"time"

	"github.com/DeAccountSystems/das_commonlib/common"
	lru "github.com/hashicorp/golang-lru"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2021
 * FileName: normal_cell
 * Author:   LinGuanHong
 * Date:     2021/3/24 11:52
 * Description:
 */

type NormalCellRepeater struct {
	syncLocker          sync.Mutex
	secsOverdue         int64
	normalCellRecorder  *lru.Cache
}

func NewNormalCellRepeater(secsOverdue int64) *NormalCellRepeater {
	recorder, _ := lru.New(1000)
	return &NormalCellRepeater{
		syncLocker:          sync.Mutex{},
		secsOverdue:         secsOverdue,
		normalCellRecorder:  recorder,
	}
}

func (c *NormalCellRepeater) LoadLiveCells(client rpc.Client, searchKey *indexer.SearchKey, capNeed uint64) ([]indexer.LiveCell, uint64, error) {
	c.syncLocker.Lock()
	defer c.syncLocker.Unlock()
	return common.LoadLiveNormalCells(client, searchKey, capNeed, func(cell *indexer.LiveCell) bool {
		// return ture means will use this one
		hexBytes, _ := cell.OutPoint.Serialize()
		return c.canUse(hexBytes)
	})
}

func (c *NormalCellRepeater) canUse(hexBytes []byte) bool {
	if value, ok := c.normalCellRecorder.Get(hex.EncodeToString(hexBytes)); ok && value != nil {
		return time.Now().Unix()-value.(int64) > c.secsOverdue
	}
	return true
}

func (c *NormalCellRepeater) Record(outPoints []*types.OutPoint) {
	c.syncLocker.Lock()
	defer c.syncLocker.Unlock()
	for _, outPoint := range outPoints {
		keyBytes, _ := outPoint.Serialize()
		c.normalCellRecorder.Add(hex.EncodeToString(keyBytes),time.Now().Unix())
	}
}
