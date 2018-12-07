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
firstname: "Werner",
lastname: "Vogels",
nationality: "Dutch",
company: "Amazon Web Services",
department: "Administration",
age: 60,
selfie: "https://en.wikipedia.org/wiki/Werner_Vogels#/media/File:WernerVogels.JPG",
fields: "Distributed computing",
twitter: "@Werner"
}
 */

func Sum(x int, y int) int {
	return x + y
}




func DecodeR(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("Reponse -> ", user)
	fmt.Fprintf(w, "%s %s is %d years old, who's from %s", user.Firstname, user.Lastname, user.Age, user.Company)
}

func EncodeU(w http.ResponseWriter, r *http.Request) {
	peter := User{
		Firstname: "Chuan",
		Lastname:  "CHen",
		Nationality: "Dutch",
		Company:  "Amazon Web Services",
		Department: "Administration",
		Age:       30,
		Selfie: "https://en.wikipedia.org/wiki/Werner_Vogels#/media/File:WernerVogels.JPG",
		Fields: "Distributed computing",
		Twitter: "@Werner",
	}
	fmt.Println("Reponse -> ", peter)
	json.NewEncoder(w).Encode(peter)
}


func main() {

	http.HandleFunc("/decode", DecodeR)
	http.HandleFunc("/encode", EncodeU)

	ver := flag.String("ver","-1.0.0", "application version.")
	flag.Parse()

	fmt.Println("Server Version: ", *ver)
	fmt.Println("Sever is ready at port [8080] ...")
	http.ListenAndServe(":8080", nil)
}
