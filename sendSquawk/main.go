package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type RequestParams struct {
	Message   string      `json:"message"`
	To    string 	`json:"to"` // can be a topic (eg. `/topics/key_test` ) or device registration token id
}

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (Response, error) {

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(MustGetenv("AWS_REGION")))
    if err != nil {
		return Response{}, fmt.Errorf("unable to load SDK config, %v", err)
    }
	svc := s3.NewFromConfig(cfg)

	rawObject, err := svc.GetObject(
		ctx,&s3.GetObjectInput{
			Bucket: aws.String(MustGetenv("THUMBNAILS_S3_BUCKET")), 
			Key:    aws.String("squawker-5b498-firebase-adminsdk-wdxa6-9b77127a4c.json"),
		},
	)
	if err != nil {
		return Response{StatusCode: 500}, fmt.Errorf("error initializing app: %v", err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(rawObject.Body)	
	// myFileContentAsString := buf.String()
	// log.Println(myFileContentAsString)

	opt := option.WithCredentialsJSON(buf.Bytes())
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return Response{StatusCode: 500}, fmt.Errorf("error initializing app: %v", err)
	}

	rp := &RequestParams{}
	json.Unmarshal([]byte(req.Body), rp)

	client, err := app.Messaging(ctx)
	if err != nil {
		return Response{StatusCode: 500}, err
	}

	//send to many devices at once
	// client.SendMulticast(ctx, &messaging.MulticastMessage{
	// 	Tokens: [],
	// })

	_, err = client.Send(ctx, &messaging.Message{
		Token: rp.To, 
		Data: map[string]string{
			"author":    "TestAccount",
			"message":   rp.Message,
			"date":      fmt.Sprintf("%d", time.Since(time.Unix(0, 0)).Milliseconds()), //we chose to use our date in milliseconds
			"authorKey": "key_test",
			"url": "reactndsquawker://about", //about page to test url :)
		},
		Android: &messaging.AndroidConfig{
			// Required for background/quit data-only messages on Android
			// See: https://rnfirebase.io/messaging/usage#data-only-messages
			Priority: "high",
		},
		APNS: &messaging.APNSConfig{
			// required headers for data-only messages iOS
			Headers: map[string]string{
				"apns-push-type": "background",
				"apns-priority": "5",
				"apns-topic": "com.reactndsquawker", // your app bundle identifier
			},
			// required Aps for data-only message 
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					ContentAvailable: true,
				},
			},
			
		},
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}

	body := []byte(`{}`)
	resp := Response{
		StatusCode:      200,
		Body: string(body),
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}


func MustGetenv (k string) string {
	v, ok := os.LookupEnv(k)
	if !ok {
		log.Fatalf("Warning: %s environment variable is not set.", k)
	}
	return v
}