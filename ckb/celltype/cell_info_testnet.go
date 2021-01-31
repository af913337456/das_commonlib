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
	DasAnyOneCanSendCellInfo = DASCellBaseInfoOut{
		CodeHash:     types.HexToHash(""),
		CodeHashType: "",
		Args:         nil,
	}
	DasConfigCellScript = DASCellBaseInfo{
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
	DasStateCellScript = DASCellBaseInfo{
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
	DasActionCellScript = DASCellBaseInfo{
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
	DasApplyRegisterCellScript = DASCellBaseInfo{
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
	DasRegisterCellScript = DASCellBaseInfo{
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
	DasOwnerCellScript = DASCellBaseInfo{
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
	DasPreAccountCellScript = DASCellBaseInfo{
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
	DasProposeCellScript = DASCellBaseInfo{
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
	DasAccountCellScript = DASCellBaseInfo{
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
	DasPaymentCellScript = DASCellBaseInfo{
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
	DasQuoteCellScript = DASCellBaseInfo{
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
	DasTimeCellScript = DASCellBaseInfo{
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
