package main

import (
	"crawler/df"
	"crawler/feed"
	"crawler/kns"
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

var region = os.Getenv("MY_AWS_REGION")
var stream = os.Getenv("MY_AWS_KINESIS_NAME")

func main() {

	if region != "" && stream != "" {
		cities := df.LoadCities("china-cities.json")
		var wg sync.WaitGroup
		wg.Add(1)

		go schedule(cities, 30*1000*time.Millisecond)

		wg.Wait()
	} else {
		log.Fatal("Check out environments: MY_AWS_REGION & MY_AWS_KINESIS_NAME")
	}


}

func pullAirData(city string) feed.AirQuality {
	var ori feed.OriginAirQuality
	var apiError feed.ApiError
	var air feed.AirQuality

	cf := feed.CityFeed(city)
	if cf != nil {

		err := json.Unmarshal(cf, &ori)
		if err != nil {
			log.Println(err)
		}
		if ori.Status == "error" {
			err2 := json.Unmarshal(cf, &apiError)
			if err2 != nil {
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

	for {

		for _, c := range cities {

			log.Println("City -> ", c)
			data := pullAirData(c)
			b, err := json.Marshal(data)
			if err != nil {
				log.Println(err)
				continue
			}
			kns.Push2Kinesis(region, stream, b)

		}

		log.Println("...sleep...", delay)
		time.Sleep(delay)

	}

}
