package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// movie.ID = strconv.Itoa(rand.Intn(10000000))
	movie.ID = strconv.Itoa(len(movies) + 1)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var m *Movie
	var nm Movie

	json.NewDecoder(r.Body).Decode(&nm)

	for _, item := range movies {
		if item.ID == params["id"] {
			m = &item
		}
	}
	*m = Movie{
		ID:       m.ID,
		Title:    nm.Title,
		ISBN:     nm.ISBN,
		Director: nm.Director,
	}

	json.NewEncoder(w).Encode(*m)
}

var movies []Movie

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{
		ID:    "1",
		ISBN:  "438227",
		Title: "Movie One",
		Director: &Director{
			FirstName: "Anaet Hossain",
			LastName:  "Rezve",
		},
	}, Movie{
		ID:    "2",
		ISBN:  "432447",
		Title: "Movie Two",
		Director: &Director{
			FirstName: "John",
			LastName:  "Doe",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server At Port: 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
