package main

import (
	"api_resources/src/repositories"
	"api_resources/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Started API\nPort:8080")
	r := router.Generate()
	repositories.New()
	log.Fatal(http.ListenAndServe(":8080", r))
}
