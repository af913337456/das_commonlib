package ckb

//
// import (
// 	"encoding/hex"
// 	"fmt"
// 	"math/big"
// 	"testing"
//
// 	"github.com/DA-Services/das_commonlib/ckb/builder"
// 	"github.com/DA-Services/das_commonlib/ckb/celltype"
// 	"github.com/nervosnetwork/ckb-sdk-go/address"
// 	"github.com/nervosnetwork/ckb-sdk-go/indexer"
// 	"github.com/nervosnetwork/ckb-sdk-go/rpc"
// 	"github.com/nervosnetwork/ckb-sdk-go/types"
// 	"github.com/nervosnetwork/ckb-sdk-go/utils"
// )
//
// /**
//  * Copyright (C), 2019-2020
//  * FileName: test_tx
//  * Author:   LinGuanHong
//  * Date:     2020/12/15 12:15
//  * Description:
//  */
//
// func Test_TxBuilder(t *testing.T) {
// 	client, err := rpc.DialWithIndexer("http://127.0.0.1:8114", "http://127.0.0.1:8116")
// 	//client,err := rpc.DialWithIndexer("http://:8114","http://:8116")
// 	if err != nil {
// 		panic(err)
// 	}
// 	systemScripts, err := utils.NewSystemScripts(client)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	fromAddress, err := address.Parse("ckt1qyqq62ckhmdqmq2ucpars0763ra7k2d7cdcs0y8kk2")
// 	if err != nil {
// 		panic(err)
// 	}
// 	searchKey := &indexer.SearchKey{
// 		Script:     fromAddress.Script,
// 		ScriptType: indexer.ScriptTypeLock,
// 	}
// 	collector := utils.NewLiveCellCollector(
// 		client, searchKey, indexer.SearchOrderAsc, 1000, "",
// 		utils.NewCapacityLiveCellProcessor(celltype.CkbTxMinOutputCKBValue))
// 	result, err := collector.Collect()
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	cellData1 := celltype.NewConfigCellDataBuilder().Build()
// 	// cellData2 := celltype.NewActionCellDataBuilder().Build()
// 	// cellData3 := celltype.NewAccountCellDataBuilder().Build()
// 	fmt.Println(cellData1)
// 	txBuilder :=
// 		builder.NewTransactionBuilder2(nil, 10000).
// 			AddDasSpecOutput(celltype.NewAccountCell(nil))
// 		// AddDasSpecOutput(celltype.NewActionCell(celltype.TestNetActionCell(&cellData2))).
// 		// AddDasSpecOutput(celltype.NewAccountCell(celltype.TestNetAccountCell(&cellData3)))
//
// 	// 计算需要的 input
// 	if err := txBuilder.AddInputAutoComputeItems(result); err != nil {
// 		panic(err.Error())
// 	}
//
// 	// 找零，SecpSingleSigCell 找零给个人
// 	txBuilder.AddChargeOutput(txBuilder.FromScript(), systemScripts.SecpSingleSigCell)
//
// 	t.Log("tx ====> ", txBuilder.Log())
// 	txHashStr, err := txBuilder.TxHash()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	t.Log(txHashStr)
// }
//
// func Test_StateCellTx(t *testing.T) {
// 	// client,err := rpc.DialWithIndexer("http://127.0.0.1:8114","http://127.0.0.1:8116")
// 	client, err := rpc.DialWithIndexer("http://:8114", "http://:8116")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	systemScripts, err := utils.NewSystemScripts(client)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	fromAddress, err := address.Parse("ckt1qyqq62ckhmdqmq2ucpars0763ra7k2d7cdcs0y8kk2")
// 	if err != nil {
// 		panic(err)
// 	}
// 	searchKey := &indexer.SearchKey{
// 		Script:     fromAddress.Script,
// 		ScriptType: indexer.ScriptTypeLock,
// 	}
// 	collector := utils.NewLiveCellCollector(
// 		client, searchKey, indexer.SearchOrderAsc, 1000, "",
// 		utils.NewCapacityLiveCellProcessor(celltype.CkbTxMinOutputCKBValue))
// 	_, err = collector.Collect()
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	txBuilder, _ := builder.NewTransactionBuilder1("", 1000000)
// 	txBuilder.AddCellDep(&types.CellDep{
// 		OutPoint: &types.OutPoint{
// 			TxHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnAggron),
// 			Index:  0,
// 		},
// 		DepType: types.DepTypeDepGroup,
// 	})
// 	txBuilder.AddCellDep(&types.CellDep{ // state_cell
// 		OutPoint: systemScripts.SecpSingleSigCell.OutPoint,
// 		DepType:  types.DepTypeDepGroup,
// 	})
// 	// txBuilder.AddInput(&types.CellInput{
// 	// 	Since:          0,
// 	// 	PreviousOutput: result.LiveCells[0].OutPoint,
// 	// })
// 	// totalCap := result.Capacity
// 	// txBuilder.addOutputAutoComputeCap(fromAddress.Script, nil, nil, nil)
// 	txHashStr, err := txBuilder.TxHash()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	t.Log(txHashStr)
// }
//
// func Test_OutputCell(t *testing.T) {
// 	/**
// 	{
// 	      "capacity": "0x34e62ce00",
// 	      "lock": {
// 	        "code_hash": "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
// 	        "hash_type": "type",
// 	        "args": "0x7eb507d383e1d4ce0b0d1cada812cf4fd6a936e1"
// 	      },
// 	      "type": {
// 	        "code_hash": "0xe7f93d7120de3ca8548b34d2ab9c40fe662eec35023f07e143797789895b4869",
// 	        "hash_type": "data",
// 	        "args": "0x4135551dbf738dde43d1b641d0e0c2d3ac8da62996b52c497790b675e300d967"
// 	      }
// 	    }
// 	*/
// 	arg1, _ := hex.DecodeString("7eb507d383e1d4ce0b0d1cada812cf4fd6a936e1")
// 	arg2, _ := hex.DecodeString("4135551dbf738dde43d1b641d0e0c2d3ac8da62996b52c497790b675e300d967")
// 	b, _ := new(big.Int).SetString("34e62ce00", 16)
// 	t.Log(b.Uint64())
// 	o := types.CellOutput{
// 		Capacity: b.Uint64(),
// 		Lock: &types.Script{
// 			CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
// 			HashType: types.HashTypeType,
// 			Args:     arg1,
// 		},
// 		Type: &types.Script{
// 			CodeHash: types.HexToHash("0xe7f93d7120de3ca8548b34d2ab9c40fe662eec35023f07e143797789895b4869"),
// 			HashType: types.HashTypeData,
// 			Args:     arg2,
// 		},
// 	}
// 	data, _ := hex.DecodeString("a0860100000000000000000000000000")
// 	t.Log(o.OccupiedCapacity(data))
// }
//
// func Test_OutputCell2(t *testing.T) {
// 	arg1, err := hex.DecodeString("7eb507d383e1d4ce0b0d1cada812cf4fd6a936e1")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	t.Log(len(arg1))
// 	b, _ := new(big.Int).SetString("34e62ce00", 16)
// 	t.Log(b.Uint64())
// 	o := types.CellOutput{
// 		Capacity: b.Uint64(),
// 		Lock: &types.Script{
// 			CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
// 			HashType: types.HashTypeType,
// 			Args:     arg1,
// 		},
// 		Type: nil,
// 	}
// 	data, _ := hex.DecodeString("0x")
// 	t.Log(o.OccupiedCapacity(data))
// }
