import { CfnOutput, RemovalPolicy, Stack, StackProps } from 'aws-cdk-lib';
import { AttributeType } from 'aws-cdk-lib/aws-dynamodb';
import { Table } from 'aws-cdk-lib/aws-dynamodb';
import { Role, ServicePrincipal, ManagedPolicy } from 'aws-cdk-lib/aws-iam';
import { Construct } from 'constructs';
import * as go from "@aws-cdk/aws-lambda-go-alpha"
import { LambdaRestApi } from 'aws-cdk-lib/aws-apigateway';
import path = require('path');

export class LambdaStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    // DynamoDB Table
    const dynamoTable = new Table(this, 'MyDynamoTable', {
      partitionKey: { name: 'email', type: AttributeType.STRING },
      sortKey: { name: 'signedup_at', type: AttributeType.STRING },
      removalPolicy: RemovalPolicy.DESTROY  // Only for dev/test environments
    });

    // IAM Role for Lambda to access DynamoDB
    const lambdaRole = new Role(this, 'LambdaRole', {
      assumedBy: new ServicePrincipal('lambda.amazonaws.com')
    });

    lambdaRole.addManagedPolicy(ManagedPolicy.fromAwsManagedPolicyName('AmazonDynamoDBFullAccess'));

    // Lambda function
    const lambdaFunction = new go.GoFunction(this, 'handler', {
      entry: path.join(__dirname, '../lambdas/dynamo'),
      environment: {
        "TABLE_NAME": dynamoTable.tableName
      }
    });

    const api = new LambdaRestApi(this, 'Endpoint', {
      handler: lambdaFunction,
      defaultCorsPreflightOptions: {
        allowOrigins: ['*'], // Adjust this as needed
      }
    });

    // Grant Lambda permissions to write to the DynamoDB table
    dynamoTable.grantFullAccess(lambdaFunction);

    new CfnOutput(this, 'ApiUrl', {
      value: api.url,
      description: 'The URL of the API Gateway associated with the Lambda function',
      exportName: 'ApiUrl',
    });

    new CfnOutput(this, 'TableName', {
      value: dynamoTable.tableName,
    })
  }
}

