package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"gophercises/quiz_game/quiz"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "csv file with format of question,answer")
	flag.Parse()

	problems, err := quiz.Start(csvFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "quiz_game: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Starting game: ")

	var correctAns int
	in := bufio.NewScanner(os.Stdin)
	for c, problem := range *problems {
		fmt.Printf("Prob #%d: %s: ", c+1, problem.Ques)
		in.Scan()
		if problem.Ans == in.Text() {
			correctAns++
		}
	}

	fmt.Printf("\n%d out of %d are correct.\n", correctAns, len(*problems))
}
