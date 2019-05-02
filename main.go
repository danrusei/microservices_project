package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

var (
	indexTmpl = template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))
)

type logo struct {
	Link string
}

type indexdata struct {
	Logos       []logo
	Style       string
	RequestTime string
}

func main() {
	http.HandleFunc("/", indexHandler)

	// Serve static files out of the public directory.
	// By configuring a static handler in app.yaml, App Engine serves all the
	// static content itself. As a result, the following two lines are in
	// effect for development only.
	public := http.StripPrefix("/public/", http.FileServer(http.Dir("public")))
	http.Handle("/public/", public)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := indexdata{
		Logos: []logo{
			{Link: "/public/gopher.jpg"},
			{Link: "/public/google_cloud.jpg"},
			{Link: "/public/tensorflow.jpg"},
		},
		Style:       "/public/style.css",
		RequestTime: time.Now().Format(time.RFC822),
	}

	if err := indexTmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
