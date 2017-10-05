package main

import (
	"flag"
	"fmt"

	embedtext "github.com/brimstone/go-embed-text"
)

func main() {

	messagePtr := flag.String("msg", "", "a string")
	flag.Parse()

	// If the user doesn't specify a message
	if *messagePtr == "" {
		message := embedtext.Read()
		if message == "" {
			fmt.Println("No message")
		} else {
			fmt.Printf("Message is: %s\n", message)
		}
		return
	}

	fmt.Printf("Storing: %s\n", *messagePtr)
	embedtext.Embed(*messagePtr)
}
