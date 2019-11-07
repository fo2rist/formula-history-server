package ergastapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fo2rist/formula-history/app/ergastapi/models"
)

func GetSeasonCalendar() string {
	response := getHTTP("http://ergast.com/api/f1/current.json")
	responseData := models.CurrentSeasonResponse{}
	if err := json.Unmarshal([]byte(response), &responseData); err != nil {
		log.Printf("Can't parse\n%.25s", response)
		return ""
	}

	calendar := responseData.MRData.Data.ToDomainModel()

	if str, err := json.MarshalIndent(calendar, "", " "); err != nil {
		log.Print("Can't serialize: ", calendar, err)
		return ""
	} else {
		return string(str)
	}
}

func GetDrivers() string {
	return getHTTP("http://ergast.com/api/f1/drivers.json?limit=900&offset=00")
}

func GetCircuits() string {
	return getHTTP("http://ergast.com/api/f1/circuits.json?limit=100&offset=00")
}

func GetTeams() string {
	return getHTTP("http://ergast.com/api/f1/constructors.json?limit=2500&offset=00")
}

func GetStandings() string {
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