package buddha

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