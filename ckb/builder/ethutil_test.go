package builder

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

/**
 * Copyright (C), 2019-2021
 * FileName: ethutil_test
 * Author:   LinGuanHong
 * Date:     2021/3/25 10:12
 * Description:
 */

func Test_ETH_ComputeHash(t *testing.T) {
	testPrivHex := "9111ecca8bd827f3568eca5a29433d27b9949dfdff7b956eb1b2a2657386a339"
	testAddrHex := "a0324794ff56ecb258220046034a363d0da98f51"
	key, _ := crypto.HexToECDSA(testPrivHex)
	addr := common.HexToAddress(testAddrHex)

	xx := crypto.Keccak256([]byte("foo"))
	fmt.Println(len(xx))
	msg, err := ETHMessageDigest(xx)
	if err != nil {
		t.Errorf("Sign error: %s", err)
	}
	sig, err := crypto.Sign(msg, key)
	if err != nil {
		t.Errorf("Sign error: %s", err)
	}
	recoveredPub, err := crypto.Ecrecover(msg, sig)
	if err != nil {
		t.Errorf("ECRecover error: %s", err)
	}
	pubKey, _ := crypto.UnmarshalPubkey(recoveredPub)
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	if addr != recoveredAddr {
		t.Errorf("Address mismatch: want: %x have: %x", addr, recoveredAddr)
	}

	// should be equal to SigToPub
	recoveredPub2, err := crypto.SigToPub(msg, sig)
	if err != nil {
		t.Errorf("ECRecover error: %s", err)
	}
	recoveredAddr2 := crypto.PubkeyToAddress(*recoveredPub2)
	if addr != recoveredAddr2 {
		t.Errorf("Address mismatch: want: %x have: %x", addr, recoveredAddr2)
	}
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
