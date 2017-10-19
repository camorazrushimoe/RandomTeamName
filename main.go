package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/shogo82148/go-shuffle"
)

func main() {

	now := time.Now()
	secs := now.Unix()

	rand.Seed(secs)
	firstWord := []string{"Веселые", "Грустные", "Несчастные", "Бодрые", "Смешные", "Известные", "Плавные"}
	secondWord := []string{"Непоседы", "Бунтари", "Агностики", "Умники", "Клерики"}

	shuffle.Strings(firstWord)
	shuffle.Strings(secondWord)

	fmt.Println("Team name:", firstWord[rand.Intn(len(firstWord))], secondWord[rand.Intn(len(secondWord))])

}
