package kns

import (
	"crypto/sha1"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"log"
)

var svc *kinesis.Kinesis

func initKinesis(region string) *kinesis.Kinesis {

	conf := aws.Config{
		Region: aws.String(region),
		//DisableSSL: aws.Bool(true),
		//Credentials: credentials.NewSharedCredentials("~/.aws/credentials", "default"),
	}

	se, err := session.NewSession(&conf)
	if err != nil {
		log.Println(err)
		return nil
	}

	return kinesis.New(se)

}

func Push2Kinesis(region string, stream string, b []byte) {

	if svc == nil {
		svc = initKinesis(region)
	}

	streamInput := kinesis.DescribeStreamInput{
		StreamName: aws.String(stream),
	}
	so, err := svc.DescribeStream(&streamInput)
	if err != nil {
		log.Println(err)
	}
	log.Println(so)

	record := kinesis.PutRecordInput{
		Data:         b,
		StreamName:   aws.String(stream),
		PartitionKey: aws.String(getHash(string(b))),
	}
	ro, err := svc.PutRecord(&record)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%v\n", ro)

}

func getHash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	return string(h.Sum(nil))
}
