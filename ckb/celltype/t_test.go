package celltype

import (
	"fmt"
	"testing"
)

/**
 * Copyright (C), 2019-2020
 * FileName: t_test
 * Author:   LinGuanHong
 * Date:     2020/12/30 4:23 下午
 * Description:
 */

func Test_NewDasWitnessData(t *testing.T) {
	obj := NewDasWitnessData(1, []byte("china"))
	t.Log(obj.ToWitness())
	das, err := NewDasWitnessDataFromSlice(obj.ToWitness())
	if err != nil {
		panic(err.Error())
	} else {
		t.Log(das.Tag, das.TableType)
	}
}

func Test_AppendProposeWitnessSliceDataObjectList(t *testing.T) {
	p := ProposeWitnessSliceDataObjectList{}
	p.Add([]byte("1"), []byte("2"), AccountCellStatus_Proposed)
	p.Add([]byte("11"), []byte("21"), AccountCellStatus_Proposed)
	fmt.Println(string(p[0].AccountId))
	fmt.Println(string(p[1].AccountId))
}
