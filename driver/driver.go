package driver

import (
	"database/sql"
	"log"
	"os"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//GetConnection from Heroku
func GetConnection() *sql.DB {

	db, err := sql.Open("postgres", os.Getenv("HEROKU_PG_URI"))
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
