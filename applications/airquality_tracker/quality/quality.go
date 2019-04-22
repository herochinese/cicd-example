package main

import (
	"log"
	"net/http"
	"quality/api"
)

func main() {

	router := api.NewRouter()
	bind := ":8080"
	log.Printf("Quality Server is ready at [ %s ]\n", bind)

	log.Fatal(http.ListenAndServe(bind, router))
}


/*
 ====================
 REQ
 ====================

{
	"index_city_v_hash": "175bf71c6370151071d139083b1a95ac3af7148d",
	"index_city": "Beijing_1451",
	"idx": 1451,
	"aqi": 137,
	"city": "Beijing",
	"city_cn": "北京",
	"lat": "39.9546",
	"lng": "116.468",
	"co": "9.1",
	"h": "33.5",
	"no2": "38.4",
	"o3": "1.3",
	"p": "1012.1",
	"pm10": "68",
	"pm25": "137",
	"so2": "6.6",
	"t": "4.7",
	"w": "0",
	"s": "2018-12-17 21:00:00",
	"tz": "+08:00",
	"v": 1545080400
}

*/