package database

var CreateRpcUrlTable = `
	CREATE TABLE IF NOT EXISTS network_table (
		network_seq			INTEGER										PRIMARY KEY AUTOINCREMENT,
		network_name		TEXT										NOT NULL,
		network_url			TEXT										NOT NULL,
		network_status		INTEGER										NOT NULL DEFAULT 1,
		created_date		DATETIME									NOT NULL DEFAULT 	CURRENT_TIMESTAMP
	)
`

var CreateAccountTable = `
	CREATE TABLE IF NOT EXISTS account_table (
		account_seq			INTEGER 									PRIMARY KEY AUTOINCREMENT,
		account_address 	TEXT										NOT NULL,
		account_password	TEXT										NOT NULL,
		account_type 		TEXT CHECK(kube_type IN ('WEB3'))		 	NOT NULL DEFAULT 'WEB3',
		account_status		INTEGER										NOT NULL DEFAULT 1,
		created_date 		DATETIME 									NOT NULL DEFAULT 	CURRENT_TIMESTAMP
	)
`

var CreatePrivateKeyTable = `
		CREATE TABLE IF NOT EXISTS private_key_table (
		private_key_seq		INTEGER 		PRIMARY KEY AUTOINCREMENT,
		account_seq			INTEGER 		NOT NULL REFERENCES account_table (account_seq) ON DELETE CASCADE,
		private_key_dir		TEXT			NOT NULL,
		created_date 		DATETIME 		NOT NULL DEFAULT 	CURRENT_TIMESTAMP
	)
`

var CreateNonceTable = `
		CREATE TABLE IF NOT EXISTS nonce_table (
		nonce_seq		INTEGER 		PRIMARY KEY AUTOINCREMENT,
		account_seq		INTEGER 		NOT NULL REFERENCES account_table (account_seq) ON DELETE CASCADE,
		nonce			INTEGER			NOT NULL,
		created_date 	DATETIME 		NOT NULL DEFAULT 	CURRENT_TIMESTAMP
	)
`

var CreateBalanceTable = `
	CREATE TABLE IF NOT EXISTS balance_table (
		balance_seq			INTEGER 		PRIMARY KEY AUTOINCREMENT,
		account_seq			INTEGER 		NOT NULL REFERENCES account_table (account_seq) ON DELETE CASCADE,
		balance				INTEGER			NOT NULL DEFAULT 	0,
		created_date 		DATETIME 		NOT NULL DEFAULT 	CURRENT_TIMESTAMP
	)
`

var CreateTransactionTable = `
	CREATE TABLE IF NOT EXISTS transaction_table (
		transaction_seq		INTEGER							PRIMARY KEY AUTOINCREMENT,
		network_seq			INTEGER							NOT NULL REFERENCES network_table	(network_seq),
		account_seq			INTEGER							NOT NULL REFERENCES account_table	(account_seq),
		to_address			TEXT							NOT NULL,
		transaction_type	TEXT CHECK('RAW', 'CONTRACT')	NOT NULL,
		transaction_hash	TEXT								NULL,
		created_date 		DATETIME 						NOT NULL DEFAULT 	CURRENT_TIMESTAMP
	)
`

var CreateTableTransactionQueue = []string{
	CreateRpcUrlTable,
	CreateAccountTable, CreatePrivateKeyTable, CreateBalanceTable, CreateNonceTable,
	CreateTransactionTable,
}
