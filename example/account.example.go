package example

import (
	"log"

	"org.donghyusn.com/chain/collector/web3"
)

func CreateAccountExample() {
	address, _ := web3.CreateAccount("example_password")
	log.Printf("Address: %s", address)
}

func LoadAccountExample() {
	web3Instance := web3.GetWeb3Instance("example rpc url")

	// account, _ := web3.CreateAccount("example_password")

	_, account, _ := web3.GetAccount("example_password")

	balance, _ := web3Instance.GetBalance(account)

	nonce, _ := web3Instance.GetTxCount(account)

	log.Printf("Account: %s, Balance: %d, Nonce: %d", account, balance, nonce)
}
