package buddha

import (
	"os"
	"fmt"
	image "image"
	tiff "golang.org/x/image/tiff"
)

type SaveOptions struct {
	RenderType int
	TiffOptions *tiff.Options
	
	SaveInterval int64
	SaveIntervalEnabled bool

	OutFolder string
}

func saveFile(img image.Image, saveOptions *SaveOptions, filename string) {
	fmt.Println("Saving: ", filename)
	os.Mkdir(saveOptions.OutFolder, 0777)
	f, _ := os.OpenFile(saveOptions.OutFolder + filename, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	
	tiff.Encode(f, img, saveOptions.TiffOptions)
}