package domain

type Payment struct {
	ID         int64
	UserID     int64
	OrderID    int64
	TotalPrice float32
}
