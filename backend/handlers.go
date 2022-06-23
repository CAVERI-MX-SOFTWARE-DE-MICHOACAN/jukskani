package main

import (
	"net/http"
	"strconv"

	"caveri.mx/jukskani/backend/models"
	"github.com/gin-gonic/gin"
)

func RelayHandler(Env *models.Environ) gin.HandlerFunc {
	return func(c *gin.Context) {
		index, _ := strconv.Atoi(c.Param("id"))
		state, _ := strconv.ParseBool(c.Query("state"))

		Env.Relays[index].Write(state)
		saveEnviron(Env)
		c.IndentedJSON(http.StatusOK, Env.Relays[index])
	}
}
func DHT22Handler(Env *models.Environ) gin.HandlerFunc {
	return func(c *gin.Context) {
		//err := Env.SensorDHT.Read()

		c.IndentedJSON(http.StatusOK, gin.H{"T": Env.SensorDHT.Temperature, "H": Env.SensorDHT.Humidity})

	}
}
