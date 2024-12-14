# Blockchain Client

## Database

* sqlite3
    * save account, network, and other data

## Used Dependencies

* sqlite3
    * manage data

```bash
    go get -u github.com/mattn/go-sqlite3
```

* go-ethereum
    * create account and other utilization
    * save keystore

```bash
    go get -u github.com/ethereum/go-ethereum
    go get -u github.com/ethereum/go-ethereum/accounts/keystore
```

* bcrypt

```bash
    go get -u golang.org/x/crypto/bcrypt
```

* godotenv

```bash
    go get -u github.com/joho/godotenv
```

---

## Network
* network name
    * network identification name
* network url
    * rpc url for connecting network/node
* network type
    * Support only web3 so far
    * Planning to support other networks
* network status
    * 1 means active network
    * 2 means deactive network
    * 0 means deleted network


## Web3

### Account

* account address
    * account address encoded in hex
* account name
    * account address identification name
* account password
    * account password generating account.
    * save as bycrypt with 10 rounds
* account type
    * Support only web3 so far
    * Planning to support other networks
* account status
    * 1 means active account
    * 2 means deactive account
    * 0 means deleted account


### Private Key

* account_seq
    * account foreign key which private key belongs to
* private_key_dir
    * key store directory
    * save as absolute path
