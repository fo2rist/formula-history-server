package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getHttp(addr string) string {
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

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, string([]byte("**Formula 1 History API!**\nProxy for ergast.com")))
	})

	router.GET("/drivers", func(c *gin.Context) {
		c.String(http.StatusOK,
			getHttp("http://ergast.com/api/f1/drivers.json?limit=160&offset=00"))
	})

	router.GET("/circuits", func(c *gin.Context) {
		c.String(http.StatusOK,
			getHttp("http://ergast.com/api/f1/circuits.json?limit=160&offset=00"))
	})

	router.GET("/teams", func(c *gin.Context) {
		c.String(http.StatusOK,
			getHttp("http://ergast.com/api/f1/constructors.json?limit=160&offset=00"))
	})

	router.GET("/standing", func(c *gin.Context) {
		c.String(http.StatusOK,
			getHttp("http://ergast.com/api/f1/current/driverStandings.json"))
	})

	router.GET("/season", func(c *gin.Context) {
		c.String(http.StatusOK,
			getHttp("http://ergast.com/api/f1/current.json"))
	})

	router.Run(":" + port)
}
