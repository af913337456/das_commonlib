package celltype

import (
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
var TestNetAccountCell = func(depIndex, oldIndex, newIndex uint32, old, new *AccountCellData) *AccountCellParam {
	acp := &AccountCellParam{
		Version:      1,
		Data:         *buildDasCommonMoleculeDataObj(depIndex, oldIndex, newIndex, nil, old, new),
		CellCodeInfo: TestNet_AccountCellScript,
		AlwaysSpendableScriptInfo: DASCellBaseInfo{
			Dep: DASCellBaseInfoDep{
				TxHash:  "0xec26b0f85ed839ece5f11c4c4e837ec359f5adc4420410f6453b1f6b60fb96a6",
				TxIndex: 0,
				DepType: types.DepTypeDepGroup,
			},
			Out: DASCellBaseInfoOut{
				CodeHash:     "0x3419a1c09eb2567f6552ee7a8ecffd64155cffe0f1796e6e61ec088d740c1356",
				CodeHashType: types.HashTypeType,
				Args:         nil,
			},
		},
	}
	return acp
}

/**
lock: <always_spendable_script>
type:
  code_hash: <nft_script>
  type: type
  args: [id] // 自己的 ID，生成算法为 hash(account)
data:
	hash(data: AccountCellData)
	account // AccountCell 为了避免数据丢失导致用户无法找回自己用户所以额外储存了 account 的明文信息

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
			TxHash: types.HexToHash(c.p.AlwaysSpendableScriptInfo.Dep.TxHash),
			Index:  c.p.AlwaysSpendableScriptInfo.Dep.TxIndex,
		},
		DepType: c.p.AlwaysSpendableScriptInfo.Dep.DepType,
	}
}
func (c *AccountCell) TypeDepCell() *types.CellDep {
	return &types.CellDep{ // state_cell
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash(c.p.CellCodeInfo.Dep.TxHash),
			Index:  c.p.CellCodeInfo.Dep.TxIndex, // state_script_tx_index
		},
		DepType: c.p.CellCodeInfo.Dep.DepType,
	}
}
func (c *AccountCell) LockScript() *types.Script {
	return &types.Script{
		CodeHash: types.HexToHash(c.p.AlwaysSpendableScriptInfo.Out.CodeHash),
		HashType: c.p.AlwaysSpendableScriptInfo.Out.CodeHashType,
		Args:     c.p.AlwaysSpendableScriptInfo.Out.Args,
	}
}
func (c *AccountCell) TypeScript() *types.Script {
	return &types.Script{
		CodeHash: types.HexToHash(c.p.CellCodeInfo.Out.CodeHash),
		HashType: c.p.CellCodeInfo.Out.CodeHashType,
		Args:     c.p.CellCodeInfo.Out.Args,
	}
}

/**
  table AccountCellData {
	  owner_cell: Hash,
	  manager_cell: Hash,
	  account: Bytes,
	  registered_at: Timestamp,
	  expired_at: Timestamp,
	  records: Records,
	}
*/

func (c *AccountCell) Data() ([]byte, error) {
	tableBys := c.TableData()
	hashBys, err := blake2b.Blake256(tableBys)
	if err != nil {
		return nil, err
	}
	return append(hashBys, tableBys...), nil
}

func (c *AccountCell) TableType() TableType {
	return TableType_ACCOUNT_CELL
}

func (c *AccountCell) TableData() []byte {
	return c.p.Data.AsSlice()
}
