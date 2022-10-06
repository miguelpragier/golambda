package apigwhandler

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golangsugar/chatty"
)

// EventFound returns the events.APIGatewayProxyRequest if present
func EventFound(eventRawJSON json.RawMessage) (events.APIGatewayV2HTTPRequest, bool) {
	var e events.APIGatewayV2HTTPRequest

	if err := json.Unmarshal(eventRawJSON, &e); err != nil {
		chatty.Debugf("error unmarshalling apigw msg: %v", err)
		return events.APIGatewayV2HTTPRequest{}, false
	}

	if e.RequestContext.HTTP.Method == "" || e.RawPath == "" {
		chatty.Debug("empty/invalid apigw event found")

		return events.APIGatewayV2HTTPRequest{}, false
	}

	chatty.Debug("apigateway event found")

	return e, true
}
