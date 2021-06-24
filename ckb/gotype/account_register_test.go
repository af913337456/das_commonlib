package gotype

import (
	"encoding/json"
	"fmt"
	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
	"testing"
	"time"
)

/**
 * Copyright (C), 2019-2021
 * FileName: account_register_test
 * Author:   LinGuanHong
 * Date:     2021/6/24 4:04
 * Description:
 */

func Test_CalAccountCellExpiredAt(t *testing.T) {
	fmt.Println(len(celltype.DasLockCellScript.Out.Args))
	param := celltype.CalAccountCellExpiredAtParam{
		Quote:             13464, // 1000 ckb = 1 usd
		AccountCellCap:    211 * celltype.OneCkb,
		PriceConfigNew:    6000000, // 10 usd
		PreAccountCellCap: 112200000000, // 566 * OneCkb,
		RefCellCap:        0,
		DiscountRate:      0,
	}
	timeSec, err := CalAccountCellExpiredAt(param, 1622085497)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("expiredAt:", timeSec)
		fmt.Println(time.Unix(int64(timeSec), 0).String())
		bys, _ := json.Marshal(param)
		t.Log(string(bys))
		// current(1617782601) + (profit(585600000000) / (price(5000000) / quote(1000) * 100_000_000)) * 365 * 86400
	}
}
