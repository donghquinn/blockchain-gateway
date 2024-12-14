package config

import (
	"fmt"
	"os"
)

type GlobalConfigStruct struct {
	AppPort   string
	AppHost   string
	AesIv     string
	AesKey    string
	SecretKey string
	// JwtKey    string
}

var GlobalConfig GlobalConfigStruct

func SetGlobalConfig() {
	GlobalConfig.AppPort = os.Getenv("APP_PORT")
	GlobalConfig.AppHost = fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	GlobalConfig.AesIv = os.Getenv("AES_IV")
	GlobalConfig.AesKey = os.Getenv("AES_KEY")
	// GlobalConfig.JwtKey = os.Getenv("JWT_KEY")
	GlobalConfig.SecretKey = os.Getenv("SECRET_KEY")
}
