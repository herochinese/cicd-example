package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var LocalCache map[string]AirQuality

func NewRouter() *mux.Router {

	LocalCache = make(map[string]AirQuality)
	ProcessMessage()

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{

	Route{
		"AllCitiesAirQuality",
		"GET",
		"/cities",
		Cities,
	},
	Route{
		"CityAirQuality",
		"GET",
		"/cities/{city}",
		City,
	},
	Route{
		"AirQualityFeed",
		"POST",
		"/feed",
		Feed,
	},
}
