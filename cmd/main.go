package main

import (
	"encoding/json"
	"fmt"
	story "gophercise-cyoa"
	"log"
	"os"
)

func main() {
	// Read file from OS to bytes
	jsonPayload, err := os.ReadFile("./gopher.json")
	if err != nil {
		log.Fatalf("could not open local json file: %v", err)
	}

	// Unmarshal JSON to type
	var parsedJson story.Chapter
	err = json.Unmarshal(jsonPayload, &parsedJson)
	if err != nil {
		log.Fatalf("could not parse json: %v", err)
	}

	// Start the story
	intro, ok := parsedJson[story.StoryStart]
	if !ok {
		log.Fatalf("story is missing an intro")
	}

	fmt.Println(intro)
}
