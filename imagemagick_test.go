package main

import (
	"fmt"
	"testing"
)

func TestCaptureImage(t *testing.T) {
	var gyazo Gyazo
	gyazo.CaptureImage()
	fmt.Println(gyazo.ImageBinary)

	gyazo.FileName = "test.png"
	fmt.Println(gyazo.ImageBinary)

}
