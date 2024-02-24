package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const storyStart = "intro"

type Chapter map[string]Page
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type Page struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

func main() {
	// Read file from OS to bytes
	jsonPayload, err := os.ReadFile("./gopher.json")
	if err != nil {
		log.Fatalf("could not open local json file: %v", err)
	}

	// Unmarshal JSON to type
	var parsedJson Chapter
	err = json.Unmarshal(jsonPayload, &parsedJson)
	if err != nil {
		log.Fatalf("could not parse json: %v", err)
	}

	// Start the story
	intro, ok := parsedJson[storyStart]
	if !ok {
		log.Fatalf("story is missing an intro")
	}

	fmt.Println(intro)
}
