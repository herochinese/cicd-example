package kns

import (
	"crypto/sha1"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"log"
	"os"
)



var svc *kinesis.Kinesis


func initKinesis() *kinesis.Kinesis {
	region := os.Getenv("AWS_REGION")
	conf := aws.Config{
		Region:      aws.String(region),
		DisableSSL: 	aws.Bool(true),
	}

	se, err := session.NewSession( &conf)
	if err != nil {
		log.Println(err)
		return nil
	}

	return kinesis.New(se)

}

func Push2Kinesis(stream string, json string) {


	if svc != nil {
		record := kinesis.PutRecordInput{
			Data: []byte(json),
			StreamName: aws.String(stream),
			PartitionKey: aws.String(getHash(json)),

		}
		output, err := svc.PutRecord(&record)
		if err != nil {
			log.Println(err)
		}
		log.Printf("%v\n", output)

	} else {
		svc = initKinesis()
	}

}

func getHash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	return  string(h.Sum(nil))
}