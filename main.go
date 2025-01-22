package main

import "fmt"

const englishHelloPrefix = "Hello, "


func Hello(name string) string {  // function returns a string
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}