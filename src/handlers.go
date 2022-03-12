package main

import (
	"net/http"
	"strconv"

	"caveri.mx/jukskani/src/models"
	"github.com/gin-gonic/gin"
)

func RelayHandler(Env *models.Environ) gin.HandlerFunc {
	return func(c *gin.Context) {
		relay, _ := strconv.Atoi(c.Param("id"))
		state, _ := strconv.Atoi(c.Query("state"))
		Env.Relays[relay].Write(state)
		c.IndentedJSON(http.StatusOK, Env.Relays[relay])
	}
}
