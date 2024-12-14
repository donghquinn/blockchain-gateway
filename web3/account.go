package web3

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"org.donghyusn.com/chain/collector/constant"
	"org.donghyusn.com/chain/collector/database"
	"org.donghyusn.com/chain/collector/utils"
)

func CreateAccount(password string) (string, error) {
	dbCon, dbErr := database.GetConnection()

	if dbErr != nil {
		return "", dbErr
	}

	address, createErr := utils.GenerateNewAccount(password, constant.PrivateKeyStoreDir)

	if createErr != nil {
		return "", createErr
	}

	go dbCon.InsertQuery(InsertAccountQuery, address, password)

	return address, nil
}

func GetAccount(keystorePath string, password string) (*ecdsa.PrivateKey, string, error) {
	privateKey, address, accountErr := LoadAccountFromKeystore(keystorePath, password)

	if accountErr != nil {
		return nil, "", accountErr
	}

	return privateKey, address, accountErr
}

// Load Account from private key (hex)
func LoadAccountFromPrivateKey(hexPrivateKey string) (*ecdsa.PrivateKey, common.Address, error) {
	// Decode the private key from hex
	privateKey, err := crypto.HexToECDSA(hexPrivateKey)

	if err != nil {
		log.Printf("[WEB3] Decode private key from hex value Error: %v", err)
		return nil, common.Address{}, fmt.Errorf("invalid private key: %v", err)
	}

	// Derive the public address from the private key
	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	return privateKey, address, nil
}

// Load Account from private key (keystore)
func LoadAccountFromKeystore(keystorePath, password string) (*ecdsa.PrivateKey, string, error) {
	fileContent, err := os.ReadFile(keystorePath)

	if err != nil {
		log.Printf("[WEB3] Load Keystore file Error: %v", err)
		return nil, "", fmt.Errorf("failed to read keystore file: %v", err)
	}

	// Decrypt the keystore to get the private key
	key, err := keystore.DecryptKey(fileContent, password)

	if err != nil {
		log.Printf("[WEB3] Decode private key from keystore Error: %v", err)
		return nil, "", fmt.Errorf("failed to decrypt keystore: %v", err)
	}

	// Get the address
	address := key.Address.Hex()

	return key.PrivateKey, address, nil
}

// Sign Transaction
func SignTransaction(privateKey *ecdsa.PrivateKey, toAddress common.Address, value *big.Int, gasLimit uint64, gasPrice *big.Int, nonce uint64, chainID *big.Int) ([]byte, error) {
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	signedTx, signErr := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

	if signErr != nil {
		log.Printf("[WEB3] Sign new transaction Error: %v", signErr)
		return nil, fmt.Errorf("failed to sign transaction: %v", signErr)
	}

	// Encode the signed transaction into RLP (ready for broadcasting)
	rawTx, encodeErr := rlp.EncodeToBytes(signedTx)

	if encodeErr != nil {
		log.Printf("[WEB3] Encode signed transaction to bytes Error: %v", signErr)
		return nil, fmt.Errorf("failed to encode transaction: %v", encodeErr)
	}

	return rawTx, nil
}
