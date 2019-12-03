package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"

	"github.com/fo2rist/formula-history/app/ergastapi"
	"github.com/fo2rist/formula-history/app/repository"
	"github.com/fo2rist/formula-history/app/storage"
)

var database *storage.Database = &storage.Database{}

func main() {
	gotenv.OverLoad()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	dbConnectionString := os.Getenv("DATABASE_URL")
	database.InitDB(dbConnectionString)

	client := ergastapi.ErgastClient{}

	repo := repository.Repository{
		WebClient: client,
		Storage:   database,
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, string([]byte("**Formula 1 History API!** Proxy for ergast.com")))
	})

	router.GET("/drivers", func(c *gin.Context) {
		c.String(http.StatusOK, client.GetDrivers())
	})

	router.GET("/circuits", func(c *gin.Context) {
		c.String(http.StatusOK, client.GetCircuits())
	})

	router.GET("/teams", func(c *gin.Context) {
		c.String(http.StatusOK, client.GetTeams())
	})

	router.GET("/standing", func(c *gin.Context) {
		c.String(http.StatusOK, client.GetStandings())
	})

	router.GET("/season", func(c *gin.Context) {
		calendar, _ := repo.GetSeasonCalendar()
		if calendar != nil {
			c.String(http.StatusOK, calendar.String())
		} else {
			c.String(http.StatusInternalServerError, "Can't fetch data")
		}
	})

	router.Run(":" + port)
}
