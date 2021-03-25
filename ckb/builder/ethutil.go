package builder

import (
	"github.com/ethereum/go-ethereum/crypto"
)

/**
 * Copyright (C), 2019-2021
 * FileName: eth
 * Author:   LinGuanHong
 * Date:     2021/3/25 6:11 下午
 * Description:
 */

func ETH_ComputeHash(rawBytes []byte) ([]byte, error) {
	// data, err := t.Serialize()
	// if err != nil {
	// 	return types.Hash{}, err
	// }
	return crypto.Keccak256(rawBytes), nil
}
