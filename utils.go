package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"caveri.mx/jukskani/models"
)

func loadEnviron() *models.Environ {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	environ_file_path := fmt.Sprintf("%s/data/environ.json", wd)
	fmt.Println("ENVIRON_PATH", environ_file_path)
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
	_ = ioutil.WriteFile("./data/environ.json", content, 0644)
}
