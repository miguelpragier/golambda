package apigwhandler

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type handlerManager struct {
	mapping []handlerMap
}

func (m handlerManager) getHandler(e events.APIGatewayV2HTTPRequest) (handlerFunc, bool) {
	for _, x := range m.mapping {
		if hf, found := x.Matches(e); found {
			return hf, true
		}
	}

	return nil, false
}

func configureHandlers() []handlerMap {
	var a []handlerMap

	a = append(a, newHandlerMap("/api/v1/ping", pingPong, http.MethodGet))
	a = append(a, newHandlerMap("/api/v1/req", requestEcho, http.MethodGet))

	return a
}

func getHandlerManager() handlerManager {
	return handlerManager{mapping: configureHandlers()}
}

func Perform(e events.APIGatewayV2HTTPRequest) map[string]interface{} {
	hm := getHandlerManager()

	if handlerRoutine, found := hm.getHandler(e); found {
		return handlerRoutine(e)
	}

	return NotFound(e)
}
