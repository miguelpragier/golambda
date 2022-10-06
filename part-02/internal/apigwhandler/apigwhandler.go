package apigwhandler

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

func pingPong(e events.APIGatewayV2HTTPRequest) map[string]interface{} {
	return httpRawResponse(http.StatusOK, "pong")
}

func requestEcho(e events.APIGatewayV2HTTPRequest) map[string]interface{} {
	return JSONResponse(http.StatusOK, e, nil)
}

func Perform(e events.APIGatewayV2HTTPRequest) map[string]interface{} {
	hm := getHandlerManager()

	if handlerRoutine, found := hm.getHandler(e); found {
		return handlerRoutine(e)
	}

	return NotFound(e)
}
