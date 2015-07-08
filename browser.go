package main

import (
	"log"
	"os/exec"
)

func (g *Gyazo) BrowserOpen() {
	if g.ContentUrl == "" {
		return
	}

	cmd := exec.Command("xdg-open", g.ContentUrl)
	err := cmd.Run()

	if err != nil {
		log.Fatalf("Browser open error: %v", err)
	}
}
