package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	var quizFile string
	flag.StringVar(&quizFile, "quizFile", "problems.csv", "Name of the quiz file")

	flag.Parse()

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.File
	file, err := os.Open(quizFile)

	// Checks for the error
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	// Closes the file
	defer file.Close()

	// The csv.NewReader() function is called in
	// which the object os.File passed as its parameter
	// and this creates a new csv.Reader that reads
	// from the file
	reader := csv.NewReader(file)

	// ReadAll reads all the questions from the CSV file
	// and Returns them as slice of slices of string
	// and an error if any
	quiz, err := reader.ReadAll()

	// Checks for the error
	if err != nil {
		fmt.Println("Error reading records")
	}

	var answer string
	var score = 0
	// Loop to iterate through
	// and print each of the string slice

	for _, line := range quiz {
		problem := line[0]
		solution := line[1]

		fmt.Printf(problem + ": \n")
		fmt.Scanf("%s\n", &answer)

		if answer == solution {
			score++
		}
	}
	fmt.Printf("You scored %d out of %d!", score, len(quiz))
}
