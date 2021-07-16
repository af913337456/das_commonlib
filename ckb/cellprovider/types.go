package cellprovider

import (
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"strings"
)

/**
 * Copyright (C), 2019-2021
 * FileName: types
 * Author:   LinGuanHong
 * Date:     2021/2/25 2:40
 * Description:
 */

type EmptyErr error

var IndexerSearchKeyHashLenParam = [2]uint64{32, 33}

const emptyMsg = "cant found on chain"

func NewEmptyErr(cellName string) EmptyErr {
	return EmptyErr(fmt.Errorf("%s %s", cellName, emptyMsg))
}

func IsEmptyErr(err error) bool {
	return err != nil && strings.HasSuffix(err.Error(), emptyMsg)
}

type LiveCellPackObj struct {
	LiveCell    *indexer.LiveCell
	Obj         interface{}
	CellCap     uint64
	WitnessData []byte
}

func (l *LiveCellPackObj) TxHash() types.Hash {
	return l.LiveCell.OutPoint.TxHash
}
func (l *LiveCellPackObj) TxHashStr() string {
	return l.LiveCell.OutPoint.TxHash.String()
}
