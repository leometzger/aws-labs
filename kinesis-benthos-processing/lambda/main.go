package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type CustomEvent struct {
	Created    string
	Timestamp  string
	X          float32
	Y          float32
	Error      uint32
	Associated bool
}

func HandleLambdaEvent(ctx context.Context, event []CustomEvent) error {
	fmt.Println("received:", len(event))
	return nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
