package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"caveri.mx/jukskani/src/models"
	device "github.com/d2r2/go-hd44780"
	"github.com/d2r2/go-i2c"
	"github.com/gin-gonic/gin"
	_ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
)

const TIME_FORMAT = "01 Jan 15:04:05"

var PIN_DHT22 = "GPIO17"
var PIN2RELES = []int{27, 22, 23, 24, 25, 16, 26, 6}
var PIN2RELESSTR = []string{"GPIO27", "GPIO22", "GPIO23", "GPIO24", "GPIO25", "GPIO16", "GPIO26", "GPIO6"}
var sensorDHT *models.SensorDHT

var Temperature, Humidity float64

var _i2c *i2c.I2C
var lcd *device.Lcd

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func lcd_init() {
	_i2c, err := i2c.NewI2C(0x23, 1)
	check(err)
	lcd, err := device.NewLcd(_i2c, device.LCD_16x2)
	lcd.BacklightOn()
	lcd.Clear()
	check(err)

}
func lcd_print(line1 string, line2 string) {
	lcd.Clear()
	lcd.Home()
	lcd.SetPosition(0, 0)
	fmt.Fprint(lcd, line1)
	lcd.SetPosition(1, 0)
	fmt.Fprint(lcd, line2)

}
func lcdDisplayRoutine() {
	for range time.Tick(1 * time.Second) {
		now := time.Now().Format(TIME_FORMAT)
		lcd_print(now, fmt.Sprintf("%.2f *C %.2f %%HR", Temperature, Humidity))
	}
}
func readDHT(sensor *models.SensorDHT) {
	for range time.Tick(5 * time.Second) {
		now := time.Now().Format(TIME_FORMAT)
		log.Print("Reading sensor...\t")

		err := sensor.Read()
		if err != nil {
			lcd_print(now, "DHT22 ERROR")
			continue
		}
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
func initEnviron() *models.Environ {

	sensorDHT = &models.SensorDHT{PinName: PIN_DHT22}
	sensorDHT.Init()
	Env, err := loadEnviron()
	log.Println("Environment", Env, err)
	if err == nil {
		for _, Rele := range Env.Relays {
			Rele.Write(Rele.State)
		}

	} else {
		var RELES []models.Relay
		for _, pin := range PIN2RELESSTR {
			Rele := models.Relay{PinName: pin}
			Rele.Write(true)
			RELES = append(RELES, Rele)
		}
		Env = &models.Environ{Relays: RELES, SensorDHT: sensorDHT}
	}
	return Env
}

func main() {

	log.Println("Init... testing LCD...")
	lcd_init()
	defer _i2c.Close()
	lcd_print("CAVERI.MX", "JUKSKANI V1.0")

	sign := make(chan os.Signal, 1)

	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	Env := initEnviron()

	//embd.InitGPIO()
	go prepareExit(sign)
	go readDHT(sensorDHT)
	go lcdDisplayRoutine()

	router := gin.Default()

	router.Static("/public", "./public")
	router.GET("/api/relays/:id", RelayHandler(Env))
	router.GET("/api/dht22", DHT22Handler(Env))

	router.Run(":8080")

}
