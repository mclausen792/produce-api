package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	. "just-ripe/config"
	. "just-ripe/dao"
	. "just-ripe/models"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

var config = Config{}
var dao = FruitsDAO{}

func AllFruitsEndPoint(w http.ResponseWriter, r *http.Request) {
	fruits, err := dao.FindAllFruit()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, fruits)
}

func AllVegetablesEndPoint(w http.ResponseWriter, r *http.Request) {
	fruits, err := dao.FindAllVegetables()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, fruits)
}

func FindFruitEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fruit, err := dao.FindFruitById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Id")
		return
	}
	respondWithJson(w, http.StatusOK, fruit)
}

func FindVegetableEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fruit, err := dao.FindVegetableById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Id")
		return
	}
	respondWithJson(w, http.StatusOK, fruit)
}

func UpdateFruitEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var fruit Fruit
	if err := json.NewDecoder(r.Body).Decode(&fruit); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdateFruit(fruit); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func UpdateVegetableEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var veggies Vegetable
	if err := json.NewDecoder(r.Body).Decode(&veggies); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdateVegetable(veggies); err != nil {
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
	dao.DialInfo = &mgo.DialInfo{
		Addrs:    []string{config.Server},
		Database: config.Database,
		Username: config.Username,
		Password: config.Password,
	}

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()

}

// Define HTTP request routes
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	r := mux.NewRouter()
	r.HandleFunc("/fruit", AllFruitsEndPoint).Methods("GET")
	r.HandleFunc("/vegetable", AllVegetablesEndPoint).Methods("GET")
	r.HandleFunc("/fruit", UpdateFruitEndPoint).Methods("PUT")
	r.HandleFunc("/vegetable", UpdateVegetableEndPoint).Methods("PUT")
	r.HandleFunc("/fruit/{id}", FindFruitEndpoint).Methods("GET")
	r.HandleFunc("/vegetable/{id}", FindVegetableEndpoint).Methods("GET")
	fmt.Println("running on Port" + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
