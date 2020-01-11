package storybook

import (
	"cor_gophercises/adventure/pkg/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const firstPage = "intro"

type BookFile map[string]BookPage

type BookPage struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []BookOptions `json:"options"`
}

type BookOptions struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func LoadBook(fileName string) (error, *BookFile) {
	var bookData BookFile

	bookFile, fileReadError := os.Open(fileName)

	if fileReadError != nil {
		logger.Error.Println("Error opening book file:", fileName)
		logger.Error.Println(fileReadError.Error())
		return fileReadError, nil
	}

	bookFileData, _ := ioutil.ReadAll(bookFile)

	jsonParseError := json.Unmarshal(bookFileData, &bookData)

	if jsonParseError != nil {
		logger.Error.Println("Error opening book file:", fileName)
		logger.Error.Println(jsonParseError)
		return jsonParseError, nil
	}

	fmt.Println(bookData)

	return nil, nil
}
