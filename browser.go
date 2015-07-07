package main

import (
	"log"
	"os/exec"
)

func (g *Gyazo) BrowserOpen() {
	cmd := exec.Command("xdg-open", g.ContentUrl)
	err := cmd.Run()

	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
