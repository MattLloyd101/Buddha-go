package main

import (
	"github.com/MattLloyd101/Buddha-go/buddha"
	tiff "golang.org/x/image/tiff"
)

func main() {
	var tiffOptions = tiff.Options{
		tiff.Uncompressed,
		false}

	var data = buddha.BuddhaData{
		0xDEADBEEF,
		3500,
		2000,
		10000,
		0xFFFFFF,
		buddha.RenderType16Greyscale,
		nil,
		0,
		&tiffOptions,
		0}

	buddha.RunBuddha(&data)	
}