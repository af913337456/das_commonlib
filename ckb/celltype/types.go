package celltype

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2020
 * FileName: types
 * Author:   LinGuanHong
 * Date:     2020/12/18 3:58 下午
 * Description:
 */

// [das, type, table]
type DASWitnessDataObj struct {
	Tag       string `json:"tag"`
	TableType uint32 `json:"table_type"`
	TableBys  []byte `json:"table_bys"`
}

/**
- [0:3] 3 个字节固定值为 `0x646173`，这是 `das` 三个字母的 ascii 编码，指明接下来的数据是 DAS 系统数据；
- [4:7] 4 个字节为小端编码的 u32 整形，它是对第 8 字节之后数据类型的标识，具体值详见[Type 常量列表](#Type 常量列表)。首先要通过这个标识判断出具体的数据类型，然后才能用 molecule 编码去解码，下文会解释什么是 molecule 编码；
- [8:] 第 8 字节开始往后的都是 molecule 编码的特殊数据结构，其整体结构如下；
*/
func NewDasWitnessDataFromSlice(rawData []byte) (*DASWitnessDataObj, error) {
	if size := len(rawData); size <= MoleculeBytesHeaderSize+8 { // header'size + min(data)'size
		return nil, fmt.Errorf("invalid rawData size: %d", size)
	}
	tag := string(rawData[MoleculeBytesHeaderSize:7])
	if tag != witnessDas {
		return nil, fmt.Errorf("invalid tag: %s", tag)
	}
	tableType, err := MoleculeU32ToGo(rawData[7:11])
	if err != nil {
		return nil, fmt.Errorf("invalid tableType err: %s", err.Error())
	}
	return &DASWitnessDataObj{
		Tag:       tag,
		TableType: tableType,
		TableBys:  rawData[11:],
	}, nil
}
func NewDasWitnessData(tableType TableType, tableBys []byte) *DASWitnessDataObj {
	return &DASWitnessDataObj{
		Tag:       witnessDas,
		TableType: uint32(tableType),
		TableBys:  tableBys,
	}
}
func (d *DASWitnessDataObj) ToWitness() []byte {
	if d.TableBys == nil {
		return nil
	}
	bytebuf := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytebuf, binary.LittleEndian, d.TableType)
	temp := append([]byte(d.Tag), bytebuf.Bytes()...)
	moBytes := GoBytesToMoleculeBytes(append(temp, d.TableBys...))
	return moBytes.AsSlice()
}

type DASCellBaseInfoDep struct {
	TxHash  string        `json:"tx_hash"`
	TxIndex uint          `json:"tx_index"`
	DepType types.DepType `json:"dep_type"`
}

func (c DASCellBaseInfoDep) ToDepCell() *types.CellDep {
	return &types.CellDep{
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash(c.TxHash),
			Index:  c.TxIndex,
		},
		DepType: c.DepType,
	}
}

type DASCellBaseInfoOut struct {
	CodeHash     string               `json:"code_hash"`
	CodeHashType types.ScriptHashType `json:"code_hash_type"`
	Args         []byte               `json:"args"`
}

func (c DASCellBaseInfoOut) SameScript(script *types.Script) bool {
	current := &types.Script{
		CodeHash: types.HexToHash(c.CodeHash),
		HashType: c.CodeHashType,
		Args:     c.Args,
	}
	return current.Equals(script)
}

type DASCellBaseInfo struct {
	Dep DASCellBaseInfoDep `json:"dep"`
	Out DASCellBaseInfoOut `json:"out"`
}

// type ActionCellParam struct {
// 	Version                   uint32          `json:"version"`
// 	Data                      *ActionCellData `json:"data"` // todo 换成 actionCellData
// 	CellCodeInfo              DASCellBaseInfo `json:"cell_code_info"`
// 	AlwaysSpendableScriptInfo DASCellBaseInfo `json:"always_spendable_script_info"`
// }
//
// type StateCellParam struct {
// 	Version                   uint32          `json:"version"`
// 	Data                      *StateCellData  `json:"data"`
// 	CellCodeInfo              DASCellBaseInfo `json:"cell_code_info"`
// 	AlwaysSpendableScriptInfo DASCellBaseInfo `json:"always_spendable_script_info"`
// }

/**
lock: <use_lock_script>
type: <owner_cell_script>
data:
  [version: u32]
  table RefCellData {
    account_id: Hash, // 由于 AccountCell 可能被单独花费，所以 RefCell 采用 ID 来表达对 AccountCell 的引用关系
}
*/
type RefcellParam struct {
	Version        uint32          `json:"version"`
	Data           string          `json:"data"`
	CellCodeInfo   DASCellBaseInfo `json:"cell_code_info"`
	UserLockScript DASCellBaseInfo `json:"user_lock_script"`
}

// type AccountCommonParam struct {
// 	InstanceId string `json:"instance_id"`
// 	// Quantity   uint64 `json:"quantity"`
// 	// TokenLogic string `json:"token_logic"`
// }
//
// func (a AccountCommonParam) ToBytes() []byte {
// 	retBytes := []byte{}
// 	instanceId := GoHexToMoleculeHash(a.InstanceId)
// 	retBytes = append(retBytes, instanceId.RawData()...)
// 	// quantity := GoUint64ToMoleculeU64(a.Quantity)
// 	// retBytes = append(retBytes, quantity.RawData()...)
// 	// tokenLogic := GoHexToMoleculeHash(a.TokenLogic)
// 	// retBytes = append(retBytes, tokenLogic.RawData()...)
// 	return retBytes
// }

// func AccountCommonParamByteLen() int {
// 	return 32 + CellVersionByteLen
// }

type ProposeCellParam struct {
	// AccountCommonParam
	Version                   uint32          `json:"version"`
	Data                      Data            `json:"data"`
	CellCodeInfo              DASCellBaseInfo `json:"cell_code_info"`
	AlwaysSpendableScriptInfo DASCellBaseInfo `json:"always_spendable_script_info"`
}

type AccountCellParam struct {
	// AccountCommonParam
	Version                   uint32          `json:"version"`
	Data                      Data            `json:"data"`
	CellCodeInfo              DASCellBaseInfo `json:"cell_code_info"`
	AlwaysSpendableScriptInfo DASCellBaseInfo `json:"always_spendable_script_info"`
}

type ParseDasWitnessBysDataObj struct {
	WitnessObj            *DASWitnessDataObj
	MoleculeData          *Data
	MoleculeOldDataEntity *DataEntity
	MoleculeNewDataEntity *DataEntity
}

type ProposeWitnessSliceDataObject struct {
	AccountId string                     `json:"account_id"`
	ItemType  ProposeWitnessDataItemType `json:"item_type"`
	Next      string                     `json:"next"`
}
