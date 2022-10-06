package apigwhandler

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golangsugar/chatty"
	"net/http"
)

func getCORSHeaders() map[string]string {
	return map[string]string{
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "*",
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Credentials": "true",
	}
}

func httpRawResponse(statusCode int, statusBody string) map[string]interface{} {
	return map[string]interface{}{
		"statusCode": statusCode,
		"headers":    getCORSHeaders(),
		"body":       statusBody,
	}
}

func responseToMap(e events.APIGatewayV2HTTPResponse) map[string]interface{} {
	data, err := json.Marshal(e)
	if err != nil {
		err = fmt.Errorf("error marshalling response: %v", err)
		chatty.ErrorErr(err)
		return httpRawResponse(http.StatusInternalServerError, err.Error())
	}

	m := make(map[string]interface{})

	if err = json.Unmarshal(data, &m); err != nil {
		err = fmt.Errorf("error unmarshalling response: %v", err)
		chatty.ErrorErr(err)
		return httpRawResponse(http.StatusInternalServerError, err.Error())
	}

	return m
}

// Response creates a well formatted response
func Response(statusCode int, statusBody string, headers map[string]string) map[string]interface{} {
	resp := events.APIGatewayV2HTTPResponse{
		StatusCode: statusCode,
		Headers:    make(map[string]string),
		Body:       statusBody,
	}

	// Add CORS headers
	corsHeaders := getCORSHeaders()
	for k, v := range corsHeaders {
		resp.Headers[k] = v
	}

	// Add the given headers
	for k, v := range headers {
		resp.Headers[k] = v
	}

	return responseToMap(resp)
}

// JSONResponse creates a well formatted json response
func JSONResponse(statusCode int, body interface{}, headers map[string]string) map[string]interface{} {
	j, err := json.Marshal(body)
	if err != nil {
		err = fmt.Errorf("error decoding response body: %v", err)
		chatty.ErrorErr(err)
		return httpRawResponse(http.StatusInternalServerError, err.Error())
	}

	jh := map[string]string{
		"Content-Type": "application/json",
	}

	// Add the given headers
	for k, v := range headers {
		jh[k] = v
	}

	return Response(statusCode, string(j), jh)
}
