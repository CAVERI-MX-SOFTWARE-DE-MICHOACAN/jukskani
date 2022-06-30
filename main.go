package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"caveri.mx/jukskani/models"
	device "github.com/d2r2/go-hd44780"
	"github.com/d2r2/go-i2c"
	"github.com/gin-gonic/gin"
	_ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
	"github.com/robfig/cron"
)

const TIME_FORMAT = "02 Jan 15:04:05"

var sensorDHT *models.SensorDHT

var Temperature, Humidity float64
var DHTError = false

var _i2c *i2c.I2C
var _lcd *device.Lcd
var Cron *cron.Cron

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func lcd_init() *device.Lcd {
	_i2c, err := i2c.NewI2C(0x23, 1)
	check(err)
	_lcd, err := device.NewLcd(_i2c, device.LCD_16x2)
	if err != nil {
		log.Fatalln("No se puede inicializar la pantalla LCD...")
		panic(err)
	}
	_lcd.BacklightOn()
	_lcd.Clear()
	check(err)
	return _lcd

}
func lcd_print(_lcd *device.Lcd, line1 string, line2 string) {
	//_lcd.Clear()
	_lcd.Home()
	_lcd.SetPosition(0, 0)
	fmt.Fprint(_lcd, line1)
	_lcd.SetPosition(1, 0)
	fmt.Fprint(_lcd, line2)
}

func lcdDisplayRoutine(lcd *device.Lcd) {
	for range time.Tick(1 * time.Second) {
		now := time.Now().Format(TIME_FORMAT)
		dhterr := " "
		if DHTError {
			dhterr = "!"
		}
		lcd_print(lcd, now, fmt.Sprintf("%04.1f*C %04.1f%%HR %s", Temperature, Humidity, dhterr))
	}
}
func readDHT(sensor *models.SensorDHT) {
	for range time.Tick(5 * time.Second) {
		log.Print("Reading sensor...\t")

		err := sensor.Read()
		if err != nil {
			DHTError = true
			log.Println(err)
			continue
		}
		DHTError = false
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

	Env := loadEnviron()
	log.Printf("Environ: %+v \n %+v \n", Env, Env.SensorDHT)
	sensorDHT = Env.SensorDHT
	sensorDHT.Init()

	for _, Rele := range Env.Relays {
		Rele.Write(Rele.State)
	}
	return Env
}

func main() {
	Cron = cron.New()
	log.Println("Init... testing LCD...")
	lcd := lcd_init()
	defer _i2c.Close()
	lcd_print(lcd, "CAVERI.MX", "JUKSKANI V1.0")

	sign := make(chan os.Signal, 1)

	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	Env := initEnviron()

	//embd.InitGPIO()
	initCronTasks(Env, Cron)
	go prepareExit(sign)
	go readDHT(sensorDHT)
	go lcdDisplayRoutine(lcd)

	router := gin.Default()

	router.Static("/", "./public")
	router.POST("/api/relays/:id", RelayHandler(Env))
	router.POST("/api/cron", AddTaskCronHandler(Env, Cron))
	router.DELETE("/api/cron/:id", DeleteTaskCronHandler(Env, Cron))
	router.GET("/api/dht22", DHT22Handler(Env))

	router.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))

}
