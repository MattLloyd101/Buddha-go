package buddha

import (
	"math"
	"math/rand"
)

const (
	xMin float32 = -2.5
	xMax float32 = 1.0
	yMin float32 = -1.0
	yMax float32 = 1.0
	uint16Max uint16 = 0xFFFF
)

const ONE_DIV_FOUR float32 = (1.0/4.0)
const ONE_DIV_SIXTEEN float32 = (1.0 / 16.0)

func isInBulb(x float32, y float32) bool {
	var a = (x - ONE_DIV_FOUR)
	var y2 = y*y
	var p = float32(math.Sqrt(float64(a*a + y2)))
	var b = (x + 1)

	return (x < p - (2*p*p) + ONE_DIV_FOUR) && ((b*b) + y2 < ONE_DIV_SIXTEEN)
}

func isInCartoid(x float32, y float32) bool {
	var a = (x - ONE_DIV_FOUR)
	var y2 = y*y
	var q = (a*a + y2)
	return (q * (q + a)) < (ONE_DIV_FOUR * y2)
}

func isValidPoint(x float32, y float32) bool {
	return !isInBulb(x, y) && !isInCartoid(x, y)
}

func createPass(iteration int64) iterationPass {
	var pointIsInvalid = true
	var dX float32 = 0.0
	var dY float32 = 0.0

	for pointIsInvalid {
		dX = scale(float32(rand.Float64()), 0, 1, xMin, xMax)
		dY = scale(float32(rand.Float64()), 0, 1, yMin, yMax)
		pointIsInvalid = !isValidPoint(dX, dY)
	}
	
	return iterationPass {iteration: iteration, dX: dX, dY: dY}
}