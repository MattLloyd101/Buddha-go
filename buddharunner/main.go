package main

import (
	"fmt"
	"time"
	"github.com/MattLloyd101/Buddha-go/buddha"
	tiff "golang.org/x/image/tiff"
)

func main() {
	var now = time.Now()
	var outFolder = fmt.Sprintf("out/%04d-%02d-%02d %02d.%02d.%02d/", 
		now.Year(), 
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second())

	const (
		DPI = 720 * (10.0/100.0)
		width = 35
		height = 20

		realWidth = width*DPI
		realHeight = height*DPI

		passCount = realWidth*realHeight*10
		MinIterations = 0x0
		MaxIterations = 0xFFFF

		logInterval = 1*time.Second

		renderIntervalEnabled = true
		renderInterval = 10*time.Minute
	)

	var renderOptions = buddha.RenderOptions {
		IntervalEnabled: renderIntervalEnabled,
		Interval: renderInterval,
		RenderType: buddha.RenderType16Greyscale,
		Invert: false,
	}

	var tiffOptions = tiff.Options{
		tiff.Deflate,
		false}

	var saveOptions = buddha.SaveOptions{
		TiffOptions: &tiffOptions,
		OutFolder: outFolder}

	var logOptions = buddha.LogOptions {
		LogInterval: logInterval}

	var options = buddha.Options{
		Seed: now.Unix(),
		Width: realWidth,
		Height: realHeight,

		PassCount: passCount,
		MinIterations: MinIterations,
		MaxIterations: MaxIterations,

		WorkerParrallelism: 6,
		MergeParrallelism: 2,

		RenderOptions: &renderOptions,
		SaveOptions: &saveOptions,
		LogOptions:	&logOptions}

	buddha.RunBuddha(&options)	
}