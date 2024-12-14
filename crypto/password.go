package crypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// 평문의 패스워드에서 단방향 해시를 생성한다
func EncryptHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Printf("[ENCODE] Encode Hash Password: %v", err)

		return "", err
	}

	return string(bytes), nil
}

// 패스워드 비교
func PasswordCompare(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Printf("[COMPARE] Compare Password Error: %v", err)

		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, err
		}
		return false, err
	}
	return true, nil
}
