package main

import (
	"testing"
)

func TestCaptureImage(t *testing.T) {
	g := NewGyazo()
	g.CaptureImage()

	g.FileName = "test.png"
	g.CaptureImage()

}
