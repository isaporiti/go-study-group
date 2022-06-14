package countdown

import (
	"testing"
	"time"
)

type SpySleepFunc struct {
	Calls    int
	Duration time.Duration
}

func (s *SpySleepFunc) Sleep(d time.Duration) {
	s.Calls++
	s.Duration = d
}

func TestConfigurableSleeper(t *testing.T) {
	sleepFunc := SpySleepFunc{}

	const duration = 3 * time.Second
	sleeper := ConfigurableSleeper{duration, sleepFunc.Sleep}

	sleeper.Sleep()

	if sleepFunc.Calls != 1 {
		t.Errorf("want 1 Sleep call, got %d", sleepFunc.Calls)
	}
	if sleepFunc.Duration != duration {
		t.Errorf("want Sleep Duration of %d, got %d", duration, sleepFunc.Duration)

	}
}
