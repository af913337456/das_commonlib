package gotype

import (
	"encoding/hex"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
)

/**
 * Copyright (C), 2019-2021
 * FileName: addr_test
 * Author:   LinGuanHong
 * Date:     2021/7/9 10:57
 * Description:
 */

func Test_Addr(t *testing.T) {
	tronAddr := Address("TQoLh9evwUmZKxpD1uhFttsZk3EBs8BksV")
	b,e := tronAddr.HexBys(types.HexToHash("1"))
	if e != nil {
		panic(e)
	}
	t.Log(hex.EncodeToString(b))
}

