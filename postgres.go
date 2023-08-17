package main

import "database/sql"

func CreateNewDB(db *sql.DB) error {

	CreateTablesQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id serial PRIMARY KEY,
			name VARCHAR (25) NOT NULL UNIQUE,
			password VARCHAR NOT NULL UNIQUE,
			is_admin boolean
		);
		CREATE TABLE IF NOT EXISTS accounts (
			id serial PRIMARY KEY,
			user_id serial NOT NULL UNIQUE,
			iban VARCHAR (34) NOT NULL UNIQUE,
			balance serial
		);
		CREATE TABLE IF NOT EXISTS payments (
			id serial PRIMARY KEY,
			account_id serial NOT NULL UNIQUE,
			amount_payment serial,
			date TIMESTAMP WITH TIME ZONE,
			reciever VARCHAR
		);
	`

	_, err := db.Exec(CreateTablesQuery)
	return err
}
