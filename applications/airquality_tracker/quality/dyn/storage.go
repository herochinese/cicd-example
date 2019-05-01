package dyn

import (
	"quality/api"
	"quality/data"
)

func RecoverData() {
	//TODO: Recover data from DynamoDB
	d := data.AirQuality{
		IndexCityVHash: "175bf71c6370151071d139083b1a95ac3af7148d",
		IndexCity: "Demo_9999",
		StationIndex: 9999,
		AQI: 999,
		City: "Demo",
		CityCN: "Demo",
		Latitude: "99.9999",
		Longitude: "99.9999",
		Co: "99.9999",
		H: "99.9999",
		No2: "99.9999",
		P: "99.9999",
		Pm10: "99.9999",
		Pm25: "99.9999",
		So2: "99.9999",
		T: "99.9999",
		W: "99.9999",
		S: "2019-12-30 21:00:00",
		TZ: "+08:00",
		V: 999999,
	}

	api.WorkQueue <- d
}


/*
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