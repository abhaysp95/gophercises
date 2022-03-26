package quiz

import (
	"encoding/csv"
	"os"
)

type Problem struct {
	Ques string
	Ans string
}

func Start(csvFile *string) (*[]Problem, error) {
	file, err := os.Open(*csvFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	csvr := csv.NewReader(file)
	lines, err := csvr.ReadAll()
	if err != nil {
		return nil, err
	}

	return parseLines(lines), nil
}

func parseLines(lines [][]string) *[]Problem {
	problems := make([]Problem, len(lines))  // is this on heap ? (cause problems is still accessible outside)
	for c, line := range lines {
		problems[c] = Problem {
			Ques: line[0],
			Ans: line[1],
		}
	}
	return &problems
}
