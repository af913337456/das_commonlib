package celltype

import (
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2020
 * FileName: cell_info
 * Author:   LinGuanHong
 * Date:     2020/12/22 3:01 下午
 * Description:
 */

var (
	DasAnyOneCanSendCellInfo = DASCellBaseInfoOut{
		CodeHash:     types.HexToHash(""),
		CodeHashType: types.HashTypeType,
		Args:         nil,
	}
	DasActionCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: types.HashTypeType,
			Args:         nil,
		},
	}
	DasWalletCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: types.HashTypeType,
			Args:         nil,
		},
	}
	DasApplyRegisterCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: types.HashTypeType,
			Args:         nil,
		},
	}
	DasRefCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: types.HashTypeType,
			Args:         nil,
		},
	}
	DasPreAccountCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: types.HashTypeType,
			Args:         nil,
		},
	}
	DasProposeCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: types.HashTypeType,
			Args:         nil,
		},
	}
	DasAccountCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: types.HashTypeType,
			Args:         nil,
		},
	}
	// DasQuoteCellScript = DASCellBaseInfo{
	// 	Dep: DASCellBaseInfoDep{
	// 		TxHash:  types.HexToHash(""),
	// 		TxIndex: 0,
	// 		DepType: "",
	// 	},
	// 	Out: DASCellBaseInfoOut{
	// 		CodeHash:     types.HexToHash(""),
	// 		CodeHashType: "",
	// 		Args:         nil,
	// 	},
	// }
	// DasConfigCellScript = DASCellBaseInfo{
	// 	Dep: DASCellBaseInfoDep{
	// 		TxHash:  types.HexToHash(""),
	// 		TxIndex: 0,
	// 		DepType: "",
	// 	},
	// 	Out: DASCellBaseInfoOut{
	// 		CodeHash:     types.HexToHash(""),
	// 		CodeHashType: "",
	// 		Args:         nil,
	// 	},
	// }
	DasTimeCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: types.HashTypeType,
			Args:         nil,
		},
	}
	SystemCodeScriptMap = map[string]*DASCellBaseInfo{}
)

func init() {
	SystemCodeScriptMap[SystemScript_ApplyRegisterCell] = &DasApplyRegisterCellScript
	SystemCodeScriptMap[SystemScript_PreAccoutnCell] = &DasPreAccountCellScript
	SystemCodeScriptMap[SystemScript_AccoutnCell] = &DasAccountCellScript
	SystemCodeScriptMap[SystemScript_ProposeCell] = &DasProposeCellScript
	SystemCodeScriptMap[SystemScript_WalletCell] = &DasWalletCellScript
	SystemCodeScriptMap[SystemScript_RefCell] = &DasRefCellScript
}

func ParseDasCellsScript(data *ConfigCellData) map[types.Hash]string {
	applyRegisterCodeHash := types.BytesToHash(data.TypeIdTable().ApplyRegisterCell().RawData())
	preAccountCellCodeHash := types.BytesToHash(data.TypeIdTable().PreAccountCell().RawData())
	accountCellCodeHash := types.BytesToHash(data.TypeIdTable().AccountCell().RawData())
	proposeCellCodeHash := types.BytesToHash(data.TypeIdTable().ProposalCell().RawData())
	walletCellCodeHash := types.BytesToHash(data.TypeIdTable().WalletCell().RawData())
	refCellCodeHash := types.BytesToHash(data.TypeIdTable().RefCell().RawData())

	retMap := map[types.Hash]string{}
	retMap[applyRegisterCodeHash] = SystemScript_ApplyRegisterCell
	retMap[preAccountCellCodeHash] = SystemScript_PreAccoutnCell
	retMap[accountCellCodeHash] = SystemScript_AccoutnCell
	retMap[proposeCellCodeHash] = SystemScript_ProposeCell
	retMap[walletCellCodeHash] = SystemScript_WalletCell
	retMap[refCellCodeHash] = SystemScript_RefCell
	return retMap
}

func SetSystemScript(scriptName string, dasCellBaseInfo *DASCellBaseInfo) error {
	if v := SystemCodeScriptMap[scriptName]; v != nil {
		*SystemCodeScriptMap[scriptName] = *dasCellBaseInfo
		return nil
	}
	return errors.New("unSupport script")
}
