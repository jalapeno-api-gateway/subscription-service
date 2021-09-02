package subscribers

func StartSubscriptionService() {
	go handleLsNodeSubscribers()
	go handleLsLinkSubscribers()
	go handlePhysicalInterfaceSubscribers()
	handleLoopbackInterfaceSubscribers()
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

func handlePhysicalInterfaceSubscribers() {
	for {
		subscriptionUpdate := <-physicalIntefaceSubscriberUpdates
		if subscriptionUpdate.Action == Subscribe {
			physicalInterfaceSubscribers = append(physicalInterfaceSubscribers, subscriptionUpdate.UpdateChannel)
		} else if subscriptionUpdate.Action == Unsubscribe {
			index := -1
			for i, updateChannel := range physicalInterfaceSubscribers { // Find subscriber index in array
				if subscriptionUpdate.UpdateChannel == updateChannel {
					index = i
				}
			}
			physicalInterfaceSubscribers = append(physicalInterfaceSubscribers[:index], physicalInterfaceSubscribers[:index+1]...) // Remove subscriber from array
		}
	}
}

func handleLoopbackInterfaceSubscribers() {
	for {
		subscriptionUpdate := <-loopbackIntefaceSubscriberUpdates
		if subscriptionUpdate.Action == Subscribe {
			loopbackInterfaceSubscribers = append(loopbackInterfaceSubscribers, subscriptionUpdate.UpdateChannel)
		} else if subscriptionUpdate.Action == Unsubscribe {
			index := -1
			for i, updateChannel := range loopbackInterfaceSubscribers { // Find subscriber index in array
				if subscriptionUpdate.UpdateChannel == updateChannel {
					index = i
				}
			}
			loopbackInterfaceSubscribers = append(loopbackInterfaceSubscribers[:index], loopbackInterfaceSubscribers[:index+1]...) // Remove subscriber from array
		}
	}
}
