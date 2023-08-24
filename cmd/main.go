package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/zaptross/godohuver"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("No arguments provided")
	}

	repository := args[0]
	count := 4
	if len(args) > 1 {
		countArg, err := strconv.Atoi(args[1])

		if err != nil {
			log.Fatalf("Invalid count: %s", args[1])
		}

		count = countArg
	}

	images, err := godohuver.GetTagsFromRepository(repository, count)

	if err != nil {
		log.Fatal(err)
	}

	for _, image := range images {
		fmt.Printf("%s:%s - %s\n", image.Name, image.Tag, image.LastPushed)
	}
}
