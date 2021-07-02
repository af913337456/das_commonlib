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
