package example

import (
	"log"
	"math/big"

	"org.donghyusn.com/chain/collector/web3"
)

func GetTransactionCount() {
	instance := web3.GetWeb3Instance(
		"example chain",
		"https://stylish-wispy-gas.ethereum-sepolia.quiknode.pro/734558df59493e4eb5e64a5809095edd60744514",
	)

	va := big.NewInt(int64(20396234))

	txCount, _ := instance.GetTransactionCountInBlock(va)

	log.Printf("Tx Count: %v", txCount)
}
