package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"srbolab/internal"
	"srbolab/internal/driver"
)

type ConfigApp struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func setupApp() (*ConfigApp, error) {

	configFile, err := os.Open("../../config.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully Opened config.json")
	defer configFile.Close()
	byteValue, _ := ioutil.ReadAll(configFile)

	configApp := ConfigApp{}
	err = json.Unmarshal([]byte(byteValue), &configApp)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Connecting to database....")

	db, err := driver.ConnectSqliteDb()
	if err != nil {
		log.Fatal("Cannot connect to database!", err)
	}

	repo := internal.NewDbHandlers(db)
	internal.NewHandlers(repo)

	return &configApp, nil
}