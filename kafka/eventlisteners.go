package kafka

import (
	"os"

	"github.com/Shopify/sarama"
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
		case msg := <-lsNodeEventsConsumer.Messages():
			LsNodeEvents <- unmarshalKafkaMessage(msg)
		case msg := <-lsLinkEventsConsumer.Messages():
			LsLinkEvents <- unmarshalKafkaMessage(msg)
		case msg := <-telemetryConsumer.Messages():
			kafkaTelemetryDataRateEvent := createKafkaTelemetryDataRateEvent(string(msg.Value))
			kafkaTelemetryEvent, err := createKafkaTelemetryEvent(string(msg.Value))
			if err == nil {
				TelemetryEvents <- kafkaTelemetryEvent
			}
			if kafkaTelemetryDataRateEvent.DataRate != -1 { //Not all telemetry messages contain a data Rate (if the data Rate is not present a event with negative datarate is created)
				TelemetryDataRateEvents <- kafkaTelemetryDataRateEvent
			}
		}
	}
}
