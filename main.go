package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Title    string    `json:"title"`
	Id       string    `json:"id"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "welcome to the Golang world")

// 	details := Movie{
// 		Title: "tiger",
// 		Id:    "123",
// 	}
// 	fmt.Fprintln(w, details.Title)
// 	fmt.Fprintln(w, details.Id)
// }

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {

		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {

		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)

}

func createMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	// ----->
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	//params
	// loop over the movies and range
	for index, item := range movies {

		// delete the movie that  id , whichever you have sent

		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			//append the new movie
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}

}

func main() {

	r := mux.NewRouter()
	movies = append(movies, Movie{Title: "Dhoom", Id: "234", Director: &Director{Firstname: "karan", Lastname: "johar"}})
	movies = append(movies, Movie{Title: "Jai ho", Id: "421", Director: &Director{Firstname: "prince", Lastname: "raj"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	// http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", r)
}
