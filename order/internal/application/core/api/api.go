package api

import (
	"errors"

	"github.com/raizaft/microservicos_grpc/order/internal/application/core/domain"
	"github.com/raizaft/microservicos_grpc/order/internal/ports"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{db: db, payment: payment}
}

func (a *Application) PlaceOrder(order domain.Order) (domain.Order, error) {

	// Parte 3: regra de negócio
	if order.TotalItems() > 50 {
		return domain.Order{}, errors.New("order with more than 50 items")
	}

	if err := a.db.Save(&order); err != nil {
		order.Status = "Canceled"
		return order, err
	}

	if err := a.payment.Charge(&order); err != nil {
		order.Status = "Canceled"
		return order, err
	}

	order.Status = "Paid"
	return order, nil
}
