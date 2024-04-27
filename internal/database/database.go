package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	cursor *sql.DB
}

func (database *Database)Init() (err error){
	connStr := "user='dekkobraze' password='admin' dbname='wbdb0'"
	database.cursor, err = sql.Open("postgres", connStr)
	if (err != nil) {
		return
	}

	_, err = database.cursor.Exec(`CREATE TABLE IF NOT EXISTS Orders (
		order_uid CHAR(128) PRIMARY KEY,
		track_number CHAR(128) UNIQUE,
		entry CHAR(32),
		locale CHAR(2),
		internal_signature CHAR(32),
		customer_id CHAR(64),
		delivery_service CHAR(16),
		shardkey CHAR(16),
		sm_id INTEGER,
		date_created TIMESTAMP,
		oof_shard CHAR(16));`)
	if (err != nil) {
		return
	}

	_, err = database.cursor.Exec(`CREATE TABLE IF NOT EXISTS Deliveries (
		id INTEGER PRIMARY KEY,
		name CHAR(64) UNIQUE,
		order_uid CHAR REFERENCES Orders,
		phone CHAR(16),
		zip CHAR(16),
		city CHAR(32),
		address CHAR(128),
		region CHAR(32),
		email CHAR(32));`)
	if (err != nil) {
		return
	}

	_, err = database.cursor.Exec(`CREATE TABLE IF NOT EXISTS Payments (
		transaction CHAR(128) PRIMARY KEY,
		order_uid CHAR REFERENCES Orders,
		request_id CHAR(32),
		currency CHAR(16),
		provider CHAR(32),
		amount INTEGER,
		payment_dt INTEGER,
		bank CHAR(32),
		delivery_cost INTEGER,
		goods_total INTEGER,
		custom_fee INTEGER);`)
	if (err != nil) {
		return
	}

	_, err = database.cursor.Exec(`CREATE TABLE IF NOT EXISTS Items (
		chrt_id INTEGER PRIMARY KEY,
		order_uid CHAR REFERENCES Orders,
		track_number CHAR(64) UNIQUE,
		price INTEGER,
		rid CHAR(64),
		name CHAR(64),
		sale INTEGER,
		size CHAR(16),
		total_price INTEGER,
		nm_id INTEGER,
		brand CHAR(64),
		status INTEGER);`)

	return
}
