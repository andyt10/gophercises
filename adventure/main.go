package main

import (
	server "cor_gophercises/adventure/pkg/bookserver"
	"cor_gophercises/adventure/pkg/logger"
	bookloader "cor_gophercises/adventure/pkg/storybook"
	"flag"
	"io/ioutil"
	"os"
)

func main() {
	//1. Read Book
	//2. Initialise Web Server
	//3. Serve Book
	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stdout)
	logger.Info.Println("Starting CYOA book")

	bookName := parseArgs()
	logger.Info.Println("Using Book File:", bookName)

	bookloader.LoadBook(bookName)

	server.StartServer()

}

func parseArgs() string {
	logger.Trace.Println("Parsing arguments")
	fileName := flag.String("bookdata", "book_data.json", "A JSON file (to required spec) of a CYOA book.")

	logger.Trace.Println("Arguments parsed")
	return *fileName
}
