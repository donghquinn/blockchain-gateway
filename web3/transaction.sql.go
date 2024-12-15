package web3

var InsertTransaction = `
	INSERT INTO transaction_table
		(network_seq, account_seq, transaction_type)
	VALUES
		((SELECT network_seq FROM network_table WHERE network_name = ?), (SELECT account_seq FROM account_table WHERE account_address = ?), ?)
`

var UpdateTransactionStatus = `
	UPDATE transaction_table
	SET transaction_status = ?
	WHERE transaction_seq = ?
`

var InsertTransactionData = `
	INSERT INTO transaction_data_table
		(transaction_seq, from_address, to_address, value, nonce, gas_price, gas, chain_id)
	VALUES
		(?, ?, ?, ?, ?, ?, ?, ?)
`
