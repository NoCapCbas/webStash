package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title   string
	Message string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// main portfolio handler

	// parse template file
	tmpl, err := template.ParseFiles(
		"templates/index.html",
		"templates/partials/nav.html",
		"templates/partials/hero.html",
		"templates/partials/about.html",
		"templates/partials/skills.html",
		"templates/partials/projects.html",
		"templates/partials/contact.html",

		"templates/projects/savviURL.html",
		"templates/projects/savviVerifyEmail.html",
		"templates/projects/savviPingPong.html",
	)
	if err != nil {
		log.Println("Error parsing templates: ", err)
		http.Error(w, "Could not load template: ", http.StatusInternalServerError)
		return
	}

	// create some data to pass to the template
	data := PageData{
		Title:   "title",
		Message: "message",
	}

	// execute the template and pass the data
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template: ", err)
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}

}

func main() {
	// serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// pass porfolio handler to server
	http.HandleFunc("/", indexHandler)

	// start server
	log.Println("Listening on port :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
