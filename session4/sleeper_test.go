package countdown

import (
	"testing"
	"time"
)

func TestConfigurableSleeper(t *testing.T) {
	var calls int
	var duration time.Duration
	sleep := func(d time.Duration) { calls++; duration = d }

	sleeper := ConfigurableSleeper{Duration: 3 * time.Second, SleepFunc: sleep}

	sleeper.Sleep()

	if calls != 1 {
		t.Errorf("want 1 Sleep call, got %d", calls)
	}
	if sleeper.Duration != duration {
		t.Errorf("want Sleep duration of %d, got %d", sleeper.Duration, duration)
	}
}
