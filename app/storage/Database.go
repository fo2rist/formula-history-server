package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func (s *Database) InitDB(url string) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

//StoreSeason update the storage with latest season info
func (s *Database) StoreSeason() {

}
