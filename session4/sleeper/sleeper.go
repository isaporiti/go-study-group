package sleeper

import "time"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(d time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func NewConfigurableSleeper(duration time.Duration, sleep func(d time.Duration)) *ConfigurableSleeper {
	return &ConfigurableSleeper{duration: duration, sleep: sleep}
}
