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

func createKafkaTelemetryEvent(telemetryString string) KafkaTelemetryEventMessage {
	indexOfDataRate := strings.Index(telemetryString, "data_rates/output_data_rate")
	indexOfIpAddress := strings.Index(telemetryString, "ip_information/ip_address")
	indexOfTotalPacketsSent := strings.Index(telemetryString, "interface_statistics/full_interface_stats/packets_sent")
	indexOfTotalPacketsReceived := strings.Index(telemetryString, "interface_statistics/full_interface_stats/packets_received")

	//if any of the attributes dataRate, ip or totalPacketsSent is not in string its an update for a loopback address
	//loopback update messaged seem to contain no valuable metrics
	if indexOfDataRate == -1 || indexOfIpAddress == -1 || indexOfTotalPacketsSent == -1 || indexOfTotalPacketsReceived == -1 { // if the dataRate or IP is not contained in the telemetry message return an empty Event
		return KafkaTelemetryEventMessage{"", -1, -1, -1, false} //empty message
	}

	dataRate := getDataRateFromTelemetryData(telemetryString)
	ipAddress := getIpAddressFromTelemetryData(telemetryString)
	totalPacketsSent := getTotalPacketsSentFromTelemetryData(telemetryString)
	totalPacketsReceived := getTotalPacketsReceivedFromTelemetryData(telemetryString)

	return KafkaTelemetryEventMessage{
		IpAddress:            ipAddress,
		DataRate:             int64(dataRate),
		TotalPacketsSent:     int64(totalPacketsSent),
		TotalPacketsReceived: int64(totalPacketsReceived),
		ContainsData:         true,
	}
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
	ipAddressRemovedApostrophe := strings.Trim(ipAddress, "\"") //every ip address contains a leading and tailing apostrophe
	return ipAddressRemovedApostrophe
}

func getDataRateFromTelemetryData(telemetryString string) int {
	indexOfDataRate := strings.Index(telemetryString, "data_rates/output_data_rate")
	substring1 := telemetryString[indexOfDataRate:]
	indexOfComma := strings.Index(substring1, ",")
	substring2 := substring1[:indexOfComma]
	split := strings.Split(substring2, "=")
	datarate := split[1]
	datarateRemovedI := datarate[:len(datarate)-1] //every int value contains an i at the end
	dataRateInt, err := strconv.Atoi(datarateRemovedI)
	if err != nil {
		log.Fatalf("Failed to convert string to int: %v", err)
	}
	return dataRateInt
}

func getTotalPacketsSentFromTelemetryData(telemetryString string) int {
	indexOfTotalPacketsSent := strings.Index(telemetryString, "interface_statistics/full_interface_stats/packets_sent")
	substring1 := telemetryString[indexOfTotalPacketsSent:]
	indexOfComma := strings.Index(substring1, ",")
	substring2 := substring1[:indexOfComma]
	split := strings.Split(substring2, "=")
	totalPacketsSent := split[1]
	totalPacketsSentRemoveI := totalPacketsSent[:len(totalPacketsSent)-1] //every int value contains an i at the end
	totalPacketsSentInt, err := strconv.Atoi(totalPacketsSentRemoveI)
	if err != nil {
		log.Fatalf("Failed to convert string to int: %v", err)
	}
	return totalPacketsSentInt
}

func getTotalPacketsReceivedFromTelemetryData(telemetryString string) int {
	indexOfTotalPacketsReceived := strings.Index(telemetryString, "interface_statistics/full_interface_stats/packets_received")
	substring1 := telemetryString[indexOfTotalPacketsReceived:]
	indexOfComma := strings.Index(substring1, ",")
	substring2 := substring1[:indexOfComma]
	split := strings.Split(substring2, "=")
	totalPacketsReceived := split[1]
	totalPacketsReceivedRemoveI := totalPacketsReceived[:len(totalPacketsReceived)-1] //every int value contains an i at the end
	totalPacketsReceivedInt, err := strconv.Atoi(totalPacketsReceivedRemoveI)
	if err != nil {
		log.Fatalf("Failed to convert string to int: %v", err)
	}
	return totalPacketsReceivedInt
}
