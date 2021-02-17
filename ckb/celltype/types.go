package celltype

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
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

type TypeInputCell struct {
	Input    types.CellInput `json:"input"`
	LockType LockScriptType  `json:"lock_type"`
}

type BuildTransactionRet struct {
	Group      []int              `json:"group"`
	WitnessArg *types.WitnessArgs `json:"witness_arg"`
}

type InputWithWitness struct {
	CellCap        uint64                 `json:"cell_cap"`
	CellInput      TypeInputCell          `json:"cell_input"`
	GetWitnessData CellDepWithWitnessFunc `json:"-"`
}

type CellDepWithWitnessFunc func(index uint32) ([]byte, error)

type CellDepWithWitness struct {
	CellDep        *types.CellDep
	GetWitnessData CellDepWithWitnessFunc
}

// [das, type, table]
type DASWitnessDataObj struct {
	Tag       string    `json:"tag"`
	TableType TableType `json:"table_type"`
	TableBys  []byte    `json:"table_bys"`
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
		TableType: TableType(tableType),
		TableBys:  rawData[11:],
	}, nil
}
func NewDasWitnessData(tableType TableType, tableBys []byte) *DASWitnessDataObj {
	return &DASWitnessDataObj{
		Tag:       witnessDas,
		TableType: tableType,
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
	TxHash  types.Hash    `json:"tx_hash"`
	TxIndex uint          `json:"tx_index"`
	DepType types.DepType `json:"dep_type"`
}

func (c DASCellBaseInfoDep) ToDepCell() *types.CellDep {
	return &types.CellDep{
		OutPoint: &types.OutPoint{
			TxHash: c.TxHash,
			Index:  c.TxIndex,
		},
		DepType: c.DepType,
	}
}

type DASCellBaseInfoOut struct {
	CodeHash     types.Hash           `json:"code_hash"`
	CodeHashType types.ScriptHashType `json:"code_hash_type"`
	Args         []byte               `json:"args"`
}

func DASCellBaseInfoOutFromScript(script *types.Script) DASCellBaseInfoOut {
	return DASCellBaseInfoOut{
		CodeHash:     script.CodeHash,
		CodeHashType: script.HashType,
		Args:         script.Args,
	}
}

func (c DASCellBaseInfoOut) SameScript(script *types.Script) bool {
	current := &types.Script{
		CodeHash: c.CodeHash,
		HashType: c.CodeHashType,
		Args:     c.Args,
	}
	return current.Equals(script)
}

type DASCellBaseInfo struct {
	Dep DASCellBaseInfoDep `json:"dep"`
	Out DASCellBaseInfoOut `json:"out"`
}

type ApplyRegisterCellParam struct {
	Version              uint32          `json:"version"`
	PubkeyHash           string          `json:"pubkey_hash"`
	Account              string          `json:"account"`
	Timestamp            uint64          `json:"timestamp"`
	CellCodeInfo         DASCellBaseInfo `json:"cell_code_info"`
	SenderLockScriptInfo DASCellBaseInfo `json:"sender_lock_script_info"`
}

type PreAccountCellDatas struct {
	DepAccountCellData *PreAccountCellData `json:"-"`
	OldAccountCellData *PreAccountCellData `json:"-"`
	NewAccountCellData *PreAccountCellData `json:"-"`
}
type PreAccountCellCellParam struct {
	Version                   uint32              `json:"version"`
	Data                      Data                `json:"data"`
	PreAccountCellDatas       PreAccountCellDatas `json:"-"`
	CellCodeInfo              DASCellBaseInfo     `json:"cell_code_info"`
	AlwaysSpendableScriptInfo DASCellBaseInfo     `json:"always_spendable_script_info"`
}

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

type AccountCellFullData struct {
	NextAccountId []byte          `json:"next_account_id"`
	RegisteredAt  uint64          `json:"registered_at"`
	ExpiredAt     uint64          `json:"expired_at"`
	AccountInfo   AccountCellData `json:"-"`
}

type AccountCellDatas struct {
	DepAccountCellData *AccountCellData     `json:"-"`
	OldAccountCellData *AccountCellData     `json:"-"`
	NewAccountCellData *AccountCellFullData `json:"-"`
}
type AccountCellParam struct {
	AccountCellDatas          AccountCellDatas `json:"-"`
	Version                   uint32           `json:"version"`
	Data                      Data             `json:"data"`
	CellCodeInfo              DASCellBaseInfo  `json:"cell_code_info"`
	AlwaysSpendableScriptInfo DASCellBaseInfo  `json:"always_spendable_script_info"`
}

type ParseDasWitnessBysDataObj struct {
	WitnessObj            *DASWitnessDataObj
	MoleculeData          *Data
	MoleculeDepDataEntity *DataEntity
	MoleculeOldDataEntity *DataEntity
	MoleculeNewDataEntity *DataEntity
}

type ProposeWitnessSliceDataObject struct {
	AccountId []byte            `json:"account_id"`
	ItemType  AccountCellStatus `json:"item_type"`
	Next      []byte            `json:"next"`
}

func (p ProposeWitnessSliceDataObject) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(
		`{"account_id":"%s","item_type":"%s","next":"%s"}`,
		hex.EncodeToString(p.AccountId),
		p.ItemType.Str(), hex.EncodeToString(p.Next))), nil
}

