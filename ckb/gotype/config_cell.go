package gotype

import (
	"errors"
	"fmt"

	"github.com/DeAccountSystems/das_commonlib/ckb/celltype"
	"github.com/DeAccountSystems/das_commonlib/ckb/gotype/configcells"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/syncmap"
)

/**
 * Copyright (C), 2019-2021
 * FileName: config_cell
 * Author:   LinGuanHong
 * Date:     2021/1/25 12:35
 * Description:
 */

type ConfigCell struct {
	ConfigCellChildMap syncmap.Map
}

func (c *ConfigCell) Ready() bool {
	ready := true
	c.ConfigCellChildMap.Range(func(key, value interface{}) bool {
		item := value.(configcells.IConfigChild)
		if !item.Ready() {
			ready = false
			return false
		}
		return true
	})
	return ready
}

func NewDefaultConfigCell(neyType celltype.DasNetType) *ConfigCell {
	c := &ConfigCell{
		ConfigCellChildMap: syncmap.Map{}, // map[celltype.CfgCellType]configcells.IConfigChild{},
	}

	c.ConfigCellChildMap.Store(celltype.TableType_ConfigCell_Main, &configcells.CfgMain{})
	c.ConfigCellChildMap.Store(celltype.TableType_ConfigCell_Price, &configcells.CfgPrice{})
	c.ConfigCellChildMap.Store(celltype.TableType_ConfigCell_CharSetEmoji, &configcells.CfgChatSetEmoji{})
	c.ConfigCellChildMap.Store(celltype.TableType_ConfigCell_CharSetDigit, &configcells.CfgChatSetDigit{})
	c.ConfigCellChildMap.Store(celltype.TableType_ConfigCell_CharSetEn, &configcells.CfgChatSetEn{})
	if neyType != celltype.DasNetType_Mainnet {
		c.ConfigCellChildMap.Store(celltype.TableType_ConfigCell_CharSetHanS, &configcells.CfgChatSetHans{})
		c.ConfigCellChildMap.Store(celltype.TableType_ConfigCell_CharSetHanT, &configcells.CfgChatSetHant{})
	}
	c.ConfigCellChildMap.Store(celltype.TableType_ConfigCell_Apply, &configcells.CfgApply{})
	c.ConfigCellChildMap.Store(celltype.TableType_ConfigCell_ProfitRate, &configcells.CfgProfitRate{})
	c.ConfigCellChildMap.Store(celltype.TableType_ConfigCell_Account, &configcells.CfgAccount{})
	c.ConfigCellChildMap.Store(celltype.TableType_ConfigCell_Proposal, &configcells.CfgProposal{})
	c.ConfigCellChildMap.Store(celltype.TableType_ConfigCell_Income, &configcells.CfgIncome{})
	c.ConfigCellChildMap.Store(celltype.TableType_ConfigCell_RecordNamespace, &configcells.CfgNameSpace{})
	c.ConfigCellChildMap.Store(celltype.TableType_ConfigCell_Release, &configcells.CfgRelease{})

	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount00, "PreservedAccount00"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount01, "PreservedAccount01"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount02, "PreservedAccount02"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount03, "PreservedAccount03"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount04, "PreservedAccount04"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount05, "PreservedAccount05"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount06, "PreservedAccount06"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount07, "PreservedAccount07"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount08, "PreservedAccount08"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount09, "PreservedAccount09"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount10, "PreservedAccount10"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount11, "PreservedAccount11"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount12, "PreservedAccount12"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount13, "PreservedAccount13"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount14, "PreservedAccount14"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount15, "PreservedAccount15"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount16, "PreservedAccount16"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount17, "PreservedAccount17"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount18, "PreservedAccount18"))
	c.StorePreservedAccountMap(configcells.NewCfgPreservedAccount(celltype.TableType_ConfigCell_PreservedAccount19, "PreservedAccount19"))

	return c
}

func (c *ConfigCell) RemoveConfigFromMap(tableType celltype.TableType) {
	c.ConfigCellChildMap.Delete(tableType)
}

func (c *ConfigCell) StorePreservedAccountMap(xx *configcells.CfgPreservedAccountXX) {
	c.ConfigCellChildMap.Store(xx.TableType, xx)
}

