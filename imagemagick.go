package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)

func (g *Gyazo) CaptureImage() {
	var err error
	tmpFile := g.Config.HistDir + "/" + time.Now().Format("20060102150405") + ".png"

	if g.FileName == "" {
		cmd := exec.Command("import", tmpFile)

		err = cmd.Run()

		if err != nil {
			log.Fatalf("ImageMagick import error: %v", err)
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
}
