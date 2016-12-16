package talk_training

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	DB_NAME     = "tokopedia-dev-db"
	DB_USER     = "sm161114"
	DB_PASSWORD = "Abcd1234"
)

var MainDB *sqlx.DB

func connection() {
	db, err := sqlx.Open("postgres", "user="+DB_USER+" password="+DB_PASSWORD+" dbname="+DB_NAME+" host=192.168.100.126 port=5432 sslmode=disable")

	if err != nil {
		log.Fatal("Gagal koneksi")
	}
	MainDB = db
}
