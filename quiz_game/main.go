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
var dur = flag.Duration("dur", time.Duration(5 * time.Second), " ")

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
	timer := time.NewTimer(*limit)
	ansCh := make(chan string)
	// updateAnsCh := make(chan bool)
	// go askQuestion(problems, updateAnsCh, finishCh)

	for c, problem := range *problems {
		select {
		case <-timer.C:
			fmt.Println("Timer Over!!!")
			return
		default:
			go askQuestion(io, c, problem, ansCh)
		}
	}

/* loop:
	for {
		select {
		case <-timer.C:
			fmt.Println("Time Over!!!")
			break loop
		case <-finishCh:
			fmt.Println("Game Over!!!")
			return
		default:
		}
	} */
	fmt.Printf("\n%d out of %d are correct.\n", correctAns, len(*problems))
}

func askQuestion(io *bufio.Scanner, c int, problem quiz.Problem, ansCh chan<- string) {
	fmt.Printf("Prob #%d: %s: ", c+1, problem.Ques)
	timer := time.NewTimer(*dur)
	for {
		select {
		case <-timer.C:
		}
	}
}

/* func askQuestion(problems *[]quiz.Problem, updateAns chan<- bool, finish chan<- struct{}) {
	in := bufio.NewScanner(os.Stdin)
	ansCh := make(chan string)
	for c, problem := range *problems {
		fmt.Printf("Prob #%d: %s: ", c+1, problem.Ques)
		timer := time.NewTimer(*dur)
		getAns(in, ansCh)
		select {
		case <-timer.C:
			fmt.Printf("\ntime up\n")
			continue
		case ans := <-ansCh:
			fmt.Println("update ans")
			if ans == problem.Ans {
				updateAns<- true
			} else {
				updateAns<- false
			}
		}
	}
}

func getAns(in *bufio.Scanner, ans chan<- string) {
	in.Scan()
	ans<- in.Text()
} */
