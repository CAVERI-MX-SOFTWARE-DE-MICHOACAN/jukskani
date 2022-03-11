package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
)

var PIN_DHT22 = 17
var PIN2RELES = []int{27, 22, 23, 24, 25, 16, 26, 6}
var RELES []embd.DigitalPin

func main() {
	token := "43114589"
	h := sha256.New()
	h.Write([]byte(token))
	encoded := hex.EncodeToString(h.Sum(nil))
	fmt.Println(encoded)
	embd.InitGPIO()
	defer embd.CloseGPIO()
	router := gin.Default()
	router.GET("/api/relays/:id", RelayHandler)
	router.Run(":8000")
	// for _, pin := range PIN2RELES {
	// 	Rele, _ := embd.NewDigitalPin(pin)
	// 	Rele.SetDirection(embd.Out)
	// 	Rele.Write(embd.High)
	// 	RELES = append(RELES, Rele)
	// }
	// for {
	// 	for _, pin := range RELES {
	// 		pin.Write(embd.High)
	// 		time.Sleep(500 * time.Millisecond)
	// 		pin.Write(embd.Low)
	// 		time.Sleep(500 * time.Millisecond)
	// 		fmt.Print("*")
	// 	}
	// }
}
