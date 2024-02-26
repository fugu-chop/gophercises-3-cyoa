package main

import (
	"flag"
	story "gophercise-cyoa"
	"log"
	"net/http"
	"text/template"
)

type storyHandler struct {
	StoryData story.Chapter
	Template  *template.Template
}

func (s storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Template.Execute(w, s.StoryData[story.StoryStart])
}

func main() {
	// Use flags to allow user input file location
	fileName := flag.String("filename", "gopher.json", "file location of story JSON")
	templateName := flag.String("template", "story_template.html", "file location of HTML template")
	flag.Parse()

	jsonStoryData, err := story.ParseJSON(fileName)
	if err != nil {
		log.Fatalf("could not open local json file: %v", err)
	}

	// Populate HTML Template
	template, err := populateTemplate(*templateName)
	if err != nil {
		log.Fatalf("could not open parse HTML template: %v", err)
	}

	// Create handler with data
	storyDataHandler := &storyHandler{
		StoryData: *jsonStoryData,
		Template:  template,
	}

	// start HTTP server
	startServer(storyDataHandler)
}

func startServer(handler *storyHandler) {
	mux := http.NewServeMux()
	mux.Handle("/home", *handler)
	http.ListenAndServe(":8080", mux)
}

func populateTemplate(templateLocation string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(templateLocation)
	if err != nil {
		log.Printf("could not open local json file: %v", err)
		return nil, err
	}

	return tmpl, nil
}
