package countdown

import (
	"testing"
	"time"
)

func TestConfigurableSleeper(t *testing.T) {
	var calls int
	var duration time.Duration
	sleep := func(d time.Duration) { calls++; duration = d }
	sleeper := NewConfigurableSleeper(3*time.Second, sleep)

	sleeper.Sleep()

	if calls != 1 {
		t.Errorf("want 1 Sleep call, got %d", calls)
	}
	if sleeper.duration != duration {
		t.Errorf("want Sleep duration of %d, got %d", sleeper.duration, duration)
	}
}
