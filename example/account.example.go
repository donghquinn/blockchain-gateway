package example

import (
	"log"
	"time"

	"org.donghyusn.com/chain/collector/web3"
)

func Web3Example() {
	CreateNetworkExample()

	CreateAccountExample()

	time.Sleep(time.Second * 3)

	LoadAccountExample()

	GetBalance()
	GetNonce()
}

func CreateAccountExample() {
	address, _ := web3.CreateAccount("example account", "example_password")
	log.Printf("Address: %s", address)
}

func LoadAccountExample() {
	_, account, _ := web3.GetAccount("example account", "example_password")

	log.Printf("Account: %s", account)
}

func CreateNetworkExample() {
	web3.CreateNewNetwork("example_network", "https://example-network.com")
}

func GetBalance() {
	network, _ := web3.GetRpcUrlByNetworkName("example_network")

	web3Instance := web3.GetWeb3Instance("example chain", network.NetworkUrl)
	_, account, _ := web3.GetAccount("example account", "example_password")

	balance, _ := web3Instance.GetBalance(account)

	log.Printf("Account: %s, Balance: %d", account, balance)

}

func GetNonce() {
	network, _ := web3.GetRpcUrlByNetworkName("example_network")

	web3Instance := web3.GetWeb3Instance("example chain", network.NetworkUrl)
	_, account, _ := web3.GetAccount("example account", "example_password")

	nonce, _ := web3Instance.GetTxCount(account)

	log.Printf("Account: %s,  Nonce: %d", account, nonce)

}
