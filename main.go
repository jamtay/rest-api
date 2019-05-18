package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest-api/resources"
)

func main() {
	router := mux.NewRouter()
	setUpPeopleApi(router)
	setUpMoviesApi(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func setUpMoviesApi(router *mux.Router) {
	router.HandleFunc("/movies", resources.AllMoviesEndPoint).Methods("GET")
	router.HandleFunc("/movies", resources.CreateMovieEndPoint).Methods("POST")
	router.HandleFunc("/movies", resources.UpdateMovieEndPoint).Methods("PUT")
	router.HandleFunc("/movies", resources.DeleteMovieEndPoint).Methods("DELETE")
	router.HandleFunc("/movies/{id}", resources.FindMovieEndPoint).Methods("GET")
}

func setUpPeopleApi(router *mux.Router) {
	router.HandleFunc("/people", resources.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", resources.GetPerson).Methods("GET")
	router.HandleFunc("/people", resources.CreatePerson).Methods("POST")
	router.HandleFunc("/people", resources.DeletePerson).Methods("DELETE")
}