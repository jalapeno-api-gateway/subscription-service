package subscribers

var lsNodeSubscribers []chan LsNodeEvent
var lsLinkSubscribers []chan LsLinkEvent
var physicalInterfaceSubscribers []chan PhysicalInterfaceEvent
var loopbackInterfaceSubscribers []chan LoopbackInterfaceEvent

var lsNodeSubscriberUpdates = make(chan lsNodeSubscriberUpdate)
var lsLinkSubscriberUpdates = make(chan lsLinkSubscriberUpdate)
var physicalIntefaceSubscriberUpdates = make(chan physicalInterfaceSubscriberUpdate)
var loopbackIntefaceSubscriberUpdates = make(chan loopbackInterfaceSubscriberUpdate)

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

func SubscribeToPhysicalInterfaceEvents(updateChannel chan PhysicalInterfaceEvent) {
	physicalIntefaceSubscriberUpdates <- physicalInterfaceSubscriberUpdate{Action: Subscribe, UpdateChannel: updateChannel}
}

func UnsubscribeFromPhysicalInterfaceEvents(updateChannel chan PhysicalInterfaceEvent) {
	//TODO: Test unsubscribe
	physicalIntefaceSubscriberUpdates <- physicalInterfaceSubscriberUpdate{Action: Unsubscribe, UpdateChannel: updateChannel}
}

func SubscribeToLoopbackInterfaceEvents(updateChannel chan LoopbackInterfaceEvent) {
	loopbackIntefaceSubscriberUpdates <- loopbackInterfaceSubscriberUpdate{Action: Subscribe, UpdateChannel: updateChannel}
}

func UnsubscribeFromLoopbackInterfaceEvents(updateChannel chan LoopbackInterfaceEvent) {
	loopbackIntefaceSubscriberUpdates <- loopbackInterfaceSubscriberUpdate{Action: Unsubscribe, UpdateChannel: updateChannel}
}
