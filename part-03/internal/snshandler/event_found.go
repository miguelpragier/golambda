package snshandler

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golangsugar/chatty"
)

// SNSEventFound returns the events.SNSEvent if present
func SNSEventFound(eventRawJSON json.RawMessage) (events.SNSEvent, bool) {
	var e events.SNSEvent

	if err := json.Unmarshal(eventRawJSON, &e); err != nil {
		return events.SNSEvent{}, false
	}

	if len(e.Records) < 1 || e.Records[0].EventSource != "aws:sns" {
		return events.SNSEvent{}, false
	}

	chatty.Debug("sns event found")

	return e, true
}
