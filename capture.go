package main

import (
	"log"
)

func (g *Gyazo) CaptureImage() {
	switch g.CaptureModule {
	case "imagemagick":
		g.CaptureImageMagick()
	case "gnome-screenshot":
		g.CaptureGnomeScreenshot()
	default:
		log.Fatalf("error:not found this capture option.")
	}
}
