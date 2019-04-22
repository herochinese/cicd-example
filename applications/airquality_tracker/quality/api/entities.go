package api

type AirQuality struct {
	IndexCityVHash string `json:"index_city_v_hash"`
	IndexCity string `json:"index_city"`
	StationIndex int `json:"idx"`
	AQI int `json:"aqi"`
	City string `json:"city"`
	CityCN string `json:"city_cn"`
	Latitude string `json:"lat"`
	Longitude string `json:"lng"`
	Co string `json:"co"`
	H string `json:"h"`
	No2 string `json:"no2"`
	O3 string `json:"o3"`
	P string `json:"p"`
	Pm10 string `json:"pm10"`
	Pm25 string `json:"pm25"`
	So2 string `json:"so2"`
	T string `json:"t"`
	W string `json:"w"`
	S string `json:"s"` //Local measurement time
	TZ string `json:"tz"` //Station timezone
	V int `json:"v"`
}


type OriginAirQuality struct {
	Status string `json:"status"`
	Data OriginData `json:"data"`
}

type OriginData struct {
	AQI int `json:"aqi"`
	StationIndex int `json:"idx"`
	City OriginCity `json:"city"`
	IAQI OriginIAQI	`json:"iaqi"`
	OriginTime OriginTime `json:"time"`
}

type OriginCity struct {
	Geo []float64 `json:"geo"`
	Name string `json:"name"`
}

type OriginIAQI struct {
	Co OValue `json:"co"`
	H OValue `json:"h"`
	No2 OValue `json:"no2"`
	O3 OValue `json:"o3"`
	P OValue `json:"p"`
	Pm10 OValue `json:"pm10"`
	Pm25 OValue `json:"pm25"`
	So2 OValue `json:"so2"`
	T OValue `json:"t"`
	W OValue `json:"w"`
}

type OValue struct {
	V float64 `json:"v"`
}

type OriginTime struct {
	S string `json:"s"` //Local measurement time
	TZ string `json:"tz"` //Station timezone
	V int `json:"v"`
}
