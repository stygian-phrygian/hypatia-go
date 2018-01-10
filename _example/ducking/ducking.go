package main

import (
	"fmt"
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
	bpm := 120.0
	quarterNoteDuration := time.Duration(float64(time.Second) * (60.0 / bpm))

	// load samples and set parts to the samples
	// also turn on sends and set reverb/delay on them
	part1 := 1
	part2 := 2
	fxSend1 := 1
	fxSend2 := 2
	startTime := 0.0
	engine.Perform(
		// load samples at sample slots
		h.LoadSample(startTime, 1, "../samples/707kick.wav"),
		h.LoadSample(startTime, 2, "../samples/707closedhat.wav"),
		// set samples slots to parts
		h.SetPartSample(startTime, part1, 1),
		h.SetPartSample(startTime, part2, 2),
		// turn on part 1's send
		h.SetPartSendDestination(startTime, part1, float64(fxSend1)),
		h.SetPartSendWet(startTime, part1, 1),
		// turn on part 2's send (and simulate noise)
		h.SetPartSendDestination(startTime, part2, float64(fxSend2)),
		h.SetPartSendWet(startTime, part2, 1),
		h.SetFXSendReverbWet(startTime, fxSend2, 1))

	// give CSound time to load the samples
	time.Sleep(100 * time.Millisecond)

	// loop (without side chain compression)
	fmt.Printf("\n\n\nloop (without side chain compression)\n\n\n")
	for i := 0; i < 8; i++ {
		d := quarterNoteDuration
		time.Sleep(time.Duration(d))
		engine.Perform(
			h.PlayPart(startTime, 1, part1, 0),
			h.PlayPart(startTime+0.25, 0.5, part2, 0),
			h.PlayPart(startTime, 0.5, part2, 0))
	}

	// loop (with side chain compression)
	fmt.Printf("\n\n\nloop (with side chain compression)\n\n\n")
	engine.Perform(
		h.SetFXSendCompressorSideChain(startTime, fxSend2, float64(fxSend1)),
		h.SetFXSendCompressorRatio(startTime, fxSend2, 8),
		h.SetFXSendCompressorThreshold(startTime, fxSend2, -10),
		h.SetFXSendCompressorAttack(startTime, fxSend2, 0.001),
		h.SetFXSendCompressorRelease(startTime, fxSend2, 0.1),
		h.SetFXSendCompressorGain(startTime, fxSend2, 10))
	for i := 0; i < 8; i++ {
		d := quarterNoteDuration
		time.Sleep(time.Duration(d))
		engine.Perform(
			h.PlayPart(startTime, 1, part1, 0),
			h.PlayPart(startTime+0.25, 0.5, part2, 0),
			h.PlayPart(startTime, 0.5, part2, 0))
	}
}
