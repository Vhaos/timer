<p align="center">
  <a href="https://github.com/vhaos/timer"><img width="100" src="https://raw.githubusercontent.com/Vhaos/timer/master/assets/icon.png" alt="timer logo" /></a>
</h1>

<h2 align="center">
  Timer
</h2>

<p align="center">
  A simple CLI tool for setting timers
</p>

<p align="center">
	<a href="https://pkg.go.dev/github.com/vhaos/timer">
		<img src="https://godoc.org/github.com/vhaos/timer?status.svg">
	</a>
	<a href="https://goreportcard.com/badge/github.com/vhaos/timer">
		<img src="https://goreportcard.com/badge/github.com/vhaos/timer">
	</a>
	<a href="https://github.com/vhaos/timer/raw/master/LICENSE">
		<img src="https://img.shields.io/badge/license-MIT-blue">
	</a>
</p>

## About

Timer is CLI timer built with [cobra](https://github.com/spf13/cobra) (cli framework) and [beeep](https://github.com/gen2brain/beeep/) (cross-platform native notifications).

## Install

Assuming you already have go installed:

```bash
go get -u github.com/vhaos/timer
```

## Examples
```
# Start a 10 seconds timer as a background process.
timer set 10s &

# Start a timer for 3 minutes and 31 seconds.
timer set 3m21s

# Start a timer for 1 hour and 1 minute.
timer set 1h1m
```
