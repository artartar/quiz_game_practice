package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	const defaultQuiz = "problems.csv"

	// Create flag so user can provide quiz file with custom name
	var quizFlag = flag.String("file", defaultQuiz, "name of the custom quiz file")

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

	var score = 0
	var userInput string

	for i := range len(questions) {
		fmt.Printf("Question %d: %v", i+1, questions[i][0])
		fmt.Scan(&userInput)
		fmt.Print("\n")

		if userInput == questions[i][1] {
			score++
		}
	}
	fmt.Printf("Total questions %d\nYour score: %d/%d", len(questions), score, len(questions))
}

// Changes for practice!!!

// Add a function to parse file input
