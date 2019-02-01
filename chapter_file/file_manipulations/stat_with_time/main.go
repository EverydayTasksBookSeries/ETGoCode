package main

import (
	"log"

	times "gopkg.in/djherbis/times.v1"
)

func main() {
	t, err := times.Stat("D:/sonnet18.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(t.AccessTime())
	log.Println(t.ModTime())

	if t.HasChangeTime() {
		log.Println(t.ChangeTime())
	}

	if t.HasBirthTime() {
		log.Println(t.BirthTime())
	}
}
