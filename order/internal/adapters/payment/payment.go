package payment

import (
	"context"

	"github.com/raizaft/microservicos_grpc/order/internal/application/core/domain"
	pb "github.com/raizaft/microservicos_grpc_proto/golang/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	client pb.PaymentClient
}

func NewAdapter(url string) (*Adapter, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Adapter{client: pb.NewPaymentClient(conn)}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.client.Create(context.Background(), &pb.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
	return err
}
