package main

import (
	"encoding/json"
	"flag"
	"fmt"
	story "gophercise-cyoa"
	"log"
	"os"
)

func main() {
	// Use flags to allow user input file location
	fileName := flag.String("filename", "gopher.json", "file location of story JSON")
	flag.Parse()

	// Read file from OS to bytes
	jsonPayload, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("could not open local json file: %v", err)
	}

	// Unmarshal JSON to type
	var parsedJson story.Chapter
	decoder := json.NewDecoder(jsonPayload)
	err = decoder.Decode(&parsedJson)
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
