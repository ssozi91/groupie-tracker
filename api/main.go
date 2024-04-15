package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Artists struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Members      []string  `json:"members"`
	Image        string    `json:"image"`
	CreationDate int       `json:"creationDate"`
	FirstAlbum   string    `json:"firstAlbum"`
	Relation     Relations `json:"relations"`
	NextPage     int       `json:"-"`
	PreviousPage int       `json:"-"`
}

type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	// Get the page number from the query parameters
	pageParam := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	// Define the number of artists per page
	artistsPerPage := 9

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var artistsObject []Artists
	json.Unmarshal(responseData, &artistsObject)

	// Get the search query from the query parameters
	searchQuery := r.URL.Query().Get("search")

	// Filter the artists based on the search query
	if searchQuery != "" {
		filteredArtists := []Artists{}
		for _, artist := range artistsObject {
			if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(searchQuery)) {
				filteredArtists = append(filteredArtists, artist)
			}
		}
		if len(filteredArtists) > 0 {
			artistsObject = filteredArtists
		} else {
			tmpl, err := template.ParseFiles("../static/no_results.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	// Calculate the total number of pages
	totalPages := (len(artistsObject) + artistsPerPage - 1) / artistsPerPage

	// Create a slice for the page numbers
	pages := make([]int, totalPages)
	for i := range pages {
		pages[i] = i + 1
	}

	start := (page - 1) * artistsPerPage
	end := start + artistsPerPage
	if end > len(artistsObject) {
		end = len(artistsObject)
	}

	tmpl, err := template.ParseFiles("../static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, struct {
		Artists     []Artists
		Pages       []int
		CurrentPage int
	}{
		Artists:     artistsObject[start:end],
		Pages:       pages,
		CurrentPage: page,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	// Get the artist's ID from the URL
	id := strings.TrimPrefix(r.URL.Path, "/artists/")

	artistID, err := strconv.Atoi(id)
	if err != nil || artistID < 1 || artistID > 52 {
		http.Error(w, "NOT FOUND", http.StatusBadRequest)
		return
	}

	baseURL := "https://groupietrackers.herokuapp.com/api/artists/"
	url := baseURL + id

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var artistsObject Artists
	json.Unmarshal(responseData, &artistsObject)

	relationURL := "https://groupietrackers.herokuapp.com/api/relation/"
	url_relation := relationURL + id

	response_rel, err := http.Get(url_relation)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer response_rel.Body.Close()

	responseData_rel, err := ioutil.ReadAll(response_rel.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Check the status code of the response
	if response_rel.StatusCode == http.StatusNotFound {
		fmt.Println("Received 404 response code")
		http.Error(w, "not found", http.StatusNotFound)
		return
	} else if response_rel.StatusCode != http.StatusOK {
		fmt.Println("Received non-200 response code")
		http.Error(w, "Failed to get relation data", http.StatusInternalServerError)
		return
	}

	var artistsObject_rel Relations
	json.Unmarshal(responseData_rel, &artistsObject_rel)
	artistsObject.Relation = artistsObject_rel

	tmpl, err := template.ParseFiles("../static/artists.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artistsObject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/artists/", artistHandler)
	http.ListenAndServe(":8080", nil)
}
