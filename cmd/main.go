package main

import (
	"flag"
	"fmt"
	story "gophercise-cyoa"
	"log"
)

func main() {
	// Use flags to allow user input file location
	fileName := flag.String("filename", "gopher.json", "file location of story JSON")
	flag.Parse()

	storyData, err := story.ParseJSON(fileName)
	if err != nil {
		log.Fatalf("could not open local json file: %v", err)
	}

	// Start the story
	intro, ok := storyData[story.StoryStart]
	if !ok {
		log.Fatalf("story is missing an intro")
	}

	fmt.Println(intro)
}
