package database

import (
	"database/sql"
	"time"
	
	_ "github.com/lib/pq"
)

const (
	dbUser     = "dekkobraze"
	dbPassword = "admin"
	dbName     = "wbdb0"
)

type dbItem struct {
	id int
	data []byte
}

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
	Date_created       time.Time   `json:"date_created"`
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

	_, err = database.cursor.Exec(`CREATE TABLE IF NOT EXISTS Orders (order_uid CHAR(64) PRIMARY KEY, data JSONB);`)

	return
}

func (database *Database) InsertOrder(byteData []byte, unmarshaledData Order) (err error) {
	_, err = database.cursor.Exec(`INSERT INTO Orders (order_uid, data) VALUES ($1, $2)`, unmarshaledData.Order_uid, byteData)

	return
}

func (database *Database) SelectOrders() (orders [][]byte, err error) {
	rows, err := database.cursor.Query(`SELECT * FROM Orders`)
	if err != nil {
        return
    }
    defer rows.Close()

    for rows.Next(){
        dbItem := dbItem{}
        err = rows.Scan(&dbItem.id, &dbItem.data)
        if err != nil {
			return
		}
        orders = append(orders, dbItem.data)
    }

	return
}
