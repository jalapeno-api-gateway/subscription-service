package kafka

import (
	"context"

	protocol "github.com/influxdata/line-protocol"
	"github.com/jalapeno-api-gateway/jagw-core/arango"
	"github.com/jalapeno-api-gateway/jagw-core/model/class"
	"github.com/jalapeno-api-gateway/subscription-service/events"
	"github.com/jalapeno-api-gateway/subscription-service/helpers"
	"github.com/jalapeno-api-gateway/subscription-service/pubsub"
	"github.com/sirupsen/logrus"
)

const (
	MeasurementIdentifier = "measurement"
)

func handleTopologyEvent(msg KafkaEventMessage, className class.Class) {
	logger := logrus.WithFields(logrus.Fields{"id": msg.Id, "key": msg.Key, "action": msg.Action, "className": className})
	logger.Trace("Handling incoming Topology event from Kafka.")

	ctx := context.Background()
	document := fetchDocument(ctx, logger, msg, className)
	event := events.TopologyEvent{Action: msg.Action, Key: msg.Key, Document: document}
	publishTopologyEvent(logger, event, className)
}

func fetchDocument(ctx context.Context, logger *logrus.Entry, msg KafkaEventMessage, className class.Class) interface{} {
	if msg.Action == "del" {
		switch className {
			case class.LsNode: return arango.LSNode{}
			case class.LsLink: return arango.LSLink{}
			case class.LsPrefix: return arango.LSPrefix{}
			case class.LsSrv6Sid: return arango.LSSRv6SID{}
			case class.LsNodeEdge: return arango.LSNode_Edge{}
			default: 
				logger.Panic("ClassName not implemented.")
			
		}
	}

	switch className {
		case class.LsNode: return arango.FetchLsNode(ctx, msg.Key)
		case class.LsLink: return arango.FetchLsLink(ctx, msg.Key)
		case class.LsPrefix: return arango.FetchLsPrefix(ctx, msg.Key)
		case class.LsSrv6Sid: return arango.FetchLsSrv6Sid(ctx, msg.Key)
		case class.LsNodeEdge: return arango.FetchLsNodeEdge(ctx, msg.Key)
		default:
			logger.Panic("ClassName not implemented.")
		}
	return nil
}

func publishTopologyEvent(logger *logrus.Entry, event events.TopologyEvent, className class.Class) {
	switch className {
		case class.LsNode: pubsub.LsNodeTopic.Publish(event)
		case class.LsLink: pubsub.LsLinkTopic.Publish(event)
		case class.LsPrefix: pubsub.LsPrefixTopic.Publish(event)
		case class.LsSrv6Sid: pubsub.LsSrv6SidTopic.Publish(event)
		case class.LsNodeEdge: pubsub.LsNodeEdgeTopic.Publish(event)
		default:
			logger.Panic("ClassName not implemented.")
		}
}

func handleTelemetryEvent(telemetryString string) {
	metric := parseTelemetryString(telemetryString)
	telemetryEvent := events.TelemetryEvent{Measurement: metric.Name(), Metric: metric}
	pubsub.PublishTelemetry(metric.Name(), telemetryEvent)
}

func parseTelemetryString(telemetryString string) protocol.Metric {
	handler := protocol.NewMetricHandler()
	parser := protocol.NewParser(handler)
	metric, err := parser.Parse([]byte(telemetryString))
	if err != nil {
		logrus.WithError(err).Panic("Failed to parse telemetry string.")
	}
	return metric[0]
}
