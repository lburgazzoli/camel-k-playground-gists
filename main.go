package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v32/github"
)

func main() {
	client := github.NewClient(nil)
	g, _, err := client.Gists.Get(context.Background(), os.Args[1])
	if err != nil {
		panic(err)
	}

	for _, v := range g.Files {
		fmt.Printf("%s", *v.Filename)
		if v.Language != nil {
			fmt.Printf(", language=%s", *v.Language)
		}
		fmt.Print("\n")
	}
}
