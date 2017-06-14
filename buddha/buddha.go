package buddha

import (
	"fmt"
	"math/rand"
)

const (
	xMin float64 = -2.5
	xMax float64 = 1.0
	yMin float64 = -1.0
	yMax float64 = 1.0
	uint16Max uint16 = 0xFFFF
)

const (
	_ = iota
	RenderType16Greyscale
)

type icoordinate struct {
	x float64
	y float64
}

func initalizeData(options *Options) *internalState {
	var state = internalState{
		Options: options,
		RawData: nil,
		MaxValue: 0,
		Skipped: 0,
		LastIteration: 0}

	fmt.Println("Initalising Data")
	fmt.Printf("Size: %dx%d\n", options.Width, options.Height)

	var rawData = make([][]uint32, options.Width)
	for x := 0; x < options.Width; x++ { rawData[x] = make([]uint32, options.Height) }
	state.RawData = &rawData

	fmt.Printf("Seeding(0x%X)\n", options.Seed)
	rand.Seed(options.Seed)

	return &state
}

func RunBuddha(options *Options) {
	var state = initalizeData(options)

	fmt.Println("Setting up Logger")
	setupLogger(state)

	var jobs = make(chan iterationPass)
	var results = make(chan stateDelta)
	var mergeResults = make(chan bool)

	fmt.Println("Creating Workers")
	for i := 0; i <= options.WorkerParrallelism; i++ {
		go buddhaWorker(i, state, jobs, results)
	}

	fmt.Println("Creating Mergers")
	for i := 0; i <= options.MergeParrallelism; i++ {
		go mergeWorker(i, state, results, mergeResults)
	}

	fmt.Println("Loading Jobs")
	for i := int64(0); i <= options.PassCount; i++ {
		var dX = scale(rand.Float64(), 0, 1, xMin, xMax)
		var dY = scale(rand.Float64(), 0, 1, yMin, yMax)
		var job = iterationPass {iteration: i, dX: dX, dY: dY}
		jobs <- job
	}
	close(jobs)

	render(state, "final.tiff")
}

