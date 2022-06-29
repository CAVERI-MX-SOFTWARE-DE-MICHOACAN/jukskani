package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"caveri.mx/jukskani/models"
)

func loadEnviron() *models.Environ {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	wd := filepath.Dir(ex)

	environ_file_path := fmt.Sprintf(os.Getenv("ENVIRON"), wd)
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
	_ = ioutil.WriteFile(os.Getenv("ENVIRON"), content, 0644)
}
