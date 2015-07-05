package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func (g *Gyazo) Xclip() {
	cmd := exec.Command("xclip")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Errorf("error: %v", err)
		os.Exit(1)
	}

	io.WriteString(stdin, g.ContentUrl)
	stdin.Close()
	cmd.Run()
}
