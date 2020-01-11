package main

import (
	server "cor_gophercises/adventure/pkg/bookserver"
	bookloader "cor_gophercises/adventure/pkg/storybook"
	"fmt"
)

func main() {
	//1. Read Book
	//2. Initialise Web Server
	//3. Serve Book
	fmt.Println()
	server.StartServer()
	bookloader.LoadBook()
}
