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
