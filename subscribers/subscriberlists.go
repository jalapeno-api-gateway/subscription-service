package subscribers

var lsNodeSubscribers []chan LsNodeEvent
var lsLinkSubscribers []chan LsLinkEvent
var dataRateSubscribers []chan DataRateEvent
var telemetrySubscribers []chan TelemetryEvent

var lsNodeSubscriberUpdates = make(chan lsNodeSubscriberUpdate)
var lsLinkSubscriberUpdates = make(chan lsLinkSubscriberUpdate)
var dataRateSubscriberUpdates = make(chan dataRateSubscriberUpdate)
var telemetrySubscriberUpdates = make(chan telemetrySubscriberUpdate)

func SubscribeToLsNodeEvents(updateChannel chan LsNodeEvent) {
	lsNodeSubscriberUpdates <- lsNodeSubscriberUpdate{Action: Subscribe, UpdateChannel: updateChannel}
}

func UnSubscribeFromLsNodeEvents(updateChannel chan LsNodeEvent) {
	lsNodeSubscriberUpdates <- lsNodeSubscriberUpdate{Action: Unsubscribe, UpdateChannel: updateChannel}
}

func SubscribeToLsLinkEvents(updateChannel chan LsLinkEvent) {
	lsLinkSubscriberUpdates <- lsLinkSubscriberUpdate{Action: Subscribe, UpdateChannel: updateChannel}
}

func UnSubscribeFromLsLinkEvents(updateChannel chan LsLinkEvent) {
	lsLinkSubscriberUpdates <- lsLinkSubscriberUpdate{Action: Unsubscribe, UpdateChannel: updateChannel}
}

func SubscribeToDataRateEvents(updateChannel chan DataRateEvent) {
	dataRateSubscriberUpdates <- dataRateSubscriberUpdate{Action: Subscribe, UpdateChannel: updateChannel}
}

func UnsubscribeFromDataRateEvents(updateChannel chan DataRateEvent) {
	dataRateSubscriberUpdates <- dataRateSubscriberUpdate{Action: Unsubscribe, UpdateChannel: updateChannel}
}

func SubscribeToTelemetryEvents(updateChannel chan TelemetryEvent) {
	telemetrySubscriberUpdates <- telemetrySubscriberUpdate{Action: Subscribe, UpdateChannel: updateChannel}
}

func UnsubscribeFromTelemetryEvents(updateChannel chan TelemetryEvent) {
	//TODO: Test unsubscribe
	telemetrySubscriberUpdates <- telemetrySubscriberUpdate{Action: Unsubscribe, UpdateChannel: updateChannel}
}
