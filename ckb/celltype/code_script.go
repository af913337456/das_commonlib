package celltype

import (
	"encoding/hex"
	"encoding/json"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
	"strings"
)

/**
 * Copyright (C), 2019-2020
 * FileName: cell_info
 * Author:   LinGuanHong
 * Date:     2020/12/22 3:01 下午
 * Description:
 */

var (
	TestNetLockScriptDep = DASCellBaseInfoDep{
		TxHash:  types.HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"),
		TxIndex: 0,
		DepType: types.DepTypeDepGroup,
	}

	DasETHLockCellInfo = DASCellBaseInfoOut{
		CodeHash:     types.HexToHash(PwLockTestNetCodeHash), // default
		CodeHashType: types.HashTypeType,
		Args:         emptyHexToArgsBytes(),
	}
	DasBTCLockCellInfo = DASCellBaseInfoOut{
		CodeHash:     types.HexToHash(""),
		CodeHashType: types.HashTypeType,
		Args:         emptyHexToArgsBytes(),
	}
	DasAnyOneCanSendCellInfo = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash("0x6cb20b88912311e6bba89a5fcfc53cfebcc39b99c3cce0796ce3e485a5d47011"),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash("0xd483925160e4232b2cb29f012e8380b7b612d71cf4e79991476b6bcf610735f6"), // default
			CodeHashType: types.HashTypeData,
			Args:         emptyHexToArgsBytes(),
		},
	}

	DasAnyOneCanPayCellInfo = DASCellBaseInfoOut{
		CodeHash:     types.HexToHash(utils.AnyoneCanPayCodeHashOnAggron), // default
		CodeHashType: types.HashTypeType,
		Args:         emptyHexToArgsBytes(),
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
			Args:         emptyHexToArgsBytes(),
		},
	}
	DasWalletCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash("0x71d4d91e3a416fbab4c7d53e4db3b46cc473b265d85d3c317c6f335532e5285c"),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash("0x9878b226df9465c215fd3c94dc9f9bf6648d5bea48a24579cf83274fe13801d2"),
			CodeHashType: types.HashTypeType,
			Args:         emptyHexToArgsBytes(),
		},
	}
	DasApplyRegisterCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash("0xc9b182b156be56b4420f6e78621604fa0dc71e85418171402cddc8cdc5fec932"),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash("0xa2c3a2b18da897bd24391a921956e45d245b46169d6acc9a0663316d15b51cb1"),
			CodeHashType: types.HashTypeType,
			Args:         emptyHexToArgsBytes(),
		},
	}
	DasRefCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash("0x628da7b8d6d2c6736c63efa4b517f5f2ada6a8c58d49b3244b89c01f1b7e8277"),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash("0xe79953f024552e6130220a03d2497dc7c2f784f4297c69ba21d0c423915350e5"),
			CodeHashType: types.HashTypeType,
			Args:         emptyHexToArgsBytes(),
		},
	}
	DasPreAccountCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash("0x42642847df624fe7147428f996cc8d579a8652582847fd875be6c2b84056dece"),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash("0x92d6a9525b9a054222982ab4740be6fe4281e65fff52ab252e7daf9306e12e3f"),
			CodeHashType: types.HashTypeType,
			Args:         emptyHexToArgsBytes(),
		},
	}
	DasProposeCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash("0x3141859166855bed53973e5fec772b41e2dba4c88cf79fa33cc2f98da05f97fa"),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash("0x4154b5f9114b8d2dd8323eead5d5e71d0959a2dc73f0672e829ae4dabffdb2d8"),
			CodeHashType: types.HashTypeType,
			Args:         emptyHexToArgsBytes(),
		},
	}
	DasAccountCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash("0xf81ad628f4ce019643ad2f8dbcac010bd8aebe802213814254f1a79301d956c4"),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash("0x274775e475c1252b5333c20e1512b7b1296c4c5b52a25aa2ebd6e41f5894c41f"),
			CodeHashType: types.HashTypeType,
			Args:         emptyHexToArgsBytes(),
		},
	}
	DasBiddingCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash("0x123"),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash("0x123"),
			CodeHashType: types.HashTypeType,
			Args:         emptyHexToArgsBytes(),
		},
	}
	DasOnSaleCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash("0x123"),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash("0x123"),
			CodeHashType: types.HashTypeType,
			Args:         emptyHexToArgsBytes(),
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
	DasConfigCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash(""),
			TxIndex: 0,
			DepType: "",
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash("0x489ff2195ed41aac9a9265c653d8ca57c825b22db765b9e08d537572ff2cbc1b"),
			CodeHashType: types.HashTypeType,
			Args:         emptyHexToArgsBytes(),
		},
	}
	DasHeightCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash("0x711bb5cec27b3a5c00da3a6dc0772be8651f7f92fd9bf09d77578b29227c1748"),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash("0x5f6a4cc2cd6369dbcf38ddfbc4323cf4695c2e8c20aed572b5db6adc2faf9d50"),
			CodeHashType: types.HashTypeType,
			Args:         hexToArgsBytes("0xe1a958a4c112af95a1220c6fee5f969972a3d8ce13fb7b3211f71abb5db1824102000000"),
		},
	}
	DasTimeCellScript = DASCellBaseInfo{
		Dep: DASCellBaseInfoDep{
			TxHash:  types.HexToHash("0xf3c13ffbaa1d34b8fac6cd848fa04db2e6b4e2c967c3c178295be2e7cdd77164"),
			TxIndex: 0,
			DepType: types.DepTypeCode,
		},
		Out: DASCellBaseInfoOut{
			CodeHash:     types.HexToHash("0xe4fd6f46ab1fd3d5b377df9e2d4ea77e3b52f53ac3319595bb38d097ea051cfd"),
			CodeHashType: types.HashTypeType,
			Args:         hexToArgsBytes("0xd0c1c7156f2e310a12822e2cc336398ec4ef194abc1f96023b743f3249f09e2102000000"),
		},
	}
	SystemCodeScriptMap = map[types.Hash]*DASCellBaseInfo{}
)

