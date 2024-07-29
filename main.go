package main

import "fmt"

func main() {
	fmt.Println("hello to my new program")
	fmt.Println(greeter("John"))
}

func greeter(name string) string {
	return "Hello " + name + ". How are you?"
}
