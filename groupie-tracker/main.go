package main

import (
	"fmt"
	"groupie-trackers/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/artists", controllers.ArtistHandler)
	//http.HandleFunc("/locations/", controllers.LocationHandler)
	//http.HandleFunc("/dates/", controllers.DateHandler)

	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(port, nil)
}

/* func main() {
	functions.GetArtistFunc()
	locations := functions.GetLocation(1)
	dates := functions.GetDate(1)
	for i, _ := range locations {
		fmt.Printf("%v: %v\n", locations[i], dates[i])
	}

} */
