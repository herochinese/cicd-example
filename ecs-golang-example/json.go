package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}


/*
{
	"firstname": "John",
	"lastname":  "Doe",
	"age":       25
}
 */
func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Println("Reponse -> ", user)
		fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
	})


	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "John",
			Lastname:  "Doe",
			Age:       25,
		}
		fmt.Println("Reponse -> ", peter)
		json.NewEncoder(w).Encode(peter)

	})

	fmt.Println("Sever is ready at port-8080 ...")
	http.ListenAndServe(":8080", nil)
}
