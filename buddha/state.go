package buddha

import (
	// "fmt"
	"sync"
	"image"
)
type internalState struct {
	Options *Options
	RawData *[][]uint32
	MaxValue uint32
	LastIteration int64
	LastMerged int64
}

type stateDelta struct {
	iteration int64
	iterationCount int
	coordinates []icoordinate
}


func mergeWorker(id int, state *internalState, deltas chan stateDelta, waitGroup *sync.WaitGroup) {
	var options = state.Options
	var width = options.Width
	var height = options.Height

	for delta := range deltas {
		var realCoords = imaginaryToImagePoint(delta.coordinates, width, height)
		combine(realCoords, delta.iterationCount, state)
		
		// fmt.Printf("[%d] merging %X deltas(%d)\n", id, delta.iteration, len(deltas))
		state.LastMerged = delta.iteration
	}
	waitGroup.Done()
}

func combine(coordinates []image.Point, iterationCount int, state *internalState) {
	var rawData = state.RawData

	// we only go over iterationCount as coordinates isn't perfectly sized for efficiency.
	for i := 0; i < iterationCount; i++ {
		var coordinate = coordinates[i]
		
		// if exists within the bounds increment
		if (coordinate.X > 0 && coordinate.Y > 0 && coordinate.X < state.Options.Width && coordinate.Y < state.Options.Height) {
			(*rawData)[coordinate.X][coordinate.Y]++
		
			var value = (*rawData)[coordinate.X][coordinate.Y]
			if(value > state.MaxValue) { state.MaxValue = value }
		} 
	}
}