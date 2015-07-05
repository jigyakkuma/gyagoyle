package main

import (
	"fmt"
	"os/exec"
)

func (g *Gyazo) BrowserOpen() {
	cmd := exec.Command("xdg-open", g.ContentUrl)
	err := cmd.Run()

	if err != nil {
		fmt.Errorf("error: %v", err)
	}
}
