package ports

import "github.com/raizaft/microservicos_grpc/order/internal/application/core/domain"

type DBPort interface {
	Save(order *domain.Order) error
}
