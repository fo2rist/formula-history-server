package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/fo2rist/formula-history/app/models"
)

type Database struct {
	handle *sql.DB
}

func (db *Database) InitDB(url string) {
	var err error
	db.handle, err = sql.Open("postgres", url)
	if err != nil {
		log.Panic(err)
	}

	if err = db.handle.Ping(); err != nil {
		log.Panic(err)
	}
}

//StoreSeason update the storage with latest season info
func (db *Database) StoreSeason(season *models.Season) error {
	for _, race := range season.Races {
		_, err := db.handle.Exec("INSERT INTO races VALUES (default, $1, $2, $3, $4, $5, $6)", race.RaceName, season.Year, race.URL, race.Date, race.Time, race.Circuit.CircuitID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *Database) LoadSeason(year int) (*models.Season, error) {
	rows, err := db.handle.Query("SELECT name, year, url, date, time, circuit_id from races WHERE year=$1", year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	races := make([]models.Race, 0)
	for rows.Next() {
		var race models.Race
		var circuitID string
		if err := rows.Scan(&race.RaceName, &race.Year, &race.URL, &race.Date, &race.Time, &circuitID); err != nil {
			return nil, err
		}
		race.Circuit = models.Circuit{CircuitID: circuitID}

		races = append(races, race)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	if len(races) == 0 {
		return nil, nil
	}
	season := models.Season{
		Year:  races[0].Year,
		Races: races,
	}
	return &season, nil
}
