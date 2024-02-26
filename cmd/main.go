package main

import (
	"flag"
	"fmt"
	story "gophercise-cyoa"
	"log"
	"net/http"
)

type storyHandler struct {
	StoryData story.Chapter
}

func (s storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(s.StoryData[story.StoryStart])
}

func main() {
	// Use flags to allow user input file location
	fileName := flag.String("filename", "gopher.json", "file location of story JSON")
	flag.Parse()

	jsonStoryData, err := story.ParseJSON(fileName)
	if err != nil {
		log.Fatalf("could not open local json file: %v", err)
	}

	// Create handler with data
	storyDataHandler := &storyHandler{
		StoryData: *jsonStoryData,
	}

	// start HTTP server
	mux := http.NewServeMux()
	mux.Handle("/home", *storyDataHandler)
	http.ListenAndServe(":8080", mux)
}
