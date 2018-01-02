package hypatia

import (
	"os"
	"os/exec"
	"strings"

	"github.com/hypebeast/go-osc/osc"
	"github.com/spf13/viper"
)

const (
	configFileName = "hypatia-config.yaml"
	csoundFileName = "main.csd"
	oscAddress     = "/score"
)

// the running csound instance and an OSC client into it
type Engine struct {
	OSCClient *osc.Client
	Command   *exec.Cmd
}

// boots up csound in the current working directory,
// with configuration provided by a config file (should it exist)
// defaults are used if the config doesn't exist
// returns the engine
func New() (*Engine, error) {

	// read config file (if it exists)
	// otherwise use defaults (which are set below)
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile(configFileName)
	v.SetConfigType("yaml")
	v.ReadInConfig() // if it errors, it doesn't matter (we have defaults)
	v.SetDefault("input", "adc")
	v.SetDefault("output", "dac")
	v.SetDefault("sample-rate", "44100")
	v.SetDefault("ksmps", "128")
	v.SetDefault("number-of-parts", "16")
	v.SetDefault("number-of-fx-sends", "2")
	v.SetDefault("osc-listen-port", "8080")
	// v.SetDefault("logfile", "null")
	v.SetDefault("other-flags", "") // <-- option to add further csound flags

	// convert values read from the config file into csound flags
	flags := []string{
		"--input=" + v.GetString("input"),
		"--output=" + v.GetString("output"),
		"--sample-rate=" + v.GetString("sample-rate"),
		"--ksmps=" + v.GetString("ksmps"),
		"--omacro:NUMBER_OF_PARTS=" + v.GetString("number-of-parts"),
		"--omacro:NUMBER_OF_FX_SENDS=" + v.GetString("number-of-fx-sends"),
		"--omacro:OSC_LISTEN_PORT=" + v.GetString("osc-listen-port"),
		// "--logfile=" + v.GetString("logfile"),
	}
	if v.GetString("other-flags") != "" {
		flags = append(flags, v.GetString("other-flags"))
	}

	// create OSC client
	c := osc.NewClient("localhost", v.GetInt("osc-listen-port"))

	// get the csound csd file asset
	// create a temp file with the csound csd file asset data
	// the temp file is removed in func Close() beneath
	err := RestoreAsset("./", csoundFileName)
	if err != nil {
		return nil, err
	}

	// create arguments for csound
	args := append([]string{csoundFileName}, flags...)

	// create the csound command
	// which utilizes the recently created csound csd temp file and flags
	p := exec.Command("csound", args...)

	// return the "engine"
	return &Engine{
		OSCClient: c,
		Command:   p,
	}, nil
}

func (e *Engine) Start() error {
	// start the command
	return e.Command.Start()

}

// closes the underlying csound process
func (e *Engine) Close() error {
	// remove the temp csound file
	os.Remove(csoundFileName)
	// end the csound process
	return e.Command.Process.Kill()
}

// sends csound score data (strings) into the running csound instance
func (e *Engine) Perform(scoreLines ...string) {
	e.OSCClient.Send(osc.NewMessage(
		oscAddress, // <--- magic constant (look in main.csd)
		strings.Join(scoreLines, "")))
}
