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

type setIndex func()

func index(f setIndex) {
	f()
}

func read(p *AccountCellDataPreObj_Old_New) {
	fmt.Println("-->", p.InputIndex)
}

func Test_AccountCellDataPreObj_Old_New(t *testing.T) {
	v := &AccountCellDataPreObj_Old_New{}
	index(func() {
		v.InputIndex = 1001
	})
	read(v)
}

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

}

func Test_SliceValue(t *testing.T) {
}
