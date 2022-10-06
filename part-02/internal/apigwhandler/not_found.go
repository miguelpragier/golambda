package apigwhandler

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/golangsugar/chatty"
	"net/http"
)

// NotFound returns a valid http 404 response
func NotFound(e events.APIGatewayV2HTTPRequest) map[string]interface{} {
	msg := "can't handle requested path " + e.RequestContext.HTTP.Method + " " + e.RawPath

	chatty.Warn(msg)

	return httpRawResponse(http.StatusNotFound, msg)
}
