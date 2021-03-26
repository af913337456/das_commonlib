package eth_chain

import (
	"fmt"
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
