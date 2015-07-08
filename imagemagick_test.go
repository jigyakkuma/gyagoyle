package main

import (
	"testing"
)

func TestCaptureImage(t *testing.T) {
	var gyazo Gyazo
	gyazo.Config.Init()
	gyazo.CaptureImage()

	gyazo.FileName = "test.png"
	gyazo.CaptureImage()

}
