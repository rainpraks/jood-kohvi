package main

import (
	"ascii-art-web/asciiart"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	// reads one or more template and parses them to create a template.
	tpl, _ = template.ParseGlob("templates/*.html")
	//set up a request handler for a specific URL path in GO web server.
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	// Starts HTTP server and begins listening for incoming HTTP requests.
	http.ListenAndServe(":8080", nil)
}

// IndexHandler function is responsible for generating the response to HTTP requests
func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		errorHandler(w, "Page not found", http.StatusNotFound)
		return
	}
	tpl.ExecuteTemplate(w, "webpage.html", nil)
}

// ErrorHandler function is responsible to generate error messages 500,404,400.
func errorHandler(w http.ResponseWriter, message string, statuscode int) {
	data := struct {
		StatusCode int
		Message    string
	}{
		StatusCode: statuscode,
		Message:    message,
	}
	w.WriteHeader(statuscode)
	tpl.ExecuteTemplate(w, "error.html", data)
}

// AsciiArtHandler function is a key component of web application responsible for handling incoming HTTP POST requests.
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	//Is it POST method? Post method is used to send data to the server to be processed
	if r.Method == http.MethodPost {

		asciiText := r.FormValue("asciiText")
		for _, ch := range asciiText {
			if ch > 126 {
				errorHandler(w, "Bad request", http.StatusBadRequest)
				return
			}
		}
		asciiFont := r.FormValue("asciiFont")

		asciiArt := asciiart.GenerateAsciiArt(asciiText, asciiFont)
		if asciiArt == "" {

			errorHandler(w, "Error loading character map", http.StatusInternalServerError)
			return
		}
		tpl.ExecuteTemplate(w, "webpage.html", asciiArt)

	}
}
