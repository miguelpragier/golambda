package centralhandler

result, customError := custom_event_handlers.HandleCustomEvent(rawEventJSON,
custom_event_handlers.HandlersMap,
custom_event_handlers.DefaultHandlerParametersMap)
if customError != nil {
return result, nil
}

