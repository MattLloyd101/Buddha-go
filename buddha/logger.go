package buddha

import (
	"time"
	"math"
	"fmt"
)

type LogOptions struct {
	LogInterval time.Duration
}

type Logger struct {
	firstTimestamp int64
	iteration *int64
	passCount *int64
	ticker *time.Ticker
}

func setupLogger(state *internalState) *Logger {
	var options = state.Options
	var now = time.Now().UnixNano()
	var ticker = time.NewTicker(options.LogOptions.LogInterval)
	var logger = Logger {
		firstTimestamp: now,
		iteration: &state.LastIteration,
		passCount: &options.PassCount,
		ticker: ticker}

    go func() {
        for range ticker.C {
            logger.showProgress()
        }
    }()

	return &logger
}

func (logger *Logger) showProgress() {
	var now = time.Now().UnixNano()
	var nanoDiff = now - logger.firstTimestamp
	var seconds = float64(nanoDiff) / float64(1000000000.0)

	var iteration = *(logger.iteration)
	var passCount = *(logger.passCount)

	var secondsPerPass = seconds / float64(iteration)
	var passesLeft = passCount - iteration
	var secondsLeft = secondsPerPass * float64(passesLeft)
	var minutesLeft = secondsLeft / 60
	var hoursLeft = minutesLeft / 60
	// pretty sure this is really stupid math that can be more efficient...
	var hoursPart = int(math.Floor(hoursLeft))
	var minsPart = int(math.Floor(minutesLeft - float64(hoursPart * 60)))
	var secondsPart = int(math.Floor(secondsLeft - float64(minsPart * 60) - float64(hoursPart * 60 * 60)))

	var percent = 100 * (float32(iteration) / float32(passCount))
	fmt.Printf("%06X/%06X – %02d hours %02d mins %02d seconds remain (%0.2f%%)\n", iteration, passCount, hoursPart, minsPart, secondsPart, percent)
}

func (logger *Logger) Stop() {
	logger.ticker.Stop();
}