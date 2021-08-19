package kafka

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
)

func StartEventConsumption() {
	consumer := newSaramaConsumer()
	lsNodeEventsConsumer := newPartitionConsumer(consumer, os.Getenv("LSNODE_KAFKA_TOPIC"))
	lsLinkEventsConsumer := newPartitionConsumer(consumer, os.Getenv("LSLINK_KAFKA_TOPIC"))
	go consumeMessages(consumer, lsNodeEventsConsumer, lsLinkEventsConsumer)
}

func consumeMessages(consumer sarama.Consumer, lsNodeEventsConsumer sarama.PartitionConsumer, lsLinkEventsConsumer sarama.PartitionConsumer) {
	defer func() {
		log.Print("consumers closed")
		closeConsumers(consumer, lsNodeEventsConsumer, lsLinkEventsConsumer)
	}()

	for {
		select {
		case msg := <-lsNodeEventsConsumer.Messages():
			LsNodeEvents <- unmarshalKafkaMessage(msg)
		case msg := <-lsLinkEventsConsumer.Messages():
			log.Print("event received")
			LsLinkEvents <- unmarshalKafkaMessage(msg)
		}
	}
}