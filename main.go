package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"time"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetMainPage show templates/index.html page
func GetMainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	//fmt.Println(posts)

	t.ExecuteTemplate(w, "index", nil)
}

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
	router.HandleFunc("/", GetMainPage).Methods("GET")
	router.HandleFunc("/random-name", GetTeamName).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
