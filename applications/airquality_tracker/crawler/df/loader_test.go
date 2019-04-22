package df_test

import (
	"crawler/df"
	"testing"
)

func TestLoadCities(t *testing.T) {
	f := "../china-cities.json"
	cities := df.LoadCities(f)
	if cities != nil {

		if len(cities) != 2 {
			t.Errorf("length of city[] expected %s got %d\n", "2", len(cities))
		}

		if cities[0] != "beijing" {

		}
		if cities[1] != "shanghai" {
			t.Errorf("city[1] expected %s got %s\n", "shanghai", cities[1])
		}
	} else {
		t.Errorf("loading [ %s ] was failed.\n", f)
	}

}
