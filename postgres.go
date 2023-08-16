package main

import "database/sql"

func CreateNewDB(db *sql.DB) error {

	CreateTablesQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id serial PRIMARY KEY,
			name VARCHAR (25) NOT NULL UNIQUE,
			password VARCHAR NOT NULL UNIQUE
			is_admin boolean
		);
		CREATE TABLE IF NOT EXISTS accaunt (
			id serial PRIMARY KEY,
			user_id serial,
			iban VARCHAR (34),
			AccountBalance serial
		);
		CREATE TABLE IF NOT EXISTS payment (
			id serial PRIMARY KEY,
			accaunt_id serial,
			amount_payment serial,
			date TIMESTAMP WITH TIME ZONE
		);
	`

	_, err := db.Exec(CreateTablesQuery)
	return err
}
