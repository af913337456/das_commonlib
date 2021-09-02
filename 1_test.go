package das_commonlib

import (
	"encoding/hex"
	"fmt"
	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
	"github.com/DeAccountSystems/das_commonlib/ckb/gotype/configcells"
	ckbtype "github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
)

/**
 * Copyright (C), 2019-2021
 * FileName: 1_test
 * Author:   LinGuanHong
 * Date:     2021/8/31 10:57 上午
 * Description:
 */

func Test_ParseConfigCellAccount(t *testing.T) {
	hexStr := "0x6461736800000001020000100000001100000035010000012401000024000000440000006400000084000000a4000000c4000000e400000004010000334540e23ec513f691cdd9490818237cbc9675861e4f19c480e0c520c715fd348b0ab9073521cc7c7beb7a15368e8600188012594c4e6449316f6ecbf07a1da1000000000000000000000000000000000000000000000000000000000000000055faf1d18e77c640de0e4a1e0193884ed7bda3ebdf1f34281e4f83f198d960960000000000000000000000000000000000000000000000000000000000000000179934b1f21db4c99081e1c585f83b7410a4cf94eaefa1ba3eb320cd35cc5f02c8a23512a45e45ece5b2358a80d9a837501de622ba482960b6b578fffd2de36619ff9b7d2956f4b59ef30b196949d7fc0a284c3472c716bc95ea07c7bec38cbecc000000180000003c0000006000000084000000a8000000209b35208da7d20d882f0871f3979c68c53981bcc4caa71274c035449074d08200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b035c200bf759537d3796edf49b5d6a8ec5f5d78326713f987f31ad24d0b0171000000007dc4ae8fe597045fbd7fe78f2bd26435644a69b755de3824a856f681bacb732b00000000"
	rawWitnessBytes,err := hex.DecodeString(hexStr[2:])
	if err != nil {
		fmt.Println(1,err.Error())
	}
	dasObj, err := celltype.ParseTxWitnessToDasWitnessObj(rawWitnessBytes)
	if err != nil {
		fmt.Println(2,err.Error())
	}
	fmt.Println(dasObj.WitnessObj.TableType)
	cellData := dasObj.MoleculeNewDataEntity.Entity().RawData()
	child := configcells.CfgMain{}
	err = child.NotifyData(&configcells.ConfigCellChildDataObj{
		CellDep:      ckbtype.CellDep{},
		WitnessData:  rawWitnessBytes,
		MoleculeData: cellData,
	})
	if err != nil {
		fmt.Println(3,err.Error())
	}
}