package chainpool

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"testing"
	"time"
)

/**
 * Copyright (C), 2019-2021
 * FileName: pool_test
 * Author:   LinGuanHong
 * Date:     2021/3/30 11:50
 * Description:
 */

func rpcClient() rpc.Client {
	rpcClient, err := rpc.DialWithIndexerContext(
		context.TODO(),
		"http://:8114",
		"http://:8116")
	if err != nil {
		panic(fmt.Errorf("init rpcClient failed: %s", err.Error()))
	}
	return rpcClient
}

func Test_PoolLoop(t *testing.T) {
	client := rpcClient()
	chainPool :=  NewChainTxPool()
	go chainPool.StartLoopGetRawTxPool(nil,client,time.Second*5)
	go func() {
		for {
			t.Log("pool info:",chainPool.Print())
			time.Sleep(time.Second * 3)
		}
	}()
	select {}
}














