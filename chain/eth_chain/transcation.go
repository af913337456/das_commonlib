package eth_chain

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func NewTranscation(ctx context.Context, client *ethclient.Client, rpcAddr, private, to string, amount *big.Int, data []byte) (tx *types.Transaction, err error) {
	return nil, nil
}
