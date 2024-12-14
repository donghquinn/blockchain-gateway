package utils

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateNewAccount(password, keystoreDir string) (string, error) {
	privateKey, genErr := crypto.GenerateKey()

	if genErr != nil {
		return "", fmt.Errorf("failed to generate private key: %v", genErr)
	}

	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	if keystoreDir != "" {
		ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

		account, importErr := ks.ImportECDSA(privateKey, password)

		if importErr != nil {
			return "", fmt.Errorf("failed to save key to keystore: %v", importErr)
		}

		log.Printf("Keystore file created: %s\n", account.URL.Path)
	}

	return publicAddress, nil
}
