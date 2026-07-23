package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connectStr := "host=db port=5432 user=postgres password=1234 dbname=cocktails_db sslmode=disable"

	var db *sql.DB
	var err error

	for i := 1; i <= 10; i++ {
		db, err = sql.Open("postgres", connectStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				fmt.Printf("\n[Успешное подключение к PostgreSQL в Docker (попытка #%d)]\n", i)
				break
			}
		}

		log.Printf("База данных еще не готова (попытка #%d). Ждем 2 секунды...", i)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных: ", err)
	}

	query := `
    CREATE TABLE IF NOT EXISTS cocktails (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        category VARCHAR(100) NOT NULL,
        ingredients TEXT NOT NULL,
        method TEXT NOT NULL,
        serving TEXT NOT NULL,
        is_iba BOOLEAN DEFAULT FALSE,
        image_path TEXT DEFAULT '',
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Ошибка при создании таблицы: ", err)
	}

	queryInventory := `
	CREATE TABLE IF NOT EXISTS inventory (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) UNIQUE NOT NULL,
		in_stock BOOLEAN DEFAULT FALSE
	);`

	_, err = db.Exec(queryInventory)
	if err != nil {
		log.Fatal("Ошибка при создании таблицы инвентаря: ", err)
	}

	fmt.Println("Таблица cocktails создана и готова к работе!")
	return db
}
