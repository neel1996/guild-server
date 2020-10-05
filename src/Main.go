package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neel1996/guild-server/src/api"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getsocialdata", api.SocialAPI).Methods("GET")
	http.Handle("/", router)
	err := http.ListenAndServe(":3000", router)

	if err != nil {
		log.Fatal(err)
	}
}