func (c *ConfigCell) main() *celltype.ConfigCellMain {
	v, _ := c.ConfigCellChildMap.Load(celltype.TableType_ConfigCell_Main)
	return (v.(configcells.IConfigChild)).MocluObj().(*celltype.ConfigCellMain)
}

func (c *ConfigCell) apply() *celltype.ConfigCellApply {
	v, _ := c.ConfigCellChildMap.Load(celltype.TableType_ConfigCell_Apply)
	return (v.(configcells.IConfigChild)).MocluObj().(*celltype.ConfigCellApply)
}

func (c *ConfigCell) price() *celltype.ConfigCellPrice {
	v, _ := c.ConfigCellChildMap.Load(celltype.TableType_ConfigCell_Price)
	return (v.(configcells.IConfigChild)).MocluObj().(*celltype.ConfigCellPrice)
}

func (c *ConfigCell) proposal() *celltype.ConfigCellProposal {
	v, _ := c.ConfigCellChildMap.Load(celltype.TableType_ConfigCell_Proposal)
	return (v.(configcells.IConfigChild)).MocluObj().(*celltype.ConfigCellProposal)
}

func (c *ConfigCell) account() *celltype.ConfigCellAccount {
	v, _ := c.ConfigCellChildMap.Load(celltype.TableType_ConfigCell_Account)
	return (v.(configcells.IConfigChild)).MocluObj().(*celltype.ConfigCellAccount)
}

func (c *ConfigCell) income() *celltype.ConfigCellIncome {
	v, _ := c.ConfigCellChildMap.Load(celltype.TableType_ConfigCell_Income)
	return (v.(configcells.IConfigChild)).MocluObj().(*celltype.ConfigCellIncome)
}

func (c *ConfigCell) profitRate() *celltype.ConfigCellProfitRate {
	v, _ := c.ConfigCellChildMap.Load(celltype.TableType_ConfigCell_ProfitRate)
	return (v.(configcells.IConfigChild)).MocluObj().(*celltype.ConfigCellProfitRate)
}

func (c *ConfigCell) release() *celltype.ConfigCellRelease {
	v, _ := c.ConfigCellChildMap.Load(celltype.TableType_ConfigCell_Release)
	return (v.(configcells.IConfigChild)).MocluObj().(*celltype.ConfigCellRelease)
}

func (c *ConfigCell) CKBSingleSoCellDepHash() types.Hash {
	return types.BytesToHash(c.main().DasLockOutPointTable().CkbSignall().TxHash().RawData())
}

func (c *ConfigCell) CKBMultiSoCellDepHash() types.Hash {
	return types.BytesToHash(c.main().DasLockOutPointTable().CkbMultisign().TxHash().RawData())
}

func (c *ConfigCell) ETHSoCellDepHash() types.Hash {
	return types.BytesToHash(c.main().DasLockOutPointTable().Eth().TxHash().RawData())
}

func (c *ConfigCell) TRONSoCellDepHash() types.Hash {
	return types.BytesToHash(c.main().DasLockOutPointTable().Tron().TxHash().RawData())
}

func (c *ConfigCell) OpenStatus() (uint8, error) {
	return celltype.MoleculeU8ToGo(c.main().Status().RawData())
}

