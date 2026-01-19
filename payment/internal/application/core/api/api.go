package api

import (
	"context"

	"github.com/raizaft/microservicos_grpc/payment/internal/application/core/domain"
	"github.com/raizaft/microservicos_grpc/payment/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a *Application) Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	if payment.TotalPrice > 1000 {
		return domain.Payment{}, status.Error(codes.InvalidArgument, "payment over 1000")
	}
	err := a.db.Save(ctx, &payment)
	return payment, err
}