func init() {
	SystemCodeScriptMap[DasApplyRegisterCellScript.Out.CodeHash] = &DasApplyRegisterCellScript
	SystemCodeScriptMap[DasPreAccountCellScript.Out.CodeHash] = &DasPreAccountCellScript
	SystemCodeScriptMap[DasBiddingCellScript.Out.CodeHash] = &DasBiddingCellScript
	SystemCodeScriptMap[DasAccountCellScript.Out.CodeHash] = &DasAccountCellScript
	SystemCodeScriptMap[DasOnSaleCellScript.Out.CodeHash] = &DasOnSaleCellScript
	SystemCodeScriptMap[DasProposeCellScript.Out.CodeHash] = &DasProposeCellScript
	SystemCodeScriptMap[DasWalletCellScript.Out.CodeHash] = &DasWalletCellScript
	SystemCodeScriptMap[DasRefCellScript.Out.CodeHash] = &DasRefCellScript
}

func SetSystemCodeScriptOutPoint(typeId types.Hash, point types.OutPoint) *DASCellBaseInfo {
	if _, ok := SystemCodeScriptMap[typeId]; !ok {
		return nil
	}
	SystemCodeScriptMap[typeId].Dep.TxHash = point.TxHash
	SystemCodeScriptMap[typeId].Dep.TxIndex = point.Index
	return SystemCodeScriptMap[typeId]
}

func emptyHexToArgsBytes() []byte {
	return []byte{}
}

func hexToArgsBytes(hexStr string) []byte {
	if strings.HasPrefix(hexStr, "0x") {
		hexStr = hexStr[2:]
	}
	bys, _ := hex.DecodeString(hexStr)
	return bys
}

func IsSystemCodeScriptReady() bool {
	for _, item := range SystemCodeScriptMap {
		if item.Out.CodeHash.Hex() == "0x" {
			return false
		}
	}
	return true
}

func SystemCodeScriptBytes() ([]byte, error) {
	return json.Marshal(SystemCodeScriptMap)
}

func SystemCodeScriptFromBytes(bys []byte) error {
	if err := json.Unmarshal(bys, &SystemCodeScriptMap); err != nil {
		return err
	}
	return nil
}

//
// func ParseDasCellsScript(data *ConfigCellMain) map[types.Hash]string {
// 	applyRegisterCodeHash := types.BytesToHash(data.TypeIdTable().ApplyRegisterCell().RawData())
// 	preAccountCellCodeHash := types.BytesToHash(data.TypeIdTable().PreAccountCell().RawData())
// 	biddingCellCodeHash := types.BytesToHash(data.TypeIdTable().BiddingCell().RawData())
// 	accountCellCodeHash := types.BytesToHash(data.TypeIdTable().AccountCell().RawData())
// 	proposeCellCodeHash := types.BytesToHash(data.TypeIdTable().ProposalCell().RawData())
// 	onSaleCellCodeHash := types.BytesToHash(data.TypeIdTable().OnSaleCell().RawData())
// 	walletCellCodeHash := types.BytesToHash(data.TypeIdTable().WalletCell().RawData())
// 	refCellCodeHash := types.BytesToHash(data.TypeIdTable().RefCell().RawData())
//
// 	retMap := map[types.Hash]string{}
// 	retMap[applyRegisterCodeHash] = SystemScript_ApplyRegisterCell
// 	retMap[preAccountCellCodeHash] = SystemScript_PreAccoutnCell
// 	retMap[biddingCellCodeHash] = SystemScript_BiddingCell
// 	retMap[accountCellCodeHash] = SystemScript_AccoutnCell
// 	retMap[proposeCellCodeHash] = SystemScript_ProposeCell
// 	retMap[onSaleCellCodeHash] = SystemScript_OnSaleCell
// 	retMap[walletCellCodeHash] = SystemScript_WalletCell
// 	retMap[refCellCodeHash] = SystemScript_RefCell
// 	return retMap
// }
//
// func SetSystemScript(scriptName string, dasCellBaseInfo *DASCellBaseInfo) error {
// 	if v := SystemCodeScriptMap[scriptName]; v != nil {
// 		*SystemCodeScriptMap[scriptName] = *dasCellBaseInfo
// 		return nil
// 	}
// 	return errors.New("unSupport script")
// }
