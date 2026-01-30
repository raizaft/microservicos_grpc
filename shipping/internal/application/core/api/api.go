package api

import "github.com/raizaft/microservicos_grpc/shipping/internal/application/core/domain"

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Calculate(s domain.Shipping) int32 {
	return s.DeliveryDays()
}
