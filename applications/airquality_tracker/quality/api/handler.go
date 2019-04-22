package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Cities(w http.ResponseWriter, r *http.Request) {
	values := []AirQuality{}
	for _, value := range LocalCache {
		values = append(values, value)
	}
	json.NewEncoder(w).Encode(values)
}

func City(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]
	if len(LocalCache[city].IndexCityVHash) > 0 {
		json.NewEncoder(w).Encode(LocalCache[city])
	} else {
		fmt.Fprintf(w, `{"status":"error", "description":"%s wasn't existed in the cache.'"}`, city)
	}

}

var WorkQueue = make(chan AirQuality, 200)

func Feed(w http.ResponseWriter, r *http.Request) {
	var air AirQuality
	err := json.NewDecoder(r.Body).Decode(&air)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, `{"status":"error", "description":"%s"}`, err)
		return
	}

	//LocalCache[air.City]=air
	WorkQueue <- air
	fmt.Fprintf(w, `{"status":"success", "description":"Air Quality data of %s has been cached."}`, air.City)

}

func ProcessMessage() {
	go func() {
		for {
			select {
			case work := <-WorkQueue:
				LocalCache[work.City] = work
				//case <-w.QuitChan:
				//	// We have been asked to stop.
				//	fmt.Printf("worker%d stopping\n", w.ID)
				//	return
			}
		}
	}()
}
