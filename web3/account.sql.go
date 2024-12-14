package web3

var InsertAccountQuery = `
	INSERT INTO account_table
		(account_address, account_password)
	VALUES
		(?, ?)
`

var InsertKeyStore = `
	INSERT INTO private_key_table
		(account_seq, private_key_dir)
	VALUES
		(?, ?)
`

var SelectAccountKeyStore = `
	SELECT pk.private_key_dir AS keystore
	FROM private_key_table pk
	LEFT JOIN account_table a ON a.account_seq = pk.account_seq
	WHERE a.account_password = ?
`
