package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jalapeno-api-gateway/jagw-core/arango"
	"github.com/jalapeno-api-gateway/subscription-service/helpers"
	"github.com/jalapeno-api-gateway/subscription-service/kafka"
	"github.com/jalapeno-api-gateway/subscription-service/subscriptionservice"
	"github.com/jalapeno-api-gateway/subscription-service/pubsub"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"google.golang.org/grpc"
)

func main() {
	log.Print("Starting Subscription Service ...")
	arango.InitializeArangoDbAdapter(getDefaultArangoDbConfig())
	pubsub.InitializeTopics()
	kafka.StartEventConsumption()

	serverAddress := os.Getenv("APP_SERVER_ADDRESS")
	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", serverAddress, err)
	}

	grpcServer := grpc.NewServer()

	signals := helpers.WatchInterruptSignals()
	go func() {
		<-signals
		grpcServer.Stop()
	}()

	jagw.RegisterSubscriptionServiceServer(grpcServer, subscriptionservice.NewServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

func getDefaultArangoDbConfig() arango.ArangoDbConfig {
	return arango.ArangoDbConfig{
		Server: fmt.Sprintf("http://%s", os.Getenv("ARANGO_ADDRESS")),
		User: os.Getenv("ARANGO_DB_USER"),
		Password: os.Getenv("ARANGO_DB_PASSWORD"),
		DbName: os.Getenv("ARANGO_DB_NAME"),
	}
}