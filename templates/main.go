package main

import (
	"html/template"
	"log"
	"os"
)

const STATIC_SRC = "/src/github.com/joyrexus/sandbox/templates/static"

var STATIC = os.Getenv("GOPATH") + STATIC_SRC
var templates = template.Must(template.ParseFiles(
	STATIC + "/study_view.html",
	STATIC + "/study_edit.html",
))

type Study struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
}

func main() {
	study := &Study{"study_a", "description of study_a"}

	t := "study_view.html"
	if err := templates.ExecuteTemplate(os.Stdout, t, study); err != nil {
		log.Fatalf("couldn't execute template %q: %v", t, err)
	}

	t = "study_edit.html"
	if err := templates.ExecuteTemplate(os.Stdout, t, study); err != nil {
		log.Fatalf("couldn't execute template %q: %v", t, err)
	}
}
