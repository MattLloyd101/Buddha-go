package buddha

import (
	"fmt"
	"time"
	"image"
	// Damn Americans and their insistence on spelling things wrong!
	colour "image/color"
)

const (
	_ = iota
	RenderType16Greyscale
)

type RenderOptions struct {
	IntervalEnabled bool	
	Interval time.Duration
	RenderType int
	Invert bool
}

type Renderer struct {
	renderType int
	invert bool
	ticker *time.Ticker
}

func setupRenderInterval(renderer *Renderer, saver *Saver, state *internalState) *time.Ticker {
	var options = state.Options
	var renderOptions = options.RenderOptions

	var ticker = time.NewTicker(renderOptions.Interval)
    go func() {
        for range ticker.C {
            var img = renderer.render(state)
            saver.save(img, state)
        }
    }()

    return ticker
}

func setupRenderer(state *internalState, saver *Saver) *Renderer {
	var options = state.Options
	var renderOptions = options.RenderOptions

	var renderer = Renderer {
		renderType: renderOptions.RenderType,
		invert: renderOptions.Invert}

	if (renderOptions.IntervalEnabled) {
		renderer.ticker = setupRenderInterval(&renderer, saver, state)
	}

	return &renderer
}

func (renderer *Renderer) render16bitGreyscale(state *internalState) image.Image {
	var options = state.Options
	var img = image.NewGray16(image.Rect(0, 0, options.Width, options.Height))
	// transer data
	for x := 0; x < options.Width; x++ {
		for y := 0; y < options.Height; y++ {
			var raw = (*state.RawData)[x][y]
			var value = scale(float32(raw), 
				0, 
				float32(state.MaxValue), 
				0, 
				float32(uint16Max))

			var value16Bit uint16
			// var value = clamp(float32(raw), 0, float32(uint16Max))
			if (renderer.invert) {
				value16Bit = uint16Max - uint16(value)
			} else {
				value16Bit = uint16(value)
			}
			
			
			var colourValue = colour.Gray16{value16Bit}
			img.SetGray16(x, y, colourValue)
		}
	}

	return img
}

func (renderer *Renderer) render(state *internalState) image.Image {
	fmt.Printf("Renering State at: %X\n", state.LastMerged)
	switch renderer.renderType {
	case RenderType16Greyscale:
		return renderer.render16bitGreyscale(state)
	}

	fmt.Println("[WARN] Unrecognised Render type! Falling back to 16bit Greyscale!")
	return renderer.render16bitGreyscale(state)
}

func (renderer *Renderer) Stop() {
	if renderer.ticker != nil {
		renderer.ticker.Stop();
	}
}
