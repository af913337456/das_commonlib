package ckb_rocksdb_parser

import (
	"fmt"
)

/**
 * Copyright (C), 2019-2021
 * FileName: rocksdb_parser
 * Author:   LinGuanHong
 * Date:     2021/3/13 10:28
 * Description:
 */

type MyTestLogger struct{}

func (l *MyTestLogger) Info(args ...interface{})  { fmt.Println(args...) }
func (l *MyTestLogger) Error(args ...interface{}) { fmt.Println(args...) }
func (l *MyTestLogger) Warn(args ...interface{})  { fmt.Println(args...) }

// func NewTestRocksdbBlockParser() *scanner.BlockScanner {
// 	rpcClient := test.RpcClientWithoutSync()
// 	ckbChain := NewCKBBlockChainWithRpcClient(context.TODO(), rpcClient, nil)
// 	rocksDb := NewCKBRocksDb("data", MsgHandler{
// 		Receive: func(info *keeperTypes.TxMsgData) error {
// 			fmt.Println("receiver msg:", info.BlockBaseInfo.BlockNumber)
// 			return nil
// 		},
// 		Close: nil,
// 	})
// 	return scanner.NewBlockScanner(scanner.InitBlockScanner{
// 		Chain: ckbChain,
// 		Db:    rocksDb,
// 		Log:   new(MyTestLogger),
// 		Control: types.DelayControl{
// 			RoundDelay: time.Millisecond,
// 			CatchDelay: time.Second,
// 		},
// 		FrontNumber: 3,
// 	})
// }

// func Test_Blockparser(t *testing.T) {
// 	parser := NewTestRocksdbBlockParser()
// 	_ = parser.SetStartScannerHeight(1859963)
// 	if err := parser.Start(); err != nil {
// 		panic(err)
// 	}
// 	select {}
// }
