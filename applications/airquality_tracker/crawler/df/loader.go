package df

import (
	"crawler/data"
	"encoding/json"
	"io/ioutil"
	"log"
)

func LoadCities(file string) []string {
	var pr []data.Province
	body, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}

	err2 := json.Unmarshal(body, &pr)
	if err2 != nil {
		log.Println(err)
	}
	cs := make([]string, len(pr))

	for i, p := range pr {
		cs[i] = p.City[0].County[0].Name
	}
	return cs
}
