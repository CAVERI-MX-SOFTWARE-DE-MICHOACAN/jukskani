package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	device "github.com/d2r2/go-hd44780"
	i2c "github.com/d2r2/go-i2c"
	"caveri.mx/jukskani/src/models"
	"github.com/gin-gonic/gin"
	_ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
)


var PIN_DHT22 = "GPIO17"
var PIN2RELES = []int{27, 22, 23, 24, 25, 16, 26, 6}
var PIN2RELESSTR = []string{"GPIO27", "GPIO22", "GPIO23", "GPIO24", "GPIO25", "GPIO16", "GPIO26", "GPIO6"}
var sensorDHT  *models.SensorDHT

var Temperature, Humidity float64

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func lcd_print(string [] line1) {
	i2c, err := i2c.NewI2C(0x23, 1)
	check(err)
	defer i2c.Close()
	lcd, err := device.NewLcd(i2c, device.LCD_16x2)
	check(err)
	lcd.BacklightOn()
	lcd.Clear()
	for {
		lcd.Home()
		t := time.Now()
		lcd.SetPosition(0, 0)
		fmt.Fprint(lcd, t.Format("Monday Jan 2"))
		lcd.SetPosition(1, 0)
		fmt.Fprint(lcd, t.Format("15:04:05 2006"))
		//		lcd.SetPosition(4, 0)
		//		fmt.Fprint(lcd, "i2c, VGA, and Go")
		time.Sleep(333 * time.Millisecond)
	}
  }

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
func initEnviron() *models.Environ{
	
	
	sensorDHT = &models.SensorDHT{PinName: PIN_DHT22}
	sensorDHT.Init()
	Env, err:=loadEnviron();
	log.Println("Environment",Env, err)
	if  err==nil{
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

	lcd_test()
	log.Println("Init...")

	
	sign := make(chan os.Signal, 1)


	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	Env := initEnviron()

	//embd.InitGPIO()
	go prepareExit(sign)
	go readDHT(sensorDHT)

	

	router := gin.Default()

	router.Static("/public", "./public")
	router.GET("/api/relays/:id", RelayHandler(Env))
	router.GET("/api/dht22", DHT22Handler(Env))

	log.Println("Listening in port 80")

	router.Run(":80")


}
