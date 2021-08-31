package configcells

import (
	"fmt"
	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
)

/**
 * Copyright (C), 2019-2021
 * FileName: main
 * Author:   LinGuanHong
 * Date:     2021/8/31 9:58
 * Description:
 */

type CfgSecondaryMarket struct {
	Data *ConfigCellChildDataObj
	MocluData *celltype.ConfigCellSecondaryMarket
}

func (c *CfgSecondaryMarket) Ready() bool{
	return c.Data != nil && c.MocluData != nil && c.MocluData.FieldCount() > 0
}

func (c *CfgSecondaryMarket) Name() string {
	return "configCellSecondaryMarket:"
}

func (c *CfgSecondaryMarket) NotifyData(Data *ConfigCellChildDataObj) error {
	c.Data = Data
	if len(c.Data.MoleculeData) == 0 {
		temp := celltype.ConfigCellSecondaryMarketDefault()
		c.MocluData = &temp
		return nil
	}
	obj, err := celltype.ConfigCellSecondaryMarketFromSlice(c.Data.MoleculeData, false)
	if err != nil {
		return fmt.Errorf("ConfigCellSecondaryMarketFromSlice %s",err.Error())
	}
	c.MocluData = obj
	return nil
}

func (c *CfgSecondaryMarket) MocluObj() interface{} {
	return c.MocluData
}

func (c *CfgSecondaryMarket) Tag() celltype.TableType {
	return celltype.TableType_ConfigCell_SecondaryMarket
}

func (c *CfgSecondaryMarket) Witness() *celltype.CellDepWithWitness {
	return &celltype.CellDepWithWitness{
		CellDep: &c.Data.CellDep,
		GetWitnessData: func(index uint32) ([]byte, error) {
			return c.Data.WitnessData, nil
		}}
}