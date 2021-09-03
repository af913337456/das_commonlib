package rule712

import (
	"encoding/hex"
	"fmt"
	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
	"github.com/DeAccountSystems/das_commonlib/ckb/gotype"
)

/**
 * Copyright (C), 2019-2021
 * FileName: 712
 * Author:   LinGuanHong
 * Date:     2021/9/3 10:50
 * Description:
 */

var MMJsonA = `{
  "types": {
    "EIP712Domain": [
      {"name": "chainId", "type": "uint256"},
      {"name": "name", "type": "string"},
      {"name": "verifyingContract", "type": "address"},
      {"name": "version", "type": "string"}
    ],
    "Action": [
      {"name": "action", "type": "string"},
      {"name": "params", "type": "string"}
    ],
    "Cell": [
      {"name": "capacity", "type": "string"},
      {"name": "lock", "type": "string"},
      {"name": "type", "type": "string"},
      {"name": "data", "type": "string"},
      {"name": "extraData", "type": "string"}
    ],
    "Transaction": [
      {"name": "plainText", "type": "string"},
      {"name": "inputsCapacity", "type": "string"},
      {"name": "outputsCapacity", "type": "string"},
      {"name": "fee", "type": "string"},
      {"name": "action", "type": "Action"},
      {"name": "inputs", "type": "Cell[]"},
      {"name": "outputs", "type": "Cell[]"},
      {"name": "digest", "type": "bytes32"}
    ]
  },
  "primaryType": "Transaction",
  "domain": {
    "chainId": 1,
    "name": "da.systems",
    "verifyingContract": "0xb3dc32341ee4bae03c85cd663311de0b1b122955",
    "version": "1"
  },
  "message": {
    "plainText": {{ plainText }},
    "inputsCapacity": {{ inputsCapacity }},
    "outputsCapacity": {{ outputsCapacity }},
    "fee": {{ fee }},
    "action": {{ action }},
    "inputs": {{ inputs }},
    "outputs": {{ outputs }},
    "digest": %s
  }
}`

type ActionParam712 struct {
	Action string `json:"action"`
	Params string `json:"params"`
}
func Create712Action(action string,isOwner bool) ActionParam712 {
	param := "0x01"
	if isOwner {
		param = "0x00"
	}
	return ActionParam712{
		Action: action,
		Params: param,
	}
}

// Transfer the account xxxxxxxxxx.bit from ETH:0x11111111111111 to TRX:0x22222222222222.
var transferAccountPlainText = "Transfer the account %s from %s:%s to %s:%s."
func CreateTransferAccountPlainText(accountCell gotype.AccountCell,newOwnerParam celltype.DasLockArgsPairParam) string {
	originOwnerIndexType := celltype.DasLockCodeHashIndexType(accountCell.DasLockArgs[0])
	originOwnerAddrBytes := accountCell.DasLockArgs[1:celltype.DasLockArgsMinBytesLen/2]
	newOwnerAddrBytes := newOwnerParam.Script.Args[1:celltype.DasLockArgsMinBytesLen/2]
	return fmt.Sprintf(
		transferAccountPlainText,
		celltype.AccountFromOutputData(accountCell.Data),
		originOwnerIndexType.ChainType().String(),
		hex.EncodeToString(originOwnerAddrBytes),
		newOwnerParam.HashIndexType.ChainType().String(),
		hex.EncodeToString(newOwnerAddrBytes))
}

// Transfer from ckb1xxxx(111.111 CKB), ckb1yyyy(222.222 CKB) to ckb1zzzz(333 CKB), ckb1zzzz(0.333 CKB).
func CreateWithdrawPlainText(text string) string {
	return fmt.Sprintf("Transfer from %s.",text)
}

func CreateMMJsonB(txDigestHexStr string) string {
	return MMJsonA
}














