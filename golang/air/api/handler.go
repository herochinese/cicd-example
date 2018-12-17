package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"../feed"
)


func Cities(w http.ResponseWriter, r *http.Request) {
	values := []feed.AirQuality{}
	for _, value := range LocalCache {
		values = append(values, value)
	}
	json.NewEncoder(w).Encode(values)
}

func City(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]
	if len(LocalCache[city].IndexCityVHash)>0 {
		json.NewEncoder(w).Encode(LocalCache[city])
	} else {
		fmt.Fprintf(w, `{"status":"error", "description":"%s wasn't existed in the cache.'"}`, city)
	}

}

func Feed(w http.ResponseWriter, r *http.Request) {
	var air feed.AirQuality
	err := json.NewDecoder(r.Body).Decode(&air)
	if err!=nil {
		log.Println(err)
		fmt.Fprintf(w, `{"status":"error", "description":"%s"}`, err)
		return
	}

	LocalCache[air.City]=air
	fmt.Fprintf(w, `{"status":"success", "description":"Air Quality data of %s has been cached."}`, air.City)

}