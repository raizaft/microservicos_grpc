package grpcadapter

import (
	"context"

	"github.com/raizaft/microservicos_grpc/shipping/internal/application/core/api"
	"github.com/raizaft/microservicos_grpc/shipping/internal/application/core/domain"
	pb "github.com/raizaft/microservicos_grpc_proto/golang/shipping"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedShippingServer
	app *api.Application
}

func Register(s *grpc.Server, app *api.Application) {
	pb.RegisterShippingServer(s, &Server{app: app})
}

func (s *Server) Create(ctx context.Context, req *pb.CreateShippingRequest) (*pb.CreateShippingResponse, error) {
	var items []domain.Item
	for _, it := range req.Items {
		items = append(items, domain.Item{
			ProductCode: it.ProductCode,
			Quantity:    it.Quantity,
		})
	}

	ship := domain.Shipping{
		OrderID: req.OrderId,
		Items:   items,
	}

	days := s.app.Calculate(ship)

	return &pb.CreateShippingResponse{
		DeliveryDays: days,
	}, nil
}
