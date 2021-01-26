package wallet

import "testing"

/**
 * Copyright (C), 2019-2020
 * FileName: ckb_wallet_test
 * Author:   LinGuanHong
 * Date:     2020/12/24 9:56 上午
 * Description:
 */

func Test_InitWallet(t *testing.T) {
	key, err := InitCkbWallet("", nil)
	if err != nil {
		panic(err)
	}
	t.Log(key.Secp256k1Key.PubKey())
}
