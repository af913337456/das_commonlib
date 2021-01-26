package celltype

import "github.com/nervosnetwork/ckb-sdk-go/types"

/**
 * Copyright (C), 2019-2020
 * FileName: cell_info
 * Author:   LinGuanHong
 * Date:     2020/12/22 3:01 下午
 * Description:
 */

var (
	TestNet_ConfigCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: "",
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: "",
			Args:         nil,
		},
	}
	TestNet_StateCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: "",
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: "",
			Args:         nil,
		},
	}
	TestNet_ActionCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: "",
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: "",
			Args:         nil,
		},
	}
	TestNet_ApplyRegisterCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: "",
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: "",
			Args:         nil,
		},
	}
	TestNet_RegisterCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: "",
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: "",
			Args:         nil,
		},
	}
	TestNet_OwnerCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: "",
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: "",
			Args:         nil, // [id, account_id]
		},
	}
	TestNet_PreAccountCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: "",
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: "",
			Args:         nil,
		},
	}
	TestNet_ProposeCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: "",
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: "",
			Args:         nil,
		},
	}
	TestNet_AccountCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: "",
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: "",
			Args:         nil,
		},
	}
	TestNet_PaymentCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: "",
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: "",
			Args:         nil,
		},
	}
	TestNet_QuoteCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: "",
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: "",
			Args:         nil,
		},
	}
	TestNet_TimeCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: "",
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash(""),
			CodeHashType: "",
			Args:         nil,
		},
	}
)
