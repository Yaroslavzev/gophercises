package main

import (
	"cyoa"
	"cyoa/handlers"
	"cyoa/templates"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	storeFile := flag.String("stories", "gophers.json", "path to file with stories")
	flag.Parse()

	file, err := os.Open(*storeFile)
	if err != nil {
		panic(err)
	}

	stories, err := cyoa.ParseJSON(file)

	if err != nil {
		panic(err)
	}
	tmp := template.Must(template.New("").Parse(templates.StoryHandlerTmpl))
	h := handlers.NewHandler(stories,
		handlers.WithTemplate(tmp),
		handlers.WithPathFunc(PathFn))

	mux := http.NewServeMux()
	mux.Handle("/story/", h)

	fmt.Println("Starting the server on port: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func PathFn(r *http.Request) string {
	path := r.URL.Path

	return path[len("/story/"):]
}
