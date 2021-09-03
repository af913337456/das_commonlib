package configcells

import (
	"fmt"
	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
)

/**
 * Copyright (C), 2019-2021
 * FileName: main
 * Author:   LinGuanHong
 * Date:     2021/5/17 2:51
 * Description:
 */

type CfgUnavailable struct {
	Data *ConfigCellChildDataObj
	MocluData *celltype.ConfigCellAccount
}

func (c *CfgUnavailable) Ready() bool{
	return c.Data != nil && c.MocluData != nil && c.MocluData.FieldCount() > 0
}

func (c *CfgUnavailable) Name() string {
	return "configCellAccount:"
}

func (c *CfgUnavailable) NotifyData(Data *ConfigCellChildDataObj) error {
	c.Data = Data
	if len(c.Data.MoleculeData) == 0 {
		temp := celltype.ConfigCellAccountDefault()
		c.MocluData = &temp
		return nil
	}
	obj, err := celltype.ConfigCellAccountFromSlice(c.Data.MoleculeData, false)
	if err != nil {
		return fmt.Errorf("ConfigCellAccountFromSlice %s",err.Error())
	}
	c.MocluData = obj
	return nil
}

func (c *CfgUnavailable) MocluObj() interface{} {
	return c.MocluData
}

func (c *CfgUnavailable) Tag() celltype.TableType {
	return celltype.TableType_ConfigCell_Account
}

func (c *CfgUnavailable) Witness() *celltype.CellDepWithWitness {
	return &celltype.CellDepWithWitness{
		CellDep: &c.Data.CellDep,
		GetWitnessData: func(index uint32) ([]byte, error) {
			return c.Data.WitnessData, nil
		}}
}