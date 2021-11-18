package rule712

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/DeAccountSystems/das_commonlib/ckb/builder"
	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
	"github.com/DeAccountSystems/das_commonlib/ckb/gotype"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"math/big"
	"strings"
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
      {"name": "DAS_MESSAGE", "type": "string"},
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
    "chainId": %d,
    "name": "da.systems",
    "verifyingContract": "0x0000000000000000000000000000000020210722",
    "version": "1"
  },
  "message": {
    "DAS_MESSAGE": "%s",
    "inputsCapacity": "%s CKB",
    "outputsCapacity": "%s CKB",
    "fee": "%s CKB",
    "action": %s,
    "inputs": %s,
    "outputs": %s,
    "digest": "%s"
  }
}`

const maxHashLen = 20

type MMJsonObj struct {
	Types struct {
		EIP712Domain []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"EIP712Domain"`
		Action []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"Action"`
		Cell []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"Cell"`
		Transaction []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"Transaction"`
	} `json:"types"`
	PrimaryType string `json:"primaryType"`
	Domain      struct {
		ChainID           int    `json:"chainId"`
		Name              string `json:"name"`
		VerifyingContract string `json:"verifyingContract"`
		Version           string `json:"version"`
	} `json:"domain"`
	Message struct {
		DasMessage      string                `json:"DAS_MESSAGE"`
		InputsCapacity  string                `json:"inputsCapacity"`
		OutputsCapacity string                `json:"outputsCapacity"`
		Fee             string                `json:"fee"`
		Action          interface{}           `json:"action"`
		Inputs          []inputOutputParam712 `json:"inputs"`
		Outputs         []inputOutputParam712 `json:"outputs"`
		Digest          string                `json:"digest"`
	} `json:"message"`
}

type MMJson struct {
	dasMessage      string `json:"DAS_MESSAGE"`
	inputsCapacity  uint64 `json:"inputsCapacity"`
	outputsCapacity uint64 `json:"outputsCapacity"`
	fee             uint64 `json:"fee"`
	action          string `json:"action"`
	inputs          string `json:"inputs"`
	outputs         string `json:"outputs"`
	digest          string `json:"digest"`
}

func (m *MMJson) FillInputs(inputList InputOutputParam712List, accountData *celltype.AccountCellData) error {
	inputStr, err := inputList.To712Json(accountData)
	if err != nil {
		return err
	}
	m.inputs = inputStr
	return nil
}

func (m *MMJson) FillOutputs(outputList InputOutputParam712List, accountData *celltype.AccountCellData) error {
	outputStr, err := outputList.To712Json(accountData)
	if err != nil {
		return err
	}
	m.outputs = outputStr
	return nil
}

func (m *MMJson) Build(evmChainId int64, net celltype.DasNetType) (*MMJsonObj, error) {
	if evmChainId == 0 {
		evmChainId = 1
		if net != celltype.DasNetType_Mainnet {
			evmChainId = 5
		}
	}
	tmp := fmt.Sprintf(MMJsonA, evmChainId,
		m.dasMessage,
		removeSuffixZeroChar(ckbValueStr(m.inputsCapacity)),
		removeSuffixZeroChar(ckbValueStr(m.outputsCapacity)),
		removeSuffixZeroChar(ckbValueStr(m.fee)), m.action,
		m.inputs, m.outputs, m.digest)
	ret := MMJsonObj{}
	if err := json.Unmarshal([]byte(tmp), &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

type InputOutputParam712List []InputOutputParam712
type InputOutputParam712 struct {
	Capacity uint64        `json:"-"`
	Lock     *types.Script `json:"-"`
	Type     *types.Script `json:"-"`
	Data     []byte        `json:"-"`
}

type inputOutputParam712 struct {
	Capacity  string `json:"capacity"`
	LockStr   string `json:"lock"`
	TypeStr   string `json:"type"`
	Data      string `json:"data"`
	ExtraData string `json:"extraData"`
}

func append0xEncode(b []byte) string {
	return "0x" + hex.EncodeToString(b)
}

func (i InputOutputParam712) parseData() string {
	if len(i.Data) > maxHashLen {
		return append0xEncode(i.Data[:maxHashLen]) + "..."
	} else {
		if len(i.Data) == 0 {
			return ""
		}
		return append0xEncode(i.Data[:])
	}
}

/**
计算 inputs 和 outputs
因为 inputs 和 outputs 实际上都是 Cell 的数组，所以这里也就采用了一套基于 Cell 的统一个转换规则。首先是对 Cell 进行分类：
- 不含 type 和 outputs_data 的 Cell 视为普通 Cell ，这种 Cell 就直接略过了，为了尽可能精简 JSON 的内容；
- 含有 type 或者 outputs_data ，但是不属于 DAS 的某种 Cell ，这种 Cell 只会把其 outputs_data 当作字节进行转换，并只保留前 20 字节，至于 type 也是采用通用的转换规则；
- 属于 DAS 的某种 Cell ，capacity, lock, type 都采用统一的转换规则，data 和 extra_data 根据不同的 Cell 采用特有的规则；
然后是逐一对 Cell 的各个字段进行转换：
- capacity 按照 capacity 的语义化规则转换即可；
- lock 的语义化分为以下几步：
  - 将 code_hash 对比 DAS 各个脚本的 type ID ，如果返回 true 则使用对应脚本名；
  - 如果对比全部返回 false 就保留前 20 字节，并转为 hex 以 ... 为结尾；
  - hash_type 转为 hex；
  - args 保留前 20 字节转为 hex ，如果 args 超出 20 字节结尾需要加上 ...；
  - 最后将 code_hash, hash_type, args 三个部分以 , 相连拼接在一起；
- type 采用同 lock 一样的转换方式；
- outputs_data 重命名为 data ，根据不同的业务采用可读的语义化转换。
- 额外增加一个 extraData 字段，用来根据业务存放跟 Cell 强关联的额外数据，完全由业务来定义。

// 这是一个 DAS 中的 AccountCell ，按照上述规则转换后就得到以下 JSON 结构
{
  // 以下三个字段就是对任何 Cell 而言转换规则都完全一样
  "capacity": "999.99 CKB",
  "lock": "das-lock,0x01,0x0000000000000000000000000000000000000011",
  "type": "account-cell-type,0x01,0x",
  // 以下两个字段就是 AccountCell 特有的转换规则，注意字段的内容并不是 JSON 只是为了方便查看组织得比较像 JSON
  "data": "{ account: das00001.bit, expired_at: 1642649600 }",
  "extraData": "{ status: 0, records_hash: 55478d76900611eb079b22088081124ed6c8bae21a05dd1a0d197efcc7c114ce }"
}
*/
func parseScript(script *types.Script, dasCellInfo *celltype.DASCellBaseInfo) string {
	suffix := ""
	if len(script.Args) > maxHashLen {
		suffix = append0xEncode(script.Args[:maxHashLen]) + "..."
	} else {
		suffix = append0xEncode(script.Args[:])
	}
	parseStr := fmt.Sprintf("%s,0x01,%s", dasCellInfo.Name, suffix)
	return parseStr
}
func (list *InputOutputParam712List) AppendAccountCellInput(accountCell *gotype.AccountCell) {
	*list = append(*list, InputOutputParam712{
		Capacity: accountCell.CellCap,
		Lock: &types.Script{
			CodeHash: celltype.DasLockCellScript.Out.CodeHash,
			HashType: types.HashTypeType,
			Args:     accountCell.DasLockArgs,
		},
		Type: &types.Script{
			CodeHash: celltype.DasAccountCellScript.Out.CodeHash,
			HashType: types.HashTypeType,
		},
		Data: accountCell.Data,
	})
}
func (list *InputOutputParam712List) AppendAccountCellOutput(cellCap uint64, data []byte, accountCell *celltype.AccountCell) {
	*list = append(*list, InputOutputParam712{
		Capacity: cellCap,
		Lock:     accountCell.LockScript(),
		Type:     accountCell.TypeScript(),
		Data:     data,
	})
}
func (list InputOutputParam712List) To712Json(accountCellData *celltype.AccountCellData) (string, error) {
	size := len(list)
	retList := make([]inputOutputParam712, 0, size)
	for i := 0; i < size; i++ {
		item := list[i]
		if item.Data == nil && item.Type == nil {
			continue
		}
		retItem := inputOutputParam712{}
		retItem.Capacity = removeSuffixZeroChar(ckbValueStr(item.Capacity)) + " CKB"
		if item.Data != nil || item.Type != nil {
			if item.Type != nil {
				if typeInfo, ok := celltype.SystemCodeScriptMap.Load(item.Type.CodeHash); ok {
					// das-type
					retItem.TypeStr = parseScript(item.Type, typeInfo.(*celltype.DASCellBaseInfo))
					if item.Lock != nil {
						if lockInfo, ok := celltype.SystemCodeScriptMap.Load(item.Lock.CodeHash); ok {
							// das-lock
							retItem.LockStr = parseScript(item.Lock, lockInfo.(*celltype.DASCellBaseInfo))
						}
					}
					if item.Type.CodeHash == celltype.DasAccountCellScript.Out.CodeHash {
						expiredAt, err := celltype.ExpiredAtFromOutputData(item.Data)
						if err != nil {
							return "", fmt.Errorf("ExpiredAtFromOutputData err: %s", err.Error())
						}
						status, err := celltype.MoleculeU8ToGo(accountCellData.Status().RawData())
						if err != nil {
							return "", fmt.Errorf("parse accountCell's status err: %s", err.Error())
						}
						rawAccount := celltype.AccountCharsToAccount(*accountCellData.Account())
						retItem.Data = fmt.Sprintf("{ account: %s, expired_at: %d }", rawAccount, expiredAt)
						recordsHashBytes, err := blake2b.Blake256(accountCellData.Records().AsSlice())
						if err != nil {
							return "", fmt.Errorf("parse accountCell's Records err: %s", err.Error())
						}
						retItem.ExtraData = fmt.Sprintf("{ status: %d, records_hash: %s }", status, types.BytesToHash(recordsHashBytes).String())
					}
				} else {
					// not das-type
					retItem.Data = item.parseData()
				}
			} else {
				// type is nil, but data not empty
				retItem.Data = item.parseData()
			}
			retList = append(retList, retItem)
		}
	}
	jsonBytes, err := json.Marshal(retList)
	if err != nil {
		return "", err
	}
	fmt.Println("jsonBytes:", string(jsonBytes))
	return string(jsonBytes), nil
}

// func (m *MMJson) FillDigest(digest string){
// 	if !strings.HasPrefix(digest,"0x") {
// 		digest = "0x" + digest
// 	}
// 	m.digest = digest
// }

func (m *MMJson) Fill712Action(action string, isOwner bool) {
	param := "0x01"
	if isOwner {
		param = "0x00"
	}
	m.action = fmt.Sprintf(`{"action": "%s","params": "%s"}`, action, param)
}

func (m *MMJson) Fill712Capacity(txBuilder *builder.TransactionBuilder) error {
	fee, inputCap, outputCap, err := txBuilder.InputsOutputsFeeCapacity()
	if err != nil {
		return fmt.Errorf("InputsOutputsFeeCapacity err: %s", err.Error())
	}
	m.fee = fee
	m.inputsCapacity = inputCap
	m.outputsCapacity = outputCap
	return nil
}

var transferAccountDasMessage = "TRANSFER THE ACCOUNT %s TO %s:%s"

func (m *MMJson) FillTransferAccountDasMessage(isTestNet bool, accountCell *gotype.AccountCell, newOwnerParam celltype.DasLockArgsPairParam) {
	//originOwnerIndexType := celltype.DasLockCodeHashIndexType(accountCell.DasLockArgs[0])
	//originOwnerAddrBytes := accountCell.DasLockArgs[1 : celltype.DasLockArgsMinBytesLen/2]
	newOwnerAddrBytes := newOwnerParam.Script.Args[1 : celltype.DasLockArgsMinBytesLen/2]
	account, _ := celltype.AccountFromOutputData(accountCell.Data)
	//originOwnerAddr := gotype.PubkeyHashToAddress(isTestNet, originOwnerIndexType.ChainType(), hex.EncodeToString(originOwnerAddrBytes))
	newOwnerAddr := gotype.PubkeyHashToAddress(isTestNet, newOwnerParam.HashIndexType.ChainType(), hex.EncodeToString(newOwnerAddrBytes))
	m.dasMessage = fmt.Sprintf(
		transferAccountDasMessage,
		account,
		newOwnerParam.HashIndexType.ChainType().String(),
		newOwnerAddr)
}

// Transfer from ckb1xxxx(111.111 CKB), ckb1yyyy(222.222 CKB) to ckb1zzzz(333 CKB), ckb1zzzz(0.333 CKB).
type WithdrawPlainTextOutputParam struct {
	ReceiverCkbScript types.Script
	Amount            uint64
}

func ckbValueStr(cellCap uint64) string {
	first := new(big.Rat).SetInt(new(big.Int).SetUint64(cellCap))
	second := new(big.Rat).SetInt(new(big.Int).SetUint64(celltype.OneCkb))
	return new(big.Rat).Quo(first, second).FloatString(8)
}

func getChainStr(ct celltype.ChainType) string {
	switch ct {
	case celltype.ChainType_ETH:
		return "ETH"
	case celltype.ChainType_CKB:
		return "CKB"
	case celltype.ChainType_TRON:
		return "TRX"
	}
	return "ETH"
}

// now, always withdraw all the money, so there is no change cell
func (m *MMJson) FillWithdrawDasMessage(isTestNet bool, inputs []gotype.WithdrawDasLockCell, output WithdrawPlainTextOutputParam) {
	inputStr := ""
	var mapInputs = make(map[string]gotype.WithdrawDasLockCell)
	for i, v := range inputs {
		args := hex.EncodeToString(v.LockScriptArgs)
		if mi, ok := mapInputs[args]; ok {
			mi.CellCap += v.CellCap
			mapInputs[args] = mi
		} else {
			mapInputs[args] = inputs[i]
		}
	}
	for _, v := range mapInputs {
		item := v
		hashIndex := celltype.DasLockCodeHashIndexType(item.LockScriptArgs[0])
		str := gotype.PubkeyHashToAddress(isTestNet, hashIndex.ChainType(), hex.EncodeToString(item.LockScriptArgs[1:celltype.DasLockArgsMinBytesLen/2]))
		ChainStr := getChainStr(hashIndex.ChainType())
		inputStr = inputStr + fmt.Sprintf("%s:%s(%s CKB), ", ChainStr, str, removeSuffixZeroChar(ckbValueStr(item.CellCap)))
	}
	inputStr = strings.TrimSuffix(inputStr, ", ") + " "

	if output.ReceiverCkbScript.CodeHash == celltype.DasLockCellScript.Out.CodeHash {
		hashIndex := celltype.DasLockCodeHashIndexType(output.ReceiverCkbScript.Args[0])
		str := gotype.PubkeyHashToAddress(isTestNet, hashIndex.ChainType(), hex.EncodeToString(output.ReceiverCkbScript.Args[1:celltype.DasLockArgsMinBytesLen/2]))
		ChainStr := getChainStr(hashIndex.ChainType())
		inputStr = inputStr + fmt.Sprintf("TO %s:%s(%s CKB)", ChainStr, str, removeSuffixZeroChar(ckbValueStr(output.Amount)))
	} else {
		mod := address.Mainnet
		if isTestNet {
			mod = address.Testnet
		}
		receiverAddr, _ := address.Generate(mod, &output.ReceiverCkbScript)
		//receiverAddr := gotype.PubkeyHashToAddress(isTestNet, celltype.ChainType_CKB, hex.EncodeToString(output.ReceiverCkbScript.Args))
		inputStr = inputStr + fmt.Sprintf("TO CKB:%s(%s CKB)", receiverAddr, removeSuffixZeroChar(ckbValueStr(output.Amount)))
	}
	m.dasMessage = fmt.Sprintf("TRANSFER FROM %s", inputStr)
}

func (m *MMJson) FillEditRecordDasMessage(account celltype.DasAccount) {
	m.dasMessage = fmt.Sprintf("EDIT RECORDS OF ACCOUNT %s", account)
}

func (m *MMJson) FillEditManagerDasMessage(account celltype.DasAccount) {
	m.dasMessage = fmt.Sprintf("EDIT MANAGER OF ACCOUNT %s", account)
}

func CreateMMJsonB(txDigestHexStr string) string {
	return MMJsonA
}

func removeSuffixZeroChar(ckbValueStr string) string {
	size := len(ckbValueStr)
	index := 0
	for i := size - 1; i >= 0; i-- {
		if ckbValueStr[i] == '0' {
			index++
		} else {
			break
		}
	}
	last := size - index
	if ckbValueStr[last-1] == '.' {
		last = last - 1
	}
	return ckbValueStr[0:last]
}
