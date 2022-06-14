package main

import (
	"countdown"
	"os"
	"time"
)

func main() {
	sleeper := &countdown.ConfigurableSleeper{Duration: 5 * time.Second, SleepFunc: time.Sleep}
	countdown.Countdown(os.Stdout, sleeper)
}
