package main

import (
	"log"
	"net"

	"github.com/raizaft/microservicos_grpc/payment/internal/adapters/db"
	grpcadapter "github.com/raizaft/microservicos_grpc/payment/internal/adapters/grpc"
	"github.com/raizaft/microservicos_grpc/payment/internal/application/core/api"
	"google.golang.org/grpc"
)

func main() {
	db := db.NewMemoryDB()
	app := api.NewApplication(db)

	server := grpc.NewServer()
	grpcadapter.Register(server, app)

	listener, _ := net.Listen("tcp", ":50051")
	log.Println("💳 Payment rodando na porta 50051")
	server.Serve(listener)
}
