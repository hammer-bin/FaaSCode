package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestModel struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, request RequestModel) (string, error) {
	if len(request.Name) > 0 {
		return fmt.Sprintf("Hello %s!", request.Name), nil
	}
	return "Hello world", nil
}

func main() {
	lambda.Start(HandleRequest)
}
