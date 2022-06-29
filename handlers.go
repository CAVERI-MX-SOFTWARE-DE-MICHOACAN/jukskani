package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"caveri.mx/jukskani/models"
	"github.com/gin-gonic/gin"
)

/*
 Esta funcion cambia el estado de los relevadores.
 El valor state=1|0 donde 0 enciende/apaga el relevador
 identificado con el parametro id en la ruta, por ejemplo:
 curl http://127.0.0.1:8080/api/relays/0?state=1
 apagará el primer relevador registrado en el
 archivo environ.json.
*/
func RelayHandler(Env *models.Environ) gin.HandlerFunc {
	return func(c *gin.Context) {
		var relay models.Relay
		err := json.NewDecoder(c.Request.Body).Decode(&relay)
		if err != nil {
			log.Println("BAD_REQUEST", err)
			c.IndentedJSON(http.StatusBadRequest, err)
		}

		index, _ := strconv.Atoi(c.Param("id"))
		state := relay.State

		if index >= len(Env.Relays) {
			c.Status(http.StatusBadRequest)
			fmt.Fprintf(c.Writer, "Index del relevador fuera de rango permitido [%d, %d].", 0, len(Env.Relays)-1)
		} else {
			Env.Relays[index].Write(state)
			saveEnviron(Env)
			c.IndentedJSON(http.StatusOK, Env.Relays[index])
		}
	}
}
func DHT22Handler(Env *models.Environ) gin.HandlerFunc {
	return func(c *gin.Context) {
		//err := Env.SensorDHT.Read()

		c.IndentedJSON(http.StatusOK, gin.H{"T": Env.SensorDHT.Temperature, "H": Env.SensorDHT.Humidity, "LastRead": Env.SensorDHT.LastRead})

	}
}
