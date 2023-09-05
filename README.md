# Setup

1. Add your key to the `docker-compose`
1. `npm install`
1. `make up`
1. `npm run bootstrap && npm run deploy`
1. `curl -d "email=example@domain.com" -H "Content-Type: application/x-www-form-urlencoded" -X POST <URL>`

```
2023/09/05 18:36:44 failed to put item: operation error DynamoDB: PutItem, https response error StatusCode: 400, RequestID: c823ba61-3c9a-402d-ac71-a5679c6a35e1, ResourceNotFoundException: Cannot do operations on a non-existent table
```
