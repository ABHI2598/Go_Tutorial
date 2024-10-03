package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ABHI2598/mongoapi/router"
)

func main() {
	fmt.Println("Hello Mongo db api")

	r := router.Router()

	fmt.Println("Running at localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
