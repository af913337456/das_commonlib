package celltype

import (
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2020
 * FileName: refcell
 * Author:   LinGuanHong
 * Date:     2020/12/27 11:17 上午
 * Description:
 */

var TestNetRefCell = func(lockScript *types.Script) *RefcellParam {
	return &RefcellParam{
		Version:      1,
		Data:         "",
		CellCodeInfo: TestNet_OwnerCellScript,
		UserLockScript: DASCellBaseInfo{
			Dep: DASCellBaseInfoDep{
				// secp256h1
				TxHash:  "0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37",
				TxIndex: 0,
				DepType: types.DepTypeDepGroup,
			},
			Out: DASCellBaseInfoOut{
				CodeHash:     lockScript.CodeHash.String(),
				CodeHashType: lockScript.HashType,
				Args:         lockScript.Args,
			},
		},
	}
}

type Refcell struct {
	p *RefcellParam
}

func NewRefcell(p *RefcellParam) *Refcell {
	return &Refcell{p: p}
}

func (c *Refcell) LockDepCell() *types.CellDep {
	return &types.CellDep{
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash(c.p.UserLockScript.Dep.TxHash),
			Index:  c.p.UserLockScript.Dep.TxIndex,
		},
		DepType: c.p.UserLockScript.Dep.DepType,
	}
}
func (c *Refcell) TypeDepCell() *types.CellDep {
	return &types.CellDep{ // state_cell
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash(c.p.CellCodeInfo.Dep.TxHash),
			Index:  c.p.CellCodeInfo.Dep.TxIndex, // state_script_tx_index
		},
		DepType: c.p.CellCodeInfo.Dep.DepType,
	}
}
func (c *Refcell) LockScript() *types.Script {
	return &types.Script{
		CodeHash: types.HexToHash(c.p.UserLockScript.Out.CodeHash),
		HashType: c.p.UserLockScript.Out.CodeHashType,
		Args:     c.p.UserLockScript.Out.Args,
	}
}
func (c *Refcell) TypeScript() *types.Script {
	return &types.Script{
		CodeHash: types.HexToHash(c.p.CellCodeInfo.Out.CodeHash),
		HashType: c.p.CellCodeInfo.Out.CodeHashType,
		Args:     c.p.CellCodeInfo.Out.Args,
	}
}

func (c *Refcell) Data() ([]byte, error) {
	return nil, nil
}

func (c *Refcell) TableType() TableType {
	return 0
}

func (c *Refcell) TableData() []byte {
	return nil
}
