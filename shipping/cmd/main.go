package main

import (
	"log"
	"net"

	grpcadapter "github.com/raizaft/microservicos_grpc/shipping/internal/adapters/grpc"
	"github.com/raizaft/microservicos_grpc/shipping/internal/application/core/api"
	"google.golang.org/grpc"
)

func main() {
	app := api.NewApplication()

	server := grpc.NewServer()
	grpcadapter.Register(server, app)

	lis, _ := net.Listen("tcp", ":50053")
	log.Println("🚚 Shipping rodando na porta 50053")
	server.Serve(lis)
}
