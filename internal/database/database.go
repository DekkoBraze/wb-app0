package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbUser     = "dekkobraze"
	dbPassword = "admin"
	dbName     = "wbdb0"
)

type Order struct {
	Order_uid          string   `json:"order_uid"`
	Track_number       string   `json:"track_number"`
	Entry              string   `json:"entry"`
	Delivery           Delivery `json:"delivery"`
	Payment            Payment  `json:"payment"`
	Items              []Item   `json:"items"`
	Locale             string   `json:"locale"`
	Internal_signature string   `json:"internal_signature"`
	Customer_id        string   `json:"customer_id"`
	Delivery_service   string   `json:"delivery_service"`
	Shardkey           string   `json:"shardkey"`
	Sm_id              int      `json:"sm_id"`
	Date_created       string   `json:"date_created"`
	Oof_shard          string   `json:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction   string `json:"transaction"`
	Request_id    string `json:"request_id"`
	Currency      string `json:"currency"`
	Provider      string `json:"provider"`
	Amount        int    `json:"amount"`
	Payment_dt    int    `json:"payment_dt"`
	Bank          string `json:"bank"`
	Delivery_cost int    `json:"delivery_cost"`
	Goods_total   int    `json:"goods_total"`
	Custom_fee    int    `json:"custom_fee"`
}

type Item struct {
	Chrt_id      int    `json:"chrt_id"`
	Track_number string `json:"track_number"`
	Price        int    `json:"price"`
	Rid          string `json:"rid"`
	Name         string `json:"name"`
	Sale         int    `json:"sale"`
	Size         string `json:"size"`
	Total_price  int    `json:"total_price"`
	Nm_id        int    `json:"nm_id"`
	Brand        string `json:"brand"`
	Status       int    `json:"status"`
}

type Database struct {
	cursor *sql.DB
}

func (database *Database) Init() (err error) {
	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName
	database.cursor, err = sql.Open("postgres", connStr)
	if err != nil {
		return
	}

	_, err = database.cursor.Exec(`CREATE TABLE IF NOT EXISTS Orders (id serial PRIMARY KEY, data JSONB);`)

	return
}

func (database *Database) InsertJson(data []byte) (err error) {
	_, err = database.cursor.Exec(`INSERT INTO Orders (data) VALUES ($1)`, data)
	log.Print("data in db!")

	return
}
