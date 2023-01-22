package main

import (
	"fmt"
	g "greeter/greeter"
)

func main() {
	greeting := g.Greeter("Go Study Group", g.SPANISH)
	fmt.Println(greeting)
}
