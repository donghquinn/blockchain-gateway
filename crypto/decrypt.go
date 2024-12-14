package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"

	"org.donghyusn.com/chain/collector/config"
)

// 복호화 요청 #회원가입 #로그인
func DecryptString(cipherText string) (string, error) {
	if strings.TrimSpace(cipherText) == "" {
		return cipherText, nil
	}

	globalConfig := config.GlobalConfig

	decodedCipherText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(globalConfig.AesKey))
	if err != nil {
		return "", err
	}

	decrypter := cipher.NewCBCDecrypter(block, []byte(globalConfig.AesIv))
	plainText := make([]byte, len(decodedCipherText))

	decrypter.CryptBlocks(plainText, decodedCipherText)
	trimmedPlainText := trimPKCS5(plainText)

	return string(trimmedPlainText), nil
}

// Unpadding PKCS
func trimPKCS5(text []byte) []byte {
	padding := text[len(text)-1]
	return text[:len(text)-int(padding)]
}
