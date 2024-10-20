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
	flag.StringVar(&quizFile, "quizFile", "problems.csv", "Name of the quizCSV file")

	flag.Parse()

	file, err := os.Open(quizFile)

	// Checks for the error
	if err != nil {
		log.Fatal("Error while reading the file", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)

	// ReadAll reads all the questions from the CSV file
	// and Returns them as slice of slices of string
	// and an error if any
	quizCSV, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
		return
	}

	quiz := make(map[string]string)

	for _, line := range quizCSV {
		quiz[line[0]] = line[1]
	}

	fmt.Println(quiz)

	// Loop to iterate through
	// Add each item to a struct

	// fmt.Printf("You scored %d out of %d!", score, len(quizCSV))
}
