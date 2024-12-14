package web3

var InsertAccountQuery = `
	INSERT INTO account_table
		(account_address, account_password)
	VALUES
		(?, ?)
`
