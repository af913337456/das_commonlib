package celltype

import (
	"encoding/hex"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
)

/**
 * Copyright (C), 2019-2021
 * FileName: code_scipt_test
 * Author:   LinGuanHong
 * Date:     2021/6/8 4:07
 * Description:
 */

func Test_CalTypeIdFromScript(t *testing.T) {
	/**
	"code_hash":
	0x00000000000000000000000000000000000000000000000000545950455f4944
	"args":
	0xeedd10c7d8fee85c119daf2077fea9cf76b9a92ddca546f1f8e0031682e65aee
	*/
	bys,_ := hex.DecodeString("eedd10c7d8fee85c119daf2077fea9cf76b9a92ddca546f1f8e0031682e65aee")
	idHash := CalTypeIdFromScript(&types.Script{
		CodeHash: types.HexToHash("0x00000000000000000000000000000000000000000000000000545950455f4944"),
		HashType: types.HashTypeType,
		Args:     bys,
	})
	fmt.Println(idHash.String())
}

