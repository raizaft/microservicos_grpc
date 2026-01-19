package ports

import "github.com/raizaft/microservicos_grpc/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(order *domain.Order) error
}
