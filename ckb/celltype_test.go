package ckb
//
// import (
// 	"github.com/DA-Services/das_commonlib/ckb/celltype"
// 	"github.com/nervosnetwork/ckb-sdk-go/types"
// 	"testing"
// )
//
// /**
//  * Copyright (C), 2019-2020
//  * FileName: celltype_test
//  * Author:   LinGuanHong
//  * Date:     2020/12/18 12:20 下午
//  * Description:
//  */
//
// func Test_BuildStateCell(t *testing.T) {
// 	dataBuilder := celltype.NewConfigCellDataBuilder().Build()
// 	// StateCellDataFromSlice()
// 	/**
// 	"testnet": {
// 	      "codeHash": "0x3419a1c09eb2567f6552ee7a8ecffd64155cffe0f1796e6e61ec088d740c1356",
// 	      "hashType": "type",
// 	      "txHash": "0xec26b0f85ed839ece5f11c4c4e837ec359f5adc4420410f6453b1f6b60fb96a6",
// 	      "index": "0x0",
// 	      "depType": "depGroup"
// 	    }
// 	*/
// 	celltype.NewStateCell(&celltype.StateCellParam{
// 		Version: 1,
// 		Data:    &dataBuilder,
// 		CellCodeInfo: celltype.DASCellBaseInfo{
// 			Dep: celltype.DASCellBaseInfoDep{
// 				TxHash:  "",
// 				TxIndex: 0,
// 				DepType: types.DepTypeDepGroup,
// 			},
// 			Out: celltype.DASCellBaseInfoOut{
// 				CodeHash:     "",
// 				CodeHashType: types.HashTypeType,
// 				Args:         nil,
// 			},
// 		},
// 		AlwaysSpendableScriptInfo: celltype.DASCellBaseInfo{
// 			Dep: celltype.DASCellBaseInfoDep{
// 				TxHash:  "0xec26b0f85ed839ece5f11c4c4e837ec359f5adc4420410f6453b1f6b60fb96a6",
// 				TxIndex: 0,
// 				DepType: types.DepTypeDepGroup,
// 			},
// 			Out: celltype.DASCellBaseInfoOut{
// 				CodeHash:     "0x3419a1c09eb2567f6552ee7a8ecffd64155cffe0f1796e6e61ec088d740c1356",
// 				CodeHashType: types.HashTypeType,
// 				Args:         nil,
// 			},
// 		},
// 	})
// }
//
// func Test_BuildActionCell(t *testing.T) {
// 	celltype.NewActionCell(&celltype.ActionCellParam{
// 		Version: 1,
// 		Data:    nil,
// 		CellCodeInfo: celltype.DASCellBaseInfo{
// 			Dep: celltype.DASCellBaseInfoDep{
// 				TxHash:  "",
// 				TxIndex: 0,
// 				DepType: types.DepTypeDepGroup,
// 			},
// 			Out: celltype.DASCellBaseInfoOut{
// 				CodeHash:     "",
// 				CodeHashType: types.HashTypeType,
// 				Args:         nil,
// 			},
// 		},
// 		AlwaysSpendableScriptInfo: celltype.DASCellBaseInfo{
// 			Dep: celltype.DASCellBaseInfoDep{
// 				TxHash:  "0xec26b0f85ed839ece5f11c4c4e837ec359f5adc4420410f6453b1f6b60fb96a6",
// 				TxIndex: 0,
// 				DepType: types.DepTypeDepGroup,
// 			},
// 			Out: celltype.DASCellBaseInfoOut{
// 				CodeHash:     "0x3419a1c09eb2567f6552ee7a8ecffd64155cffe0f1796e6e61ec088d740c1356",
// 				CodeHashType: types.HashTypeType,
// 				Args:         nil,
// 			},
// 		},
// 	})
// }
