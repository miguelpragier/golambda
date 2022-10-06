package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"golambda/internal/apigwhandler"
)

func LambdaHandler(ctx context.Context, e events.APIGatewayV2HTTPRequest) (map[string]interface{}, error) {
	_ = ctx

	return apigwhandler.Perform(e), nil
}

func main() {
	lambda.Start(LambdaHandler)
}
