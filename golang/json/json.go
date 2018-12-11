package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Nationality string `json:"nationality"`
	Company	string `json:"company"`
	Department string `json:"department"`
	Age       int    `json:"age"`
	Selfie	string `json:"selfie"`
	Fields string `json:"fields"`
	Twitter string `json:"twitter"`
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
		fmt.Fprintf(w, "%s %s is %d years old, who's from %s", user.Firstname, user.Lastname, user.Age, user.Company)
	})


	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "Werner",
			Lastname:  "Vogels",
			Nationality: "Dutch",
			Company:  "Amazon Web Services",
			Department: "Administration",
			Age:       60,
			Selfie: "https://en.wikipedia.org/wiki/Werner_Vogels#/media/File:WernerVogels.JPG",
			Fields: "Distributed computing",
			Twitter: "@Werner",
		}
		fmt.Println("Reponse -> ", peter)
		json.NewEncoder(w).Encode(peter)

	})

	ver := flag.String("ver","-1.0.0", "application version.")
	flag.Parse()

	fmt.Println("Server Version: ", *ver)
	fmt.Println("Sever is ready at port [8080] ...")
	http.ListenAndServe(":8080", nil)
}
