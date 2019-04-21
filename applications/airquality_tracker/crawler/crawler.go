package main

import (
	"crawler/df"
	"crawler/feed"
	"encoding/json"
	"log"
	"sync"
	"time"
)


func main() {

	cities := df.LoadCities("china-cities.json")
	var wg sync.WaitGroup
	wg.Add(1)

		go schedule(cities, 30*1000*time.Millisecond)

	wg.Wait()

	//for _, c := range cities {
	//
	//	log.Println("City -> ", c)
	//	pullAirData(c)
	//}

}


func pullAirData(city string) feed.AirQuality {
	var ori feed.OriginAirQuality
	var apiError feed.ApiError
	var air feed.AirQuality

	cf := feed.CityFeed(city)
	if cf!=nil {

		err := json.Unmarshal(cf, &ori)
		if err!=nil {
			log.Println(err)
		}
		if ori.Status=="error" {
			err2 := json.Unmarshal(cf, &apiError)
			if err2!=nil {
				log.Println(err2)
			}
			log.Printf("Retrieve data of %s from https://api.waqi.info/ was failed due to <%s>. ", city, apiError.Data)
			return air

		}
		air = feed.Copy2AirQuality(ori)

	}
	return air

}

func schedule(cities []string, delay time.Duration) {

	for _, c := range cities {

		log.Println("City -> ", c)
		pullAirData(c)
	}

	tick := time.Tick(delay)
	for range tick {
		log.Println(".")
	}

}