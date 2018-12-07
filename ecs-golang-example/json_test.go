package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}


func TestEncode(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/encode" {
			EncodeU(rw, req)
		}

	}))

	defer server.Close()


	resp, err := http.Get("/encode")
	if err == nil {
		fmt.Println(resp.Body)
		var user User
		json.NewDecoder(resp.Body).Decode(&user)
		assert.Equal(t,"Werner", user.Firstname)

	}
}


func TestDecode(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/decode" {
			EncodeU(rw, req)
		}

	}))

	defer server.Close()

	p := User {
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

	b, e := json.Marshal(p)
	if e!=nil {

	}

	resp, err := http.Post("/decode","application/json", bytes.NewBuffer(b))
	if err == nil {
		fmt.Println(resp.Body)
		var user User
		json.NewDecoder(resp.Body).Decode(&user)
		assert.Equal(t,"Werner", user.Firstname)

	}
}

//func TestEncodePerf(t *testing.T) {
//
//	for i:=0; i<1000; i++ {
//		resp, err := http.Get("http://jsd-d-appsv-1m6m5ddl5uo85-2133984274.us-east-1.elb.amazonaws.com:8080/encode")
//		if err != nil {
//			fmt.Println(err)
//		}
//		defer resp.Body.Close()
//		body, err := ioutil.ReadAll(resp.Body)
//		fmt.Println(body)
//	}
//}