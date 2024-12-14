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
	crypt "org.donghyusn.com/chain/collector/crypto"
	"org.donghyusn.com/chain/collector/database"
	"org.donghyusn.com/chain/collector/utils"
)

func CreateAccount(accountName string, password string) (string, error) {
	encodedPassword, encodePasswordErr := crypt.EncryptHashPassword(password)

	if encodePasswordErr != nil {
		return "", encodePasswordErr
	}

	address, keystore, createErr := utils.GenerateNewAccount(password, constant.PrivateKeyStoreDir)

	if createErr != nil {
		return "", createErr
	}

	go insertAccountData(address, accountName, encodedPassword, keystore)

	return address, nil
}

func insertAccountData(address string, accountName string, encodedPassword string, keystore string) error {
	accountSeq, insertErr := insertAccount(address, accountName, encodedPassword)

	if insertErr != nil {
		log.Printf("[CREATE ACCOUNT] Account insertion failed: %v", insertErr)
		return insertErr
	}

	insertKeyErr := insertKeystore(fmt.Sprintf("%d", accountSeq), keystore)

	if insertKeyErr != nil {
		log.Printf("[CREATE ACCOUNT] Keystore insertion failed: %v", insertKeyErr)
		return insertKeyErr
	}

	log.Printf("[CREATE ACCOUNT] Keystore path inserted: %s", keystore)

	return nil
}

func insertAccount(address string, accountName string, encodedPassword string) (int64, error) {
	dbCon, err := database.GetConnection()
	if err != nil {
		return 0, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Execute the insert query
	accountSeq, err := dbCon.InsertQuery(InsertAccountQuery, address, accountName, encodedPassword)
	if err != nil {
		return 0, fmt.Errorf("failed to execute account insert query: %v", err)
	}

	return accountSeq, nil
}

func insertKeystore(accountSeq string, keystorePath string) error {
	dbCon, err := database.GetConnection()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Execute the insert query
	_, err = dbCon.InsertQuery(InsertKeyStore, accountSeq, keystorePath)
	if err != nil {
		return fmt.Errorf("failed to execute keystore insert query: %v", err)
	}

	return nil
}

// ================= GET ACCOUNT  =================
// Get account and private key from database
func GetAccount(accountName string, password string) (*ecdsa.PrivateKey, string, error) {
	dbCon, dbErr := database.GetConnection()

	if dbErr != nil {
		return nil, "", dbErr
	}

	result, queryErr := dbCon.SelectOneRow(SelectAccountKeyStore, accountName)

	if queryErr != nil {
		return nil, "", queryErr
	}

	var queryResult AccountPrivateKeyDir

	if scanErr := result.Scan(&queryResult.KeystoreDir, &queryResult.Password); scanErr != nil {
		log.Printf("[WEB3] Scan Account Private Key Dir Error: %v", scanErr)
		return nil, "", scanErr
	}

	isMatch, compareErr := crypt.PasswordCompare(queryResult.Password, password)

	if compareErr != nil {
		log.Printf("[WEB3] Password compare Error :%v", compareErr)
		return nil, "", compareErr
	}

	if !isMatch {
		log.Println("[WEB3] Password does not match")
		return nil, "", fmt.Errorf("password does not match")
	}

	privateKey, address, accountErr := LoadAccountFromKeystore(queryResult.KeystoreDir, password)

	if accountErr != nil {
		return nil, "", accountErr
	}

	return privateKey, address, accountErr
}

// ================= PRIVATE KEY =================
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
