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

	action := args[0]
	args = args[1:]

	repository := args[0]
	count := 4
	if len(args) > 1 {
		countArg, err := strconv.Atoi(args[1])

		if err != nil {
			log.Fatalf("Invalid count: %s", args[1])
		}

		count = countArg
	}

	switch action {
	case "images":
		imagesCount(repository, count)
	case "latest":
		latest(repository)
	default:
		log.Fatalf("Invalid action: %s", action)
	}
}

func imagesCount(repository string, count int) {
	images, err := godohuver.GetTagsFromRepository(repository, count)

	if err != nil {
		log.Fatal(err)
	}

	for _, image := range images {
		fmt.Printf("%s:%s - %s\n", image.Name, image.Tag, image.LastPushed)
	}
}

func latest(repository string) {
	latest, err := godohuver.GetLatestImage(repository)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Latest: %s:%s - %s\n", latest.Name, latest.Tag, latest.LastPushed)
}
