package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"caveri.mx/jukskani/models"
)

func loadEnviron() (*models.Environ, error) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/data/environ.json", wd))
	if err != nil {
		panic(err)
	}
	Env := models.Environ{}
	_ = json.Unmarshal([]byte(content), &Env)
	return &Env, nil
}
func saveEnviron(Env *models.Environ) {
	content, _ := json.MarshalIndent(Env, "", " ")
	_ = ioutil.WriteFile("./data/environ.json", content, 0644)
}
