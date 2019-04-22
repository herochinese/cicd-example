package feed

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var token = "b0e78ca32d058a9170b6907c5214c0e946534cc9"
var host = "https://api.waqi.info"

type HttpAccessLayer interface {
	Get(url string) ([]byte, error)
	Post(url string, contentType string, body []byte) ([]byte, error)
}

type HttpHAL struct {
	transport *http.Transport
	client    *http.Client
}

func NewHttpHAL() (*HttpHAL, error) {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: transport}

	hal := &HttpHAL{
		transport: transport,
		client:    client,
	}
	return hal, nil
}

func (h *HttpHAL) Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("http.NewRequest was failed due to [%s].\n", err)
		return nil, err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		log.Printf("http.Call was failed due to [%s].\n", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll was failed due to [%s].\n", err)
		return nil, err
	}
	return body, nil
}

func (h *HttpHAL) Post(url string, contentType string, b []byte) ([]byte, error) {
	resp, err1 := http.Post(url, contentType, bytes.NewReader(b))
	if err1 != nil {
		log.Printf("API call was failed from %s with Err: %s. \n", url, err1)
		return nil, err1
	}
	defer resp.Body.Close()

	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Printf("Read buffer failed.\n")
	}
	toJson("Response -> ", b)
	return b, nil
}

func Search(keyword string) []byte {
	//https://api.waqi.info/search/?token=demo&keyword=bangalore
	url := host + "/feed/here/?token=" + token + "&keyword=" + keyword
	return ApiGet(url)
}
func IPFeed() []byte {
	//https://api.waqi.info/feed/here/?token=demo
	url := host + "/feed/here/?token=" + token
	return ApiGet(url)

}

func GeoFeed(lat string, lng string) []byte {
	//https://api.waqi.info/feed/geo:10.3;20.7/?token=demo
	url := host + "/feed/geo:" + lat + ";" + lng + "/?token=" + token
	return ApiGet(url)
}

func CityFeed(city string) []byte {
	//https://api.waqi.info/fee/beijing/?token=demo
	url := host + "/feed/" + city + "/?token=" + token
	return ApiGet(url)
}

func ApiGet(url string) []byte {
	h, err := NewHttpHAL()
	b, err := h.Get(url)
	if err != nil {
		log.Printf("HttpAccessLayer.Get was failed due to [%s].\n", err)
		return nil
	}

	toJson("Resp -> ", b)
	return b
}

func toJson(title string, body []byte) {
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}
	log.Println(title, prettyJSON.String())

}
