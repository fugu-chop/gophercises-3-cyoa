package main

import (
	"flag"
	story "gophercise-cyoa/story"
	"log"
	"net/http"
	"text/template"
)

func main() {
	// Use flags to allow user input file location
	fileName := flag.String("filename", "gopher.json", "file location of story JSON")
	templateName := flag.String("template", "story_template.html", "file location of HTML template")
	flag.Parse()

	jsonStoryData, err := story.ParseJSON(fileName)
	if err != nil {
		log.Fatalf("could not open local json file: %v", err)
	}

	template, err := populateTemplate(*templateName)
	if err != nil {
		log.Fatalf("could not open parse HTML template: %v", err)
	}

	storyDataHandler := story.CreateHandler(jsonStoryData, template)

	mux := http.NewServeMux()
	mux.Handle("/", storyDataHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func populateTemplate(templateLocation string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(templateLocation)
	if err != nil {
		log.Printf("could not open local json file: %v", err)
		return nil, err
	}

	return tmpl, nil
}
