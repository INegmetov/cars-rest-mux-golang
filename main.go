package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	Id     int
	Make   string
	Moderl string
	Price  int
}

var vehicles = []Vehicle{
	{1, "Toyta", "Corolla", 10000},
	{2, "Toyta", "Camry", 15000},
	{3, "Toyta", "LandCruser", 25000},
	{4, "Lexus", "LS300", 30000},
}

func returnAllCars(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}
func returnCarByBrand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carMake := vars["make"]
	cars := &[]Vehicle{}
	for _, car := range vehicles {
		if car.Make == carMake {
			*cars = append(*cars, car)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}

func returnCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert id")
	}
	for _, car := range vehicles {
		if car.Id == carId {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(car)
		}
	}

}

func updateCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert id")
	}
	var updatedCar Vehicle
	json.NewDecoder(r.Body).Decode(&updatedCar)
	for k, v := range vehicles {
		if v.Id == carId {
			vehicles = append(vehicles[:k], vehicles[k+1:]...)
			vehicles = append(vehicles, updatedCar)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

func createCar(w http.ResponseWriter, r *http.Request) {
	var newCar Vehicle
	json.NewDecoder(r.Body).Decode(&newCar)
	vehicles = append(vehicles, newCar)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

func removeCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert id")
	}
	for k, v := range vehicles {
		if v.Id == carId {
			vehicles = append(vehicles[:k], vehicles[k+1:]...)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/cars", returnAllCars).Methods("GET")
	router.HandleFunc("/cars/make/{make}", returnCarByBrand).Methods("GET")
	router.HandleFunc("/cars/{id}", returnCarById).Methods("GET")
	router.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	router.HandleFunc("/cars", createCar).Methods("POST")
	router.HandleFunc("/cars/{id}", removeCarById).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
