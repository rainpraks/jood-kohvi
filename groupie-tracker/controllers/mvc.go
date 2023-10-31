package controllers

import (
	"groupie-trackers/functions"
	"html/template"
	"net/http"
)

var tpl *template.Template

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	artists := functions.GetArtistFunc()

	if r.Method == http.MethodPost {

		tpl.ExecuteTemplate(w, "mainpage.html", artists)
	}
}

/* func LocationHandler() {

} */

/* func DateHandler() {

}
*/
