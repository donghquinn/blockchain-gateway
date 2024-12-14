package main

import (
	"time"

	"github.com/joho/godotenv"
	"org.donghyusn.com/chain/collector/database"
	"org.donghyusn.com/chain/collector/example"
)

func main() {
	godotenv.Load(".env")
	database.InitializeDB()

	example.CreateAccountExample()

	time.Sleep(time.Second * 5)

	example.LoadAccountExample()
}
