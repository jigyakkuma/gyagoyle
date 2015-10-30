package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

func (g *Gyazo) CaptureGnomeScreenshot() {
	var err error
	tmpFile := g.Config.HistDir + "/" + time.Now().Format("20060102150405") + ".png"

	if g.FileName == "" {
		cmd := exec.Command("gnome-screenshot", "-a", "-f", tmpFile)

		err = cmd.Run()

		if err != nil {
			log.Fatalf("gnome-screenshot error: %v", err)
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
