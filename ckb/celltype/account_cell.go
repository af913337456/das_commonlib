package celltype

import (
	"fmt"
	"github.com/DA-Services/das_commonlib/common"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2020
 * FileName: publishaccountcell
 * Author:   LinGuanHong
 * Date:     2020/12/25 5:51 下午
 * Description:
 */

/**
table DataEntity {
    index: Uint32, // 表明此数据项属于 inputs/outputs 中的第几个 cell
    version: Uint32, // 表明 entity 数据结构的版本号
    entity: Bytes, // 代表具体的数据结构
}
*/
var TestNetAccountCell = func(depIndex, oldIndex, newIndex uint32, dep, old *AccountCellData, new *AccountCellFullData) *AccountCellParam {
	acp := &AccountCellParam{
		Version: 1,
		Data: *buildDasCommonMoleculeDataObj(
			depIndex, oldIndex, newIndex, dep, old, &new.AccountInfo),
		CellCodeInfo: DasAccountCellScript,
		AccountCellDatas: AccountCellDatas{
			DepAccountCellData: dep,
			OldAccountCellData: old,
			NewAccountCellData: new,
		},
		AlwaysSpendableScriptInfo: DasAnyOneCanSendCellInfo,
	}
	return acp
}

/**
lock: <always_success>
type:
  code_hash: <nft_script>
  type: type
  args: []
data:
  hash(data: AccountCellData)
  id // 自己的 ID，生成算法为 hash(account)，然后取前 10 bytes
  next // 下一个 AccountCell 的 ID
  expired_at // 小端编码的 u64 时间戳
  account // AccountCell 为了避免数据丢失导致用户无法找回自己用户所以额外储存了 account 的明文信息，不含 .bit

witness:
  table Data {
    old: table DataEntityOpt {
    	index: Uint32,
    	version: Uint32,
    	entity: AccountCellData
    },
    new: table DataEntityOpt {
      index: Uint32,
      version: Uint32,
      entity: AccountCellData
    },
  }

======
table AccountCellData {
    // The first 160 bits of the hash of account.
    id: AccountId,
    // The lock script of owner.
    owner: Script,
    // The lock script of manager.
    manager: Script,
    account: Bytes,
    // The status of the account, 0 means normal, 1 means being sold, 2 means being auctioned.
    status: Uint8,
    records: Records,
}

array AccountId [byte; 20];

option AccountIdOpt (AccountId);

table Record {
    record_type: Bytes,
    record_label: Bytes,
    record_value: Bytes,
    record_ttl: Uint32,
}

vector Records <Record>;
*/

type AccountCell struct {
	p *AccountCellParam
}

func NewAccountCell(p *AccountCellParam) *AccountCell {
	return &AccountCell{p: p}
}

func (c *AccountCell) LockDepCell() *types.CellDep {
	return &types.CellDep{
		OutPoint: &types.OutPoint{
			TxHash: c.p.AlwaysSpendableScriptInfo.Dep.TxHash,
			Index:  c.p.AlwaysSpendableScriptInfo.Dep.TxIndex,
		},
		DepType: c.p.AlwaysSpendableScriptInfo.Dep.DepType,
	}
}
func (c *AccountCell) TypeDepCell() *types.CellDep {
	return &types.CellDep{ // state_cell
		OutPoint: &types.OutPoint{
			TxHash: c.p.CellCodeInfo.Dep.TxHash,
			Index:  c.p.CellCodeInfo.Dep.TxIndex, // state_script_tx_index
		},
		DepType: c.p.CellCodeInfo.Dep.DepType,
	}
}
func (c *AccountCell) LockScript() *types.Script {
	return &types.Script{
		CodeHash: c.p.AlwaysSpendableScriptInfo.Out.CodeHash,
		HashType: c.p.AlwaysSpendableScriptInfo.Out.CodeHashType,
		Args:     c.p.AlwaysSpendableScriptInfo.Out.Args,
	}
}
func (c *AccountCell) TypeScript() *types.Script {
	return &types.Script{
		CodeHash: c.p.CellCodeInfo.Out.CodeHash,
		HashType: c.p.CellCodeInfo.Out.CodeHashType,
		Args:     nil,
	}
}

/**
  hash(data: AccountCellData)
  id // 自己的 ID，生成算法为 hash(account)，然后取前 10 bytes
  next // 下一个 AccountCell 的 ID
  expired_at // 小端编码的 u64 时间戳
  account // AccountCell 为了避免数据丢失导致用户无法找回自己用户所以额外储存了 account 的明文信息，不含 .bit

*/

func AccountIdFromOutputData(data []byte) (DasAccountId, error) {
	if size := len(data); size < HashBytesLen+dasAccountIdLen {
		return EmptyAccountId, fmt.Errorf("AccountIdFromOutputData invalid data, len not enough: %d", size)
	}
	return DasAccountIdFromBytes(data[HashBytesLen : HashBytesLen+dasAccountIdLen]), nil
}

func NextAccountIdFromOutputData(data []byte) (DasAccountId, error) {
	minLen := dasAccountIdLen + HashBytesLen
	if size := len(data); size < minLen {
		return EmptyAccountId, fmt.Errorf("invalid data, len not enough: %d", size)
	}
	return DasAccountIdFromBytes(data[HashBytesLen:minLen]), nil
}

func ExpiredAtFromOutputData(data []byte) (int64, error) {
	endLen := HashBytesLen + dasAccountIdLen*2 + 8
	if size := len(data); size < endLen {
		return 0, fmt.Errorf("ExpiredAtFromOutputData invalid data, len not enough: %d", size)
	}
	return common.BytesToInt64(data[endLen-8 : endLen]), nil
}

func IsAccountExpired(accountCellData []byte, cmpTimeSec int64) (bool, error) {
	expired, err := ExpiredAtFromOutputData(accountCellData)
	if err != nil {
		return false, err
	}
	return cmpTimeSec <= expired, nil
}

func IsAccountFrozen(accountCellData []byte, cmpTimeSec, frozenRangeSec int64) (bool, error) {
	expired, err := ExpiredAtFromOutputData(accountCellData)
	if err != nil {
		return false, err
	}
	return expired >= cmpTimeSec && expired < cmpTimeSec+frozenRangeSec, nil
}

func SetAccountCellNextAccountId(data []byte, accountId DasAccountId) []byte {
	minLen := HashBytesLen
	accountIdEndLen := HashBytesLen + dasAccountIdLen
	if size := len(data); size < minLen {
		data = append(data, EmptyDataHash[:]...)
		data = append(data, EmptyAccountId.Bytes()...)
	} else if size < accountIdEndLen {
		data = append(data, EmptyAccountId.Bytes()...)
	}
	dataLen := len(data)
	temp1 := make([]byte, 0, HashBytesLen)
	temp2 := make([]byte, 0, dataLen-minLen)
	prefix := append(temp1, data[:HashBytesLen]...)
	suffix := append(temp2, data[accountIdEndLen:]...)
	return append(append(prefix, accountId.Bytes()...), suffix...)
}

func DefaultAccountCellDataBytes(accountId, nextAccountId DasAccountId) []byte {
	holder := EmptyDataHash
	return append(append(holder, accountId.Bytes()...), nextAccountId.Bytes()...)
}

func accountCellOutputData(newData *AccountCellFullData) ([]byte, error) {
	dataBytes := []byte{}
	accountInfoDataBytes, _ := blake2b.Blake256(newData.AccountInfo.AsSlice())
	dataBytes = append(dataBytes, accountInfoDataBytes...)
	accountBytes := newData.AccountInfo.Account().AsSlice()
	accountIdBytes, _ := blake2b.Blake160(accountBytes)
	dataBytes = append(dataBytes, DasAccountIdFromBytes(accountIdBytes).Bytes()...) // id
	dataBytes = append(dataBytes, newData.NextAccountId.Bytes()...)                 // next
	dataBytes = append(dataBytes, GoUint64ToBytes(newData.ExpiredAt)...)            // expired_at
	dataBytes = append(dataBytes, accountBytes...)                                  // account
	return dataBytes, nil
}

func AccountCellCap(account string) (uint64, error) {
	output := types.CellOutput{
		Lock: &types.Script{
			CodeHash: DasAnyOneCanSendCellInfo.Out.CodeHash,
			HashType: DasAnyOneCanSendCellInfo.Out.CodeHashType,
			Args:     DasAnyOneCanSendCellInfo.Out.Args,
		},
		Type: &types.Script{
			CodeHash: DasAccountCellScript.Out.CodeHash,
			HashType: DasAccountCellScript.Out.CodeHashType,
			Args:     DasAccountCellScript.Out.Args,
		},
	}
	dataBytes := []byte{}
	expiredAtBytes := GoUint64ToBytes(0)
	accountBytes := []byte(account)

	dataBytes = append(dataBytes, EmptyDataHash...)
	dataBytes = append(dataBytes, EmptyAccountId.Bytes()...)
	dataBytes = append(dataBytes, EmptyAccountId.Bytes()...)
	dataBytes = append(dataBytes, expiredAtBytes...)
	dataBytes = append(dataBytes, accountBytes...)

	return output.OccupiedCapacity(dataBytes), nil
}

func (c *AccountCell) Data() ([]byte, error) {
	return accountCellOutputData(c.p.AccountCellDatas.NewAccountCellData)
}

func (c *AccountCell) TableType() TableType {
	return TableType_ACCOUNT_CELL
}

func (c *AccountCell) TableData() []byte {
	return c.p.Data.AsSlice()
}
