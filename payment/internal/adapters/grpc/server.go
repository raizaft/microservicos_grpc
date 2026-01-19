package grpc

import (
	"context"

	"github.com/raizaft/microservicos_grpc/payment/internal/adapters/grpc"
	"github.com/raizaft/microservicos_grpc/payment/internal/application/core/api"
	"github.com/raizaft/microservicos_grpc/payment/internal/application/core/domain"
	pb "github.com/raizaft/microservicos_grpc_proto/golang/payment"
)

type Server struct {
	pb.UnimplementedPaymentServer
	app *api.Application
}

func Register(server *grpc.Server, app *api.Application) {
	pb.RegisterPaymentServer(server, &Server{app: app})
}

func (s *Server) Create(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	payment := domain.Payment{
		UserID:     req.UserId,
		OrderID:    req.OrderId,
		TotalPrice: req.TotalPrice,
	}

	result, err := s.app.Charge(ctx, payment)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePaymentResponse{PaymentId: result.ID}, nil
}
