package buddha

import (
	"fmt"
)

type iterationPass struct {
	iteration int64
	dX float64
	dY float64
}

func buddhaWorker(id int, state *internalState, jobs chan iterationPass, results chan stateDelta) {
	var options = state.Options
	for job := range jobs {
		var escaped, stateDelta = iteration(job.dX, job.dY, options.MaxIterations)

		// we only care about things that escaped.
		if escaped {
			results <- stateDelta
		}

		fmt.Println("iteration: %X", job.iteration)
		state.LastIteration = job.iteration
	}
}


func hasEscaped(x float64, y float64, escapeDist float64) bool {
	return (x*x + y*y) > escapeDist
}

func iteration(dX float64, dY float64, maxIteration int) (bool, stateDelta) {
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

	return escaped, stateDelta {iterationCount: iteration, coordinates: coordinates}
}