package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func deleteMovie(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(movies)
}

func createMovie(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = (uuid.New()).String()
	movies = append(movies, movie)
	json.NewEncoder(res).Encode(movie)

}

func getMovie(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	fmt.Printf(params["id"])
	for _,item := range movies{
		if item.ID == params["id"]{
			
			json.NewEncoder(res).Encode(item)
			return
		}
	}
}

func getMovies(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)

}

func updateMovie(res http.ResponseWriter,req *http.Request){
	// 1) Set headers
	res.Header().Set("Content-Type", "application/json")
	// 2) get params
	params := mux.Vars(req)
	// 3) Loop through movies
	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			var movie Movie 
			_ = json.NewDecoder(req.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies,movie)
			json.NewEncoder(res).Encode(movie)
			return
		}
	}
}


func main() {
	// Create a router from mux router
	router := mux.NewRouter()
	// Add some dummy data 
	movies = append(movies, Movie{ID:"1",Title: "Harry Potter",Isbn: "163513",Director:&Director{Firstname: "Mohamed",Lastname: "Sawah"}})
	movies = append(movies, Movie{ID:"2",Title: "The lord of the rings",Isbn: "5643",Director:&Director{Firstname: "Mohamed",Lastname: "Sawah"}})
	// Routers handling	
	router.HandleFunc("/movies",getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	router.HandleFunc("/movies",createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server on port 8080")
	// Starting the server on port 8080
	log.Fatal(http.ListenAndServe(":8080",router))

}