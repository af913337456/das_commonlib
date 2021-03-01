package common

/**
 * Copyright (C), 2019-2021
 * FileName: errcode
 * Author:   LinGuanHong
 * Date:     2021/2/4 1:02 下午
 * Description:
 */

type DAS_CODE int

const DAS_SUCCESS DAS_CODE = 0

const (
	CallIndexerErr         = 20000
	InternalErr            = 20001
	AccountExpired         = 20002
	AccountFrozen          = 20003
	AccountAlreadyRegister = 20004
	AccountRecordsInvalid  = 20005
	AccountFormatInvalid   = 20006
	PubkeyHexFormatInvalid = 20007
	BaseParamInvalid       = 20008
)
