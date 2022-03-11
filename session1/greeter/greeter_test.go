package greeter

import "testing"

// t.Run helps us group use cases
func TestGreeter(t *testing.T) {
	t.Run("greets Codurance", func(t *testing.T) {
		got := Greeter("Codurance", "")
		want := "Hello, Codurance"

		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	})

	t.Run("greets Ignacio", func(t *testing.T) {
		got := Greeter("Ignacio", "")
		want := "Hello, Ignacio"

		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	})

	t.Run("greets in Spanish", func(t *testing.T) {
		got := Greeter("Ignacio", Spanish)
		want := "Hola, Ignacio"

		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	})

	t.Run("greets in Italian", func(t *testing.T) {
		got := Greeter("Ignacio", Italian)
		want := "Ciao, Ignacio"

		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	})
}

//Test tables help us reduce code repetition
func TestTableGreeter(t *testing.T) {
	tests := []struct {
		testName string
		language language
		want     string
	}{
		{testName: "greets in English", language: English, want: "Hello, Go Study Group"},
		{testName: "greets in Spanish", language: Spanish, want: "Hola, Go Study Group"},
		{testName: "greets in Italian", language: Italian, want: "Ciao, Go Study Group"},
	}
	const name = "Go Study Group"
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			got := Greeter(name, test.language)

			if got != test.want {
				t.Errorf("want %s, got %s", test.want, got)
			}
		})
	}
}
