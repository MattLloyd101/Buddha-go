package main

import (
	"fmt"
	"time"
	"github.com/MattLloyd101/Buddha-go/buddha"
	tiff "golang.org/x/image/tiff"
)

func main() {
	var now = time.Now()
	var outFolder = fmt.Sprintf("out/%d-%d-%d %d.%d.%d/", 
		now.Year(), 
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second())

	const (
		DPI = 72
		width = 35
		height = 20

		iterationCount = 0xFFFFF
		passCount = 0xFFFFF
		logInterval = 0xFF
		saveInterval = 0xFFF
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
		saveIntervalEnabled,
		outFolder}

	buddha.RunBuddha(&data)	
}