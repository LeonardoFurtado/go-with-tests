package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const portuguese = "Portuguese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Olá, "

func gettingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return gettingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", "english"))
}
