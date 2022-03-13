package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"caveri.mx/jukskani/src/models"
	"github.com/gin-gonic/gin"
	_ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
)

var PIN_DHT22 = "GPIO017"
var PIN2RELES = []int{27, 22, 23, 24, 25, 16, 26, 6}
var PIN2RELESSTR = []string{"GPIO27", "GPIO22", "GPIO23", "GPIO24", "GPIO25", "GPIO16", "GPIO26", "GPIO6"}
var Temperature, Humidity float64

func readDHT(sensor *models.SensorDHT) {

	for range time.Tick(5 * time.Second) {
		log.Print("Reading sensor...\t")
		err := sensor.Read()
		Temperature = sensor.Temperature
		Humidity = sensor.Humidity
		log.Println(Temperature, "*C", Humidity, "%HR", "err: ", err)
	}
}

func prepareExit(signal chan os.Signal) {
	sig := <-signal

	//embd.CloseGPIO()
	fmt.Println("\n", sig)
	fmt.Print("\nbye bye\n")
	os.Exit(0)
}

func main() {
	var RELES []models.Relay
	sensorDHT := &models.SensorDHT{}
	sign := make(chan os.Signal, 1)

	sensorDHT.Init(PIN_DHT22)

	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	//embd.InitGPIO()
	go prepareExit(sign)
	go readDHT(sensorDHT)

	for _, pin := range PIN2RELESSTR {
		Rele := models.Relay{PinName: pin}
		Rele.Write(true)
		RELES = append(RELES, Rele)
	}
	Env := &models.Environ{Relays: RELES, SensorDHT: sensorDHT}

	router := gin.Default()

	router.GET("/api/relays/:id", RelayHandler(Env))
	router.GET("/api/dht22", DHT22Handler(Env))
	router.Run(":8000")

}
