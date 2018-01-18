# Hypatia-Go
Golang wrapper around the CSound based sampler, Hypatia.
Kind of a hack.  Only works on unix-y systems (I guess?)

## What Is This?
The Hypatia-Go library simplifies booting Hypatia and communicating with it.
With Hypatia, you can play sound samples, fiddle with them, record new samples, etc.
Communication with Hypatia however involves CSound score which is rather verbose.
This ameliorates that issue.

## Installation
1. [CSound](http://csound.com/download.html)
2. Run `go get` on this repository.

## Usage
See `_example` directory

Warning: one must (currently) run the example golang files in the _exact_ same directory that they're located in.

## What's Going On Internally
Hypatia is just a CSound script, which listens for CSound score language
This wrapper library simplifies the usage of Hypatia by:
* booting CSound with the Hypatia script
* providing a golang API around Hypatia (which itself wants CSound score strings)

## Configuration
Hypatia-Go is configured through a yaml file placed in the same directory that your code runs in.
This file is (curiously) named `hypatia-config.yaml`.
The configuration format is derived from csound [command line flags](http://www.csounds.com/manual/html/CommandFlags.html)(which I recommend consulting the documentation regarding).
Here's an example configuration (which has the defaults that hypatia-go boots Hypatia with).
The comments above each variable offer further explanation.

```yaml
# the audio input device
# default is 'adc' however one can specify a particular device
# if multiple options are available to your system
# on linux, one can list the devices with `arecord -l`
# a specific device value would look like
# 'adc3' or ':hw:1,1'
input:              adc
# the audio output device
# default is 'dac' however one can specify a particular device
# if multiple options are available to your system
# on linux, one can list the devices with `aplay -l`
# a specific device value would look like
# 'dac3' or ':hw:1,1'
#
output:             dac
# the sample rate (obviously enough)
# should match the sample rate of the specified output device
# otherwise black magic will happen
sample-rate:        44100
# consult Hypatia documentation regarding
# the number of parts and fx-sends
number-of-parts:    16
number-of-fx-sends: 2
# the port that Hypatia receives csound score data on
# (again kind of obvious)
osc-listen-port:    8080
# you can pass whatever other CSound specific flags you want here
other-flags:        ""
```

Nota Bene, you probably *won't* need a config file as the defaults should likely work sufficiently.



