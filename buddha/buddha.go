package buddha

import (
	"fmt"
	"math/rand"
	"sync"
)

func initalizeData(options *Options) *internalState {
	var state = internalState{
		Options: options}

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
	var logger = setupLogger(state)

	fmt.Println("Setting up Saver")
	var saver = setupSaver(state)

	fmt.Println("Setting up Renderer")
	var renderer = setupRenderer(state, saver)

	var passes = make(chan iterationPass, 0xFF)
	var results = make(chan stateDelta, 0xFF)
	var workerGroup sync.WaitGroup
	var mergeGroup sync.WaitGroup

	fmt.Println("Creating Workers")
	for i := 0; i <= options.WorkerParrallelism; i++ {
		workerGroup.Add(1)
		go buddhaWorker(i, state, passes, results, &workerGroup)
	}

	fmt.Println("Creating Mergers")
	for i := 0; i <= options.MergeParrallelism; i++ {
		mergeGroup.Add(1)
		go mergeWorker(i, state, results, &mergeGroup)
	}

	fmt.Println("Loading passes")
	for i := int64(0); i <= options.PassCount; i++ {
		var pass = createPass(i)
		passes <- pass
	}
	close(passes)
	fmt.Println("Completed passes")
	
	workerGroup.Wait()
	close(results)
	fmt.Println("Workers complete")

	mergeGroup.Wait()
	fmt.Println("Mergers complete")

	logger.Stop()

	fmt.Println("Creating Final Render")
	var img = renderer.render(state)
	saver.saveWithFilename(img, "final.tiff")
	fmt.Println("Final Render Complete")
}

