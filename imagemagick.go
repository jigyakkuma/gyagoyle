package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
			fmt.Errorf("error: %v", err)
			os.Exit(1)
		}
	} else {
		tmpFile = g.FileName
	}

	if isExist(tmpFile) == false {
		fmt.Errorf("error: Can't open image file")
		os.Exit(1)
	}

	g.ImageBinary, err = ioutil.ReadFile(tmpFile)
	if err != nil {
		fmt.Errorf("error: %v", err)
		os.Exit(1)
	}
}
