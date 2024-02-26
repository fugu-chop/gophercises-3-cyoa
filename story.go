package story

import (
	"encoding/json"
	"log"
	"os"
)

const StoryStart = "intro"

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

func ParseJSON(fileLocation *string) (*Chapter, error) {
	// Read file from OS to bytes
	jsonPayload, err := os.Open(*fileLocation)
	if err != nil {
		log.Printf("could not open local json file: %v", err)
		return nil, err
	}

	// Unmarshal JSON to type
	var parsedJson Chapter
	decoder := json.NewDecoder(jsonPayload)
	err = decoder.Decode(&parsedJson)
	if err != nil {
		log.Printf("could not parse json: %v", err)
		return nil, err
	}

	return &parsedJson, nil
}
