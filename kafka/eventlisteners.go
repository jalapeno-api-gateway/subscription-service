package kafka

import (
	"os"

	"github.com/jalapeno-api-gateway/model/class"
)

func StartEventConsumption() {
	consumer := newSaramaConsumer()
	lsNodeEventsConsumer := newPartitionConsumer(consumer, os.Getenv("LSNODE_KAFKA_TOPIC"))
	lsLinkEventsConsumer := newPartitionConsumer(consumer, os.Getenv("LSLINK_KAFKA_TOPIC"))
	lsPrefixEventsConsumer := newPartitionConsumer(consumer, os.Getenv("LSPREFIX_KAFKA_TOPIC"))
	lsSRv6SIDEventsConsumer := newPartitionConsumer(consumer, os.Getenv("LSSRV6SID_KAFKA_TOPIC"))
	telemetryConsumer := newPartitionConsumer(consumer, os.Getenv("TELEMETRY_KAFKA_TOPIC"))
	
	go func() {
		defer func() {
			closeConsumers(
				consumer,
				lsNodeEventsConsumer,
				lsLinkEventsConsumer,
				lsPrefixEventsConsumer,
				lsSRv6SIDEventsConsumer,
				telemetryConsumer,
			)
		}()
		
		for {
			select {
			case msg := <-lsNodeEventsConsumer.Messages(): handleTopologyEvent(unmarshalKafkaMessage(msg), class.LSNode)
			case msg := <-lsLinkEventsConsumer.Messages(): handleTopologyEvent(unmarshalKafkaMessage(msg), class.LSLink)
			case msg := <-lsPrefixEventsConsumer.Messages(): handleTopologyEvent(unmarshalKafkaMessage(msg), class.LSPrefix)
			case msg := <-lsSRv6SIDEventsConsumer.Messages(): handleTopologyEvent(unmarshalKafkaMessage(msg), class.LSSRv6SID)
			case msg := <-telemetryConsumer.Messages(): handleTelemetryEvent(string(msg.Value))
			}
		}
	}()
}