package example

import (
	"log"
	"math/big"

	"org.donghyusn.com/chain/collector/utils"
	"org.donghyusn.com/chain/collector/web3"
)

func GetTransactionCount() {
	instance := web3.GetWeb3Instance(
		"example chain",
		"example network",
	)

	va := big.NewInt(int64(20396234))
	log.Printf("BigIntData: %v, data: %v", utils.BigIntToString(va, 16), va)
	txCount, _ := instance.GetTransactionCountInBlock(va)

	log.Printf("Tx Count: %v", txCount)
}
