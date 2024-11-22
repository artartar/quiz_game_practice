package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	// Open the file for reading by the csv package
	open_quiz_file, open_err := os.Open("problems.csv")
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

	fmt.Print(questions)
}
