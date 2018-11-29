package main

import (
	"log"
	"net/http"
	"ecs-golang-example/todo"
)

func main() {

	router := todo.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}


