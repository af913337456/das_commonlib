package eth_chain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func NewTransaction(ctx context.Context, client *ethclient.Client, rpcAddr, private, to string, amount *big.Int, data []byte) (trx *types.Transaction, err error) {
	privateKey, err := crypto.HexToECDSA(HexFormat(private))
	if err != nil {
		return nil, fmt.Errorf("HexToECDSA err:%v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("publicKey is not type ecdsa.PublicKey")
	}
	fromAddr, toAddr := crypto.PubkeyToAddress(*publicKeyECDSA), common.HexToAddress(to)
	nonce, err := PendingNonceAt(rpcAddr, fromAddr.Hex())
	if err != nil {
		return nil, fmt.Errorf("PendingNonceAt err:%v", err)
	}

	var tx *types.Transaction
	switch common.HexToAddress(to).Hex() {
	case common.HexToAddress("").Hex():
		gasPrice, gasLimit, err := EstimateGas(ctx, client, fromAddr, nil, amount, data)
		if err != nil {
			return nil, err
		}
		tx = types.NewContractCreation(nonce, amount, gasLimit, gasPrice, data)
	default:
		gasPrice, gasLimit, err := EstimateGas(ctx, client, fromAddr, &toAddr, amount, data)
		if err != nil {
			return nil, err
		}
		fmt.Println("gasPrice and gasLimit", gasPrice, gasLimit)
		tx = types.NewTransaction(nonce, common.HexToAddress(to), amount, gasLimit, gasPrice, data)
	}
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("NetworkID err:%v", err)
	}
	sigTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, fmt.Errorf("SignTx err:%v", err)
	}
	return sigTx, nil
}

// 获取 nonce
func PendingNonceAt(rpc, address string) (uint64, error) {
	var nonce string
	version := BaseResp{Result: &nonce}
	method := fmt.Sprintf(`{"jsonrpc":"2.0","method":"parity_nextNonce","params":["%s"],"id":1}`, address)
	if err := version.Request(rpc, method); err != nil || version.Error.Code != 0 {
		return 0, err
	}
	n, err := HexToUint64(nonce)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// 获取 fee ,limit
func EstimateGas(ctx context.Context, client *ethclient.Client, from common.Address, to *common.Address, value *big.Int, input []byte) (fee *big.Int, limit uint64, err error) {
	call := ethereum.CallMsg{From: from, To: to, Value: value, Data: input}
	limit, err = client.EstimateGas(ctx, call)
	if err != nil {
		return nil, 0, fmt.Errorf("EstimateGas err:%v", err)
	}
	fee, err = client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("SuggestGasPrice err:%v", err)
	}
	return fee, limit, nil
}
