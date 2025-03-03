package example

import (
	"log"
	"math/big"

	"org.donghyusn.com/chain/collector/config"
	"org.donghyusn.com/chain/collector/web3"
)

func GetTransactionCount() {
	globalConfig := config.GlobalConfig

	instance := web3.GetWeb3Instance(
		"example chain",
		globalConfig.RpcUrl,
	)

	va := big.NewInt(int64(20396234))

	txCount, _ := instance.GetTransactionCountInBlock(va)

	log.Printf("Tx Count: %v", txCount)
}
