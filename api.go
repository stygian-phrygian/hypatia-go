package hypatia

import (
	"fmt"
)

// returns csound score data which
// loads a sample into a sample slot in the csound instance
func LoadSample(startTime float64, sampleSlot int, fileName string) string {
	return fmt.Sprintf("i \"LoadSample\" %f 1 %d %q\n",
		startTime, sampleSlot, fileName)
}

// returns csound score data which
// triggers playback of a part with optional note pitch offset
// where noteOffset of 0 => no change, N => part's pitch + N semitones
// NB. duration should probably be > 0, lest it will be indefinite (according to csound score syntax)
// and indefinite instruments (currently) have no easy means of canceling them (without canceling
// every other playing part)
func PlayPart(startTime, duration float64, partNumber int, noteOffset float64) string {
	return fmt.Sprintf("i \"PlayPart\" %f %f %d %f\n",
		startTime, duration, partNumber, noteOffset)
}

// returns csound score data which
// turns off all playing parts (obviously)
func TurnOffAllPlayingParts(startTime float64) string {
	return fmt.Sprintf("i \"TurnOffAllPlayingParts\" %f 1\n", startTime)
}

// returns csound score data which
// connects audio input to an output destination for some (possibly indefinite) duration
// where
//   output destination ==     0 -> Master
//   output destination == N > 0 -> FXSend N
//
//   duration ==     0 -> indefinite duration
//   duration == N > 0 -> N seconds
//
//
// to stop monitoring, use the StopMonitoring func below
//
func MonitorInput(startTime, duration float64, outputDestination int) string {
	return fmt.Sprintf("i \"MonitorInput\" %f %f %d\n", startTime, duration, outputDestination)
}

// returns csound score data which
// stops monitoring of audio input
func StopMonitoring(startTime float64) string {
	return fmt.Sprintf("i \"StopMonitoring\" %f 1 \n", startTime)
}

// returns csound score data which
// records audio from either the Master output (resampling) or from audio input
//
// where
//   sampleSlot       ==     0 -> don't load into a sample slot
//   sampleSlot       == N > 0 -> load into sample slot N
//   recordingMode    ==     0 -> record from Master
//   recordingMode    !=     0 -> record from audio input
//   fileName         ==    "" -> auto-generate a file-name
//   fileName         !=    "" -> use the given fileName (will write over existing files)
//
// if duration == 0, it records indefinitly (until stopped with StopRecording func)
//
func RecordSample(startTime, duration float64, sampleSlot, recordingMode int, fileName string) string {
	return fmt.Sprintf("i \"RecordSample\" %f %f %d %d %q\n",
		startTime, duration, sampleSlot, recordingMode, fileName)
}

// returns csound score data which
// turns off a (possibly indefinite duration) RecordSample
func StopRecording(startTime float64) string {
	return fmt.Sprintf("i \"StopRecording\" %f 1\n", startTime)
}

// return csound score data which
// sets part parameters

func SetPartSample(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartSample\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartPitch(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartPitch\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartAmp(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartAmp\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartSampleOffset(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartSampleOffset\"  % 1f %f %f\n",
		startTime, partNumber, v)
}
func SetPartFilterCutoff(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartFilterCutoff\"  % 1f %f %f\n",
		startTime, partNumber, v)
}
func SetPartFilterResonance(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartFilterResonance\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartFilterType(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartFilterType\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartDistortion(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartDistortion\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartPan(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartPan\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartTimestretchFactor(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartTimestretchFactor\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartTimestretchWindowSize(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartTimestretchWindowSize\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartReverse(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartReverse\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartSendDestination(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartSendDestination\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartSendWet(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartSendWet\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartAmpAttack(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartAmpAttack\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartAmpDecay(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartAmpDecay\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartAmpSustain(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartAmpSustain\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartAmpRelease(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartAmpRelease\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartEnv1Attack(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartEnv1Attack\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartEnv1Decay(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartEnv1Decay\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}
func SetPartEnv1Depth(startTime, partNumber, v float64) string {
	return fmt.Sprintf("i \"SetPartEnv1Depth\" %f 1 %.f %f\n",
		startTime, partNumber, v)
}

// return csound score data which
// sets fxsend parameters

func SetFXSendEQGainLow(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendEQGainLow\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendEQGainMid(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendEQGainMid\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendEQGainHigh(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendEQGainHigh\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendEQLowCornerFrequency(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendEQLowCornerFrequency\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendEQMidPeakingFrequency(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendEQMidPeakingFrequency\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendEQHighCornerFrequency(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendEQHighCornerFrequency\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendChorusDelayTime(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendChorusDelayTime\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendChordDepth(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendChordDepth\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendChorusRate(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendChorusRate\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendChorusFeedback(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendChorusFeedback\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendChorusWet(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendChorusWet\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendDelayLeftTime(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendDelayLeftTime\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendDelayLeftFeedback(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendDelayLeftFeedback\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendDelayRightTime(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendDelayRightTime\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendDelayRightFeedback(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendDelayRightFeedback\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendDelayWet(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendDelayWet\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendRingModFrequency(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendRingModFrequency\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendReverbRoomSize(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendReverbRoomSize\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendReverbDamping(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendReverbDamping\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendReverbWet(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendReverbWet\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendBitReduction(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendBitReduction\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendCompressorRatio(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendCompressorRatio\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendCompressorThreshold(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendCompressorThreshold\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendCompressorAttack(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendCompressorAttack\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendCompressorRelease(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendCompressorRelease\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendCompressorGain(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendCompressorGain\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}
func SetFXSendGain(startTime, fxSendNumber, v float64) string {
	return fmt.Sprintf("i \"SetFXSendGain\" %f 1 %.f %f\n",
		startTime, fxSendNumber, v)
}

// return csound score data which
// sets master parameters

func SetMasterEQGainLow(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterEQGainLow\" %f 1 %f\n", startTime, v)
}
func SetMasterEQGainMid(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterEQGainMid\" %f 1 %f\n", startTime, v)
}
func SetMasterEQGainHigh(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterEQGainHigh\" %f 1 %f\n", startTime, v)
}
func SetMasterEQLowCornerFrequency(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterEQLowCornerFrequency\" %f 1 %f\n", startTime, v)
}
func SetMasterEQMidPeakingFrequency(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterEQMidPeakingFrequency\" %f 1 %f\n", startTime, v)
}
func SetMasterEQHighCornerFrequency(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterEQHighCornerFrequency\" %f 1 %f\n", startTime, v)
}
func SetMasterReverbRoomSize(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterReverbRoomSize\" %f 1 %f\n", startTime, v)
}
func SetMasterReverbDamping(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterReverbDamping\" %f 1 %f\n", startTime, v)
}
func SetMasterReverbWet(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterReverbWet\" %f 1 %f\n", startTime, v)
}
func SetMasterBitReduction(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterBitReduction\" %f 1 %f\n", startTime, v)
}
func SetMasterCompressorRatio(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterCompressorRatio\" %f 1 %f\n", startTime, v)
}
func SetMasterCompressorThreshold(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterCompressorThreshold\" %f 1 %f\n", startTime, v)
}
func SetMasterCompressorAttack(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterCompressorAttack\" %f 1 %f\n", startTime, v)
}
func SetMasterCompressorRelease(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterCompressorRelease\" %f 1 %f\n", startTime, v)
}
func SetMasterCompressorGain(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterCompressorGain\" %f 1 %f\n", startTime, v)
}
func SetMasterGain(startTime, v float64) string {
	return fmt.Sprintf("i \"SetMasterGain\" %f 1 %f\n", startTime, v)
}
