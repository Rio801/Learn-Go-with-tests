package main

import "fmt"

var HelloPrefix string

func Hello(name, language string) string {
	switch language {
	case "Spanish":
		HelloPrefix = "Hola, "
	case "French":
		HelloPrefix = "Bonjour, "
	default:
		HelloPrefix = "Hello, "
	}
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s%s", HelloPrefix, name)
}

func main() {
}
