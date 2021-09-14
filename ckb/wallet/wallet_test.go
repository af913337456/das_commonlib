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
	fmt.Println(GetShortAddressFromLockScriptArgs("",false))
	fmt.Println(GetShortAddressFromLockScriptArgs("",true))
	address,_ := GetLockScriptArgsFromShortAddress("")
	fmt.Println(address)
}














