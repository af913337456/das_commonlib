package celltype

import (
	"encoding/hex"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
)

/**
 * Copyright (C), 2019-2021
 * FileName: account_type
 * Author:   LinGuanHong
 * Date:     2021/2/1 2:21 下午
 * Description:
 */

type DasAccount string

func DasAccountFromStr(account string) DasAccount {
	return DasAccount(account)
}

func (a DasAccount) AccountId() DasAccountId {
	bys, _ := blake2b.Blake160([]byte(a))
	return bys
}

type DasAccountId []byte

func (a DasAccountId) HexStr() string {
	return hex.EncodeToString(a)
}

func (a DasAccountId) Str() string {
	return string(a)
}
