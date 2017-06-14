package buddha

import (
	"image"
)

func clamp(t float32, min float32, max float32) float32 {
	if(t < min) { return min }
	if(t > max) { return max }
	return t
}

func scale(t float32, srcMin float32, srcMax float32, targMin float32, targMax float32) float32 {
	var srcSpread = srcMax - srcMin
	var scaledT = (t - srcMin) / srcSpread
	var targSpread = targMax - targMin

	return targMin + (scaledT * targSpread)
}

func imaginaryToImagePoint(imaginaryCoords []icoordinate, width int, height int) []image.Point {
	var realCoords = make([]image.Point, len(imaginaryCoords))
	for i, imaginary := range imaginaryCoords {
		var x = scale(float32(imaginary.x), xMin, xMax, float32(0), float32(width))
		var y = scale(float32(imaginary.y), yMin, yMax, float32(0), float32(height))
		realCoords[i] = image.Point{int(x + 0.5), int(y + 0.5)}
	}
	return realCoords
}

func imaginaryToImage(x float32, y float32, width int, height int) (int, int) {
	return int(scale(x, xMin, xMax, float32(0), float32(width))), int(scale(y, yMin, yMax, float32(0), float32(height)))
}