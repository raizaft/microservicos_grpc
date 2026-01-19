package db

import "github.com/raizaft/microservicos_grpc/order/internal/application/core/domain"

type MemoryDB struct {
	data []domain.Order
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{}
}

func (db *MemoryDB) Save(order *domain.Order) error {
	order.ID = int64(len(db.data) + 1)
	db.data = append(db.data, *order)
	return nil
}
