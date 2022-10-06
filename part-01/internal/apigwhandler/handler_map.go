package apigwhandler

import "github.com/aws/aws-lambda-go/events"

type handlerFunc func(response events.APIGatewayV2HTTPRequest) map[string]interface{}

type handlerMap struct {
	Methods     []string
	Route       string
	HandlerFunc handlerFunc
}

func (m handlerMap) Matches(e events.APIGatewayV2HTTPRequest) (handlerFunc, bool) {
	for _, method := range m.Methods {
		if method == e.RequestContext.HTTP.Method && RouteMatches(m.Route, e.RawPath) {
			return m.HandlerFunc, true
		}
	}

	return nil, false
}

func newHandlerMap(route string, hf handlerFunc, methods ...string) handlerMap {
	return handlerMap{
		Methods:     methods,
		Route:       route,
		HandlerFunc: hf,
	}
}
