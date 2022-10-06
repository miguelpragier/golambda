package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"golambda/internal/centralhandler"
)

func LambdaHandler(ctx context.Context, rawEventJSON json.RawMessage) (map[string]interface{}, error) {
	// NonBlocking sentinel that waits for a timeout
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("request timeout: %w", ctx.Err())
	default:
	}

	return centralhandler.Perform(rawEventJSON)
}

func main() {
	lambda.Start(LambdaHandler)
}
