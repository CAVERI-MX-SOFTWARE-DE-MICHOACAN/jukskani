package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"caveri.mx/jukskani/models"
	"github.com/robfig/cron"
)

func initCronTasks(Env *models.Environ, Cron *cron.Cron) {
	Cron = cron.New()
	for _, _task := range Env.RelayCronTasks {
		task := _task
		log.Println("ADD CRONTASK", task.Name, task.CronSpec)
		Cron.AddFunc(task.CronSpec, func() {
			log.Println("\n\n\nCRON TASK!", task.Name, task.RelayIndex, task.State, " \n\n\n")
			Env.Relays[task.RelayIndex].Write(task.State)
			saveEnviron(Env)
		})
	}
	Cron.Start()

}
func addCronTask(Env *models.Environ, Cron *cron.Cron, task models.RelayCronTasks) {

	Cron.AddFunc(task.CronSpec, func() {
		log.Println("\n\n\nCRON TASK!\n\n\n", task.RelayIndex, task.State)
		Env.Relays[task.RelayIndex].Write(task.State)
		saveEnviron(Env)
	})
	Cron.Start()
}
func deleteCronTask(Env *models.Environ, Cron *cron.Cron, index int) {
	Env.RelayCronTasks = append(Env.RelayCronTasks[:index], Env.RelayCronTasks[index+1:]...)
	saveEnviron(Env)
	Cron.Stop()
	initCronTasks(Env, Cron)
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
