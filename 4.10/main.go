// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/edvardgospel/sandbox/4.10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	issues := map[string][]*github.Issue{}
	for _, item := range result.Items {
		if item.CreatedAt.After(time.Now().AddDate(0, -1, 0)) {
			issues["oneMonth"] = append(issues["oneMonth"], item)
		} else if item.CreatedAt.After(time.Now().AddDate(-1, 0, 0)) {
			issues["oneYear"] = append(issues["oneYear"], item)
		} else {
			issues["moreYear"] = append(issues["moreYear"], item)
		}
	}
	for k, issue := range issues {
		fmt.Println("--- ", k)
		for _, v := range issue {
			fmt.Printf("#%-5d %9.9s %.55s\n", v.Number, v.User.Login, v.Title)
		}
	}
}
