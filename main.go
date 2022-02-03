package main

import (
	"log"
	"net/http"
	"encoding/json"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token string `json:"token"`
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	dt := time.Now()

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Username != "c137@onecause.com" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "invalid username"}`))
		return
	} else if user.Password != "#th@nH@rm#y#r!$100%D0p#" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "invalid password"}`))
		return
	} else if user.Token != dt.Format("15:04") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "invalid token"}`))
		return
	} else {
		w.Write([]byte(`{"message": "login successful"}`))
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", login).Methods("POST")

	c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:8080"},
        AllowCredentials: true,
    })
	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8081", handler))
}