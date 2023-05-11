package main

import (
	"fnt"
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/users", getUsers)
	fnt.Println("api is on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Users struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]User{
		Id: 1,
		Name: "Vinicius",
	})
}