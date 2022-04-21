package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rahul.chaube/CurdeDemo/config"
	"github.com/rahul.chaube/CurdeDemo/model"
	"github.com/rahul.chaube/CurdeDemo/repository"
	"gopkg.in/mgo.v2"
)

type Movie struct {
	Id       string   `json:"id"`
	ISBN     string   `json:"isbn"`
	Title    string   `json:"title"`
	Director Director `json:"director"`
}
type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movie []Movie
var repo *repository.MovieRepository

func main() {
	var session *mgo.Session
	config := config.Mongo{
		Address: "localhost:27017",
	}

	session, err := mgo.Dial(config.Address)
	if err != nil {
		panic(err)
	}
	repo = repository.NewMovieRepository(session, "Testing")
	r := mux.NewRouter()
	movie = append(movie, Movie{Id: "1", ISBN: "112554ASDXXX", Title: "KGF 2 ", Director: Director{FirstName: "Rahul", LastName: "Chaube"}})
	r.HandleFunc("/movies", getMovies).Methods(http.MethodGet)
	r.HandleFunc("/movie/{id}", getMovie).Methods(http.MethodGet)
	r.HandleFunc("/movie/{id}", deleteMovie).Methods(http.MethodDelete)
	r.HandleFunc("/movie", updateMovie).Methods(http.MethodPut)
	r.HandleFunc("/movie", addMovie).Methods(http.MethodPost)

	fmt.Println("Starting server at port:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for _, v := range movie {
		if v.Id == id {
			json.NewEncoder(w).Encode(v)
			break
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("No Movies Found ")

}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for index, v := range movie {
		if v.Id == id {
			movie = append(movie[index:], movie[index+1:]...)
			json.NewEncoder(w).Encode(v)
			break
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("No Movies Found ")

}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m Movie
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &m)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	for index, v := range movie {
		if v.Id == m.Id {
			movie = append(movie[index:], movie[index+1:]...)
			m.Id = strconv.Itoa(rand.Intn(1000000))
			movie = append(movie, m)
			json.NewEncoder(w).Encode(m)
			break
		}
	}
	// w.WriteHeader(http.StatusNotFound)
	// json.NewEncoder(w).Encode("No Movies Found ")

}
func addMovie(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var m Movie
	err = json.Unmarshal(body, &m)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	// err = json.NewDecoder(r.Body).Decode(&movie)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// }
	// fmt.Printf("Body is %+v \n", r)
	fmt.Printf("%+v Moviees is \n", m)
	m.Id = strconv.Itoa(rand.Intn(10000000))
	mov := model.Movie{
		Id:    m.Id,
		Title: m.Title,
		Director: model.Director{
			FirstName: m.Director.FirstName,
			LastName:  m.Director.LastName,
		},
	}
	err = repo.AddMovie(mov)
	if err != nil {
		fmt.Println("Error is occured ", err)
	}
	movie = append(movie, m)
	json.NewEncoder(w).Encode(m)
}
