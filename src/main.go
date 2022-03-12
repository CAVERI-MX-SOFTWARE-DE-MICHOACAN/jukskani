package main

import (
	"fmt"

	"caveri.mx/jukskani/src/models"
	"github.com/gin-gonic/gin"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
)

var PIN_DHT22 = 17
var PIN2RELES = []int{27, 22, 23, 24, 25, 16, 26, 6}

func main() {
	var RELES []models.Relay
	sensorDHT := &models.SensorDHT{}
	sensorDHT.Init(PIN_DHT22)
	embd.InitGPIO()
	defer embd.CloseGPIO()
	fmt.Println(sensorDHT)
	for _, pin := range PIN2RELES {
		Rele := models.Relay{Pin_GPIO: pin}
		Rele.Write(1)
		RELES = append(RELES, Rele)
	}
	Env := &models.Environ{Relays: RELES, SensorDHT: sensorDHT}

	router := gin.Default()

	router.GET("/api/relays/:id", RelayHandler(Env))
	router.GET("/api/dht22", DHT22Handler(Env))
	router.Run(":8000")
}
