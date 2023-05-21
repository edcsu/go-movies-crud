package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	firstname string `json:"firstname"`
	lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "12908", Title: "Covert Ops", Director: &Director{firstname: "John", lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "12909", Title: "Dynasty In Namibia", Director: &Director{firstname: "Jane", lastname: "Patricia"}})
	
	fmt.Printf("Starting at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
