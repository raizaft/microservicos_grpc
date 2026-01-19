package main

import (
	"log"
	"net"

	"github.com/raizaft/microservicos_grpc/order/internal/adapters/db"
	grpcadapter "github.com/raizaft/microservicos_grpc/order/internal/adapters/grpc"
	payment_adapter "github.com/raizaft/microservicos_grpc/order/internal/adapters/payment"
	"github.com/raizaft/microservicos_grpc/order/internal/application/core/api"
	"google.golang.org/grpc"
)

func main() {
	db := db.NewMemoryDB()

	payment, err := payment_adapter.NewAdapter("localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	app := api.NewApplication(db, payment)

	server := grpc.NewServer()
	grpcadapter.Register(server, app)

	listener, _ := net.Listen("tcp", ":50052")
	log.Println("📦 Order rodando na porta 50052")
	server.Serve(listener)
}
