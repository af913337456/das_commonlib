package dascode

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
	Err_CallIndexer            = 20000
	Err_Internal               = 20001
	Err_AccountExpired         = 20002
	Err_AccountFrozen          = 20003
	Err_AccountAlreadyRegister = 20004
	Err_AccountRecordsInvalid  = 20005
	Err_AccountFormatInvalid   = 20006
	Err_PubkeyHexFormatInvalid = 20007
	Err_BaseParamInvalid       = 20008
)
