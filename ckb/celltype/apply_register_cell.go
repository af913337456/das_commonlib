package celltype

import (
	"encoding/hex"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2020
 * FileName: statecell
 * Author:   LinGuanHong
 * Date:     2020/12/18 3:58 下午
 * Description:
 */

var TestNetApplyRegisterCell = func(pubkey []byte, account string, height uint64, senderLockScript *types.Script) *ApplyRegisterCellParam {
	pubkeyHash, _ := blake2b.Blake160(pubkey)
	return &ApplyRegisterCellParam{
		Version:      1,
		PubkeyHash:   hex.EncodeToString(pubkeyHash),
		Account:      account,
		Height:       height,
		CellCodeInfo: DasApplyRegisterCellScript,
		SenderLockScriptInfo: DASCellBaseInfo{
			Dep: DASCellBaseInfoDep{
				TxHash:  types.HexToHash("0xec26b0f85ed839ece5f11c4c4e837ec359f5adc4420410f6453b1f6b60fb96a6"),
				TxIndex: 0,
				DepType: types.DepTypeDepGroup,
			},
			Out: DASCellBaseInfoOutFromScript(senderLockScript),
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

type ApplyRegisterCell struct {
	p *ApplyRegisterCellParam
}

func NewApplyRegisterCell(p *ApplyRegisterCellParam) *ApplyRegisterCell {
	return &ApplyRegisterCell{p: p}
}

func (c *ApplyRegisterCell) LockDepCell() *types.CellDep {
	return &types.CellDep{
		OutPoint: &types.OutPoint{
			TxHash: c.p.SenderLockScriptInfo.Dep.TxHash,
			Index:  c.p.SenderLockScriptInfo.Dep.TxIndex,
		},
		DepType: c.p.SenderLockScriptInfo.Dep.DepType,
	}
}
func (c *ApplyRegisterCell) TypeDepCell() *types.CellDep {
	return &types.CellDep{
		OutPoint: &types.OutPoint{
			TxHash: c.p.CellCodeInfo.Dep.TxHash,
			Index:  c.p.CellCodeInfo.Dep.TxIndex,
		},
		DepType: c.p.CellCodeInfo.Dep.DepType,
	}
}
func (c *ApplyRegisterCell) LockScript() *types.Script {
	return &types.Script{
		CodeHash: c.p.SenderLockScriptInfo.Out.CodeHash,
		HashType: c.p.SenderLockScriptInfo.Out.CodeHashType,
		Args:     c.p.SenderLockScriptInfo.Out.Args,
	}
}
func (c *ApplyRegisterCell) TypeScript() *types.Script {
	return &types.Script{
		CodeHash: c.p.CellCodeInfo.Out.CodeHash,
		HashType: c.p.CellCodeInfo.Out.CodeHashType,
		Args:     nil,
	}
}

func (c *ApplyRegisterCell) TableType() TableType {
	return 0
}

func (c *ApplyRegisterCell) Data() ([]byte, error) {
	idHash, err := ApplyRegisterDataId(c.p.PubkeyHash, c.p.Account)
	if err != nil {
		return nil, fmt.Errorf("ApplyRegisterDataId err: %s", err.Error())
	}
	return append(idHash, GoUint64ToBytes(c.p.Height)...), nil
}

func (c *ApplyRegisterCell) TableData() []byte {
	return nil
}

func ApplyRegisterDataId(pubKeyHex, account string) ([]byte, error) {
	return blake2b.Blake256([]byte(pubKeyHex + account))
}
