package db

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func NewConnection() *dynamo.DB {
	sess := session.Must(session.NewSession())
	return dynamo.New(sess)
}
