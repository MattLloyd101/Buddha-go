package buddha

import (
	"fmt"
	image "image"
	// Damn Americans and their insistence on spelling things wrong!
	colour "image/color"
)

func render16bitGreyscale(state *internalState) image.Image {
	var options = state.Options
	var img = image.NewGray16(image.Rect(0, 0, options.Width, options.Height))
	// transer data
	for x := 0; x < options.Width; x++ {
		for y := 0; y < options.Height; y++ {
			var raw = (*state.RawData)[x][y]
			var value = scale(float64(raw), 
				0, 
				float64(state.MaxValue), 
				0, 
				float64(uint16Max))
			// var value = clamp(float64(raw), 0, float64(uint16Max))
			var value16Bit = uint16(value)

			if(raw == state.MaxValue) { fmt.Printf("Max(%d) %d - %f - %X\n", state.MaxValue, raw, value, value16Bit) }
			
			var colourValue = colour.Gray16{value16Bit}
			img.SetGray16(x, y, colourValue)

		}
	}

	return img
}

func render(state *internalState, filename string) {
	var saveOptions = state.Options.SaveOptions
	switch saveOptions.RenderType {
	case RenderType16Greyscale:
		var img = render16bitGreyscale(state)
		saveFile(img, saveOptions, filename)
	}
}