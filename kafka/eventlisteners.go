package kafka

import (
	"os"

	"github.com/Shopify/sarama"
	"github.com/jalapeno-api-gateway/model/class"
)

func StartEventConsumption() {
	consumer := newSaramaConsumer()
	lsNodeEventsConsumer := newPartitionConsumer(consumer, os.Getenv("LSNODE_KAFKA_TOPIC"))
	lsLinkEventsConsumer := newPartitionConsumer(consumer, os.Getenv("LSLINK_KAFKA_TOPIC"))
	telemetryConsumer := newPartitionConsumer(consumer, os.Getenv("TELEMETRY_KAFKA_TOPIC"))
	go consumeMessages(consumer, lsNodeEventsConsumer, lsLinkEventsConsumer, telemetryConsumer)
}

func consumeMessages(consumer sarama.Consumer, lsNodeEventsConsumer sarama.PartitionConsumer, lsLinkEventsConsumer sarama.PartitionConsumer, telemetryConsumer sarama.PartitionConsumer) {
	defer func() {
		closeConsumers(consumer, lsNodeEventsConsumer, lsLinkEventsConsumer, telemetryConsumer)
	}()

	for {
		select {
		case msg := <-lsNodeEventsConsumer.Messages(): handleTopologyEvent(unmarshalKafkaMessage(msg), class.LSNode)
		case msg := <-lsLinkEventsConsumer.Messages(): handleTopologyEvent(unmarshalKafkaMessage(msg), class.LSLink)
		case msg := <-telemetryConsumer.Messages(): handleTelemetryEvent(string(msg.Value))
		}
	}
}
