package kns

import (
	"dsource/data"
	"dsource/dyn"
	"encoding/json"
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

func ReadRecords(region string, stream string) {
	if svc == nil {
		svc = initKinesis(region)
	}

	so, err := svc.DescribeStream(&kinesis.DescribeStreamInput{
		StreamName: aws.String(stream),
	})
	if err != nil {
		log.Println(err)
	}



	itr, err := svc.GetShardIterator(&kinesis.GetShardIteratorInput{
		ShardId:	so.StreamDescription.Shards[0].ShardId,
		StreamName:	aws.String(stream),
		ShardIteratorType: aws.String("TRIM_HORIZON"),
		// ShardIteratorType: aws.String("AT_SEQUENCE_NUMBER"),
		// ShardIteratorType: aws.String("LATEST"),
	})
	if err != nil {
		log.Println(err)
	}


	shardIterator := itr.ShardIterator
	for {
		records, err := svc.GetRecords(&kinesis.GetRecordsInput{
			ShardIterator:	shardIterator,
		})

		if len(records.Records) > 0 {
			for _,record := range records.Records {

				var air data.AirQuality
				err := json.Unmarshal(record.Data, &air)
				if err != nil {
					log.Println(err)
				} else {
					log.Println(air)
					dyn.Save2Dynamo(region, air)
				}
			}
		} else if *records.NextShardIterator == "" || itr.ShardIterator == records.NextShardIterator || err != nil {
			break
		}
		shardIterator = records.NextShardIterator
	}


}
