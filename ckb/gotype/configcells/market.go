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

type CfgMarket struct {
	Data *ConfigCellChildDataObj
	MocluData *celltype.ConfigCellSecondaryMarket
}

func (c *CfgMarket) Ready() bool{
	return c.Data != nil && c.MocluData.FieldCount() > 0
}

func (c *CfgMarket) Name() string {
	return "configCellMarket:"
}

func (c *CfgMarket) NotifyData(Data *ConfigCellChildDataObj) error {
	c.Data = Data
	if len(c.Data.MoleculeData) == 0 {
		temp := celltype.ConfigCellSecondaryMarketDefault()
		c.MocluData = &temp
		return nil
	}
	obj, err := celltype.ConfigCellSecondaryMarketFromSlice(c.Data.MoleculeData, false)
	if err != nil {
		return fmt.Errorf("ConfigCellMarketFromSlice %s",err.Error())
	}
	c.MocluData = obj
	return nil
}

func (c *CfgMarket) Tag() celltype.TableType {
	return celltype.TableType_ConfigCell_SecondaryMarket
}

func (c *CfgMarket) MocluObj() interface{} {
	return c.MocluData
}

func (c *CfgMarket) Witness() *celltype.CellDepWithWitness {
	return &celltype.CellDepWithWitness{
		CellDep: &c.Data.CellDep,
		GetWitnessData: func(index uint32) ([]byte, error) {
			return c.Data.WitnessData, nil
		}}
}