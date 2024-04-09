package main

import (
	"gocv.io/x/gocv"
	"log"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(1)
	if err != nil {
		log.Fatal("Error opening video capture device:", err)
	}
	window := gocv.NewWindow("First Demo")
	image := gocv.NewMat()
	for {
		webcam.Read(&image)
		window.IMShow(image)
		window.WaitKey(1)
	}
}
