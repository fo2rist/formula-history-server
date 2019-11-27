package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"

	"github.com/fo2rist/formula-history/app/ergastapi"
	"github.com/fo2rist/formula-history/app/storage"
)

var database *storage.Database

func main() {
	gotenv.OverLoad()

	port := os.Getenv("PORT")
	dbConnectionString := os.Getenv("DATABASE_URL")

	database.InitDB(dbConnectionString)

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, string([]byte("**Formula 1 History API!** Proxy for ergast.com")))
	})

	router.GET("/drivers", func(c *gin.Context) {
		c.String(http.StatusOK, ergastapi.GetDrivers())
	})

	router.GET("/circuits", func(c *gin.Context) {
		c.String(http.StatusOK, ergastapi.GetCircuits())
	})

	router.GET("/teams", func(c *gin.Context) {
		c.String(http.StatusOK, ergastapi.GetTeams())
	})

	router.GET("/standing", func(c *gin.Context) {
		c.String(http.StatusOK, ergastapi.GetStandings())
	})

	router.GET("/season", func(c *gin.Context) {
		c.String(http.StatusOK, ergastapi.GetSeasonCalendar())
	})

	router.Run(":" + port)
}
