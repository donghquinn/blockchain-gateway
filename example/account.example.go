package example

import (
	"log"

	"org.donghyusn.com/chain/collector/web3"
)

func CreateAccountExample() {
	address, _ := web3.CreateAccount("example account", "example_password")
	log.Printf("Address: %s", address)
}

func LoadAccountExample() {
	_, account, _ := web3.GetAccount("example account", "example_password")

	// account, _ := web3.CreateAccount("example_password")

	log.Printf("Account: %s", account)
}

func GetBalance() {
	network, _ := web3.GetRpcUrlByNetworkName("example_network")

	web3Instance := web3.GetWeb3Instance(network.NetworkUrl)
	_, account, _ := web3.GetAccount("example account", "example_password")

	balance, _ := web3Instance.GetBalance(account)

	log.Printf("Account: %s, Balance: %d", account, balance)

}

func GetNonce() {
	network, _ := web3.GetRpcUrlByNetworkName("example_network")

	web3Instance := web3.GetWeb3Instance(network.NetworkUrl)
	_, account, _ := web3.GetAccount("example account", "example_password")

	nonce, _ := web3Instance.GetTxCount(account)

	log.Printf("Account: %s,  Nonce: %d", account, nonce)

}
