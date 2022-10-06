package centralhandler

import (
	"encoding/json"
	"golambda/internal/snshandler"
)

func Perform(rawJSONEvent json.RawMessage) (map[string]interface{}, error) {
	if event, ok := snshandler.SNSEventFound(rawEventJSON); ok {
		return handleSNSEvent(event)
	}

	if event, ok := apigw.EventFound(rawEventJSON); ok {
		resp, err := handleAPIGatewayEvent(event)

		return apigw.ResponseToMap(resp), err
	}

	return map[string]interface{}{"event": string(rawEventJSON)}, nil

}
