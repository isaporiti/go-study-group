package countdown

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("prints from 3 to 1 and finally 'Go!'", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}

		Countdown(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!
`
		if got != want {
			t.Errorf("want countdown from 3 to 1, got '%s'", got)
		}
	})

	t.Run("it sleeps after printing each line", func(t *testing.T) {
		spyOperations := &SpyCountdownOperations{}

		Countdown(spyOperations, spyOperations)

		want := []CountdownOperation{
			write, sleep,
			write, sleep,
			write, sleep,
			write,
		}

		if !reflect.DeepEqual(spyOperations.Calls, want) {
			t.Errorf("want three writes and sleeps in order, got %v", spyOperations.Calls)
		}
	})
}

type SpySleeper struct {
}

func (s *SpySleeper) Sleep() {}

type SpyCountdownOperations struct {
	Calls []CountdownOperation
}

type CountdownOperation string

const (
	write CountdownOperation = "write"
	sleep CountdownOperation = "sleep"
)

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
	return
}
