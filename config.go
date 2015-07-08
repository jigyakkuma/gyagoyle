package main

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	GyazoId       string
	HomeDir       string
	HistDir       string
	IdFilename    string
	Endpoint      string
	BasicUser     string
	BasicPassword string
}

type Toml struct {
	Profile []Profiles `toml:"profile"`
}

type Profiles struct {
	BasicUser     string `toml:"basicUser"`
	BasicPassword string `toml:"basicPassword"`
	Endpoint      string `toml:"endpoint"`
	Name          string `toml:"name"`
}

func (c *Config) Init() {
	c.IdFilename = "/.gyazo.id"

	c.createHomeDir()
	c.createHistDir()
	c.getGyazoId()
}

func (c *Config) createHomeDir() {
	//home directory
	c.HomeDir = os.Getenv("HOME") + "/.gyagoyle"
	err := os.MkdirAll(c.HomeDir, 0777)
	if err != nil {
		log.Fatalf("Make a home directory is failed: %v", err)
	}

	return
}

func (c *Config) createHistDir() {
	//history directory
	c.HistDir = c.HomeDir + "/history"
	err := os.MkdirAll(c.HistDir, 0777)
	if err != nil {
		log.Fatalf("Make a history directory is failed: %v", err)
	}

	return
}

func (c *Config) getGyazoId() {
	filePath := c.HomeDir + c.IdFilename
	if isExist(filePath) == false {
	}

	id, err := ioutil.ReadFile(filePath)
	if err != nil {
		//The ID is blank is no problem
		return
	}

	c.GyazoId = string(id)

	return
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func (c *Config) SetGyazoId(id string) {
	if c.GyazoId == "" {
		ioutil.WriteFile(c.HomeDir+c.IdFilename, []byte(id), 0644)
	}

	return
}

func (c *Config) GetToml(profile string) {
	filePath := c.HomeDir + "/config.toml"
	if isExist(filePath) == false {
		log.Fatalf("config.toml not found:")
	}

	var t Toml

	_, err := toml.DecodeFile(filePath, &t)
	if err != nil {
		log.Fatalf("config.toml read error:", err)
	}

	for _, v := range t.Profile {
		if v.Name == profile {
			c.BasicUser = v.BasicUser
			c.BasicPassword = v.BasicPassword
			c.Endpoint = v.Endpoint
		}
	}

}
