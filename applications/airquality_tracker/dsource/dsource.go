package main

import (
	"dsource/kns"
	"os"
	"time"
)


func main() {
	region := os.Getenv("MY_AWS_REGION")
	stream := os.Getenv("MY_AWS_KINESIS_NAME")

	for {

		kns.ReadRecords(region, stream)
		time.Sleep(30*1000*time.Millisecond)
	}
}
