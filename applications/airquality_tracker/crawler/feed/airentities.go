package feed

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"strconv"
	"strings"
)

type ApiError struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type AirQuality struct {
	IndexCityVHash string `json:"index_city_v_hash"`
	IndexCity      string `json:"index_city"`
	StationIndex   int    `json:"idx"`
	AQI            int    `json:"aqi"`
	City           string `json:"city"`
	CityCN         string `json:"city_cn"`
	Latitude       string `json:"lat"`
	Longitude      string `json:"lng"`
	Co             string `json:"co"`
	H              string `json:"h"`
	No2            string `json:"no2"`
	O3             string `json:"o3"`
	P              string `json:"p"`
	Pm10           string `json:"pm10"`
	Pm25           string `json:"pm25"`
	So2            string `json:"so2"`
	T              string `json:"t"`
	W              string `json:"w"`
	S              string `json:"s"`  //Local measurement time
	TZ             string `json:"tz"` //Station timezone
	V              int    `json:"v"`
}

type OriginAirQuality struct {
	Status string     `json:"status"`
	Data   OriginData `json:"data"`
}

type OriginData struct {
	AQI          int        `json:"aqi"`
	StationIndex int        `json:"idx"`
	City         OriginCity `json:"city"`
	IAQI         OriginIAQI `json:"iaqi"`
	OriginTime   OriginTime `json:"time"`
}

type OriginCity struct {
	Geo  []float64 `json:"geo"`
	Name string    `json:"name"`
}

type OriginIAQI struct {
	Co   OValue `json:"co"`
	H    OValue `json:"h"`
	No2  OValue `json:"no2"`
	O3   OValue `json:"o3"`
	P    OValue `json:"p"`
	Pm10 OValue `json:"pm10"`
	Pm25 OValue `json:"pm25"`
	So2  OValue `json:"so2"`
	T    OValue `json:"t"`
	W    OValue `json:"w"`
}

type OValue struct {
	V float64 `json:"v"`
}

type OriginTime struct {
	S  string `json:"s"`  //Local measurement time
	TZ string `json:"tz"` //Station timezone
	V  int    `json:"v"`
}

func SplitName(name string) (city string, citycn string) {
	ns := strings.Split(name, "(")
	if len(ns) != 2 {
		log.Println("Input name <", name, "> wasn't matched with convention. eg: ", "Beijing (北京)")
		return name, ""
	}
	city = strings.Trim(ns[0], " ")
	citycn = strings.Trim(ns[1], ")")
	return city, citycn
}

func Copy2AirQuality(src OriginAirQuality) AirQuality {

	var dest AirQuality
	dest.StationIndex = src.Data.StationIndex
	dest.AQI = src.Data.AQI
	c, cn := SplitName(src.Data.City.Name)
	dest.City = c
	dest.CityCN = cn
	dest.Latitude = strconv.FormatFloat(src.Data.City.Geo[0], 'g', 6, 64)
	dest.Longitude = strconv.FormatFloat(src.Data.City.Geo[1], 'g', 6, 64)

	dest.Co = strconv.FormatFloat(src.Data.IAQI.Co.V, 'g', 6, 64)
	dest.H = strconv.FormatFloat(src.Data.IAQI.H.V, 'g', 6, 64)
	dest.No2 = strconv.FormatFloat(src.Data.IAQI.No2.V, 'g', 6, 64)
	dest.O3 = strconv.FormatFloat(src.Data.IAQI.O3.V, 'g', 6, 64)
	dest.P = strconv.FormatFloat(src.Data.IAQI.P.V, 'g', 6, 64)
	dest.Pm10 = strconv.FormatFloat(src.Data.IAQI.Pm10.V, 'g', 6, 64)
	dest.Pm25 = strconv.FormatFloat(src.Data.IAQI.Pm25.V, 'g', 6, 64)
	dest.So2 = strconv.FormatFloat(src.Data.IAQI.So2.V, 'g', 6, 64)
	dest.T = strconv.FormatFloat(src.Data.IAQI.T.V, 'g', 6, 64)
	dest.W = strconv.FormatFloat(src.Data.IAQI.W.V, 'g', 6, 64)

	dest.S = src.Data.OriginTime.S
	dest.TZ = src.Data.OriginTime.TZ
	dest.V = src.Data.OriginTime.V

	dest.IndexCity = "" + dest.City + "_" + strconv.Itoa(dest.StationIndex)

	h := sha1.New()
	h.Write([]byte(dest.IndexCity + "_" + strconv.Itoa(dest.V)))
	dest.IndexCityVHash = hex.EncodeToString(h.Sum(nil))
	return dest
}
