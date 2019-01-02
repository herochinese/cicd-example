package data

import (
	"air/feed"
	"air/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"time"
)

const Table = "AirQualityDataNew"
const TableTransaction = "AirQualityDataNewTransaction"

func Save2Dynamo(air feed.AirQuality) (dynamodb.PutItemOutput, time.Duration) {

	var rt dynamodb.PutItemOutput
	start := time.Now()

	svc := initSession()
	item, err := dynamodbattribute.MarshalMap(air)
	if err!=nil {
		log.Println(err)
		return rt,time.Since(start)
	}
	log.Println(item)

	input := &dynamodb.PutItemInput {
		Item: item,
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String(Table),
	}
	result, err := svc.PutItem(input)
	rt = *result
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				log.Println(dynamodb.ErrCodeConditionalCheckFailedException, aerr.Error())
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				log.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				log.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
				log.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeTransactionConflictException:
				log.Println(dynamodb.ErrCodeTransactionConflictException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				log.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				log.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				log.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Println(err.Error())
		}
		return rt,time.Since(start)
	}
	log.Println(result)
	elapsed := time.Since(start)
	log.Printf("%s took %s", "Save2Dynamo", elapsed)
	return rt, elapsed

}


func Save2DynamoPerf(air feed.AirQuality) {
	_,elapsed := Save2Dynamo(air)

	metric := &cloudwatch.PutMetricDataInput{
		Namespace: aws.String("DynamoDBFoo/Traffic"),
		MetricData: []*cloudwatch.MetricDatum{
			{
				MetricName: aws.String("ElapsedWithPut"),
				Unit:       aws.String("Seconds"),
				Value:      aws.Float64(elapsed.Seconds()),
				Dimensions: []*cloudwatch.Dimension{
					{
						Name:  aws.String("ServiceName"),
						Value: aws.String("AmazonDynamoDB"),
					},
					{
						Name:  aws.String("Method"),
						Value: aws.String("PutItem"),
					},
				},
			},
		},
	}
	util.PutMetric2CW(metric)
}

func TransactSave2Dynamo(air feed.AirQuality) (dynamodb.TransactWriteItemsOutput, time.Duration) {

	var rt dynamodb.TransactWriteItemsOutput
	start := time.Now()

	svc := initSession()
	item, err := dynamodbattribute.MarshalMap(air)
	if err != nil {
		log.Println(err)
		return rt,time.Since(start)
	}
	log.Println(item)

	result, err := svc.TransactWriteItems(&dynamodb.TransactWriteItemsInput{
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TransactItems: []*dynamodb.TransactWriteItem{
			{
				Put: &dynamodb.Put{
					Item:      item,
					TableName: aws.String(TableTransaction),
				},
			},
		},
	})
	rt = *result
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				log.Println(dynamodb.ErrCodeConditionalCheckFailedException, aerr.Error())
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				log.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				log.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
				log.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeTransactionConflictException:
				log.Println(dynamodb.ErrCodeTransactionConflictException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				log.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				log.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				log.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Println(err.Error())
		}
		return rt,time.Since(start)
	}
	log.Println(result)
	elapsed := time.Since(start)
	log.Printf("%s took %s", "TransactSave2Dynamo", elapsed)
	return rt, elapsed

}

func TransactSave2DynamoPerf(air feed.AirQuality) {

	_,elapsed := TransactSave2Dynamo(air)

	metric := &cloudwatch.PutMetricDataInput{
		Namespace: aws.String("DynamoDBFoo/Traffic"),
		MetricData: []*cloudwatch.MetricDatum{
			{
				MetricName: aws.String("ElapsedWithTrans"),
				Unit:       aws.String("Seconds"),
				Value:      aws.Float64(elapsed.Seconds()),
				Dimensions: []*cloudwatch.Dimension{
					{
						Name:  aws.String("ServiceName"),
						Value: aws.String("AmazonDynamoDB"),
					},
					{
						Name:  aws.String("Method"),
						Value: aws.String("TransactWriteItems"),
					},
				},
			},
		},
	}
	util.PutMetric2CW(metric)
}



func initSession() *dynamodb.DynamoDB {
	conf := aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("/Users/chuancc/.aws/credentials", "default"),
	}
	svc := dynamodb.New(session.New(&conf))
	log.Println("DynamoDB API : ", svc.APIVersion)
	return svc
}