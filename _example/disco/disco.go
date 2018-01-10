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
	bpm := 120.0
	quarterNoteDuration := time.Duration(float64(time.Second) * (60.0 / bpm))
	eigthNoteDuration := quarterNoteDuration / 2.0

	// load samples and set parts to the samples
	// also turn on sends and set reverb/delay on them
	part1 := 1
	part2 := 2
	part3 := 3
	part4 := 4
	fxSend1 := 1
	fxSend2 := 2
	startTime := 0.0
	engine.Perform(
		// load samples at sample slots
		h.LoadSample(startTime, 1, "../samples/707kick.wav"),
		h.LoadSample(startTime, 2, "../samples/707snare.wav"),
		h.LoadSample(startTime, 3, "../samples/707closedhat.wav"),
		h.LoadSample(startTime, 4, "../samples/bass.wav"),
		// set sample slots to parts
		h.SetPartSample(startTime, part1, 1),
		h.SetPartSample(startTime, part2, 2),
		h.SetPartSample(startTime, part3, 3),
		h.SetPartSample(startTime, part4, 4),
		// turn on part 2's send (and add reverb to the fxsend)
		h.SetPartSendDestination(0, part2, float64(fxSend1)),
		h.SetPartSendWet(0, part2, 0.4),
		h.SetFXSendReverbWet(0, fxSend1, 0.6),
		// turn on part 3's send (and add delay to the fxsend)
		h.SetPartSendDestination(0, part3, 2),
		h.SetPartSendWet(0, part3, 0.4),
		h.SetFXSendDelayWet(0, fxSend2, 0.7),
		h.SetFXSendDelayLeftFeedback(0, fxSend2, 0.2),
		h.SetFXSendDelayRightFeedback(0, fxSend2, 0.2),
		h.SetFXSendDelayLeftTime(0, fxSend2, 0.05),
		h.SetFXSendDelayRightTime(0, fxSend2, 0.1),
		// turn down the volume on part 4
		h.SetPartAmp(0, part4, 0.3))

	// kick and hat loop
	go func() {
		d := quarterNoteDuration
		for {
			time.Sleep(time.Duration(d))
			engine.Perform(
				h.PlayPart(startTime, 1, part1, 0),
				h.PlayPart(startTime, 0.5, part3, 0))
		}
	}()

	// snare loop
	go func() {
		d := 2 * quarterNoteDuration // half note duration
		for {
			time.Sleep(d)
			engine.Perform(
				h.PlayPart(startTime, 0.3, part2, 0))
		}
	}()

	// bass loop
	go func() {
		d := eigthNoteDuration
		for {
			time.Sleep(d)
			engine.Perform(
				h.PlayPart(startTime, 0.3, part4, -5))
		}
	}()

	// can hear drift between the disparate parts
	// coming in for higher times
	time.Sleep(8 * time.Second)
}
