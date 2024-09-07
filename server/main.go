package main

import (
	"fmt"
)

func main() {
	s := NewServer("localhost:8000")
	fmt.Println("Starting the server")
	s.Start()
}
