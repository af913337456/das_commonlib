package chainpool

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2021
 * FileName: tx
 * Author:   LinGuanHong
 * Date:     2021/3/30 11:36
 * Description: transaction cache pool to quickly construct transactions
 */

/**
The strategy of quickly acquiring pre occupied cells from memory pool is as follows:

1. Before checking, search whether the newly generated accountcell's outpoint's TX has the corresponding accountid;
2. If there is data, call get_pool_tx interface;
3. Compare whether there is 1 in 2, if there is, use it directly to prove that it has entered the memory pool;
4. If there is no interface synchronization, the rotation training will continue to wait, because it is faster to enter the local node memory pool;
5. After 5 rounds of training, if not, call the livecell interface.

*/

type RecentUsedTxInfoObj struct {
	TimeUnixAdd int64           `json:"time_unix_add"`
	OutPoint 	types.OutPoint  `json:"out_point"`
	DetailObj   interface{} 	`json:"detail_obj"`
}

var (
	ChainTxPoolManager 			   *ChainTxPool
	RecentAccountCellUsedTxManager *RecentUsedTxManager
	RecentManagerCellUsedTxManager *RecentUsedTxManager
	RecentIncomeCellUsedTxManager  *RecentUsedTxManager
)

func init() {
	// single instance, you can re init outside
	ChainTxPoolManager = NewChainTxPool()
	RecentAccountCellUsedTxManager = NewRecentUsedTxManager(5,time.Second * 3, 60 * 2)
	RecentManagerCellUsedTxManager = NewRecentUsedTxManager(5,time.Second * 3, 60 * 2)
	RecentIncomeCellUsedTxManager  = NewRecentUsedTxManager(6,time.Second * 3, 60 * 2)
}

type RecentUsedTxManager struct {
	UsedTx sync.Map
	RetryTime  int
	RetryDelay time.Duration
	overdueSec int64
}

func NewRecentUsedTxManager(retryTime int,retryDelay time.Duration, overdueSec int64) *RecentUsedTxManager {
	return &RecentUsedTxManager{
		UsedTx: sync.Map{},
		RetryDelay: retryDelay,
		RetryTime:  retryTime,
		overdueSec: overdueSec}
}

func (r *RecentUsedTxManager) PopOneLocalRecentUsedTx(key interface{}) *RecentUsedTxInfoObj {
	if val, ok := r.UsedTx.Load(key); ok {
		// exist
		i := val.(RecentUsedTxInfoObj)
		r.UsedTx.Delete(key)
		if time.Now().Unix() - i.TimeUnixAdd > r.overdueSec {
			return nil // too long, direct use of the chain，we should remove it
		}
		return &i
	}
	return nil
}

func (r *RecentUsedTxManager) PopOneRecentUsedTx(chainTxPool *ChainTxPool,key interface{}) *RecentUsedTxInfoObj {
	if val, ok := r.UsedTx.Load(key); ok {
		// exist
		i := val.(RecentUsedTxInfoObj)
		if time.Now().Unix() - i.TimeUnixAdd > r.overdueSec {
			r.UsedTx.Delete(key)
			return nil // too long, direct use of the chain，we should remove it
		}
		// then check pool
		loopCounter := 0
		for {
			if loopCounter >= r.RetryTime {
				return nil
			}
			if target := i.OutPoint.TxHash.String(); chainTxPool.FindOneTx(target) {
				// pool exist, means can use
				r.UsedTx.Delete(key)
				return &i
			} else {
				// wait
				time.Sleep(r.RetryDelay)
			}
			loopCounter++
		}
	} else {
		chainTxPool.Print()
	}
	return nil
}

func (r *RecentUsedTxManager) AddRecentUsedTx(accountId celltype.DasAccountId,tx RecentUsedTxInfoObj)  {
	r.UsedTx.Store(accountId,tx)
}

func (r *RecentUsedTxManager) AddRecentUsedTx2(tx RecentUsedTxInfoObj)  {
	r.UsedTx.Store(tx.OutPoint.TxHash,tx)
}

func (r *RecentUsedTxManager) RemoveOneRecentUsedTx(accountId celltype.DasAccountId,newTxHash types.Hash)  {
	if val, ok := r.UsedTx.Load(accountId); ok {
		if val.(RecentUsedTxInfoObj).OutPoint.TxHash == newTxHash {
			r.UsedTx.Delete(accountId)
		}
	}
}

func (r *RecentUsedTxManager) RemoveOneRecentUsedTxByHash(txHash types.Hash)  {
	if _, ok := r.UsedTx.Load(txHash); ok {
		r.UsedTx.Delete(txHash)
	}
}

type ChainTxPool struct {
	syncLock  sync.Mutex
	retMap    map[string][]string // this won't be very big
}

func NewChainTxPool() *ChainTxPool {
	return &ChainTxPool{syncLock:sync.Mutex{},retMap: map[string][]string{}}
}

func (c *ChainTxPool) StartLoopGetRawTxPool(ctx context.Context,rpcClient rpc.Client,duration time.Duration) {
	c.getRawTxPool(rpcClient)
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	for {
		if ctx != nil && ctx.Err() != nil {
			return
		}
		select {
		case <-ticker.C:
			c.getRawTxPool(rpcClient)
		default:
			time.Sleep(time.Second)
		}
	}
}

// each request will recover retMap
func (c *ChainTxPool) getRawTxPool(rpcClient rpc.Client) {
	c.syncLock.Lock()
	defer c.syncLock.Unlock()
	if err := rpcClient.CallContext(context.TODO(),&c.retMap,"get_raw_tx_pool",false); err != nil {

	}
}

func (c *ChainTxPool) Print() string {
	if c.retMap != nil {
		fmt.Println(len(c.retMap))
		bys,err := json.MarshalIndent(c.retMap," "," ")
		if err != nil {
			return err.Error()
		} else {
			return string(bys)
		}
	}
	return "empty"
}

func (c *ChainTxPool) FindOneTx(targetTxHash string) bool {
	if c.retMap != nil {
		c.syncLock.Lock()
		defer c.syncLock.Unlock()
		for _, array := range c.retMap {
			for _, txHash := range array {
				if txHash == targetTxHash {
					return true
				}
			}
		}
	}
	return false
}
















