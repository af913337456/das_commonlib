package common

/**
 * Copyright (C), 2019-2020
 * FileName: types.
 * Author:   LinGuanHong
 * Date:     2020/12/22 3:26 下午
 * Description:
 */

type ReqResp struct {
	ErrNo  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}
