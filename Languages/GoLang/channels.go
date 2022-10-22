package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "> Hello World welcome to Channels <" }()

	msg := <-messages

	fmt.Println(msg)
}
