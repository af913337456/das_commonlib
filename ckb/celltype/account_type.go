package celltype

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	id := &DasAccountId{}
	id.SetBytes(bys)
	return *id
}

const dasAccountIdLen = 10

type DasAccountId [dasAccountIdLen]byte

func BytesToDasAccountId(b []byte) DasAccountId {
	var h DasAccountId
	h.SetBytes(b)
	return h
}

func HexToHash(s string) DasAccountId {
	return BytesToDasAccountId(common.FromHex(s))
}

func (d *DasAccountId) SetBytes(b []byte) {
	if len(b) > len(d) {
		b = b[len(b)-dasAccountIdLen:]
	}
	copy(d[dasAccountIdLen-len(b):], b)
}

func DasAccountIdFromBytes(accountRawBytes []byte) DasAccountId {
	bys, _ := blake2b.Blake160(accountRawBytes)
	id := &DasAccountId{}
	id.SetBytes(bys)
	return *id
}

func (d DasAccountId) HexStr() string {
	return hexutil.Encode(d[:])
}

func (d DasAccountId) Str() string {
	return d.HexStr()
}

func (d DasAccountId) Bytes() []byte {
	return d[:]
}
