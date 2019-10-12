package main

import (
	"fmt"

	"github.com/augmentable-dev/gitcheck"
)

func main() {
	metrics, err := gitcheck.GetMetrics("https://github.com/augmentable-dev/tickgit", 30)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d Commits\n", len(metrics.Commits))
	fmt.Printf("%d Committers\n", len(metrics.UniqueCommiters))
}
