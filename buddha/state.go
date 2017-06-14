package buddha

import (
	"fmt"
	image "image"
)
type internalState struct {
	Options *Options
	RawData *[][]uint32
	MaxValue uint32
	Skipped int64
	LastIteration int64
	LastMerged int64
}

type stateDelta struct {
	iterationCount int
	coordinates []icoordinate
}


func mergeWorker(id int, state *internalState, deltas chan stateDelta, results chan bool) {
	var options = state.Options
	var width = options.Width
	var height = options.Height

	for delta := range deltas {
		var realCoords = imaginaryToRealCoordinates(delta.coordinates, width, height)
		combine(realCoords, delta.iterationCount, state)
		fmt.Println("merging")
		results <- true
	}
}

func imaginaryToRealCoordinates(imaginaryCoords []icoordinate, width int, height int) []image.Point {
	var realCoords = make([]image.Point, len(imaginaryCoords))
	for i, imaginary := range imaginaryCoords {
		var x = scale(imaginary.x, xMin, xMax, float64(0), float64(width))
		var y = scale(imaginary.y, yMin, yMax, float64(0), float64(height))
		realCoords[i] = image.Point{int(x + 0.5), int(y + 0.5)}
	}
	return realCoords
}

func combine(coordinates []image.Point, iterationCount int, state *internalState) {
	var rawData = state.RawData

	// fmt.Println("iterationCount", iterationCount)

	// we only go over iterationCount as coordinates isn't perfectly sized for efficiency.
	for i := 0; i < iterationCount; i++ {
		// fmt.Printf("index (%d)\n", i)
		// fmt.Printf("coordinates len(%d)\n", len(coordinates))
		var coordinate = coordinates[i]
		
		// fmt.Printf("%d %+v\n", i, coordinate)
		if (coordinate.X > 0 && coordinate.Y > 0 && coordinate.X < state.Options.Width && coordinate.Y < state.Options.Height) {
			(*rawData)[coordinate.X][coordinate.Y]++
		
			var value = (*rawData)[coordinate.X][coordinate.Y]
			if(value > state.MaxValue) { state.MaxValue = value }
		} else {
			state.Skipped++
		}
	}
}