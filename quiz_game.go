package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	const defaultQuiz = "problems.csv"

	// Create flag so user can provide quiz file with custom name
	var quizFlag = flag.String("file", defaultQuiz, "name of the custom quiz file")
	var timeFlag = flag.Int("limit", 30, "Time limit to finish the quiz.")
	flag.Parse()

	// Open the file for reading by the csv package
	open_quiz_file, open_err := os.Open(*quizFlag)
	if open_err != nil {
		log.Fatal(open_err)
	}
	defer open_quiz_file.Close()

	// Read content of file
	read_quiz := csv.NewReader(open_quiz_file)
	questions, read_err := read_quiz.ReadAll()
	if read_err != nil {
		log.Fatal(read_err)
	}

	// After adding in the flag, you will need to convert to time.Duration(flag)
	gameTimer := time.NewTimer(time.Duration(*timeFlag) * time.Second)

	var score = 0

	for i := range len(questions) {
		fmt.Printf("Question %d: %v", i+1, questions[i][0])
		inputCh := make(chan string)
		// New go routine
		go func() {
			var userInput string
			fmt.Scan(&userInput)
			inputCh <- userInput
		}()
		select {
		case <-gameTimer.C:
			fmt.Printf("Total questions %d\nYour score: %d/%d", len(questions), score, len(questions))
			return
		case userInput := <-inputCh:
			if userInput == questions[i][1] {
				score++
			}
		}
	}
}

// Changes for practice!!!

// Add a function to parse file input
