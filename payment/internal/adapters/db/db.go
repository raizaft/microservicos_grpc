package db

import (
	"context"

	"github.com/raizaft/microservicos_grpc/payment/internal/application/core/domain"
)

type MemoryDB struct {
	data []domain.Payment
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{}
}

func (db *MemoryDB) Save(ctx context.Context, payment *domain.Payment) error {
	payment.ID = int64(len(db.data) + 1)
	db.data = append(db.data, *payment)
	return nil
}
