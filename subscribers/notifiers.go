package subscribers

func NotifyLsNodeSubscribers(event LsNodeEvent) {
	for _, subscriber := range lsNodeSubscribers {
		subscriber <- event
	}
}

func NotifyLsLinkSubscribers(event LsLinkEvent) {
	for _, subscriber := range lsLinkSubscribers {
		subscriber <- event
	}
}

func NotifyPhysicalInterfaceSubscribers(event PhysicalInterfaceEvent) {
	for _, subscriber := range physicalInterfaceSubscribers {
		subscriber <- event
	}
}

func NotifyLoopbackInterfaceSubscribers(event LoopbackInterfaceEvent) {
	for _, subscriber := range loopbackInterfaceSubscribers {
		subscriber <- event
	}
}
