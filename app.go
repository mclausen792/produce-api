package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/mclausen792/produce-api/config"
	. "github.com/mclausen792/produce-api/dao"
	. "github.com/mclausen792/produce-api/models"
)

var config = Config{}
var dao = FruitsDAO{}

func AllFruitsEndPoint(w http.ResponseWriter, r *http.Request) {
	fruits, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

func FindFruitEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fruit, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Id")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}

func UpdateFruitEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var fruit Fruit
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fruit", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/fruit", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/fruit/{id}", FindMovieEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3001", r); err != nil {
		log.Fatal(err)
	}
}
