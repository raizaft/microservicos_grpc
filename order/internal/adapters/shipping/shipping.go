package shipping

import (
	"context"

	"github.com/raizaft/microservicos_grpc/order/internal/application/core/domain"
	pb "github.com/raizaft/microservicos_grpc_proto/golang/shipping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	client pb.ShippingClient
}

func NewAdapter(url string) (*Adapter, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Adapter{client: pb.NewShippingClient(conn)}, nil
}

func (a *Adapter) Calculate(order *domain.Order) error {
	var items []*pb.ShippingItem
	for _, it := range order.OrderItems {
		items = append(items, &pb.ShippingItem{
			ProductCode: it.ProductCode,
			Quantity:    it.Quantity,
		})
	}

	_, err := a.client.Create(context.Background(), &pb.CreateShippingRequest{
		OrderId: order.ID,
		Items:   items,
	})
	return err
}
