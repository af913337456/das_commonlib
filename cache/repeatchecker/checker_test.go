package repeatchecker

import (
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
)

/**
 * Copyright (C), 2019-2021
 * FileName: checker_test
 * Author:   LinGuanHong
 * Date:     2021/6/7 1:04 下午
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














