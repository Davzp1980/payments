package main

import "database/sql"

func CreateNewDB(db *sql.DB) error {

	CreateTablesQuery := `
		CREATE TABLE IF NOT EXISTS admin (
			id serial PRIMARY KEY,
			name VARCHAR (25) NOT NULL UNIQUE,
			passwordHash VARCHAR NOT NULL UNIQUE
		);
		CREATE TABLE IF NOT EXISTS users (
			id serial PRIMARY KEY,
			name VARCHAR (25) NOT NULL UNIQUE,
			password_hash VARCHAR NOT NULL UNIQUE
		);
		CREATE TABLE IF NOT EXISTS usercabinet (
			id serial PRIMARY KEY,
			user_name VARCHAR (25),
			iban VARCHAR (34),
			AccountBalance serial
		);
		CREATE TABLE IF NOT EXISTS payment (
			id serial PRIMARY KEY,
			user_name VARCHAR (25),
			number_payment VARCHAR,
			amount_payment serial,
			date TIMESTAMP WITH TIME ZONE
		);
	`

	_, err := db.Exec(CreateTablesQuery)
	return err
}
