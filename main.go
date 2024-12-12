package main

import (
	"log"

	"org.donghyusn.com/chain/collector/web3"
)

func main() {
	blockNumber, _ := web3.GetBlockNumber("https://stylish-wispy-gas.ethereum-sepolia.quiknode.pro/734558df59493e4eb5e64a5809095edd60744514")

	log.Printf("RES: %v", blockNumber)
}
