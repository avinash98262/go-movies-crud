package main

import (
	"fmt"
	"net/http"
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

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to the Golang world")

	details := Movie{
		Title: "tiger",
		Id:    "123",
	}
	fmt.Fprintln(w, details.Title)
	fmt.Fprintln(w, details.Id)
}

func main() {

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
