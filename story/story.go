package story

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

const StoryStart = "intro"

type StoryHandler struct {
	StoryData Chapter
	Template  *template.Template
}

func (s StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	chapter, _ := strings.CutPrefix(r.URL.Path, "/")
	_, ok := s.StoryData[chapter]
	if chapter == "" || !ok {
		chapter = StoryStart
	}

	err := s.Template.Execute(w, s.StoryData[chapter])
	if err != nil {
		log.Printf("error processing template: %v", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}

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
	jsonPayload, err := os.Open(*fileLocation)
	if err != nil {
		log.Printf("could not open local json file: %v", err)
		return nil, err
	}

	var parsedJson Chapter
	decoder := json.NewDecoder(jsonPayload)
	err = decoder.Decode(&parsedJson)
	if err != nil {
		log.Printf("could not parse json: %v", err)
		return nil, err
	}

	return &parsedJson, nil
}
