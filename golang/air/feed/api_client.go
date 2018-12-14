package feed

import (
	"io/ioutil"
	"log"
	"net/http"
	"../util"
)

var token = "b0e78ca32d058a9170b6907c5214c0e946534cc9"
var host = "https://api.waqi.info"



func Search(keyword string)  []byte {
	//https://api.waqi.info/search/?token=demo&keyword=bangalore
	url := host + "/feed/here/?token=" + token + "&keyword="+keyword
	return apiGet(url)
}
func IPFeed() []byte {
	//https://api.waqi.info/feed/here/?token=demo
	url := host + "/feed/here/?token=" + token
	return apiGet(url)

}

func GeoFeed(lat string, lng string) []byte {
	//https://api.waqi.info/feed/geo:10.3;20.7/?token=demo
	url := host + "/feed/geo:" + lat + ";" + lng + "/?token=" + token
	return apiGet(url)
}

func CityFeed(city string) []byte {
	//https://api.waqi.info/fee/beijing/?token=demo
	url := host + "/feed/" + city + "/?token=" + token
	return apiGet(url)
}

func apiGet(url string) []byte {

	log.Println("Request -> ", url)
	resp, err := http.Get(url)
	if err !=nil {
		log.Printf("API call was failed from %s with Err: %s. \n", url, err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read buffer failed.\n")
	}
	util.PrintJson("Response -> ", body)
	return body
}


