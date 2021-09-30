package kafka

import (
	"github.com/jalapeno-api-gateway/jagw-core/model/class"
)

func StartEventConsumption() {
	consumer := newSaramaConsumer()
	lsNodeEventsConsumer := newPartitionConsumer(consumer, LSNODE_KAFKA_TOPIC)
	lsLinkEventsConsumer := newPartitionConsumer(consumer, LSLINK_KAFKA_TOPIC)
	lsPrefixEventsConsumer := newPartitionConsumer(consumer, LSPREFIX_KAFKA_TOPIC)
	lsSRv6SIDEventsConsumer := newPartitionConsumer(consumer, LSSRV6SID_KAFKA_TOPIC)
	telemetryConsumer := newPartitionConsumer(consumer, TELEMETRY_KAFKA_TOPIC)
	
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