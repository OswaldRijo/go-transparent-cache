package main

import (
	"fmt"
	"log"
	"net/http"
	"Golang-challenge/server"
)

func main()  {

	fmt.Printf("Starting server on 8080... \n")
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))

}