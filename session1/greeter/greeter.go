package greeter

import "fmt"

type language string

const English language = "English"
const Spanish language = "Spanish"
const Italian language = "Italian"

func Greeter(name string, language language) string {
	greeting := "Hello, %s"
	if language == Spanish {
		greeting = "Hola, %s"
	}
	if language == Italian {
		greeting = "Ciao, %s"
	}
	return fmt.Sprintf(greeting, name)
}
