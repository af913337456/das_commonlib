package celltype

import "testing"

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