func (c *ConfigCell) AccountCellBaseCap() (uint64, error) {
	val, err := celltype.MoleculeU64ToGo(c.account().BasicCapacity().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *ConfigCell) PreparedFeeCapacity() (uint64, error) {
	val, err := celltype.MoleculeU64ToGo(c.account().PreparedFeeCapacity().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *ConfigCell) AccountCellPrepareCap() (uint64, error) {
	val, err := celltype.MoleculeU64ToGo(c.account().PreparedFeeCapacity().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *ConfigCell) TransferAccountFee() (uint64, error) {
	val, err := celltype.MoleculeU64ToGo(c.account().TransferAccountFee().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *ConfigCell) EditManagerFee() (uint64, error) {
	val, err := celltype.MoleculeU64ToGo(c.account().EditManagerFee().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *ConfigCell) EditRecordsFee() (uint64, error) {
	val, err := celltype.MoleculeU64ToGo(c.account().EditRecordsFee().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *ConfigCell) EditRecordsThrottle() (uint32, error) {
	val, err := celltype.MoleculeU32ToGo(c.account().EditRecordsThrottle().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *ConfigCell) EditManagerThrottle() (uint32, error) {
	val, err := celltype.MoleculeU32ToGo(c.account().EditManagerThrottle().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *ConfigCell) TransferThrottle() (uint32, error) {
	val, err := celltype.MoleculeU32ToGo(c.account().TransferAccountThrottle().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *ConfigCell) IncomeCellBaseCap() (uint64, error) {
	val, err := celltype.MoleculeU64ToGo(c.income().BasicCapacity().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *ConfigCell) IncomeCellMinTransferValue() (uint64, error) {
	val, err := celltype.MoleculeU64ToGo(c.income().MinTransferCapacity().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *ConfigCell) GetRegisterProfitConfig() *celltype.ConfigCellProfitRate {
	return c.profitRate()
}

func (c *ConfigCell) ProposalMinConfirmRequire() (uint8, error) {
	val, err := celltype.MoleculeU8ToGo(c.proposal().ProposalMinConfirmInterval().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *ConfigCell) ProposalMinExtendInterval() (uint8, error) {
	val, err := celltype.MoleculeU8ToGo(c.proposal().ProposalMinExtendInterval().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *ConfigCell) ProposalMinRecycleInterval() (uint8, error) {
	val, err := celltype.MoleculeU8ToGo(c.proposal().ProposalMinRecycleInterval().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

// proposal_max_account_affect
// func (c *ConfigCell) ProposalMaxAccountAffect() (uint32, error) {
// 	if c == nil {
// 		return 0, errors.New("ConfigCell is nil")
// 	}
// 	if c.ConfigCellData == nil {
// 		return 0, errors.New("ConfigCellData is nil")
// 	}
// 	val, err := celltype.MoleculeU32ToGo(c.ConfigCellData.ProposalMaxAccountAffect().RawData())
// 	if err != nil {
// 		return 0, err
// 	}
// 	return val, nil
// }

func (c *ConfigCell) ProposalMaxPreAccountContain() (uint32, error) {
	val, err := celltype.MoleculeU32ToGo(c.proposal().ProposalMaxPreAccountContain().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

// account_max_length
func (c *ConfigCell) AccountMaxLength() (uint32, error) {
	val, err := celltype.MoleculeU32ToGo(c.account().MaxLength().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

// apply_min_waiting_time
func (c *ConfigCell) ApplyMinWaitingBlockNumber() (uint32, error) {
	val, err := celltype.MoleculeU32ToGo(c.apply().ApplyMinWaitingBlockNumber().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

// apply_max_waiting_time
func (c *ConfigCell) ApplyMaxWaitingBlockNumber() (uint32, error) {
	val, err := celltype.MoleculeU32ToGo(c.apply().ApplyMaxWaitingBlockNumber().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

// frozen
func (c *ConfigCell) AccountExpirationGracePeriod() (uint32, error) {
	val, err := celltype.MoleculeU32ToGo(c.account().ExpirationGracePeriod().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *ConfigCell) InvitedDiscount() (uint32, error) {
	val, err := celltype.MoleculeU32ToGo(c.price().Discount().InvitedDiscount().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c *ConfigCell) InvitedDiscountFormatValue() (float64, error) {
	val, err := celltype.MoleculeU32ToGo(c.price().Discount().InvitedDiscount().RawData())
	if err != nil {
		return 0, err
	}
	return float64(val) / float64(celltype.DiscountRateBase), nil
}

func (c *ConfigCell) GetProfitOfInviter() (decimal.Decimal, error) {
	profit := c.GetRegisterProfitConfig()
	profitRateOfInviter, err := celltype.MoleculeU32ToGo(profit.Inviter().RawData())
	if err != nil {
		return decimal.Zero, err
	}
	dec := decimal.NewFromInt(int64(profitRateOfInviter))
	return dec.Div(decimal.NewFromInt(int64(celltype.DiscountRateBase))), nil
}

func (c *ConfigCell) GetProfitOfChannel() (decimal.Decimal, error) {
	profit := c.GetRegisterProfitConfig()
	profitRateOfChannel, err := celltype.MoleculeU32ToGo(profit.Channel().RawData())
	if err != nil {
		return decimal.Zero, err
	}
	dec := decimal.NewFromInt(int64(profitRateOfChannel))
	return dec.Div(decimal.NewFromInt(int64(celltype.DiscountRateBase))), nil
}

func (c *ConfigCell) InvitedDiscountFormatDiscountObj() (float64, error) {
	val, err := celltype.MoleculeU32ToGo(c.price().Discount().InvitedDiscount().RawData())
	if err != nil {
		return 0, err
	}
	return float64(val) / float64(celltype.DiscountRateBase), nil
}

func (c *ConfigCell) AccountTTL() (uint32, error) {
	val, err := celltype.MoleculeU32ToGo(c.account().RecordMinTtl().RawData())
	if err != nil {
		return 0, err
	}
	return val, nil
}

type ReleaseRuleItem struct {
	Length         uint32
	StartTimestamp uint64
	EndTimestamp   uint64
}

func (c *ConfigCell) ReleaseRules() ([]ReleaseRuleItem, error) {
	releaseRules := c.release().ReleaseRules()
	total := releaseRules.ItemCount()
	ruleIndex := uint(0)
	retList := make([]ReleaseRuleItem, 0, total)
	for ; ruleIndex < total; ruleIndex++ {
		item := releaseRules.Get(ruleIndex)
		length, err := celltype.MoleculeU32ToGo(item.Length().RawData())
		if err != nil {
			return nil, err
		}
		start, err := celltype.MoleculeU64ToGo(item.ReleaseStart().RawData())
		if err != nil {
			return nil, err
		}
		end, err := celltype.MoleculeU64ToGo(item.ReleaseEnd().RawData())
		if err != nil {
			return nil, err
		}
		retList = append(retList, ReleaseRuleItem{
			Length:         length,
			StartTimestamp: start,
			EndTimestamp:   end,
		})
	}
	return retList, nil
}

// func (c *ConfigCell) MaxSellingTime() (uint32, error) {
// 	val, err := celltype.MoleculeU32ToGo(c.market().PrimaryMarket().MaxSellingTime().RawData())
// 	if err != nil {
// 		return 0, err
// 	}
// 	return val, nil
// }

// func (c *ConfigCell) AccountCellTypeId() types.Hash {
// 	return types.BytesToHash(c.main().TypeIdTable().AccountCell().RawData())
// }

// code hash table
// func (c *ConfigCell) GetSystemCellInfoMap() (map[types.Hash]string, error) {
// 	if c == nil || c.Main == nil {
// 		return nil, errors.New("configCellMain is empty")
// 	}
// 	return celltype.ParseDasCellsScript(c.Main.Data), nil
// }

func (c *ConfigCell) GetAccountPriceConfig(account celltype.DasAccount) (*celltype.PriceConfig, error) {
	priceList := c.price().Prices()
	total := priceList.ItemCount()
	priceIndex := uint(0)
	preAccountLen := uint8(celltype.MinAccountCharsLen)
	var preItem *celltype.PriceConfig = nil
	accountCharsLen := uint8(account.AccountValidateLen()) // 字符长度
	for ; priceIndex < total; priceIndex++ {
		item := priceList.Get(priceIndex)
		accountLen, err := celltype.MoleculeU8ToGo(item.Length().RawData())
		if err != nil {
			return nil, err
		} else if accountLen < celltype.MinAccountCharsLen {
			continue
		} else if accountLen == accountCharsLen {
			return item, nil
		} else {
			preAccountLen = accountLen
			preItem = item
		}
	}
	if accountCharsLen > preAccountLen && preItem != nil {
		return preItem, nil
	}
	return nil, fmt.Errorf("account price not found, account: %s", account)
}

func (c *ConfigCell) GetAccountPrice(account celltype.DasAccount, isRenew bool) (*celltype.PriceConfig, uint64, error) {
	price, err := c.GetAccountPriceConfig(account)
	if err != nil {
		return nil, 0, err
	}
	if !isRenew {
		newPrice, err := celltype.MoleculeU64ToGo(price.New().RawData())
		if err != nil {
			return nil, 0, err
		}
		return price, newPrice, nil
	} else {
		newPrice, err := celltype.MoleculeU64ToGo(price.Renew().RawData())
		if err != nil {
			return nil, 0, err
		}
		return price, newPrice, nil
	}
}

func (c *ConfigCell) GetAccountRenewPrice(account celltype.DasAccount) (*celltype.PriceConfig, uint64, error) {
	price, err := c.GetAccountPriceConfig(account)
	if err != nil {
		return nil, 0, err
	}
	renewPrice, err := celltype.MoleculeU64ToGo(price.Renew().RawData())
	if err != nil {
		return nil, 0, err
	}
	return price, renewPrice, nil
}

func (c *ConfigCell) GetWitnessCellDep(cfgType celltype.TableType) *celltype.CellDepWithWitness {
	if obj, found := c.ConfigCellChildMap.Load(cfgType); !found {
		return nil
	} else {
		return obj.(configcells.IConfigChild).Witness()
	}
}

type ProfitRate struct {
	Invite         float64
	Channel        float64
	ProposeCreate  float64
	ProposeConfirm float64
	MergeRate      float64
}

func ParseRegisterProfitConfig(configCell *ConfigCell) (*ProfitRate, error) {
	profit := configCell.GetRegisterProfitConfig()
	inviterRate, err1 := celltype.MoleculeU32ToGoPercentage(profit.Inviter().RawData())
	channelRate, err2 := celltype.MoleculeU32ToGoPercentage(profit.Channel().RawData())
	propoCreate, err4 := celltype.MoleculeU32ToGoPercentage(profit.ProposalCreate().RawData())
	propConfirm, err5 := celltype.MoleculeU32ToGoPercentage(profit.ProposalConfirm().RawData())
	mergeFeeRat, err6 := celltype.MoleculeU32ToGoPercentage(profit.IncomeConsolidate().RawData())
	if err1 != nil || err2 != nil || err4 != nil || err5 != nil || err6 != nil {
		return nil, fmt.Errorf("parse profitRate err")
	}
	if inviterRate+channelRate+propoCreate+propConfirm+mergeFeeRat > 1 {
		return nil, fmt.Errorf("invalid profitRate, more than 100,"+
			" inviter: %f, channel: %f, creator: %f, confirm: %f, merge: %f",
			inviterRate, channelRate, propoCreate, propConfirm, mergeFeeRat)
	}
	return &ProfitRate{
		Invite:         inviterRate,
		Channel:        channelRate,
		ProposeCreate:  propoCreate,
		ProposeConfirm: propConfirm,
		MergeRate:      mergeFeeRat,
	}, nil
}

func BindConfigCellDataFromTx(tx *types.Transaction, configCell *ConfigCell) error {
	err := getTargetCellFromWitness(tx, func(rawWitnessData []byte, witnessParseObj *celltype.ParseDasWitnessBysDataObj) (bool, error) {
		witnessDataObj := witnessParseObj.WitnessObj
		if !witnessDataObj.TableType.IsConfigType() {
			return false, errors.New("skip, witness obj's tableType not configCell type")
		}
		_, index, err := witnessParseObj.NewEntity()
		if err != nil {
			return false, err
		}
		cellDep := types.CellDep{
			OutPoint: &types.OutPoint{
				TxHash: tx.Hash,
				Index:  uint(index),
			},
			DepType: types.DepTypeCode,
		}
		configTableType := witnessParseObj.WitnessObj.TableType
		cellData := witnessParseObj.MoleculeNewDataEntity.Entity().RawData()
		v, ok := configCell.ConfigCellChildMap.Load(configTableType)
		if !ok {
			return false, nil
		}
		_ = v.(configcells.IConfigChild).NotifyData(&configcells.ConfigCellChildDataObj{
			CellDep:      cellDep,
			WitnessData:  rawWitnessData,
			MoleculeData: cellData,
		})
		if configTableType == celltype.TableType_ConfigCell_Main {
			celltype.CKBSoScriptDep.TxHash = configCell.CKBSingleSoCellDepHash()
			celltype.ETHSoScriptDep.TxHash = configCell.ETHSoCellDepHash()
			celltype.TRONSoScriptDep.TxHash = configCell.TRONSoCellDepHash()
			celltype.CKBMultiSoScriptDep.TxHash = configCell.CKBMultiSoCellDepHash()
		}
		return false, nil
	})
	return err
}

func getTargetCellFromWitness(tx *types.Transaction, handle celltype.ValidHandle) error {
	return celltype.GetTargetCellFromWitness(tx, handle, func(err error) {})
}
