package greeter

import "fmt"

type language int

const (
	ENGLISH language = iota
	SPANISH
	ITALIAN
)

func Greeter(name string, language language) string {
	greeting := "Hello, %s"
	if language == SPANISH {
		greeting = "Hola, %s"
	}
	if language == ITALIAN {
		greeting = "Ciao, %s"
	}
	return fmt.Sprintf(greeting, name)
}
