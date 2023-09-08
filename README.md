# Setup

1. Add your key to the `docker-compose`
1. `npm install`
1. `docker-compose up`
1. `npm run bootstrap && npm run deploy`
1. `curl -d "email=example@domain.com" -H "Content-Type: application/x-www-form-urlencoded" -X POST <URL>`
*NOTE* the `<URL>` will be the output of `npm run deploy`, the API URL will be printed to the console as `CfnOutput`

```
2023/09/05 18:36:44 failed to put item: operation error DynamoDB: PutItem, https response error StatusCode: 400, RequestID: c823ba61-3c9a-402d-ac71-a5679c6a35e1, ResourceNotFoundException: Cannot do operations on a non-existent table
```


Table does exist

```bash
+ lambda-dynamo ➜ awslocal dynamodb list-tables --region eu-central-1
{
    "TableNames": [
        "LambdaStack-MyDynamoTableACB5AA21-2f181648"
    ]
}
```

Output from lambda

```
[2023/09/06/[$LATEST]46f4fa75dc5935dcf009edef14ad52fb]
START RequestId: 84559536-bce7-4f99-85c0-5e00b441a5b4 Version: $LATEST
2023-09-06 06:46:56
[2023/09/06/[$LATEST]46f4fa75dc5935dcf009edef14ad52fb]
2023/09/06 04:46:55 Received req events.APIGatewayProxyRequest{Resource:"/", Path:"/", HTTPMethod:"POST", Headers:map[string]string{"X-Forwarded-For":"172.24.0.1, vruqwuhlht.execute-api.localhost.localstack.cloud:4566", "accept":"*/*", "content-length":"24", "content-type":"application/x-www-form-urlencoded", "host":"vruqwuhlht.execute-api.localhost.localstack.cloud:4566", "user-agent":"curl/7.81.0", "x-localstack-edge":"https://vruqwuhlht.execute-api.localhost.localstack.cloud:4566", "x-localstack-tgt-api":"apigateway"}, MultiValueHeaders:map[string][]string{"X-Forwarded-For":[]string{"172.24.0.1, vruqwuhlht.execute-api.localhost.localstack.cloud:4566"}, "accept":[]string{"*/*"}, "content-length":[]string{"24"}, "content-type":[]string{"application/x-www-form-urlencoded"}, "host":[]string{"vruqwuhlht.execute-api.localhost.localstack.cloud:4566"}, "user-agent":[]string{"curl/7.81.0"}, "x-localstack-edge":[]string{"https://vruqwuhlht.execute-api.localhost.localstack.cloud:4566"}, "x-localstack-tgt-api":[]string{"apigateway"}}, QueryStringParameters:map[string]string(nil), MultiValueQueryStringParameters:map[string][]string(nil), PathParameters:map[string]string{}, StageVariables:map[string]string{}, RequestContext:events.APIGatewayProxyRequestContext{AccountID:"000000000000", ResourceID:"gg5i62g769", OperationName:"", Stage:"prod", DomainName:"vruqwuhlht.execute-api.localhost.localstack.cloud", DomainPrefix:"vruqwuhlht", RequestID:"1dda726d-8f65-44f2-9b86-e76e7b1cb494", ExtendedRequestID:"", Protocol:"HTTP/1.1", Identity:events.APIGatewayRequestIdentity{CognitoIdentityPoolID:"", AccountID:"000000000000", CognitoIdentityID:"", Caller:"", APIKey:"", APIKeyID:"", AccessKey:"", SourceIP:"172.24.0.1", CognitoAuthenticationType:"", CognitoAuthenticationProvider:"", UserArn:"", UserAgent:"curl/7.81.0", User:""}, ResourcePath:"/", Path:"/prod/", Authorizer:map[string]interface {}{}, HTTPMethod:"POST", RequestTime:"06/Sep/2023:04:46:54 +0000", RequestTimeEpoch:1693975614426, APIID:"vruqwuhlht"}, Body:"email=example%40domain.com", IsBase64Encoded:false}
2023-09-06 06:46:56
[2023/09/06/[$LATEST]46f4fa75dc5935dcf009edef14ad52fb]
2023/09/06 04:46:55 Received POST request with item: {Email:example@domain.com}
2023-09-06 06:46:56
[2023/09/06/[$LATEST]46f4fa75dc5935dcf009edef14ad52fb]
table is "LambdaStack-MyDynamoTableACB5AA21-2f181648"
2023-09-06 06:46:56
[2023/09/06/[$LATEST]46f4fa75dc5935dcf009edef14ad52fb]
tables are []
2023-09-06 06:46:56
[2023/09/06/[$LATEST]46f4fa75dc5935dcf009edef14ad52fb]
2023/09/06 04:46:55 failed to put item: operation error DynamoDB: PutItem, https response error StatusCode: 400, RequestID: 4bd7021a-e466-49c3-b0b9-a6e498a678ee, ResourceNotFoundException: Cannot do operations on a non-existent table
2023-09-06 06:46:56
[2023/09/06/[$LATEST]46f4fa75dc5935dcf009edef14ad52fb]
END RequestId: 84559536-bce7-4f99-85c0-5e00b441a5b4
```
