package main

import (
	"gopkg.in/mgo.v2"
	"log"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer session.Close()

	log.Println("Connected to MongoDB!")
}
