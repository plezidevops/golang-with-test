package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const haitian = "Haitian"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const haitianHelloPrefix = "Bonjou, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case haitian:
		prefix = haitianHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix

	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
