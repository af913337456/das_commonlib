package celltype

import (
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2021
 * FileName: on_sale_cell
 * Author:   LinGuanHong
 * Date:     2021/2/22 11:05
 * Description:
 */

var DefaultAccountSaleCellParam = func(startedAt,price uint64,description string, accountId DasAccountId,dasLockParam *DasLockParam) *AccountSaleCellParam {
	mAccountId := NewAccountIdBuilder().Set(GoBytesToMoleculeAccountBytes(accountId.Bytes())).Build()
	mDescription := GoStrToMoleculeBytes(description)
	return &AccountSaleCellParam{
		Version:        2,
		Price:          price,
		SaleCellData:   NewAccountSaleCellDataBuilder().
			AccountId(mAccountId).
			Description(mDescription).
			StartedAt(GoUint64ToMoleculeU64(startedAt)).
			Price(GoUint64ToMoleculeU64(price)).Build(),
		CellCodeInfo: DasAccountSaleCellScript,
		DasLock: DasLockCellScript,
	}
}

type AccountSaleCell struct {
	p *AccountSaleCellParam
}

func NewAccountSaleCell(p *AccountSaleCellParam) *AccountSaleCell {
	return &AccountSaleCell{p: p}
}

func (c *AccountSaleCell) SoDeps() []types.CellDep {
	return []types.CellDep{
		*ETHSoScriptDep.ToDepCell(),
		*CKBSoScriptDep.ToDepCell(),
		*TRONSoScriptDep.ToDepCell(),
	}
}

func (c *AccountSaleCell) LockDepCell() *types.CellDep {
	return &types.CellDep{
		OutPoint: &types.OutPoint{
			TxHash: c.p.DasLock.Dep.TxHash,
			Index:  c.p.DasLock.Dep.TxIndex,
		},
		DepType: c.p.DasLock.Dep.DepType,
	}
}
func (c *AccountSaleCell) TypeDepCell() *types.CellDep {
	return &types.CellDep{
		OutPoint: &types.OutPoint{
			TxHash: c.p.CellCodeInfo.Dep.TxHash,
			Index:  c.p.CellCodeInfo.Dep.TxIndex,
		},
		DepType: c.p.CellCodeInfo.Dep.DepType,
	}
}
func (c *AccountSaleCell) LockScript() *types.Script {
	lockScript := &types.Script{
		CodeHash: c.p.DasLock.Out.CodeHash,
		HashType: c.p.DasLock.Out.CodeHashType,
	}
	if c.p.DasLockParam != nil {
		lockScript.Args = c.p.DasLockParam.Bytes()
	}
	return lockScript
}
func (c *AccountSaleCell) TypeScript() *types.Script {
	return &types.Script{
		CodeHash: c.p.CellCodeInfo.Out.CodeHash,
		HashType: c.p.CellCodeInfo.Out.CodeHashType,
		Args:     nil,
	}
}

func (c *AccountSaleCell) Data() ([]byte, error) {
	bys, err := blake2b.Blake256(c.p.SaleCellData.AsSlice())
	if err != nil {
		return nil, err
	}
	return bys, nil
}

func (c *AccountSaleCell) TableType() TableType {
	return TableType_AccountSaleCell
}
