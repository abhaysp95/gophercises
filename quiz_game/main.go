package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"gophercises/quiz_game/quiz"
)

var csvFileName = flag.String("csv", "problems.csv", "csv file with format of question,answer")
var limit = flag.Duration("limit", time.Duration(10 * time.Second), "enter time limit for each question")

func main() {
	flag.Parse()

	problems, err := quiz.Start(csvFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "quiz_game: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Starting game: ")

	askQuiz(problems)
}

func askQuiz(problems *[]quiz.Problem) {
	var correctAns int
	in := bufio.NewScanner(os.Stdin)
	timer := time.NewTimer(*limit)
	ansCh := make(chan string)
	for c, problem := range *problems {
		fmt.Printf("Prob #%d: %s: ", c+1, problem.Ques)
		go askQuestion(in, c, problem, ansCh)
		select {
		case <-timer.C:
			fmt.Printf("\n%d out of %d are correct.\n", correctAns, len(*problems))
			return
		case ans := <-ansCh:
			if ans == problem.Ans {
				correctAns++
			}
		/* default:  // just for testing
			fmt.Println("this") */
		}
	}

	fmt.Printf("\n%d out of %d are correct.\n", correctAns, len(*problems))
}

func askQuestion(in *bufio.Scanner, c int, problem quiz.Problem, ans chan<- string) {
	in.Scan()
	ans <-in.Text()
}
