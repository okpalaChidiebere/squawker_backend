package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	fcm_server_url = "https://fcm.googleapis.com/fcm/send"
)

type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

type FcmMessage struct {
	To              string      `json:"to,omitempty"`
	RegistrationIds []string    `json:"registration_ids,omitempty"`
	Data            interface{} `json:"data,omitempty"`
}

func Handler(ctx context.Context, req Request) (Response, error) {

	var buf bytes.Buffer

	// Parse groupId variable from request url
	tId := req.PathParameters["tokenId"]

	//topic := "/topics/key_asser"

	m := FcmMessage{
		To: tId, //send squawk to a Single Device
		//To: topic, //sending this message to multiple devices subscribed to this topic. Another way is to use the "registration_ids" if you know the specific devices you want to push to
		Data: map[string]string{
			"author":    "TestAccount",
			"message":   "Hello!",
			"date":      fmt.Sprintf("%d", time.Since(time.Unix(0, 0)).Milliseconds()), //we chose to use our date in milliseconds
			"authorKey": "key_asser",
		},
	}

	out, _ := json.Marshal(m)
	rb := strings.NewReader(string(out))

	rqst, err := http.NewRequest(http.MethodPost, fcm_server_url, rb)
	if err != nil {
		fmt.Print(err)
	}
	rqst.Header.Add("Authorization", fmt.Sprintf("key=%v", os.Getenv("FIREBASE_SERVER_KEY")))
	rqst.Header.Add("Content-Type", "application/json")

	// An HTTP client for sending the request
	client := &http.Client{}
	resp, err := client.Do(rqst)

	if err != nil {
		log.Printf("Request to %s failed: %s", fcm_server_url, err.Error())
		return Response{StatusCode: resp.StatusCode}, err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Request returned status %s:\n\tHeader: %s\n\tBody %s",
		resp.Status,
		resp.Header,
		b,
	)

	json.HTMLEscape(&buf, b)

	res := Response{
		StatusCode: resp.StatusCode,
		Body:       buf.String(),
		Headers: map[string]string{
			"Content-Type":        "application/json",
			"Squawker-Func-Reply": "hello-handler",
		},
	}

	return res, nil
}

func main() {
	lambda.Start(Handler)
}
