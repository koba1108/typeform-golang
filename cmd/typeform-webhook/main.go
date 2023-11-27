package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/guregu/dynamo"
	"github.com/koba1108/typeform-golang/internal/db"
	"github.com/koba1108/typeform-golang/internal/model"
	"log"
)

var DB *dynamo.DB

func init() {
	DB = db.NewConnection()
}

func main() {
	log.Printf("lambda started")
	lambda.Start(handler)
}

func handler(ctx context.Context, req model.TypeFormEventRequest) (*model.TypeFormEventBody, error) {
	var body *model.TypeFormEventBody
	if err := json.Unmarshal([]byte(req.Body), &body); err != nil {
		return nil, err
	}
	log.Println(body)
	return body, nil
}
