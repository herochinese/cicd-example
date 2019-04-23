package dyn

import (
	"dsource/data"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"time"
)

const Table = "AirQualityDataNew"
var svc *dynamodb.DynamoDB

func initSession(region string) *dynamodb.DynamoDB {
	config := aws.Config{
		Region:      aws.String(region),
		//DisableSSL: 	aws.Bool(true),
	}
	se, err := session.NewSession(&config)
	if err != nil {
		log.Println(err)
		return nil
	}
	svc = dynamodb.New(se)
	log.Println("DynamoDB API Version -> ", svc.APIVersion)
	return svc
}


func Save2Dynamo(region string, air data.AirQuality) (*dynamodb.PutItemOutput, time.Duration) {

	start := time.Now()
	if svc == nil {
		initSession(region)
	}
	item, err := dynamodbattribute.MarshalMap(air)
	if err!=nil {
		log.Println(err)
		return nil,time.Since(start)
	}
	log.Println(item)

	input := &dynamodb.PutItemInput {
		Item: item,
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String(Table),
	}
	output, err := svc.PutItem(input)
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
		return output,time.Since(start)
	}
	log.Println(output)
	elapsed := time.Since(start)
	log.Printf("Executed Save2Dynamo took %s .", elapsed)
	return output, elapsed

}