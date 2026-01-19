package ports

import "github.com/raizaft/microservicos_grpc/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
