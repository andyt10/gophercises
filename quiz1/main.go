package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	questionData := readQuestionFile()
	totalMark := askQuestions(questionData)
	fmt.Println("Total Marks:", totalMark)
}

func askQuestions(questionData []QuestionLine) int {
	totalMark := 0
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < len(questionData); i++ {
		fmt.Println(fmt.Sprintf("What is: %v?", questionData[i].question))
		answer, _ := reader.ReadString('\n')
		answer = strings.Replace(answer, "\n", "", -1)

		if answer == questionData[i].answer {
			fmt.Println("Correct!")
			totalMark += 1
		} else {
			fmt.Println("You Suck!")
			fmt.Println(questionData[i].answer)
		}
	}

	return totalMark
}

type QuestionLine struct {
	question string
	answer   string
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
