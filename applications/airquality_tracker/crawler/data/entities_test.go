package data

import "testing"

func TestSplitName(t *testing.T) {

	tables := []struct {
		str string
		city string
		citycn string
	}{
		{"Beijing (北京)", "Beijing", "北京"},
		{"Shanghai (上海)", "Shanghai", "上海"},
		{"Guangzhou (广州)", "Guangzhou", "广州"},
	}

	for _, table := range tables {
		city, citycn := SplitName(table.str)
		if city!=table.city {
			t.Errorf("city expected %s got %s\n", table.city, city)
		}
		if citycn!=table.citycn {
			t.Errorf("citycn expected %s got %s\n", table.citycn, citycn)
		}
	}

}

func TestCopy2AirQuality(t *testing.T) {

}
