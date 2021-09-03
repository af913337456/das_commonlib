package gotype

import (
	"encoding/hex"
	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
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

func Test_PubkeyHashToAddress(t *testing.T) {
	t.Log(PubkeyHashToAddress(true, celltype.ChainType_CKB,"dc36477cf2434288a5502120ef0fd919ae37c155"))
	t.Log(PubkeyHashToAddress(false, celltype.ChainType_TRON,""))
}

