package main

import (
	"fmt"
	"log"
)

func main() {
	p, err := NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")

	if err := p.PrepareDb(); err != nil {
		log.Fatal(err)
	}

	s := NewServer("localhost:8000", p)
	fmt.Println("Starting the server")
	s.Start()
}
