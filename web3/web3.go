package web3

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"org.donghyusn.com/chain/collector/constant"
	"org.donghyusn.com/chain/collector/utils"
)

type Web3RpcRequest struct {
	// RpcUrl  string        `json:"rpcUrl"`
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type Web3RpcResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   *RPCError       `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Web3Instance struct {
	RpcUrl string
}

type TransactionParams struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas,omitempty"`
	GasPrice string `json:"gasPrice,omitempty"`
	Value    string `json:"value,omitempty"`
	Input    string `json:"input"`
}

func GetWeb3Instance(rpcUrl string) Web3Instance {
	return Web3Instance{
		RpcUrl: rpcUrl,
	}
}

// ======================= ACCOUNT =======================
// Web3 Block Number
func (instance *Web3Instance) GetBlockNumber() (*big.Int, error) {
	constant := constant.MethodConstant

	request := Web3RpcRequest{
		Jsonrpc: "2.0",
		Method:  constant["BLOCK_NUMBER"],
		Params:  []interface{}{"latest", true}, // Get the latest block with full details
		ID:      1,
	}

	res, postErr := utils.Post(instance.RpcUrl, request)

	if postErr != nil {
		return nil, postErr
	}

	var response Web3RpcResponse

	parseErr := json.Unmarshal(res, &response)

	if parseErr != nil {
		log.Printf("[WEB3] Unmarshal Block Number Response Error: %v", parseErr)
		return nil, parseErr
	}

	if response.Error != nil {
		log.Printf("[WEB3] Node RPC Response: Code: %d, Message: %s", response.Error.Code, response.Error.Message)
		return nil, fmt.Errorf("%s", response.Error.Message)
	}

	var blockNumberHex string
	unmarshalErr := json.Unmarshal(response.Result, &blockNumberHex)

	if unmarshalErr != nil {
		log.Printf("[WEB3] Unmarshal Block Number Error: %v", unmarshalErr)
		return nil, unmarshalErr
	}

	blockNumber := new(big.Int)
	blockNumber.SetString(blockNumberHex[2:], 16) // Skip "0x" prefix

	return blockNumber, nil
}

// Check Balance
func (instance *Web3Instance) GetBalance(address string) (*big.Int, error) {
	constant := constant.MethodConstant

	request := Web3RpcRequest{
		Jsonrpc: "2.0",
		Method:  constant["BALANCE"],
		Params:  []interface{}{address, "latest"},
		ID:      1,
	}

	res, postErr := utils.Post(instance.RpcUrl, request)

	if postErr != nil {
		return nil, postErr
	}

	var response Web3RpcResponse

	parseErr := json.Unmarshal(res, &response)

	if parseErr != nil {
		log.Printf("[WEB3] Unmarshal Account Balance Response Error: %v", parseErr)
		return nil, parseErr
	}

	if response.Error != nil {
		log.Printf("[WEB3] Node RPC Response: Code: %d, Message: %s", response.Error.Code, response.Error.Message)
		return nil, fmt.Errorf("%s", response.Error.Message)
	}

	var blockNumberHex string
	unmarshalErr := json.Unmarshal(response.Result, &blockNumberHex)

	if unmarshalErr != nil {
		log.Printf("[WEB3] Unmarshal Balance Error: %v", unmarshalErr)
		return nil, unmarshalErr
	}

	blockNumber := new(big.Int)
	blockNumber.SetString(blockNumberHex[2:], 16) // Skip "0x" prefix

	return blockNumber, nil
}

// Get Nonce
func (instance *Web3Instance) GetTxCount(address string) (*big.Int, error) {
	constant := constant.MethodConstant

	request := Web3RpcRequest{
		Jsonrpc: "2.0",
		Method:  constant["NONCE"],
		Params:  []interface{}{address, "latest"},
		ID:      1,
	}

	res, postErr := utils.Post(instance.RpcUrl, request)

	if postErr != nil {
		return nil, postErr
	}

	var response Web3RpcResponse

	parseErr := json.Unmarshal(res, &response)

	if parseErr != nil {
		log.Printf("[WEB3] Unmarshal Transaction Count Response Error: %v", parseErr)
		return nil, parseErr
	}

	if response.Error != nil {
		log.Printf("[WEB3] Node RPC Response: Code: %d, Message: %s", response.Error.Code, response.Error.Message)
		return nil, fmt.Errorf("%s", response.Error.Message)
	}

	var txNonce string
	unmarshalErr := json.Unmarshal(response.Result, &txNonce)

	if unmarshalErr != nil {
		log.Printf("[WEB3] Unmarshal Nonce Error: %v", unmarshalErr)
		return nil, unmarshalErr
	}

	blockNumber := new(big.Int)
	blockNumber.SetString(txNonce[2:], 16) // Skip "0x" prefix

	return blockNumber, nil
}

