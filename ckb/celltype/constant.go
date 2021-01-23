package celltype

/**
 * Copyright (C), 2019-2020
 * FileName: value
 * Author:   LinGuanHong
 * Date:     2020/12/20 3:12 下午
 * Description:
 */

const witnessDas = "das"
const CellVersionByteLen = 4
const MoleculeBytesHeaderSize = 4
const OneCkb = uint64(1e8)
const CkbTxMinOutputCKBValue = 61 * OneCkb

type TableType uint32
type AccountCellStatus uint8

const (
	TableType_ACTION       TableType = 0
	TableType_CONFIG_CELL  TableType = 1
	TableType_ACCOUNT_CELL TableType = 2
	// TableType_REGISTER_CELL TableType = 3
	TableType_ON_SALE_CELL     TableType = 3
	TableType_BIDDING_CELL     TableType = 4
	TableType_PROPOSE_CELL     TableType = 5 // todo change it
	TableType_PRE_ACCOUNT_CELL TableType = 6
)

const (
	AccountCellStatus_Exist    = 0
	AccountCellStatus_Proposed = 1
	AccountCellStatus_New      = 2
)

const (
	Action_ConfigState           = "config_state"
	Action_ApplyRegister         = "apply_register"
	Action_PreRegister           = "pre_register"
	Action_Propose               = "propose"
	Action_ExtendPropose         = "extend_propose"
	Action_ConfirmProposal       = "confirm_proposal"
	Action_Register              = "register"
	Action_VoteBiddingList       = "vote_bidding_list"
	Action_PublishAccount        = "publish_account"
	Action_RejectRegister        = "reject_register"
	Action_PublishBiddingList    = "publish_bidding_list"
	Action_BidAccount            = "bid_account"
	Action_EditManager           = "edit_manager"
	Action_EditRecords           = "edit_records"
	Action_CancelBidding         = "cancel_bidding"
	Action_CloseBidding          = "close_bidding"
	Action_QuotePriceForCkb      = "quote_price_for_ckb"
	Action_StartAccountSale      = "start_account_sale"
	Action_CancelAccountSale     = "cancel_account_sale"
	Action_StartAccountAuction   = "start_account_auction"
	Action_CancelAccountAuction  = "cancel_account_auction"
	Action_AccuseAccountRepeat   = "accuse_account_repeat"
	Action_AccuseAccountIllegal  = "accuse_account_illegal"
	Action_RecycalExpiredAccount = "recycal_expired_account"
)
