package kafka

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/Shopify/sarama"
)

func unmarshalKafkaMessage(msg *sarama.ConsumerMessage) KafkaEventMessage {
	var event KafkaEventMessage
	err := json.Unmarshal(msg.Value, &event)
	if err != nil {
		log.Fatalf("Could not unmarshal kafka message, %v", err)
	}
	return event
}

func createKafkaTelemetryDataRateEvent(telemetryString string) KafkaTelemetryDataRateEventMessage {
	//telemetryString contains all telelmetry attributes separated by comma
	//e.g: Cisco-IOS-XR-pfi-im-cmd-oper:interfaces/interface-xr/interface,host=telegraf,interface_name=GigabitEthernet0/0/0/0
	indexOfDataRate := strings.Index(telemetryString, "data_rates/output_data_rate")
	indexOfIpAddress := strings.Index(telemetryString, "ip_information/ip_address")
	if indexOfDataRate == -1 || indexOfIpAddress == -1 { // if the dataRate or IP is not contained in the telemetry message return an empty Event
		return KafkaTelemetryDataRateEventMessage{"", -1}
	}
	dataRate := getDataRateFromTelemetryData(telemetryString)
	ipAddress := getIpAddressFromTelemetryData(telemetryString)
	return KafkaTelemetryDataRateEventMessage{ipAddress, int64(dataRate)}
}

func getIpAddressFromTelemetryData(telemetryString string) string {
	indexOfIpAddress := strings.Index(telemetryString, "ip_information/ip_address")
	substring1 := telemetryString[indexOfIpAddress:]
	indexOfComma := strings.Index(substring1, ",")
	substring2 := substring1[:indexOfComma]
	split := strings.Split(substring2, "=")
	ipAddress := split[1]
	ipAddressRemovedApostrophe := strings.Trim(ipAddress, "\"")
	return ipAddressRemovedApostrophe
}

func getDataRateFromTelemetryData(telemetryString string) int {
	indexOfDataRate := strings.Index(telemetryString, "data_rates/output_data_rate")
	substring1 := telemetryString[indexOfDataRate:]
	indexOfComma := strings.Index(substring1, ",")
	substring2 := substring1[:indexOfComma]
	split := strings.Split(substring2, "=")
	datarate := split[1]
	datarateRemovedI := datarate[:len(datarate)-1]
	dataRateInt, err := strconv.Atoi(datarateRemovedI)
	if err != nil {
		log.Fatalf("Failed to convert string to int: %v", err)
	}
	return dataRateInt
}
