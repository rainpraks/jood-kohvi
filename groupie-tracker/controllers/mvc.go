package controllers

import (
	"fmt"
	"groupie-trackers/functions"
	"net/http"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	functions.GetArtistFunc()

	fmt.Fprintln(w)

}

/* func LocationHandler() {

} */

/* func DateHandler() {

}
*/