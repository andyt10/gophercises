package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	runQuiz()
}

func runQuiz() {
	fileName, limit := paseArgs()
	questionData := readQuestionFile(fileName)
	totalMark := askQuestions(questionData, limit)
	fmt.Println("Total Marks:", totalMark)
}

func paseArgs() (string, int) {
	fileName := flag.String("csv", "problems.csv", "A CSV File of Questions")
	timeLimitSeconds := flag.Int("time", 30, "Time limit for quiz")
	flag.Parse()
	return *fileName, *timeLimitSeconds
}

func askQuestions(questionData []QuestionLine, timeLimitSeconds int) int {
	timer1 := time.NewTimer(time.Duration(timeLimitSeconds) * time.Second)
	//Quiz
	totalMark := 0
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < len(questionData); i++ {
		fmt.Println(fmt.Sprintf("What is: %v?", questionData[i].question))

		answerData := make(chan string)

		go func() {
			answer, _ := reader.ReadString('\n')
			answer = strings.Replace(answer, "\n", "", -1)
			answerData <- answer
		}()

		select {
		case <-timer1.C:
			fmt.Println("-- Out Of Time -- ")
			return totalMark
		case answer := <-answerData:
			if answer == questionData[i].answer {
				totalMark += 1
			}
		}
	}

	return totalMark
}

type QuestionLine struct {
	question string
	answer   string
}

func readQuestionFile(fileName string) []QuestionLine {

	csvFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error Opening Questions File: ", err)
		os.Exit(1)
	}

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
