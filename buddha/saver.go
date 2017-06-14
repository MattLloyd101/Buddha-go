package buddha

import (
	"os"
	"fmt"
	"image"
	tiff "golang.org/x/image/tiff"
)

type SaveOptions struct {
	OutFolder string
	TiffOptions *tiff.Options
}

type Saver struct {
	outFolder string
	tiffOptions *tiff.Options
}

func setupSaver(state *internalState) *Saver {
	var options = state.Options
	var saveOptions = options.SaveOptions

	var saver = Saver {
		outFolder: saveOptions.OutFolder,
		tiffOptions: saveOptions.TiffOptions}

	return &saver
}

// TODO: this should be part of a Saver Strategy.
func (saver *Saver) generateFilename(state *internalState) string {
	return fmt.Sprintf("iteration-%X.tiff", state.LastMerged)
}

func (saver *Saver) saveWithFilename(img image.Image, filename string) {
	fmt.Println("Saving: ", filename)

	os.Mkdir(saver.outFolder, 0777)
	f, _ := os.OpenFile(saver.outFolder + filename, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	
	tiff.Encode(f, img, saver.tiffOptions)
}

func (saver *Saver) save(img image.Image, state *internalState) {
	var filename = saver.generateFilename(state)
	saver.saveWithFilename(img, filename)
}