package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"caveri.mx/jukskani/models"
	"github.com/robfig/cron"
)

func initCronTasks(Env *models.Environ) {
	Cron = cron.New()
	for _, task := range Env.RelayCronTasks {
		Cron.AddFunc(task.CronSpec, func() {
			log.Println("\n\n\nCRON TASK!\n\n\n", task.RelayIndex, task.State)
			Env.Relays[task.RelayIndex].Write(task.State)
		})
	}
	Cron.Start()
}
func initCronTask(Env *models.Environ, task models.RelayCronTasks) {
	Cron = cron.New()
	Cron.AddFunc(task.CronSpec, func() {
		log.Println("\n\n\nCRON TASK!\n\n\n", task.RelayIndex, task.State)
		Env.Relays[task.RelayIndex].Write(task.State)
	})
	Cron.Start()
}

func loadEnviron() *models.Environ {
	environ_file_path := os.Getenv("ENVIRON_JSON")
	log.Println("Reading ENVIRON_PATH", environ_file_path)
	content, err := ioutil.ReadFile(environ_file_path)
	if err != nil {
		panic(err)
	}
	Env := models.Environ{}
	_ = json.Unmarshal([]byte(content), &Env)
	return &Env
}
func saveEnviron(Env *models.Environ) {
	content, _ := json.MarshalIndent(Env, "", " ")
	_ = ioutil.WriteFile(os.Getenv("ENVIRON_JSON"), content, 0644)
}
