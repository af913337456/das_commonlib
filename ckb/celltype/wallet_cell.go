package celltype

import (
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2021
 * FileName: wallet_cell
 * Author:   LinGuanHong
 * Date:     2021/2/17 12:30 下午
 * Description:
 */

var TestNetWalletCellCell = func() *WalletCellParam {
	return &WalletCellParam{
		CellCodeInfo: DasWalletCellScript,
		AnyoneCanPayScriptInfo: DASCellBaseInfo{
			Dep: DASCellBaseInfoDep{
				TxHash:  types.HexToHash("0xec26b0f85ed839ece5f11c4c4e837ec359f5adc4420410f6453b1f6b60fb96a6"),
				TxIndex: 0,
				DepType: types.DepTypeDepGroup,
			},
			Out: DasAnyOneCanSendCellInfo,
		},
	}
}

/**
lock: <lock_script>
type: <apply_register_script>
data:
  hash(pubkey_hash + account)
  Timestamp // cell 创建时 TimeCell 的时间
*/

type WalletCell struct {
	p *WalletCellParam
}

func NewWalletCell(p *WalletCellParam) *WalletCell {
	return &WalletCell{p: p}
}

func (c *WalletCell) LockDepCell() *types.CellDep {
	return &types.CellDep{
		OutPoint: &types.OutPoint{
			TxHash: c.p.AnyoneCanPayScriptInfo.Dep.TxHash,
			Index:  c.p.AnyoneCanPayScriptInfo.Dep.TxIndex,
		},
		DepType: c.p.AnyoneCanPayScriptInfo.Dep.DepType,
	}
}
func (c *WalletCell) TypeDepCell() *types.CellDep {
	return &types.CellDep{
		OutPoint: &types.OutPoint{
			TxHash: c.p.CellCodeInfo.Dep.TxHash,
			Index:  c.p.CellCodeInfo.Dep.TxIndex,
		},
		DepType: c.p.CellCodeInfo.Dep.DepType,
	}
}
func (c *WalletCell) LockScript() *types.Script {
	return &types.Script{
		CodeHash: c.p.AnyoneCanPayScriptInfo.Out.CodeHash,
		HashType: c.p.AnyoneCanPayScriptInfo.Out.CodeHashType,
		Args:     c.p.AnyoneCanPayScriptInfo.Out.Args,
	}
}
func (c *WalletCell) TypeScript() *types.Script {
	return &types.Script{
		CodeHash: c.p.CellCodeInfo.Out.CodeHash,
		HashType: c.p.CellCodeInfo.Out.CodeHashType,
		Args:     c.p.CellCodeInfo.Out.Args,
	}
}

func (c *WalletCell) TableType() TableType {
	return 0
}

func (c *WalletCell) Data() ([]byte, error) {
	return nil, nil
}

func (c *WalletCell) TableData() []byte {
	return nil
}