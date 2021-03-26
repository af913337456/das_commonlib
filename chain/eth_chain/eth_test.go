package eth_chain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
)

func TestSign(t *testing.T) {
	msg := "aaasdada"
	hexKey := ""
	sig, err := SignByPrivateKey(hexKey, msg)
	if err != nil {
		t.Fatal(err)
	}

	hexAddr := "0xc9f53b1d85356B60453F867610888D89a0B667Ad"
	ok, err := VerifySign(hexAddr, msg, sig)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ok)
}

func TestNewTransaction(t *testing.T) {
	rpcAddr := "http://47.52.147.232:8585"
	wsAddr := "ws://47.52.147.232:8586"
	client, err := ethclient.Dial(wsAddr)
	if err != nil {
		t.Fatal(err)
	}
	private := "" //
	to := "0xc9f53b1d85356B60453F867610888D89a0B667Ad"
	amount := new(big.Int).SetUint64(5670000000000)
	tx, err := NewTransaction(context.Background(), client, rpcAddr, private, to, amount, []byte("123456abcdef"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
	if err := client.SendTransaction(context.Background(), tx); err != nil {
		t.Fatal(err)
	}
}
