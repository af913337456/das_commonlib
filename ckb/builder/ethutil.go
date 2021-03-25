package builder

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2021
 * FileName: eth
 * Author:   LinGuanHong
 * Date:     2021/3/25 6:11 下午
 * Description:
 */

func ETH_ComputeTxHash(t *types.Transaction) (types.Hash, error) {
	data, err := t.Serialize()
	if err != nil {
		return types.Hash{}, err
	}
	return types.BytesToHash(crypto.Keccak256(data)), nil
}

func ETH_ComputeHash(rawBytes []byte) ([]byte, error) {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", 32, rawBytes)
	return crypto.Keccak256([]byte(msg)), nil
}
