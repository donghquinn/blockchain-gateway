package main

import (
	"org.donghyusn.com/chain/collector/database"
	"org.donghyusn.com/chain/collector/example"
)

func main() {
	database.InitializeDB()

	example.CreateAccountExample()
	example.LoadAccountExample()
}
