package buddha

import (
	"math"
	"math/rand"
)

const (
	xMin float64 = -2.5
	xMax float64 = 1.0
	yMin float64 = -1.0
	yMax float64 = 1.0
	uint16Max uint16 = 0xFFFF
)

const ONE_DIV_FOUR float64 = (1.0/4.0)
const ONE_DIV_SIXTEEN float64 = (1.0 / 16.0)

func isInBulb(x float64, y float64) bool {
	var a = (x - ONE_DIV_FOUR)
	var y2 = y*y
	var p = math.Sqrt(a*a + y2)
	var b = (x + 1)

	return (x < p - (2*p*p) + ONE_DIV_FOUR) && ((b*b) + y2 < ONE_DIV_SIXTEEN)
}

func isInCartoid(x float64, y float64) bool {
	var a = (x - ONE_DIV_FOUR)
	var y2 = y*y
	var q = (a*a + y2)
	return (q * (q + a)) < (ONE_DIV_FOUR * y2)
}

func isValidPoint(x float64, y float64) bool {
	return !isInBulb(x, y) && !isInCartoid(x, y)
}

func createPass(iteration int64) iterationPass {
	var pointIsInvalid = true
	var dX float64 = 0.0
	var dY float64 = 0.0

	for pointIsInvalid {
		dX = scale(rand.Float64(), 0, 1, xMin, xMax)
		dY = scale(rand.Float64(), 0, 1, yMin, yMax)
		pointIsInvalid = !isValidPoint(dX, dY)
	}
	
	return iterationPass {iteration: iteration, dX: dX, dY: dY}
}