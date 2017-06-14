package buddha

import (
	// "fmt"
	"sync"
)

type icoordinate struct {
	x float32
	y float32
}

type iterationPass struct {
	iteration int64
	dX float32
	dY float32
}

func buddhaWorker(id int, state *internalState, passes chan iterationPass, results chan stateDelta, waitGroup *sync.WaitGroup) {
	var options = state.Options
	for pass := range passes {
		var escaped, stateDelta = iteration(pass.dX, pass.dY, options.MinIterations, options.MaxIterations)

		// we only care about things that escaped.
		if escaped {
			stateDelta.iteration = pass.iteration
			results <- stateDelta
		}

		// fmt.Printf("[%X] iteration: %X\n", id, pass.iteration)
		// unstable but don't care it's only used for feedback.
		state.LastIteration = pass.iteration
	}
	waitGroup.Done()
}


func iteration(dX float32, dY float32, minIteration int, maxIteration int) (bool, stateDelta) {
	var fdx float32 = float32(dX)
	var fdy float32 = float32(dY)
	var x float32 = 0.0
	var y float32 = 0.0
	var iteration int = 0
	const escapeDist = 2*2
	var escaped = (x*x + y*y) > escapeDist

	// fixed size, going for the more memory intensive option
	// in favor of not re-scaling the array each iteration.
	var coordinates = make([]icoordinate, maxIteration)

	for (!escaped && iteration < maxIteration) {
		var xtemp = x*x - y*y + fdx
		y = 2*x*y + fdy
		x = xtemp

		escaped = (x*x + y*y) > escapeDist
		coordinates[iteration] = icoordinate{x, y}
		iteration += 1
	}

	// if we haven't met the minimum iterations let's skip over.
	if (iteration < minIteration) {
		// fmt.Printf("ignoring pass due to low iteration count: %d\n", iteration)
		escaped = false
	}

	return escaped, stateDelta {iterationCount: iteration, coordinates: coordinates}
}