package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // THis dep is not used... why imported?
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
		internal_id INTEGER PRIMARY KEY AUTOINCREMENT,
		id TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		first_name TEXT NOT NULL,
		surname TEXT,
		password TEXT NOT NULL,
		company_id TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS companies (
		internal_id INTEGER PRIMARY KEY AUTOINCREMENT,
		id TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL,
		website TEXT,
		email TEXT,
		owning_user TEXT NOT NULL,
		plan_tier INTEGER
	);
	CREATE TABLE IF NOT EXISTS endpoints (
		internal_id INTEGER PRIMARY KEY AUTOINCREMENT,
		id TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL,
		credential_id TEXT NOT NULL,
		company_id TEXT NOT NULL,
		endpoint_type TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS company_permissions (
		internal_id INTEGER PRIMARY KEY AUTOINCREMENT,
		id TEXT NOT NULL UNIQUE,
		company_id TEXT NOT NULL,
		permission_id TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS endpoint_types (
		internal_id INTEGER PRIMARY KEY AUTOINCREMENT,
		id TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS permissions (
		internal_id INTEGER PRIMARY KEY AUTOINCREMENT,
		id TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS test_results (
		internal_id INTEGER PRIMARY KEY AUTOINCREMENT,
		id TEXT NOT NULL UNIQUE,
		test_plan_id TEXT NOT NULL,
		test_result TEXT
	);
	CREATE TABLE IF NOT EXISTS test_suites (
		internal_id INTEGER PRIMARY KEY AUTOINCREMENT,
		id TEXT NOT NULL UNIQUE,
		endpoint_id TEXT NOT NULL,
		name TEXT NOT NULL,
		most_recent_result TEXT
	);
	CREATE TABLE IF NOT EXISTS tests (
		internal_id INTEGER PRIMARY KEY AUTOINCREMENT,
		id TEXT NOT NULL UNIQUE,
		test_plan_id TEXT NOT NULL,
		test_data TEXT
	);
	CREATE TABLE IF NOT EXISTS user_permissions (
		internal_id INTEGER PRIMARY KEY AUTOINCREMENT,
		id TEXT NOT NULL UNIQUE,
		user_id TEXT NOT NULL,
		permission_id TEXT NOT NULL
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to create tables:", err)
	}
}
