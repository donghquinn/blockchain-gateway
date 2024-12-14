package constant

// https://ethereum.org/en/developers/docs/apis/json-rpc/#top
var MethodConstant = map[string]string{
	// acount
	"ACCOUNT_LIST": "eth_accounts",
	"BALANCE":      "eth_getBalance",
	"NONCE":        "eth_getTransactionCount",

	// chain
	"BLOCK_NUMBER": "eth_blockNumber",

	// Transaction
	"SEND_TX":     "eth_sendTransaction",
	"SEND_RAW_TX": "eth_sendRawTransaction",
}
