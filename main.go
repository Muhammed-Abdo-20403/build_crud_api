package main 

import (
	"fmt"
	"log"
	"math/rand"
	"github.com/gorilla/mux"
	"strconv"
	"net/http"
	"encoding/json"
)

type Movie struct {
	ID string 	 				`json:"id"`
	Isbn string	 				`json:"isbn`
	Title string 				`json:"title"`
	Director *Director  `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string  `json:"lastname"`
}

var movies []Movie


func getmovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deletemovie(w http.ResponseWriter, r http.Request)  {
	w.NewEncoder().Set("Content-Type", "application/json")
	params := mux.Vars(w)
	for index, item := range mivies {
		for item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]... )
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}


func getmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return 
		}
	} 	
}

func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Movie Movie
	_ := json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updatemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies{
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var Movie Movie
			_ := json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main()  {
	route := mux.NewRouter()

	movies = append(movies, Movie{ID:"1", Isbn:"451288", Title:"Movie one", Director: &Director(Firstname:"Mo", Lastname:"Abdo")})
	movies = append(movies, Movie{ID:"2", Isbn:"986532", Title:"Movie Two", Director: &Director(Firstname:"Husaam", Lastname: "Ahmed")})
	
	route.Handlefuc("/movies", getmovies).Methods("GET") 
	route.Handlefuc("/movies/{id}", getmovie).Methods("GET")
	route.Handlefuc("/movies", createmovie).Methods("POST")
	route.Handlefuc("/movies/{id}", updatemovie).Methods("PUT")
	route.Handlefuc("/movies/{id}", deletemovie).Methods("DELETE")

	fmt.Prinf("Starting server at port 8000\n")
	log.Fatal(http.ListAndServe(":8000", route))
}