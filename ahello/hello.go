package main

import "fmt"

const portuguese = "Portuguese"
const french = "French"
const portugueseHelloPrefix = "Olá, "
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	// prefix is a named return value, it is created in func scope
	// and it is assigned "zero" value i.e. ""
	switch language {
	case portuguese:
		prefix = portugueseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return // implicitly returns named return val
}

func main() {
	fmt.Println(Hello("world", ""))
	// Println is a side effect: printing to stdout
}
