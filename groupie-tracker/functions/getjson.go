package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var client *http.Client

// interface{}- can pass any value of any data type as an argument to this function
func GetJson(url string, target interface{}) error {
	fmt.Println(url)
	fmt.Println("found getjson")
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func GetArtistFunc() []Artist {
	client = &http.Client{Timeout: 10 * time.Second}
	fmt.Println("get artist")
	url := "https://groupietrackers.herokuapp.com/api/artists"
	var artists []Artist
	err := GetJson(url, &artists)
	if err != nil {
		fmt.Printf("error getting Artist: %s\n", err.Error())
	}
	return artists
}

type Location struct {
	Locations []string `json:"locations"`
}

func GetLocation(n int) []string {
	client = &http.Client{Timeout: 10 * time.Second}
	var locations Location
	url := "https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(n)

	err := GetJson(url, &locations)

	if err != nil {
		fmt.Printf("error getting location: %s\n", err.Error())
	} else {
		for _, l := range locations.Locations {
			fmt.Println(l)
		}
	}
	return locations.Locations

}

type Date struct {
	Dates []string `json:"dates"`
}

func GetDate(n int) []string {
	client = &http.Client{Timeout: 10 * time.Second}
	var dates Date
	url := "https://groupietrackers.herokuapp.com/api/dates/" + strconv.Itoa(n)
	err := GetJson(url, &dates)

	if err != nil {
		fmt.Printf("error getting dates: %s\n", err.Error())
	} else {
		for _, l := range dates.Dates {
			fmt.Println(l)
		}
	}
	return dates.Dates
}
