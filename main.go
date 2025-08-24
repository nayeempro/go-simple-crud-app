package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	//"math/rand"
	//"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movie []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movie {
		if item.ID == params["id"] {
			movie = append(movie[:index], movie[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movie)
}

// Get a single movie by id
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range movie {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// create movie
func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	// if you need random id
	// newMovie.ID = strconv.Itoa(rand.Intn(10000000))
	newMovie.ID = strconv.Itoa(len(movie) + 1)
	movie = append(movie, newMovie)
	json.NewEncoder(w).Encode(newMovie)
}

// update movie
func updateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range movie {
		if item.ID == params["id"] {
			movie = append(movie[:index], movie[index+1:]...)
			var updateMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&updateMovie)
			updateMovie.ID = params["id"]
			movie = append(movie, updateMovie)
			json.NewEncoder(w).Encode(updateMovie)
			return
		}
	}
}
func main() {
	fmt.Println("hello")

	route := mux.NewRouter()

	route.HandleFunc("/movies", getMovies).Methods("GET")
	route.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	route.HandleFunc("/movies", createMovies).Methods("POST")
	route.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	route.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Println("Server start on the prot 4000")
	log.Fatal(http.ListenAndServe(":4000", route))
}

func init() {
	mv1 := Movie{
		ID:       "1",
		Isbn:     "1234",
		Title:    "First Movie",
		Director: &Director{FirstName: "Jon", LastName: "Doe"},
	}
	mv2 := Movie{
		ID:       "2",
		Isbn:     "1234",
		Title:    "Second Movie",
		Director: &Director{FirstName: "Jon", LastName: "Doe"},
	}
	mv3 := Movie{
		ID:       "3",
		Isbn:     "1234",
		Title:    "Third Movie",
		Director: &Director{FirstName: "Jon", LastName: "Doe"},
	}

	movie = append(movie, mv1)
	movie = append(movie, mv2)
	movie = append(movie, mv3)
}
