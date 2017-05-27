package main

import (
	"github.com/MattLloyd101/Buddha-go/buddha"
	tiff "golang.org/x/image/tiff"
)

func main() {
	const (
		DPI = 72
		width = 35
		height = 20

		iterationCount = 10000
		passCount = 0xFFFFFF
		logInterval = 0xFFFF
		saveInterval = 0xFFFFF
		saveIntervalEnabled = true
	)

	var tiffOptions = tiff.Options{
		tiff.Uncompressed,
		false}

	var data = buddha.BuddhaData{
		0xDEADBEEF,
		width*DPI,
		height*DPI,
		iterationCount,
		passCount,
		buddha.RenderType16Greyscale,
		nil,
		0,
		&tiffOptions,
		0,
		logInterval,
		saveInterval,
		saveIntervalEnabled}

	buddha.RunBuddha(&data)	
}