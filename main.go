package main

import (
	"log"
	"net"
	"os"

	"github.com/Jalapeno-API-Gateway/subscription-service/helpers"
	"github.com/Jalapeno-API-Gateway/subscription-service/kafka"
	"github.com/Jalapeno-API-Gateway/subscription-service/proto/subscriptionservice"
	"github.com/Jalapeno-API-Gateway/subscription-service/pubsub"
	"google.golang.org/grpc"
)

func main() {
	log.Print("Starting Subscription Service ...")
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
