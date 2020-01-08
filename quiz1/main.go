package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {

	questionData := readQuestionFile()
}

type QuestionLine struct {
	Question string
	Answer   string
}

func readQuestionFile() []QuestionLine {
	fileName := "problems.csv"
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var questionData []QuestionLine
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		questionData = append(questionData, QuestionLine{line[0], line[1]})
	}
	return questionData
}
