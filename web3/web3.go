package web3

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"

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

// Web3 Block Number
func GetBlockNumber(rpcUrl string) (*big.Int, error) {
	constant := constant.MethodConstant

	request := Web3RpcRequest{
		Jsonrpc: "2.0",
		Method:  constant["BLOCK_NUMBER"],
		Params:  []interface{}{"latest", true}, // Get the latest block with full details
		ID:      1,
	}

	res, postErr := utils.Post(rpcUrl, request)

	if postErr != nil {
		return nil, postErr
	}

	var response Web3RpcResponse

	parseErr := json.Unmarshal(res, &response)

	if parseErr != nil {
		log.Printf("[WEB3] Unmarshal Response Error: %v", parseErr)
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
