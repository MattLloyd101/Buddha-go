package buddha

import (
	"os"
	"time"
	"fmt"
	"math"
	"math/rand"
	image "image"
	// Damn Americans and their insistence on spelling things wrong!
	colour "image/color"
	tiff "golang.org/x/image/tiff"
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

type BuddhaData struct {
	Seed int64
	Width int
	Height int
	Iterations []int
	PassCount int64
	RenderType int
	RawData *[][]uint32
	MaxValue uint32
	TiffOptions *tiff.Options
	Skipped int64
	LogInterval int64
	SaveInterval int64
	SaveIntervalEnabled bool
	OutFolder string
	Parrallelism int
}

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

func imaginaryToRealCoordinates(imaginaryCoords []icoordinate, width int, height int) []image.Point {
	var realCoords = make([]image.Point, len(imaginaryCoords))
	for i, imaginary := range imaginaryCoords {
		var x = scale(imaginary.x, xMin, xMax, float64(0), float64(width))
		var y = scale(imaginary.y, yMin, yMax, float64(0), float64(height))
		realCoords[i] = image.Point{int(x + 0.5), int(y + 0.5)}
	}
	return realCoords
}

func combine(coordinates []image.Point, iterationCount int, data *BuddhaData) {
	var rawData = data.RawData

	// fmt.Println("iterationCount", iterationCount)

	// we only go over iterationCount as coordinates isn't perfectly sized for efficiency.
	for i := 0; i < iterationCount; i++ {
		// fmt.Printf("index (%d)\n", i)
		// fmt.Printf("coordinates len(%d)\n", len(coordinates))
		var coordinate = coordinates[i]
		
		// fmt.Printf("%d %+v\n", i, coordinate)
		if (coordinate.X > 0 && coordinate.Y > 0 && coordinate.X < data.Width && coordinate.Y < data.Height) {
			(*rawData)[coordinate.X][coordinate.Y]++
		
			var value = (*rawData)[coordinate.X][coordinate.Y]
			if(value > data.MaxValue) { data.MaxValue = value }
		} else {
			data.Skipped++
		}
	}
}

func showProgress(int64 i, int64 firstTimestamp, int64 passCount) {
	var now = time.Now().UnixNano()
	var nanoDiff = now - firstTimestamp
	var seconds = float64(nanoDiff) / float64(1000000000.0)

	
	var secondsPerPass = seconds / float64(i)
	var passesLeft = passCount - i
	var secondsLeft = secondsPerPass * float64(passesLeft)
	var minutesLeft = secondsLeft / 60
	var hoursLeft = minutesLeft / 60
	// pretty sure this is really stupid math that can be more efficient...
	var hoursPart = int(math.Floor(hoursLeft))
	var minsPart = int(math.Floor(minutesLeft - float64(hoursPart * 60)))
	var secondsPart = int(math.Floor(secondsLeft - float64(minsPart * 60) - float64(hoursPart * 60 * 60)))

	fmt.Printf("%X/%X – %d hours %d mins %d seconds remain\n", i, passCount, hoursPart, minsPart, secondsPart)
}

func RunBuddha(data *BuddhaData) {

	fmt.Println("Initalising Data")
	fmt.Printf("Size: %dx%d\n", data.Width, data.Height)

	var rawData = make([][]uint32, data.Width)
	for x := 0; x < data.Width; x++ { rawData[x] = make([]uint32, data.Height) }
	data.RawData = &rawData

	fmt.Printf("Seeding(0x%X)\n", data.Seed)
	rand.Seed(data.Seed)

	fmt.Println("Beginning Iteration")

	var firstTimestamp = time.Now().UnixNano()
	var channel = make(chan bool, data.Parrallelism)
	// offset by 1 so we don't get a huge save on pass 0.
	for i := int64(1); i <= data.PassCount; i++ {
		var dX = scale(rand.Float64(), 0, 1, xMin, xMax)
		var dY = scale(rand.Float64(), 0, 1, yMin, yMax)

		// for _, maxIterations := range data.Iterations {
		// 	channel <- true
		// 	go runPass(dX, dY, maxIterations, data, channel)
		// }
		channel <- true
		go runPass(dX, dY, int(data.PassCount - i), data, channel)

		if(i % data.LogInterval == 0) {
			showProgress(i, firstTimestamp, data.PassCount)
		}

		if(data.SaveIntervalEnabled && i % data.SaveInterval == 0) {
			render(data, fmt.Sprintf("iteration-%X.tiff", i))
		}
	}

	render(data, "final.tiff")
}

func render16bitGreyscale(data *BuddhaData) image.Image {

	var img = image.NewGray16(image.Rect(0, 0, data.Width, data.Height))
	// transer data
	for x := 0; x < data.Width; x++ {
		for y := 0; y < data.Height; y++ {
			var raw = (*data.RawData)[x][y]
			var value = scale(float64(raw), 
				0, 
				float64(data.MaxValue), 
				0, 
				float64(uint16Max))
			// var value = clamp(float64(raw), 0, float64(uint16Max))
			var value16Bit = uint16(value)

			if(raw == data.MaxValue) { fmt.Printf("Max(%d) %d - %f - %X\n", data.MaxValue, raw, value, value16Bit) }
			
			var colourValue = colour.Gray16{value16Bit}
			img.SetGray16(x, y, colourValue)

		}
	}

	return img
}

func saveFile(img image.Image, data *BuddhaData, filename string) {
	fmt.Println("Saving: ", filename)
	os.Mkdir(data.OutFolder, 0777)
	f, _ := os.OpenFile(data.OutFolder + filename, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	
	tiff.Encode(f, img, data.TiffOptions)
}

func render(data *BuddhaData, filename string) {
	switch data.RenderType {
	case RenderType16Greyscale:
		var img = render16bitGreyscale(data)
		saveFile(img, data, filename)
	}
}

func runPass(dX float64, dY float64, maxIteration int, data *BuddhaData, channel chan bool) {
	var iterationCount, coordinates = iteration(dX, dY, maxIteration)
	// bailout for never escaping nodes.
	if len(coordinates) == 0 {
		<-channel 
		return 
	}

	var realCoords = imaginaryToRealCoordinates(coordinates, data.Width, data.Height)
	combine(realCoords, iterationCount, data)
	<-channel
}

func hasEscaped(x float64, y float64, escapeDist float64) bool {
	return (x*x + y*y) > escapeDist
}

func iteration(dX float64, dY float64, maxIteration int) (int, []icoordinate) {
	var x float64 = 0.0
	var y float64 = 0.0
	var iteration int = 0
	const escapeDist = 2*2
	var escaped = hasEscaped(x, y, escapeDist)

	// fixed size, going for the more memory intensive option
	// in favor of not re-scaling the array each iteration.
	var coordinates = make([]icoordinate, maxIteration)

	for (!escaped && iteration < maxIteration) {
		// Unable to split this out in to a separate function without pointers to x/y
		var xtemp = x*x - y*y + dX
		y = 2*x*y + dY
		x = xtemp

		escaped = hasEscaped(x, y, escapeDist)
		coordinates[iteration] = icoordinate{x, y}
		iteration += 1
	}

	// we only care about the coords if the iteration escaped.
	if(escaped) {
		return iteration, coordinates
	} else {
		return iteration, make([]icoordinate, 0)
	}
}