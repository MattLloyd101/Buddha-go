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
		DPI = 72/2
		width = 35
		height = 20

		passCount = 0xFFFFF
		MinIterations = 0x400
		MaxIterations = 0xFFFF

		logInterval = 1*time.Second
		saveInterval = 0xFFFF
		saveIntervalEnabled = true
		parrallelism = 16
	)

	var tiffOptions = tiff.Options{
		tiff.Deflate,
		false}

	var saveOptions = buddha.SaveOptions{
		RenderType: buddha.RenderType16Greyscale,
		TiffOptions: &tiffOptions,
		SaveInterval: saveInterval,
		SaveIntervalEnabled: saveIntervalEnabled,
		OutFolder: outFolder}

	var logOptions = buddha.LogOptions {
		LogInterval: logInterval}

	var options = buddha.Options{
		Seed: now.Unix(),
		Width: width*DPI,
		Height: height*DPI,

		PassCount: passCount,
		MinIterations: MinIterations,
		MaxIterations: MaxIterations,

		WorkerParrallelism: 6,
		MergeParrallelism: 2,

		SaveOptions: &saveOptions,
		LogOptions:	&logOptions}

	buddha.RunBuddha(&options)	
}