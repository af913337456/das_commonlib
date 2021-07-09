package tron_chain

import "testing"

/**
 * Copyright (C), 2019-2021
 * FileName: addr_test
 * Author:   LinGuanHong
 * Date:     2021/7/9 10:40
 * Description:
 */

func Test_PubkeyHexFromBase58(t *testing.T) {
	t.Log(PubkeyHexFromBase58("TQoLh9evwUmZKxpD1uhFttsZk3EBs8BksV"))
	t.Log(PubkeyHexFromBase58("1TQoLh9evwUmZKxpD1uhFttsZk3EBs8BksV"))
}
