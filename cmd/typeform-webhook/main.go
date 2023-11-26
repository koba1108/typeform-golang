package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/guregu/dynamo"
	"github.com/koba1108/typeform-golang/internal/db"
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

type TypeformWebhookRequest struct {
	FormResponse FormResponse `json:"form_response"`
}

type FormResponse struct {
	FormID      string      `json:"form_id"`
	Token       string      `json:"token"`
	LandedAt    string      `json:"landed_at"`
	SubmittedAt string      `json:"submitted_at"`
	Metadata    Metadata    `json:"metadata"`
	Definition  Definition  `json:"definition"`
	Answers     []Answers   `json:"answers"`
	Hidden      interface{} `json:"hidden"`
}

type Metadata struct {
	UserAgent string `json:"user_agent"`
	Platform  string `json:"platform"`
	Referer   string `json:"referer"`
	NetworkID string `json:"network_id"`
}

type Definition struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Fields []Field `json:"fields"`
}

type Field struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

type Answers struct {
	Type    string   `json:"type"`
	Field   Field    `json:"field"`
	Choice  Choice   `json:"choice"`
	Text    string   `json:"text"`
	Email   string   `json:"email"`
	Url     string   `json:"url"`
	Number  float64  `json:"number"`
	Date    string   `json:"date"`
	Bool    bool     `json:"bool"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Label string `json:"label"`
}

func handler(ctx context.Context, req TypeformWebhookRequest) error {
	j, err := json.Marshal(req)
	if err != nil {
		return err
	}
	log.Println(string(j))
	return nil
}
