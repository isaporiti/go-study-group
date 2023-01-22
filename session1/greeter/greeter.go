package greeter

import "fmt"

type language string

const (
	ENGLISH language = "English"
	SPANISH language = "Spanish"
	ITALIAN language = "Italian"
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
