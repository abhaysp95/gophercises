package main

import (
	"flag"
	"fmt"
	"gophercises/cyoa"
	"os"
)

func main() {
	var fileName = flag.String("file", "story.json", "provide the path to json file containing story")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	story, err := cyoa.ReadStory(file)

	fmt.Printf("%T\n%+[1]v\n", story)
}
