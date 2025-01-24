package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchhHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"


func Hello(name string, language string) string {  // function returns a string

	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix;

	switch language {
		case spanish: 
			prefix = spanishHelloPrefix
		case french: 
			prefix = frenchhHelloPrefix
	}

return prefix + name


}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}