// ======================= CHAIN =======================
// Get Transaction Number in a Block
func (instance *Web3Instance) GetTransactionCountInBlock(blockNumber *big.Int) (*big.Int, error) {

	constant := constant.MethodConstant

	request := Web3RpcRequest{
		Jsonrpc: "2.0",
		Method:  constant["BLOCK_TX_COUNT"],
		Params:  []interface{}{utils.BigIntToString(blockNumber), "latest"},
		ID:      1,
	}

	res, postErr := utils.Post(instance.RpcUrl, request)

	if postErr != nil {
		return nil, postErr
	}

	var response Web3RpcResponse

	parseErr := json.Unmarshal(res, &response)

	if parseErr != nil {
		log.Printf("[WEB3] Unmarshal Transaction Count Response Error: %v", parseErr)
		return nil, parseErr
	}

	if response.Error != nil {
		log.Printf("[WEB3] Node RPC Response: Code: %d, Message: %s", response.Error.Code, response.Error.Message)
		return nil, fmt.Errorf("%s", response.Error.Message)
	}

	var blockTxResponse string

	unmarshalErr := json.Unmarshal(response.Result, &blockTxResponse)

	if unmarshalErr != nil {
		log.Printf("[WEB3] Unmarshal Nonce Error: %v", unmarshalErr)
		return nil, unmarshalErr
	}

	bigInt := utils.StringToBigInt(blockTxResponse)

	return bigInt, nil
}

// ======================= TRANSACTION =======================
// Send Raw Tx
func (instance *Web3Instance) SendRawTransaction(address string, privateKey *ecdsa.PrivateKey, toAddress common.Address, value *big.Int, gasLimit uint64, gasPrice *big.Int, chainID *big.Int) (string, error) {

	nonce, nonceErr := instance.GetTxCount(address)

	if nonceErr != nil {
		return "", nonceErr
	}

	signedTx, signErr := SignTransaction(privateKey, toAddress, value, gasLimit, gasPrice, nonce.Uint64(), chainID)

	if signErr != nil {
		return "", signErr
	}

	constant := constant.MethodConstant

	request := Web3RpcRequest{
		Jsonrpc: "2.0",
		Method:  constant["SEND_RAW_TX"],
		Params:  []interface{}{signedTx},
		ID:      1,
	}

	res, postErr := utils.Post(instance.RpcUrl, request)

	if postErr != nil {
		return "", postErr
	}

	var response Web3RpcResponse

	parseErr := json.Unmarshal(res, &response)

	if parseErr != nil {
		log.Printf("[WEB3] Unmarshal Send Raw Transaction Response Error: %v", parseErr)
		return "", parseErr
	}

	if response.Error != nil {
		log.Printf("[WEB3] Node RPC Response: Code: %d, Message: %s", response.Error.Code, response.Error.Message)
		return "", fmt.Errorf("%s", response.Error.Message)
	}

	var transactionHash string

	unmarshalErr := json.Unmarshal(response.Result, &transactionHash)

	if unmarshalErr != nil {
		log.Printf("[WEB3] Unmarshal Raw Tx Receipt Error: %v", unmarshalErr)
		return "", unmarshalErr
	}

	return transactionHash, nil
}

// Send Contract Tx
// func (instance *Web3Instance) SendTransaction(keystorePath string, password string, toAddress common.Address, value *big.Int, gasLimit uint64, gasPrice *big.Int, nonce uint64, chainID *big.Int) (string, error) {
// 	privateKey, account, accountErr := LoadAccountFromKeystore(keystorePath, password)

// 	if accountErr != nil {
// 		return "", accountErr
// 	}

// 	constant := constant.MethodConstant

// 	request := Web3RpcRequest{
// 		Jsonrpc: "2.0",
// 		Method:  constant["SEND_TX"],
// 		Params: []interface{}{
// 			TransactionParams{
// 				From:  account,
// 				To:    toAddress.Hex(),
// 				Gas:   gasPrice.String(),
// 				Value: value.String(),
// 				Input: inpu,
// 			},
// 		},
// 		ID: 1,
// 	}

// 	res, postErr := utils.Post(instance.RpcUrl, request)

// 	if postErr != nil {
// 		return "", postErr
// 	}

// 	var response Web3RpcResponse

// 	parseErr := json.Unmarshal(res, &response)

// 	if parseErr != nil {
// 		log.Printf("[WEB3] Unmarshal Response Error: %v", parseErr)
// 		return "", parseErr
// 	}

// 	if response.Error != nil {
// 		log.Printf("[WEB3] Node RPC Response: Code: %d, Message: %s", response.Error.Code, response.Error.Message)
// 		return "", fmt.Errorf("%s", response.Error.Message)
// 	}

// 	var transactionHash string

// 	unmarshalErr := json.Unmarshal(response.Result, &transactionHash)

// 	if unmarshalErr != nil {
// 		log.Printf("[WEB3] Unmarshal Send Tx Error: %v", unmarshalErr)
// 		return "", unmarshalErr
// 	}

// 	return transactionHash, nil
// }
