package app

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"fmt"
	"log"
)

// TODO: Should store config in yaml or other config file without being tracked by git
const (
	userName     = "root"
	password     = "root"
	network      = "tcp"
	server       = "127.0.0.1"
	port         = 3306
	database     = "cu"
	databaseType = "mysql"
)

const (
	createDataBaseSQL = `CREATE DATABASE IF NOT EXISTS cu;`
	useDataBaseSQL    = `USE cu;`
	createCuTableSQL  = `
	CREATE TABLE IF NOT EXISTS data(
		id INT KEY NOT NULL,
		lat DECIMAL NOT NULL,
		lng DECIMAL NOT NULL,
		date_added BIGINT(20) NOT NULL
	);`
)

type Database struct{}

// NewDB returns the mysql db with initialize database and table for cu.
func NewDB() *sql.DB {
	DBConfig := fmt.Sprintf("%s:%s@%s(%s:%d)/", userName, password, network, server, port)
	db, err := sql.Open(databaseType, DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := createDatabase(db); err != nil {
		log.Fatal(err)
	}

	if err := useDatabase(db); err != nil {
		log.Fatal(err)
	}

	if err := createTable(db); err != nil {
		log.Fatal(err)
	}

	return db
}

func createDatabase(db *sql.DB) error {
	if _, err := db.Exec(createDataBaseSQL); err != nil {
		return fmt.Errorf("Create database failed: %w", err)
	}
	return nil
}

func useDatabase(db *sql.DB) error {
	if _, err := db.Exec(useDataBaseSQL); err != nil {
		return fmt.Errorf("Use database failed: %w", err)
	}
	return nil
}

func createTable(db *sql.DB) error {
	if _, err := db.Exec(createCuTableSQL); err != nil {
		return fmt.Errorf("Create table failed: %w", err)
	}
	return nil
}
