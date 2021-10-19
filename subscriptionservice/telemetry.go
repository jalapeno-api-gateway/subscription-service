package subscriptionservice

import (
	"github.com/jalapeno-api-gateway/jagw-core/model/property"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"github.com/jalapeno-api-gateway/subscription-service/events"
	"google.golang.org/protobuf/proto"
)

func convertPhysicalInterfaceEvent(event events.PhysicalInterfaceEvent, propertyNames []string) *jagw.TelemetryEvent {
	if len(propertyNames) == 0 { // If no propertyNames were provided, all Properties are returned to the SR-App
		propertyNames = property.AllPhysicalInterfaceProperties
	}

	telemetryEvent := jagw.TelemetryEvent{
		InterfaceId: &jagw.InterfaceIdentifier{Hostname: &event.Hostname, LinkId: &event.LinkID},
	}

	for _, propertyName := range propertyNames {
		switch propertyName {
		case property.DataRate:
			telemetryEvent.DataRate = proto.Int64(event.DataRate)
		case property.PacketsSent:
			telemetryEvent.PacketsSent = proto.Int64(event.PacketsSent)
		case property.PacketsReceived:
			telemetryEvent.PacketsReceived = proto.Int64(event.PacketsReceived)
		}
	}

	return &telemetryEvent
}

func convertLoopbackInterfaceEvent(event events.LoopbackInterfaceEvent, propertyNames []string) *jagw.TelemetryEvent {
	if len(propertyNames) == 0 { // If no propertyNames were provided, all Properties are returned to the SR-App
		propertyNames = property.AllLoopbackInterfaceProperties
	}

	telemetryEvent := jagw.TelemetryEvent{
		InterfaceId: &jagw.InterfaceIdentifier{Hostname: &event.Hostname, LinkId: &event.LinkID},
	}

	for _, propertyName := range propertyNames {
		switch propertyName {
		case property.State:
			telemetryEvent.State = proto.String(event.State)
		case property.LastStateTransitionTime:
			telemetryEvent.LastStateTransitionTime = proto.Int64(event.LastStateTransitionTime)
		}
	}

	return &telemetryEvent
}