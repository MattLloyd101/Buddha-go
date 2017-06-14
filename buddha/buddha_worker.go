package buddha

import (
	// "fmt"
	"sync"
)

type icoordinate struct {
	x float64
	y float64
}

type iterationPass struct {
	iteration int64
	dX float64
	dY float64
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


func hasEscaped(x float64, y float64, escapeDist float64) bool {
	return (x*x + y*y) > escapeDist
}

func iteration(dX float64, dY float64, minIteration int, maxIteration int) (bool, stateDelta) {
	var x float64 = 0.0
	var y float64 = 0.0
	var iteration int = 0
	const escapeDist = 2*2
	var escaped = hasEscaped(x, y, escapeDist)

	// fixed size, going for the more memory intensive option
	// in favor of not re-scaling the array each iteration.
	var coordinates = make([]icoordinate, maxIteration)

	for (!escaped && iteration < maxIteration) {
		var xtemp = x*x - y*y + dX
		y = 2*x*y + dY
		x = xtemp

		escaped = hasEscaped(x, y, escapeDist)
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