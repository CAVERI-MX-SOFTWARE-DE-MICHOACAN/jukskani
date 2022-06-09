package models

import (
	"fmt"

	"github.com/MichaelS11/go-dht"
)

type SensorDHT struct {
	PinName               string 	`json:PinName`
	Temperature			  float64 	`json:Temperature`
	Humidity			  float64 	`json:Humidity`
	Sensor                *dht.DHT	
}

func (DHT *SensorDHT) Init() {
	err := dht.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return
	}

	dht, err := dht.NewDHT(DHT.PinName, dht.Celsius, "")
	if err != nil {
		fmt.Println("NewDHT error:", err, DHT)
		return
	}
	DHT.Sensor = dht

}
func (DHT *SensorDHT) Read() error {
	humidity, temperature, err := DHT.Sensor.ReadRetry(3)

	DHT.Humidity = humidity
	DHT.Temperature = temperature
	return err
}
