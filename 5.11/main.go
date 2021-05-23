package main

import (
	"fmt"
	"log"
	"sort"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},

	"linear algebra": {"calculus"},
}

func main() {

	sorted, err := topoSort(prereqs)
	if err != nil {
		log.Fatal(err)
	}
	for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string) error
	visitAll = func(items []string) error {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				err := visitAll(m[item])
				if err != nil {
					return err
				}
				order = append(order, item)
			} else {
				cyclic := true
				for _, v := range order {
					if item == v {
						cyclic = false
					}
				}
				if cyclic {
					return fmt.Errorf("error, acyclic item: %s", item)
				}
			}
		}
		return nil
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	err := visitAll(keys)
	if err != nil {
		return nil, err
	}
	return order, nil
}
