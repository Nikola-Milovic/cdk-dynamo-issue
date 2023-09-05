package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
)

type SignupRequest struct {
	Email string `json:"email"`
}

func processPost(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.Headers["content-type"] != "application/x-www-form-urlencoded" {
		log.Printf("Unsupported content type: %s", req.Headers["content-type"])
		return clientError(http.StatusUnprocessableEntity)
	}

	values, err := url.ParseQuery(req.Body)
	if err != nil {
		log.Printf("Can't parse form data: %v", err)
		return clientError(http.StatusUnprocessableEntity)
	}
	email := values.Get("email")
	if email == "" {
		log.Printf("Email not found in form data")
		return clientError(http.StatusBadRequest)
	}

	request := SignupRequest{
		Email: email,
	}

	log.Printf("Received POST request with item: %+v", request)

	if err := createSignup(ctx, request); err != nil {
		return serverError(err)
	}

	fmt.Println("Created signup, email: ", request.Email)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       "{\"status\": \"success\"}",
	}, nil
}

func router(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received req %#v", req)

	switch req.HTTPMethod {
	case "POST":
		return processPost(ctx, req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       http.StatusText(status),
		StatusCode: status,
	}, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	log.Println(err.Error())

	return events.APIGatewayProxyResponse{
		Body:       http.StatusText(http.StatusInternalServerError),
		StatusCode: http.StatusInternalServerError,
	}, nil
}
