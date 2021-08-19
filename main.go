package main

import (
	"log"
	"net"
	"os"

	"gitlab.ost.ch/ins/jalapeno-api/push-service/events"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/helpers"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/kafka"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/proto/pushservice"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/subscribers"
	"google.golang.org/grpc"
)

func main() {
	log.Print("Starting Push Service ...")
	kafka.StartEventConsumption()
	go events.StartEventProcessing()
	go subscribers.StartSubscriptionService()
	
	serverAddress := os.Getenv("APP_SERVER_ADDRESS")
	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", serverAddress, err)
	}

	grpcServer := grpc.NewServer()
	
	signals := helpers.WatchInterruptSignals()
	go func(){
		<-signals
		grpcServer.Stop()
	}()

	pushservice.RegisterPushServiceServer(grpcServer, pushservice.NewServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
