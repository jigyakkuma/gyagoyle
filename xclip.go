package main

import (
	"io"
	"log"
	"os/exec"
)

func (g *Gyazo) Xclip() {
	cmd := exec.Command("xclip")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	io.WriteString(stdin, g.ContentUrl)
	stdin.Close()
	cmd.Run()
}
