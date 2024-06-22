package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	createTables()
}

func createTables() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		company_id INTEGER
	);
	CREATE TABLE IF NOT EXISTS companies (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		bill_status TEXT
	);
	CREATE TABLE IF NOT EXISTS endpoints (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT,
		store_name TEXT,
		token TEXT,
		company_id INTEGER
	);
	CREATE TABLE IF NOT EXISTS tests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		data TEXT,
		test_plan_id INTEGER
	);
	CREATE TABLE IF NOT EXISTS test_plans (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		endpoint_id INTEGER,
		company_id INTEGER
	);
	CREATE TABLE IF NOT EXISTS test_results (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		test_id INTEGER,
		status TEXT
	);
	CREATE TABLE IF NOT EXISTS test_plan_results (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		test_plan_id INTEGER,
		status TEXT
	);`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to create tables:", err)
	}
}
