package main

import (
	"air/data"
	"air/feed"
	"air/util"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)


type save func(air feed.AirQuality)

type Province  struct {
	Name   string `json:"name_en"`
	NameCN string `json:"name"`
	City   []City   `json:"city"`
}
type City struct {
	Name string `json:"name"`
	County []County `json:"county"`
}

type County struct {
	Name   string `json:"name_en"`
	Code   string `json:"code"`
	NameCN string `json:"name"`
}

/**
1. Setup environment - AIRDEX_API_URL as the base url of API server.
2. Running API server before AIRDEX
 */
func main() {
	city := flag.String("city","beijing", "name of the city")
	flag.Parse()

	if flag.NFlag() < 1 {
		log.Printf("City wasn't setting and loading city list from <%s>", "ChinaCityList.json")
		cs := loadCityList("ChinaCityList.json")
		//cs := []string{"beijing"}

		var wg sync.WaitGroup
		wg.Add(len(cs))
		for x := range cs {
			log.Printf("Go routine -> %s", cs[x])
			go schedule([]save {data.TransactSave2DynamoPerf, data.Save2DynamoPerf}, cs[x], 30*1000*time.Millisecond)
		}
		wg.Wait()

	} else {
		b,err := json.Marshal(getAir(*city))
		if err!=nil {
			log.Println(err)
			return
		}
		util.PrintJson("Air Quality of " + strings.ToUpper(*city) + ": ", b)
	}


}

func loadCityList(file string) []string {
	var pr []Province
	body, err := ioutil.ReadFile(file)
	if err!=nil {
		log.Fatal(err)
	}

	err2 := json.Unmarshal(body, &pr)
	if err2!=nil {
		log.Println(err)
	}
	cs := make([]string, len(pr))

	for i, p := range pr {
		cs[i] = p.City[0].County[0].Name
	}
	return cs
}




func getAir(city string) feed.AirQuality {
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

func schedule(what []save, city string, delay time.Duration) {
	tick := time.Tick(delay)
	airdexapi :=os.Getenv("AIRDEX_API_URL")
	for range tick {
		air := getAir(city)
		if &air.IndexCityVHash!=nil && len(air.IndexCityVHash)>0 {
			for _,w := range what {
				b, err :=json.Marshal(air)
				if err!=nil {
					log.Println(err)
				} else {

					feed.ApiPost(airdexapi, "application/json", b)
				}
				w(air)
			}
		}

	}

}
