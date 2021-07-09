package wallet

import (
	"fmt"
	"testing"
)

/**
 * Copyright (C), 2019-2021
 * FileName: wallet_test
 * Author:   LinGuanHong
 * Date:     2021/7/2 10:59
 * Description:
 */

func Test_RecoverAddrFromLockScriptArgs(t *testing.T) {
	fmt.Println(GetShortAddressFromLockScriptArgs("0x59ddf0d5d61386b6b3de4b4e8a74c0045f0a410a",false))
	fmt.Println(GetShortAddressFromLockScriptArgs("0x59ddf0d5d61386b6b3de4b4e8a74c0045f0a410a",true))
}

func Test_CreateWallet(t *testing.T) {
	j,_ := CreateCKBWallet(false)
	// {"PriKeyHex":"bd9af72d4d2243ddfec1e05c8409a1ff78f1f13975d49f70ea5259b55d763d11","PubKeyHex":"02c22cb57e628f1639277b79702f9f434878e8e20c88d10e14e846a97f28bb7014","AddressHex":"ckb1qyqg2n9cnmmujkmvrp6xx04dcx7nlw705ezst9lupq"}
	t.Log(j.Json())
}














