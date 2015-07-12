package main

import (
	"flag"
	"fmt"
	"os"
)

type Gyazo struct {
	ImageBinary []byte
	Config      Config
	FileName    string
	ContentUrl  string
	NoSave      bool
}

func main() {
	var varFlag bool
	var profile string
	g := NewGyazo()

	//flag
	flag.StringVar(&g.Config.Endpoint, "endpoint", "http://gyazo.com/upload.cgi", "Set the original endpoint.")
	flag.StringVar(&g.FileName, "file", "", "Specify if you want to upload the captured file.")
	flag.StringVar(&profile, "profile", "", "Specify a profile to use the configuration toml file.")
	flag.BoolVar(&g.NoSave, "no-save", false, "It is an option that does not save the image file to the history directory.")
	flag.BoolVar(&varFlag, "version", false, "Display version")
	flag.Parse()

	if varFlag == true {
		fmt.Println(Name, ": ", Version)
		os.Exit(0)
	}

	if profile != "" {
		g.Config.GetToml(profile)
	}

	g.Run()
}

func (g *Gyazo) Run() {
	g.CaptureImage()
	g.Upload()
	g.Xclip()
	g.BrowserOpen()
}
