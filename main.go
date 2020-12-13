package main

import (
	"Golang-challenge/server"
	"fmt"
	"log"
	"net/http"
)

func main()  {
	s := server.New()
	fmt.Printf("Server started on 8080... \n")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))

}