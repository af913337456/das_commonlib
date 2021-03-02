package celltype

import (
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2021
 * FileName: on_sale_cell
 * Author:   LinGuanHong
 * Date:     2021/2/22 11:05 上午
 * Description:
 */

var TestNetOnSaleCell = func(newIndex uint32, price uint64, accountId DasAccountId) *OnSaleCellParam {
	onSaleMoleData := NewOnSaleCellDataBuilder().Price(GoUint64ToMoleculeU64(price)).Build()
	return &OnSaleCellParam{
		Version: 1,
		Price:   price,
		Data: *buildDasCommonMoleculeDataObj(
			0, 0, newIndex, nil, nil, &onSaleMoleData),
		AccountId:    accountId,
		CellCodeInfo: DasOnSaleCellScript,
		AlwaysSpendableScriptInfo: DASCellBaseInfo{
			Dep: DASCellBaseInfoDep{
				TxHash:  types.HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"),
				TxIndex: 0,
				DepType: types.DepTypeDepGroup,
			},
			Out: DasAnyOneCanSendCellInfo,
		},
	}
}

type OnSaleCell struct {
	p *OnSaleCellParam
}

func NewOnSaleCell(p *OnSaleCellParam) *OnSaleCell {
	return &OnSaleCell{p: p}
}

func (c *OnSaleCell) LockDepCell() *types.CellDep {
	return &types.CellDep{
		OutPoint: &types.OutPoint{
			TxHash: c.p.AlwaysSpendableScriptInfo.Dep.TxHash,
			Index:  c.p.AlwaysSpendableScriptInfo.Dep.TxIndex,
		},
		DepType: c.p.AlwaysSpendableScriptInfo.Dep.DepType,
	}
}
func (c *OnSaleCell) TypeDepCell() *types.CellDep {
	return &types.CellDep{
		OutPoint: &types.OutPoint{
			TxHash: c.p.CellCodeInfo.Dep.TxHash,
			Index:  c.p.CellCodeInfo.Dep.TxIndex,
		},
		DepType: c.p.CellCodeInfo.Dep.DepType,
	}
}
func (c *OnSaleCell) LockScript() *types.Script {
	return &types.Script{
		CodeHash: c.p.AlwaysSpendableScriptInfo.Out.CodeHash,
		HashType: c.p.AlwaysSpendableScriptInfo.Out.CodeHashType,
		Args:     c.p.AlwaysSpendableScriptInfo.Out.Args,
	}
}
func (c *OnSaleCell) TypeScript() *types.Script {
	return &types.Script{
		CodeHash: c.p.CellCodeInfo.Out.CodeHash,
		HashType: c.p.CellCodeInfo.Out.CodeHashType,
		Args:     c.p.AccountId.Bytes(),
	}
}

func (c *OnSaleCell) Data() ([]byte, error) {
	bys, err := blake2b.Blake256(c.TableData())
	if err != nil {
		return nil, err
	}
	return bys, nil
}

func (c *OnSaleCell) TableType() TableType {
	return TableType_ON_SALE_CELL
}

func (c *OnSaleCell) TableData() []byte {
	return c.p.Data.AsSlice()
}
