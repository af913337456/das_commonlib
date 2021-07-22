package repeatchecker

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
	"github.com/DeAccountSystems/das_commonlib/common"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
	"testing"
)

/**
 * Copyright (C), 2019-2021
 * FileName: checker_test
 * Author:   LinGuanHong
 * Date:     2021/6/7 1:04
 * Description:
 */

func Test_NormalCellRepeater(t *testing.T) {
	repeater := NewNormalCellRepeater(120)
	point := &types.OutPoint{
		TxHash: types.HexToHash("123"),
		Index:  0,
	}
	point2 := &types.OutPoint{
		TxHash: types.HexToHash("1234"),
		Index:  0,
	}
	repeater.Record([]*types.OutPoint{point})
	pointBytes,_ := point.Serialize()
	fmt.Println(repeater.canUse(pointBytes))

	pointBytes2,_ := point2.Serialize()
	fmt.Println(repeater.canUse(pointBytes2))
}

func Test_LoadNormalLiveCell(t *testing.T) {
	// 20e23005afba0652e6ddf0649f984f78bc3980eb1b83ef104d84ef137b487cf8
	rpcClient, err := rpc.DialWithIndexerContext(
		context.TODO(),
		"http://47.244.126.25:8114", // 47.242.53.82:8114
		"http://47.244.126.25:8116")
	if err != nil {
		panic(fmt.Errorf("init rpcClient failed: %s", err.Error()))
	}
	systemScript, err := utils.NewSystemScripts(rpcClient)
	if err != nil {
		panic(err)
	}
	bys, _ := hex.DecodeString("d6465f6faf694011c770bb4e5aeed5b32865183b")
	lockScript := types.Script{
		CodeHash: systemScript.SecpSingleSigCell.CellHash,
		HashType: "type",
		Args:     bys,
	}
	searchNormalInputKey := &indexer.SearchKey{Script: &lockScript, ScriptType: indexer.ScriptTypeLock}
	// checker := repeatchecker.NewNormalCellRepeater(120)
	liveCells, totalCap, err := common.LoadLiveCellsWithSize(rpcClient, searchNormalInputKey, 142*celltype.OneCkb,1,true,true,nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(liveCells), totalCap)
	//for index, item := range liveCells {
	//	fmt.Println(item.OutPoint.TxHash.String(), index)
	//}
}












