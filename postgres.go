package main

import "database/sql"

func CreateNewDB(db *sql.DB) error {

	CreateTablesQuery := `
		CREATE TABLE IF NOT EXISTS admin (
			id serial PRYMARY KEY,
			name VARCHAR (25) NOT NULL UNIQUE,
			passwordHash VARCHAR NOT NULL UNIQUE
		);
		CREATE TABLE IF NOT EXISTS user (
			id serial PRYMARY KEY,
			name VARCHAR (25) NOT NULL UNIQUE,
			password_hash VARCHAR NOT NULL UNIQUE
		);
		CREATE TABLE IF NOT EXISTS payment (
			number serial,
			amount_payment serial,
			date TIMESTAMP WITH TIME ZONE
		);
		CREATE TABLE IF NOT EXISTS payment (
			id serial PRYMARY KEY,
			iban VARCHAR,
			account_balance TIMESTAMP WITH TIME ZONE
		);

	`

	_, err := db.Exec(CreateTablesQuery)
	return err
}
