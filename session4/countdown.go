package countdown

import (
	"fmt"
	"io"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
)

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(w, finalWord)
}
