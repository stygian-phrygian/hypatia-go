package main

import (
	h "hypatia-go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// create an engine
	engine, err := h.New()
	if err != nil {
		log.Fatal(err)
	}

	// hook up stderr/stdout
	engine.Command.Stdout = os.Stdout // debug
	engine.Command.Stderr = os.Stderr // debug

	// start the engine
	err = engine.Start()
	if err != nil {
		log.Fatal(err)
	}
	// defer its closing
	defer engine.Close()
	// register a signal handler to make sure
	// the engine is closed in signal interrupts
	// otherwise the engine (csound) will keep running obnoxiously
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		engine.Close()
		os.Exit(0)
	}()

	// give CSound time to parse the script and boot
	time.Sleep(1 * time.Second)

	// set up tempo durations (assuming 4/4)
	bpm := 230.0
	quarterNoteDuration := time.Duration(float64(time.Second) * (60.0 / bpm))

	// find where sample directory is located (relative to GOPATH)
	sampleDirectory := os.Getenv("GOPATH") + "src/github.com/stygian-phrygian/hypatia-go/_example/samples/"

	// load samples and set parts to the samples
	// also turn on sends and set reverb/delay on them
	part1 := 1
	fxSend1 := 1
	startTime := 0.0
	engine.Perform(
		h.LoadSample(startTime, 1, sampleDirectory+"909kick.wav"),
		h.SetPartSample(startTime, part1, 1),
		h.SetPartSendDestination(startTime, part1, float64(fxSend1)),
		h.SetPartSendWet(startTime, part1, 1),
		h.SetFXSendBitDepth(startTime, fxSend1, 10),
		h.SetFXSendDistortion(startTime, fxSend1, 0.8))

	quarterNoteDurationInSeconds := float64(quarterNoteDuration) / float64(time.Second)
	maxNumberOfSteps := 16

	// play (with bit reduction)
	for i := 0; i < maxNumberOfSteps; i++ {
		engine.Perform(
			h.PlayPart(
				startTime,
				quarterNoteDurationInSeconds,
				part1,
				0))
		time.Sleep(quarterNoteDuration)
	}
}
