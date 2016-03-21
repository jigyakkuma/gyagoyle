package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

type Command struct {
	captureCommand string
	captureOption  []string
	commandName    string
	fileName       string
}

func (g *Gyazo) CaptureImage() {
	var command Command
	command.commandName = g.CaptureModule
	tmpFile := g.Config.HistDir + "/" + time.Now().Format("20060102150405") + ".png"
	command.fileName = tmpFile

	switch command.commandName {
	case "imagemagick":
		command.captureCommand = "import"
		command.captureOption = []string{tmpFile}
	case "gnome-screenshot":
		command.captureCommand = "gnome-screenshot"
		command.captureOption = []string{"-a", "-f", tmpFile}
	default:
		log.Fatalf("error:not found this capture option.")
	}

	g.CaptureScreenshot(&command)
}

func (g *Gyazo) CaptureScreenshot(command *Command) {
	var err error
	var tmpFile string

	if g.FileName == "" {
		tmpFile = command.fileName
		cmd := exec.Command(command.captureCommand, command.captureOption...)

		err = cmd.Run()

		if err != nil {
			log.Fatalf("%s error: %v", command.commandName, err)
		}
	} else {
		tmpFile = g.FileName
	}

	if isExist(tmpFile) == false {
		log.Fatalf("error: Can't open image file")
	}

	g.ImageBinary, err = ioutil.ReadFile(tmpFile)
	if err != nil {
		log.Fatalf("Image file read error: %v", err)
	}

	if g.NoSave == true {
		err = os.Remove(tmpFile)
		if err != nil {
			log.Fatalf("Capture file delete error: %v", err)
		}
	}
}
