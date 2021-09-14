package kafka

import (
	"os"

	"github.com/Shopify/sarama"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/model"
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
		case msg := <-lsNodeEventsConsumer.Messages(): handleTopologyEvent(unmarshalKafkaMessage(msg), model.LsNodeEvent)
		case msg := <-lsLinkEventsConsumer.Messages(): handleTopologyEvent(unmarshalKafkaMessage(msg), model.LsLinkEvent)
		case msg := <-telemetryConsumer.Messages(): handleTelemetryEvent(string(msg.Value))
		}
	}
}
