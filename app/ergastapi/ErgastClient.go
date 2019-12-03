package ergastapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fo2rist/formula-history/app/ergastapi/models"
	domain "github.com/fo2rist/formula-history/app/models"
)

type ErgastClient struct {}

func (ErgastClient) GetSeasonCalendar() (*domain.Season, error) {
	response := getHTTP("http://ergast.com/api/f1/current.json")
	responseData := models.CurrentSeasonResponse{}
	if err := json.Unmarshal([]byte(response), &responseData); err != nil {
		log.Printf("Can't parse\n%.25s", response)
		return nil, err
	}

	calendar := responseData.MRData.Data.ToDomainModel()
	log.Printf("Parsed: %v", calendar)
	return calendar, nil
}

func (ErgastClient) GetDrivers() string {
	return getHTTP("http://ergast.com/api/f1/drivers.json?limit=900&offset=00")
}

func (ErgastClient) GetCircuits() string {
	return getHTTP("http://ergast.com/api/f1/circuits.json?limit=100&offset=00")
}

func (ErgastClient) GetTeams() string {
	return getHTTP("http://ergast.com/api/f1/constructors.json?limit=2500&offset=00")
}

func (ErgastClient) GetStandings() string {
	return getHTTP("http://ergast.com/api/f1/current/driverStandings.json")
}

func getHTTP(addr string) string {
	resp, err := http.Get(addr)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}