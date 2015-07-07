package main

import (
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	GyazoId    string
	HomeDir    string
	HistDir    string
	IdFilename string
}

func (config *Config) Init() {
	config.IdFilename = "/.gyazo.id"

	config.createHomeDir()
	config.createHistDir()
	config.getGyazoId()
}

func (config *Config) createHomeDir() {
	//home directory
	config.HomeDir = os.Getenv("HOME") + "/.gyagoyle"
	err := os.MkdirAll(config.HomeDir, 0777)
	if err != nil {
		log.Fatalf("Make a home directory is failed: %v", err)
	}

	return
}

func (config *Config) createHistDir() {
	//history directory
	config.HistDir = config.HomeDir + "/history"
	err := os.MkdirAll(config.HistDir, 0777)
	if err != nil {
		log.Fatalf("Make a history directory is failed: %v", err)
	}

	return
}

func (config *Config) getGyazoId() {
	filePath := config.HomeDir + config.IdFilename
	if isExist(filePath) == false {
	}

	id, err := ioutil.ReadFile(filePath)
	if err != nil {
		//The ID is blank is no problem
		return
	}

	config.GyazoId = string(id)

	return
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func (config *Config) SetGyazoId(id string) {
	if config.GyazoId == "" {
		ioutil.WriteFile(config.HomeDir+config.IdFilename, []byte(id), 0644)
	}

	return
}
