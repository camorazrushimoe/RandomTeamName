package main

import (
	"fmt"
	"math/rand"
	"time"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetTeamName function bla
func GetTeamName(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	secs := now.Unix()

	rand.Seed(secs)
	firstWord := []string{"Веселые", "Грустные", "Несчастные", "Бодрые", "Смешные", "Известные", "Плавные"}
	secondWord := []string{"Непоседы", "Бунтари", "Агностики", "Умники", "Клерики"}

	teamName := firstWord[rand.Intn(len(firstWord))] + " " + secondWord[rand.Intn(len(secondWord))]

	fmt.Fprintf(w, teamName)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetTeamName).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
