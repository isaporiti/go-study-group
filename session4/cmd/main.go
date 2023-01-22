package main

import (
	"countdown/countdown"
	"countdown/sleeper"
	"os"
	"time"
)

func main() {
	sleeper := sleeper.NewConfigurableSleeper(5*time.Second, time.Sleep)
	countdown.Countdown(os.Stdout, sleeper)
}
