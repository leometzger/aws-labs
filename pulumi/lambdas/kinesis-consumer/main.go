package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleLambdaEvent(ctx context.Context, event events.KinesisEvent) error {
	for _, record := range event.Records {
		fmt.Println("EventID", record.EventID)
	}

	return nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