type ProposeWitnessSliceDataObjectList []ProposeWitnessSliceDataObject

func (p *ProposeWitnessSliceDataObjectList) Add(accountId, nextId []byte, status AccountCellStatus) {
	*p = append(*p, ProposeWitnessSliceDataObject{AccountId: accountId, Next: nextId, ItemType: status})
}

func ProposeWitnessSliceDataObjectListFromBytes(bys []byte) ([]ProposeWitnessSliceDataObjectList, error) {
	proposeCellData, err := ProposalCellDataFromSlice(bys, false)
	if err != nil {
		return nil, err
	}
	retList := []ProposeWitnessSliceDataObjectList{}
	sliceList := proposeCellData.Slices()
	index := uint(0)
	for sl := sliceList.Get(index); sl != nil && !sl.IsEmpty(); index++ {
		proposeItemIndex := uint(0)
		list := []ProposeWitnessSliceDataObject{}
		for propose := sl.Get(proposeItemIndex); propose != nil && !propose.IsEmpty(); proposeItemIndex++ {
			itemTypeUint8, err := MoleculeU8ToGo(propose.ItemType().inner)
			if err != nil {
				return nil, err
			}
			list = append(list, ProposeWitnessSliceDataObject{
				AccountId: propose.AccountId().AsSlice(),
				ItemType:  AccountCellStatus(itemTypeUint8),
				Next:      propose.Next().AsSlice(),
			})
		}
		retList = append(retList, list)
	}
	return retList, nil
}

type ProposeWitnessSliceDataObjectLL []ProposeWitnessSliceDataObjectList

func (p ProposeWitnessSliceDataObjectLL) ToMoleculeProposalCellData(incomeLockScript *types.Script, proposerWalletId []byte) ProposalCellData {
	sliceList := make([]SL, 0, len(p))
	for _, slice := range p {
		proposeItemList := make([]ProposalItem, 0, len(slice))
		for _, item := range slice {
			accountId := NewAccountIdBuilder().Set(GoBytesToMoleculeAccountBytes(item.AccountId)).Build()
			nextAccountId := NewAccountIdBuilder().Set(GoBytesToMoleculeAccountBytes(item.Next)).Build()
			proposeItem := NewProposalItemBuilder().
				AccountId(accountId).
				Next(nextAccountId).
				ItemType(GoUint8ToMoleculeU8(uint8(item.ItemType))).
				Build()
			proposeItemList = append(proposeItemList, proposeItem)
		}
		sliceList = append(sliceList, NewSLBuilder().Set(proposeItemList).Build())
	}
	proposalCellData := NewProposalCellDataBuilder().
		ProposerLock(GoCkbScriptToMoleculeScript(*incomeLockScript)).
		ProposerWallet(GoBytesToMoleculeHash(proposerWalletId)).
		Slices(NewSliceListBuilder().Set(sliceList).Build()).
		Build()
	return proposalCellData
}

type CalAccountCellExpiredAtParam struct {
	Quote          uint64 `json:"quote"`
	AccountCellCap uint64 `json:"account_cell_cap"`
	PriceConfigNew uint64 `json:"price_config_new"`
	// AccountBytesLen    uint32 `json:"account_bytes_len"`
	PreAccountCellCap uint64 `json:"pre_account_cell_cap"`
}
