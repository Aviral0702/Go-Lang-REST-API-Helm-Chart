package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("hello to my new program")
	fmt.Println(greeter("John"))
	info()
}

func greeter(name string) string {
	return "Hello " + name + ". How are you?"

}

func info() {
	resp, err := http.Get("http://api.weatherapi.com/v1/current.json?key=9a340ab233594c87982144702242707&q=London&aqi=no")
	if err != nil {
		fmt.Println("Error in fetching data")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(body)
	fmt.Println(content)

}
