package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	// Return a greeting in the form of "Hello, <name>"
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World"))
}
