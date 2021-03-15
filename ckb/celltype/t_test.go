package celltype

import (
	"encoding/hex"
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

func Test_CalPreAccountCellCap(t *testing.T) {
	t.Log(CalPreAccountCellCap(1, 5000000, 1000, "12345678.bit"))
}

func Test_Rat(t *testing.T) {
	v := GoUint32ToMoleculeU32(3000)
	fmt.Println(MoleculeU32ToGoPercentage(v.RawData()))

	fmt.Println(CalDasAwardCap(100*OneCkb, 0.3))

	fmt.Println(EmptyAccountId.Bytes())

	bys := []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 4, 4, 3, 2, 1, 3, 3, 4, 4}
	a := SetAccountCellNextAccountId(bys, EmptyAccountId)
	fmt.Println(a)
}

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
	t.Log(hex.EncodeToString(obj.ToWitness()))
	das, err := NewDasWitnessDataFromSlice(obj.ToWitness())
	if err != nil {
		panic(err.Error())
	} else {
		t.Log(hex.EncodeToString(das.ToWitness()))
		t.Log(das.Tag, das.TableType)
	}
}

func Test_AppendProposeWitnessSliceDataObjectList(t *testing.T) {

}

func Test_SliceValue(t *testing.T) {
}
