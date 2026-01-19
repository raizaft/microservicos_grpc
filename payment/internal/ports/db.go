package ports

import (
	"context"

	"github.com/raizaft/microservicos_grpc/payment/internal/application/core/domain"
)

type DBPort interface {
	Save(ctx context.Context, payment *domain.Payment) error
}
