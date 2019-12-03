package repository

import (
	"log"
	"time"

	"github.com/fo2rist/formula-history/app/models"
)

type WebClient interface {
	GetSeasonCalendar() (*models.Season, error)
}

type Storage interface {
	StoreSeason(season *models.Season) error
	LoadSeason(year int) (*models.Season, error)
}

type Repository struct {
	WebClient WebClient
	Storage Storage
}

func (repo *Repository) GetSeasonCalendar() (*models.Season, error) {
	//get from database if possible
	currentYear := time.Now().Year()
	cachedSeasonCalendar, _ := repo.Storage.LoadSeason(currentYear)
	if !(cachedSeasonCalendar == nil) {
		log.Println("Loaded from DB")
		return cachedSeasonCalendar, nil
	} else {
		log.Println("Loading from internet")
		//if not present (how to distinguish incorrect data and missing data?)
		//then load from network
		if seasonCalendar, err := repo.WebClient.GetSeasonCalendar(); err == nil {
			repo.Storage.StoreSeason(seasonCalendar)
			log.Println(seasonCalendar)
			return seasonCalendar, nil
		} else {
			log.Println(err)
			return nil, err
		}
	}
}
