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

	_, account, _ := web3.GetAccount("keystore/UTC--2024-12-14T03-13-40.515685000Z--cd5172dd6fa457c3bde87171031d48f00448e8d9", "example_password")

	balance, _ := web3Instance.GetBalance(account)

	nonce, _ := web3Instance.GetTxCount(account)

	log.Printf("Account: %s, Balance: %d, Nonce: %d", account, balance, nonce)
}
