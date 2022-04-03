package main

import (
	"gophercises/task/cmd"
	"gophercises/task/db"
	"log"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	cmd.Execute()
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
