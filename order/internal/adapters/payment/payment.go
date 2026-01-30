package payment

import (
	"context"
	"log"
	"time"

	"github.com/raizaft/microservicos_grpc/order/internal/application/core/domain"
	pb "github.com/raizaft/microservicos_grpc_proto/golang/payment"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	client pb.PaymentClient
}

func NewAdapter(url string) (*Adapter, error) {

	// ---------- RETRY CONFIG ----------
	var opts []grpc.DialOption

	opts = append(opts,
		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithCodes(
					codes.Unavailable,
					codes.ResourceExhausted,
				),
				grpc_retry.WithMax(5),
				grpc_retry.WithBackoff(grpc_retry.BackoffLinear(1*time.Second)),
			),
		),
	)

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(url, opts...)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		client: pb.NewPaymentClient(conn),
	}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {

	// ---------- TIMEOUT ----------
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := a.client.Create(ctx, &pb.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})

	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.DeadlineExceeded {
			log.Println("⏰ Timeout ao chamar Payment")
		}
		return err
	}

	return nil
}
