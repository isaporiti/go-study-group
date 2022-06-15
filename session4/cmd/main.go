package main

import (
	"countdown"
	"os"
	"time"
)

func main() {
	sleeper := countdown.NewConfigurableSleeper(5*time.Second, time.Sleep)
	countdown.Countdown(os.Stdout, sleeper)
}
