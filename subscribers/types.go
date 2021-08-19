package subscribers

import (
	"gitlab.ost.ch/ins/jalapeno-api/push-service/arangodb"
)

//
// ---> ACTIONS <---
//

type Action string

const (
	Subscribe	Action = "subscribe"
	Unsubscribe	Action = "unsubscribe"
)

//
// ---> EVENTS <---
//

type LsNodeEvent struct {
	Action			string
	Key				string
	LsNodeDocument	arangodb.LsNodeDocument
}

type LsLinkEvent struct {
	Action			string
	Key				string
	LsLinkDocument	arangodb.LsLinkDocument
}

//
// ---> SUBSCRIBER UPDATES <---
//

type lsNodeSubscriberUpdate struct {
	Action			Action
	UpdateChannel	chan LsNodeEvent
}

type lsLinkSubscriberUpdate struct {
	Action			Action
	UpdateChannel	chan LsLinkEvent
}