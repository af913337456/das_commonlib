package builder

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

/**
 * Copyright (C), 2019-2021
 * FileName: ethutil_test
 * Author:   LinGuanHong
 * Date:     2021/3/25 10:12 下午
 * Description:
 */

func Test_ETH_ComputeHash(t *testing.T) {
	msg := []byte("hello")
	private, err := crypto.HexToECDSA("9111ecca8bd827f3568eca5a29433d27b9949dfdff7b956eb1b2a2657386a339")
	if err != nil {
		panic(err)
	}
	cmpMsgBytes, err := ETH_ComputeHash(msg)
	if err != nil {
		panic(err)
	}
	sig, err := crypto.Sign(cmpMsgBytes, private)
	if err != nil {
		panic(err)
	}
	if err := checkSign("0xa0324794ff56ecb258220046034a363d0da98f51", "hello", sig); err != nil {
		panic(err)
	}
	t.Log("success")
}

func checkSign(address, data string, signBytes []byte) error {
	msg := crypto.Keccak256([]byte(data))
	addr := common.HexToAddress(address)
	recoveredPub, err := crypto.Ecrecover(msg, signBytes)
	if err != nil {
		return err
	}
	pubKey, _ := crypto.UnmarshalPubkey(recoveredPub)
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	if addr != recoveredAddr {
		return errors.New("invalid sign address")
	}
	return nil
}
