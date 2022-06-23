package main

import (
	"encoding/json"
	"io/ioutil"
	"caveri.mx/jukskani/src/models"
)
func loadEnviron() (*models.Environ, error) {
	content, err := ioutil.ReadFile("./data/environ.json")
	if (err!=nil){
		return &models.Environ{}, err
	}
	Env := models.Environ{}
	_=json.Unmarshal([]byte(content),&Env)
	return &Env, nil
}
func saveEnviron(Env *models.Environ){
	content, _ := json.MarshalIndent(Env, ""," ")
	_ = ioutil.WriteFile("./data/environ.json", content, 0644)
}

