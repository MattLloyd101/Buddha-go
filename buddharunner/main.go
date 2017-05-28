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

	var iterations = []int{0xFF, 0xFFF, 0x5FFF}
	const (
		DPI = 72
		width = 35
		height = 20

		passCount = 0xFFFFF
		logInterval = 0xFFF
		saveInterval = 0xFFFF
		saveIntervalEnabled = true
		parrallelism = 16
	)

	var tiffOptions = tiff.Options{
		tiff.Deflate,
		false}

	var data = buddha.BuddhaData{
		now.Unix(),
		width*DPI,
		height*DPI,
		iterations,
		passCount,
		buddha.RenderType16Greyscale,
		nil,
		0,
		&tiffOptions,
		0,
		logInterval,
		saveInterval,
		saveIntervalEnabled,
		outFolder,
		parrallelism}

	buddha.RunBuddha(&data)	
}