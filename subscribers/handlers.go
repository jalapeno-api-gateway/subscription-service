package subscribers

func StartSubscriptionService() {
	go handleLsNodeSubscribers()
	go handleLsLinkSubscribers()
	handleTelemetrySubscriber()
}

func handleTelemetrySubscriber() {
	for {
		subscriptionUpdate := <-dataRateSubscriberUpdates
		if subscriptionUpdate.Action == Subscribe {
			dataRateSubscribers = append(dataRateSubscribers, subscriptionUpdate.UpdateChannel)
		} else if subscriptionUpdate.Action == Unsubscribe {
			index := -1
			for i, updateChannel := range dataRateSubscribers { // Find subscriber index in array
				if subscriptionUpdate.UpdateChannel == updateChannel {
					index = i
				}
			}
			dataRateSubscribers = append(dataRateSubscribers[:index], dataRateSubscribers[:index+1]...) // Remove subscriber from array
		}
	}
}

func handleLsNodeSubscribers() {
	for {
		subscriptionUpdate := <-lsNodeSubscriberUpdates
		if subscriptionUpdate.Action == Subscribe {
			lsNodeSubscribers = append(lsNodeSubscribers, subscriptionUpdate.UpdateChannel)
		} else if subscriptionUpdate.Action == Unsubscribe {
			index := -1
			for i, updateChannel := range lsNodeSubscribers { // Find subscriber index in array
				if subscriptionUpdate.UpdateChannel == updateChannel {
					index = i
				}
			}
			lsNodeSubscribers = append(lsNodeSubscribers[:index], lsNodeSubscribers[:index+1]...) // Remove subscriber from array
		}
	}
}

func handleLsLinkSubscribers() {
	for {
		subscriptionUpdate := <-lsLinkSubscriberUpdates
		if subscriptionUpdate.Action == Subscribe {
			lsLinkSubscribers = append(lsLinkSubscribers, subscriptionUpdate.UpdateChannel)
		} else if subscriptionUpdate.Action == Unsubscribe {
			index := -1
			for i, updateChannel := range lsLinkSubscribers { // Find subscriber index in array
				if subscriptionUpdate.UpdateChannel == updateChannel {
					index = i
				}
			}
			// Remove subscriber from array
			lsLinkSubscribers[index] = lsLinkSubscribers[len(lsLinkSubscribers)-1]
			lsLinkSubscribers = lsLinkSubscribers[:len(lsLinkSubscribers)-1]
		}
	}
}
