package main

import (
	"fmt"
	"groupie-trackers/controllers"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	http.Handle("/view/", http.StripPrefix("/view/", http.FileServer(http.Dir("view"))))
	// reads one or more template and parses them to create a template.
	tpl, _ = template.ParseGlob("view/*.html")
	//set up a request handler for a specific URL path in GO web server.
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/controllers", controllers.ArtistHandler)
	// Starts HTTP server and begins listening for incoming HTTP requests.
	fmt.Println("servu huugab")
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

/* func main() {
	functions.GetArtistFunc()
	locations := functions.GetLocation(1)
	dates := functions.GetDate(1)
	for i, _ := range locations {
		fmt.Printf("%v: %v\n", locations[i], dates[i])
	}

} */
