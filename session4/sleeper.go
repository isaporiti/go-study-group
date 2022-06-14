package countdown

import "time"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	Duration  time.Duration
	SleepFunc func(d time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.SleepFunc(s.Duration)
}
