package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Signup struct {
	Email string    `json:"email" dynamodbav:"email"`
	Date  time.Time `json:"date" dynamodbav:"signedup_at"`
}

var TableName = os.Getenv("TABLE_NAME")

var db dynamodb.Client

func createSignup(ctx context.Context, req SignupRequest) error {
	signup := Signup{
		Email: req.Email,
		Date:  time.Now(),
	}

	item, err := attributevalue.MarshalMap(signup)
	if err != nil {
		return fmt.Errorf("failed to DynamoDB marshal Record, %v", err)
	}

	fmt.Printf("table is %q\n", TableName)

	res, err := db.ListTables(ctx, &dynamodb.ListTablesInput{})
	if err != nil {
		return fmt.Errorf("failed to list tables: %v", err)
	}

	fmt.Printf("tables are %v\n", res.TableNames)

	input := &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
		Item:      item,
	}

	if _, err := db.PutItem(ctx, input); err != nil {
		return fmt.Errorf("failed to put item: %v", err)
	}

	return nil
}
