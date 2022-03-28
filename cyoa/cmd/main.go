package main

import (
	"flag"
	"fmt"
	"gophercises/cyoa"
	"log"
	"net/http"
	"os"
)

func main() {
	var fileName = flag.String("file", "story.json", "provide the path to json file containing story")
	var port = flag.Int("port", 3030, "provide port to run the program")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	story, err := cyoa.ReadStory(file)

	h := cyoa.NewHandler(story)
	fmt.Printf("Running on port :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
