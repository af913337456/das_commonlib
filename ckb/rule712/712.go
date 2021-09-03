package rule712

import (
	"encoding/hex"
	"fmt"
	"github.com/DeAccountSystems/das_commonlib/ckb/builder"
	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
	"github.com/DeAccountSystems/das_commonlib/ckb/gotype"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"math/big"
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

type MMJson struct {
	TemplateStr string `json:"template_str"`
	action string `json:"action"`
	fee uint64
	inputsCapacity  uint64
	outputsCapacity uint64
	plainText string
	digest string
}

type ActionParam712 struct {
	Action string `json:"action"`
	Params string `json:"params"`
}

func (m *MMJson) FillDigest(digest string){
	m.digest = digest
}

func (m *MMJson) Fill712Action(action string,isOwner bool) {
	param := "0x01"
	if isOwner {
		param = "0x00"
	}
	m.action = fmt.Sprintf(`{
		"action": %s,
		"params": %s
	}`, action, param)
}

func (m *MMJson) Fill712Capacity(txBuilder *builder.TransactionBuilder) error {
	fee,inputCap,outputCap,err := txBuilder.InputsOutputsFeeCapacity()
	if err != nil {
		return fmt.Errorf("InputsOutputsFeeCapacity err: %s",err.Error())
	}
	m.fee = fee
	m.inputsCapacity = inputCap
	m.outputsCapacity = outputCap
	return nil
}

// Transfer the account xxxxxxxxxx.bit from ETH:0x11111111111111 to TRX:0x22222222222222.
var transferAccountPlainText = "Transfer the account %s from %s:%s to %s:%s."
func (m *MMJson) FillTransferAccountPlainText(accountCell gotype.AccountCell,newOwnerParam celltype.DasLockArgsPairParam) {
	originOwnerIndexType := celltype.DasLockCodeHashIndexType(accountCell.DasLockArgs[0])
	originOwnerAddrBytes := accountCell.DasLockArgs[1:celltype.DasLockArgsMinBytesLen/2]
	newOwnerAddrBytes := newOwnerParam.Script.Args[1:celltype.DasLockArgsMinBytesLen/2]
	account,_ := celltype.AccountFromOutputData(accountCell.Data)
	m.plainText = fmt.Sprintf(
		transferAccountPlainText,
		account,
		originOwnerIndexType.ChainType().String(),
		hex.EncodeToString(originOwnerAddrBytes),
		newOwnerParam.HashIndexType.ChainType().String(),
		hex.EncodeToString(newOwnerAddrBytes))
}

// Transfer from ckb1xxxx(111.111 CKB), ckb1yyyy(222.222 CKB) to ckb1zzzz(333 CKB), ckb1zzzz(0.333 CKB).
type WithdrawPlainTextOutputParam struct {
	ReceiverCkbScript types.Script
	Amount  uint64
}
// now, always withdraw all the money, so there is no change cell
func (m *MMJson) FillWithdrawPlainText(isTestNet bool,inputs []gotype.WithdrawDasLockCell, output WithdrawPlainTextOutputParam) {
	inputStr := ""
	inputSize := len(inputs)
	ckbValueStr := func(cellCap uint64) string {
		first  := new(big.Rat).SetInt(new(big.Int).SetUint64(cellCap))
		second := new(big.Rat).SetInt(new(big.Int).SetUint64(celltype.OneCkb))
		return new(big.Rat).Quo(first,second).FloatString(8)
	}
	for i :=0; i<inputSize; i++ {
		item := inputs[i]
		hashIndex := celltype.DasLockCodeHashIndexType(item.LockScriptArgs[0])
		str := gotype.PubkeyHashToAddress(isTestNet,hashIndex.ChainType(),hex.EncodeToString(item.LockScriptArgs[1:celltype.DasLockArgsMinBytesLen/2]))
		ckbValueStr := ckbValueStr(item.CellCap)
		if i == inputSize - 1 {
			inputStr = inputStr + fmt.Sprintf("%s(%s CKB) ",str,removeSuffixZeroChar(ckbValueStr))
		} else {
			inputStr = inputStr + fmt.Sprintf("%s(%s CKB), ",str,removeSuffixZeroChar(ckbValueStr))
		}
	}
	receiverAddr := gotype.PubkeyHashToAddress(isTestNet,celltype.ChainType_CKB,hex.EncodeToString(output.ReceiverCkbScript.Args))
	inputStr = inputStr + fmt.Sprintf("to %s(%s CKB)",receiverAddr,ckbValueStr(output.Amount))
	m.plainText = fmt.Sprintf("Transfer from %s.",inputStr)
}

func CreateMMJsonB(txDigestHexStr string) string {
	return MMJsonA
}

func removeSuffixZeroChar(ckbValueStr string) string {
	size := len(ckbValueStr)
	index := 0
	for i :=size-1; i >= 0; i-- {
		if ckbValueStr[i] == '0' {
			index ++
		} else {
			break
		}
	}
	return ckbValueStr[0:size-index]
}











