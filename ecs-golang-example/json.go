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
	Company	string `json:"company"`
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
		fmt.Fprintf(w, "%s %s is %d years old, who's from %s", user.Firstname, user.Lastname, user.Age, user.Company)
	})


	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "John",
			Lastname:  "Doe",
			Company:  "Amazon Web Services",
			Age:       25,
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
