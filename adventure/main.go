package main

import (
	server "cor_gophercises/adventure/pkg/bookserver"
	"cor_gophercises/adventure/pkg/logger"
	"flag"
	"io/ioutil"
	"os"
)

func main() {
	//1. Read Book
	//2. Initialise Web Server
	//3. Serve Book
	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stdout)

	/*
		logger.Info.Println("Starting CYOA book")

		bookName := parseArgs()
		logger.Info.Println("Using Book File:", bookName)

		err, bookData := bookloader.LoadBook(bookName)

		if err != nil {
			logger.Error.Println("Exiting due to book read error.")
			os.Exit(1)
		}

		fmt.Println(bookData)
	*/

	server.StartServer()

}

func parseArgs() string {
	logger.Trace.Println("Parsing arguments")
	fileName := flag.String("bookdata", "book_data.json", "A JSON file (to required spec) of a CYOA book.")

	logger.Trace.Println("Arguments parsed")
	return *fileName
}
