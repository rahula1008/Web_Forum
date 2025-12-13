package initializers

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sqlx.DB

func ConnectToDB() {
	var err error

	dsn := os.Getenv("DB_URI")
	DB, err = sqlx.Connect("pgx", dsn)

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

}
