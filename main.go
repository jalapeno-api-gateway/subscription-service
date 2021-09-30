package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jalapeno-api-gateway/jagw-core/arango"
	"github.com/jalapeno-api-gateway/subscription-service/helpers"
	"github.com/jalapeno-api-gateway/subscription-service/kafka"
	"github.com/jalapeno-api-gateway/subscription-service/proto/subscriptionservice"
	"github.com/jalapeno-api-gateway/subscription-service/pubsub"
	"google.golang.org/grpc"
)

const ARANGO_DB_PORT = 30852

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

	subscriptionservice.RegisterSubscriptionServiceServer(grpcServer, subscriptionservice.NewServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

func getDefaultArangoDbConfig() arango.ArangoDbConfig {
	return arango.ArangoDbConfig{
		Server: fmt.Sprintf("http://%s:%d", os.Getenv("JALAPENO_SERVER"), ARANGO_DB_PORT),
		User: os.Getenv("ARANGO_DB_USER"),
		Password: os.Getenv("ARANGO_DB_PASSWORD"),
		DbName: os.Getenv("ARANGO_DB_NAME"),
	}
}