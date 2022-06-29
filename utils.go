package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"caveri.mx/jukskani/models"
)

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
