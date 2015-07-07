package main

import (
	"flag"
	"fmt"
	"os"
)

type Gyazo struct {
	ImageBinary []byte
	Config      Config
	Endpoint    string
	FileName    string
	ContentUrl  string
}

func main() {
	var g Gyazo
	var varFlag bool
	g.Config.Init()

	//flag
	flag.StringVar(&g.Endpoint, "endpoint", "http://gyazo.com/upload.cgi", "Set the original endpoint")
	flag.StringVar(&g.FileName, "file", "", "Specify if you want to upload the captured file")
	flag.BoolVar(&varFlag, "version", false, "Display version")
	flag.Parse()

	if varFlag == true {
		fmt.Println(Name, ": ", Version)
		os.Exit(0)
	}

	g.Run()
}

func (g *Gyazo) Run() {
	g.CaptureImage()
	g.Upload()
	g.Xclip()
	g.BrowserOpen()
}
