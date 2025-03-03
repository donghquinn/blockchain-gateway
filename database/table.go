package database

var CreateRpcUrlTable = `
	CREATE TABLE IF NOT EXISTS network_table (
		network_seq			INTEGER										PRIMARY KEY AUTOINCREMENT,
		network_name		TEXT										NOT NULL,
		network_url			TEXT										NOT NULL,
		network_type		TEXT CHECK(network_type IN ('WEB3'))		NOT NULL DEFAULT 'WEB3',
		network_status		INTEGER										NOT NULL DEFAULT 1,
		created_date		DATETIME									NOT NULL DEFAULT 	CURRENT_TIMESTAMP
	)
`
var CreateRpcUrlIndex = `
	CREATE INDEX rpc_idx ON network_table (network_name, network_type, network_status)
`

var CreateAccountTable = `
	CREATE TABLE IF NOT EXISTS account_table (
		account_seq			INTEGER 									PRIMARY KEY AUTOINCREMENT,
		account_address 	TEXT										NOT NULL UNIQUE,
		account_name		TEXT										NOT NULL UNIQUE,
		account_password	TEXT										NOT NULL,
		account_type 		TEXT CHECK(account_type IN ('WEB3'))		NOT NULL DEFAULT 'WEB3',
		account_status		INTEGER										NOT NULL DEFAULT 1,
		created_date 		DATETIME 									NOT NULL DEFAULT 	CURRENT_TIMESTAMP
	)
`

var CreateAccountIndex = `
	CREATE INDEX account_idx ON account_table (account_status, account_type)
`

var CreatePrivateKeyTable = `
	CREATE TABLE IF NOT EXISTS private_key_table (
		private_key_seq		INTEGER 		PRIMARY KEY AUTOINCREMENT,
		account_seq			INTEGER 		NOT NULL REFERENCES account_table (account_seq) ON DELETE CASCADE,
		private_key_dir		TEXT			NOT NULL,
		private_key_status	INTEGER			NOT NULL DEFAULT 1,
		created_date 		DATETIME 		NOT NULL DEFAULT 	CURRENT_TIMESTAMP
	)
`

var CreatePrivateKeyIndex = `
	CREATE INDEX private_key_idx ON private_key_idx(private_key_status)
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
		transaction_seq		INTEGER													PRIMARY KEY AUTOINCREMENT,
		account_seq			INTEGER													NOT NULL REFERENCES account_table (account_seq),
		network_seq			INTEGER													NOT NULL REFERENCES network_table (network_seq),
		transaction_status  INTEGER													NOT NULL DEFAULT 0,
		transaction_type	TEXT CHECK(transaction_type IN ('RAW', 'CONTRACT'))		NOT NULL,
		transaction_hash	TEXT														NULL UNIQUE,
		created_date 		DATETIME 												NOT NULL DEFAULT 	CURRENT_TIMESTAMP
	)
`

var CreateTransactionIndex = `
	CREATE INDEX transaction_idx ON transaction_table (transaction_status)
`

var CreateTransactionDataTable = `
	CREATE TABLE IF NOT EXISTS transaction_data_table (
		transaction_data_seq	INTEGER			PRIMARY KEY AUTOINCREMENT,
		transaction_seq			INTEGER			NOT NULL REFERENCES transaction_table (transaction_seq),
		from_address			TEXT			NOT NULL,
		to_address				TEXT			NOT NULL,
		value					INTEGER			NOT NULL,
		nonce					INTEGER			NOT NULL,
		gas_price				INTEGER			NOT NULL,
		gas						INTEGER			NOT NULL,
		chain_id				INTEGER			NOT NULL
	)

`

var CreateTableTransactionQueue = []string{
	CreateRpcUrlTable,
	CreateAccountTable, CreatePrivateKeyTable, CreateBalanceTable, CreateNonceTable,
	CreateTransactionTable, CreateTransactionDataTable,

	CreateAccountIndex, CreatePrivateKeyIndex, CreateRpcUrlIndex, CreateTransactionIndex,
}
