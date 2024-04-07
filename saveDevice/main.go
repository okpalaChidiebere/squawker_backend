package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestParams struct {
	Device struct {
		DeviceID string    `json:"deviceId"`
    	Token    string    `json:"token"`
    	Active   bool      `json:"active"`
	} `json:"device"`
}

type Response events.APIGatewayV2HTTPResponse


func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (Response, error) {
	rp := &RequestParams{}
	json.Unmarshal([]byte(req.Body), rp)

	r, _ := json.MarshalIndent(rp, "", " ")
	log.Printf("Saving Device: %s", r)

	body := []byte(`{}`)
	return Response{StatusCode: 200, Body: string(body) }, nil
}

func main() {
	lambda.Start(Handler)
}
