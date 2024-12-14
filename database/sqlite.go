package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"org.donghyusn.com/chain/collector/constant"
	"org.donghyusn.com/chain/collector/utils"
)

type DataBaseSqlite struct {
	*sql.DB
}

func InitializeDB() error {
	// DB 저장 경로 생성
	utils.CreateDir(constant.DatabaseDir)
	db, err := sql.Open("sqlite3", constant.DatabaseDir+constant.DatabaseFileNAme)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	dbInstance := DataBaseSqlite{db}

	err = dbInstance.InsertMultipleQuery(CreateTableTransactionQueue)

	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	return nil
}

func GetConnection() (DataBaseSqlite, error) {
	db, err := sql.Open("sqlite3", constant.DatabaseDir+constant.DatabaseFileNAme)

	if err != nil {
		return DataBaseSqlite{}, fmt.Errorf("failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return DataBaseSqlite{}, fmt.Errorf("failed to connect to database: %v", err)
	}

	dbInstance := DataBaseSqlite{db}

	return dbInstance, nil
}

func (db *DataBaseSqlite) SelectMultipleRows(queryString string, param ...string) (*sql.Rows, error) {
	var queryList []interface{}

	for _, query := range param {
		queryList = append(queryList, query)
	}

	queryResult, queryErr := db.Query(queryString, queryList...)

	if queryErr != nil {
		log.Printf("failed to execute query: %v", queryErr)
		return nil, queryErr
	}

	defer db.Close()

	return queryResult, nil
}

func (db *DataBaseSqlite) SelectOneRow(queryString string, param ...string) (*sql.Row, error) {
	var paramList []interface{}

	for _, argument := range param {
		paramList = append(paramList, argument)
	}

	queryResult := db.QueryRow(queryString, paramList...)

	if queryResult.Err() != nil {
		log.Printf("failed to execute query: %v", queryResult.Err())
		return nil, queryResult.Err()
	}

	defer db.Close()

	return queryResult, nil
}

func (db *DataBaseSqlite) InsertQuery(queryString string, params ...string) (int64, error) {

	var paramList []interface{}

	for _, param := range params {
		paramList = append(paramList, param)
	}

	result, execErr := db.Exec(queryString, paramList...)

	if execErr != nil {
		log.Printf("failed to execute Insert query: %v", execErr)
		return -1, execErr
	}

	defer db.Close()

	insertId, insertIdErr := result.LastInsertId()

	if insertIdErr != nil {
		log.Printf("Failed to get last insert ID: %v", insertIdErr)
	}

	return insertId, nil
}

func (db *DataBaseSqlite) InsertMultipleQuery(queryString []string) error {
	tx, err := db.Begin()

	if err != nil {
		log.Printf("failed to begin transaction: %v", err)
		return err
	}

	defer tx.Rollback()

	for _, query := range queryString {
		_, execErr := tx.Exec(query)

		if execErr != nil {
			log.Printf("failed to execute query: %v", execErr)
		}
	}

	commitErr := tx.Commit()

	if commitErr != nil {
		log.Printf("failed to commit transaction: %v", commitErr)
		return commitErr
	}

	return nil
}
