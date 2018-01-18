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
	bpm := 90.0
	quarterNoteDuration := time.Duration(float64(time.Second) * (60.0 / bpm))
	eigthNoteDuration := quarterNoteDuration / 2.0
	sixteenthNoteDuration := eigthNoteDuration / 2.0

	// find where sample directory is located (relative to GOPATH)
	sampleDirectory := os.Getenv("GOPATH") + "src/github.com/stygian-phrygian/hypatia-go/_example/samples/"

	// load samples and set parts to the samples
	// also turn on sends and set reverb/delay on them
	bassPart := 1
	synthPart := 2
	fxSend1 := 1
	startTime := 0.0
	engine.Perform(
		// load samples at sample slots
		h.LoadSample(startTime, 1, sampleDirectory+"synthBass.wav"),
		h.LoadSample(startTime, 2, sampleDirectory+"synth.wav"),
		// set parts to sample slots
		h.SetPartSample(startTime, bassPart, 1),
		h.SetPartSample(startTime, synthPart, 2),
		// change synth part's parameters
		h.SetPartAmpAttack(startTime, synthPart, 0.02),
		h.SetPartAmpDecay(startTime, synthPart, 0.1),
		h.SetPartAmpSustain(startTime, synthPart, 0.6),
		// turn on synth part's send and add effects
		h.SetPartSendDestination(startTime, synthPart, float64(fxSend1)),
		h.SetPartSendWet(startTime, synthPart, 1),
		h.SetFXSendEQLowCornerFrequency(startTime, fxSend1, 2000),
		h.SetFXSendEQGainLow(startTime, fxSend1, 0.1),
		h.SetFXSendChorusWet(startTime, fxSend1, 0.8),
		h.SetFXSendChorusDepth(startTime, fxSend1, 1),
		h.SetFXSendChorusRate(startTime, fxSend1, 0.2),
		h.SetFXSendChorusFeedback(startTime, fxSend1, 0.1),
		h.SetFXSendReverbRoomSize(startTime, fxSend1, 0.6),
		h.SetFXSendReverbWet(startTime, fxSend1, 1),
		h.SetFXSendGain(startTime, fxSend1, 1.1))

	d := sixteenthNoteDuration
	sixteenthNoteDurationInSeconds := float64(sixteenthNoteDuration) / float64(time.Second)
	maxNumberOfSteps := 64
	arpeggio := []int{0, 4, 7, 11, 12, 11, 7, 4}
	// bass := []int{0, 0, 0, 4, 7, 7, 7, -5}
	for i := 0; i < maxNumberOfSteps; i++ {
		engine.Perform(
			h.PlayPart(
				startTime,
				sixteenthNoteDurationInSeconds,
				bassPart,
				float64(arpeggio[i%len(arpeggio)])),
			h.PlayPart(
				startTime,
				sixteenthNoteDurationInSeconds,
				synthPart,
				float64(arpeggio[i%len(arpeggio)])))
		time.Sleep(time.Duration(d))
	}
}
