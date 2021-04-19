package main

import (
	"fmt"
	"log"
	"os"

	"github.com/edvardgospel/sandbox/4.11/github"
)

func main() {
	issue, err := github.Create()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)

	}
	fmt.Println(issue)
	issues, err := github.ReadAll()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(issues)
}
