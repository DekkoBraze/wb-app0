package database

import (
	"app0/internal/structs"
	"database/sql"

	_ "github.com/lib/pq"
)

// Данные для логина в бд
const (
	dbUser     = "dekkobraze"
	dbPassword = "admin"
	dbName     = "wbdb0"
)

type Database struct {
	cursor *sql.DB
}

// Инициализация
func (database *Database) Init() (err error) {
	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName
	database.cursor, err = sql.Open("postgres", connStr)
	if err != nil {
		return
	}

	_, err = database.cursor.Exec(`CREATE TABLE IF NOT EXISTS Orders (order_uid CHAR(64) PRIMARY KEY, data JSONB);`)

	return
}

// Создание заказа
func (database *Database) InsertOrder(byteData []byte, unmarshaledData structs.Order) (err error) {
	_, err = database.cursor.Exec(`INSERT INTO Orders (order_uid, data) VALUES ($1, $2)`, unmarshaledData.Order_uid, byteData)

	return
}

// Взятие заказа
func (database *Database) SelectOrders() (orders [][]byte, err error) {
	rows, err := database.cursor.Query(`SELECT * FROM Orders`)
	if err != nil {
		return
	}
	defer rows.Close()

	type dbItem struct {
		id   []byte
		data []byte
	}

	for rows.Next() {
		dbItem := dbItem{}
		err = rows.Scan(&dbItem.id, &dbItem.data)
		if err != nil {
			return
		}
		orders = append(orders, dbItem.data)
	}

	return
}
