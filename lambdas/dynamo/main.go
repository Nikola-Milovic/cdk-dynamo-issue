package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type credentialsProvider struct{}

func (c credentialsProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     "test",
		SecretAccessKey: "test",
		CanExpire:       false,
	}, nil
}

func init() {
	endpointUrl := "http://" + os.Getenv("LOCALSTACK_HOSTNAME") + ":" + os.Getenv("EDGE_PORT")
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentialsProvider{}),
		config.WithRegion("eu-central-1"),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{URL: endpointUrl}, nil
				}),
		))
	if err != nil {
		log.Fatal(err)
	}

	db = *dynamodb.NewFromConfig(sdkConfig, func(o *dynamodb.Options) {})
}

func main() {
	lambda.Start(router)
}
