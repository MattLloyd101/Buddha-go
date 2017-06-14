package buddha

func clamp(t float64, min float64, max float64) float64 {
	if(t < min) { return min }
	if(t > max) { return max }
	return t
}

func scale(t float64, srcMin float64, srcMax float64, targMin float64, targMax float64) float64 {
	var srcSpread = srcMax - srcMin
	var scaledT = (t - srcMin) / srcSpread
	var targSpread = targMax - targMin

	return targMin + (scaledT * targSpread)
}