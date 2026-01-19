package grpc

import (
	"context"

	"github.com/raizaft/microservicos_grpc/order/internal/adapters/grpc"
	"github.com/raizaft/microservicos_grpc/order/internal/application/core/api"
	"github.com/raizaft/microservicos_grpc/order/internal/application/core/domain"
	pb "github.com/raizaft/microservicos_grpc_proto/golang/order"
)

type Server struct {
	pb.UnimplementedOrderServer
	app *api.Application
}

func Register(server *grpc.Server, app *api.Application) {
	pb.RegisterOrderServer(server, &Server{app: app})
}

func (s *Server) Create(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	var items []domain.OrderItem

	for _, item := range req.OrderItems {
		items = append(items, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	order := domain.NewOrder(int64(req.CostumerId), items)

	result, err := s.app.PlaceOrder(order)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{OrderId: int32(result.ID)}, nil
}
