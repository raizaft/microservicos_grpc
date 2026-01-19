package domain

import "time"

type OrderItem struct {
	ProductCode string
	UnitPrice   float32
	Quantity    int32
}

type Order struct {
	ID         int64
	CustomerID int64
	Status     string
	OrderItems []OrderItem
	CreatedAt  int64
}

func NewOrder(customerID int64, items []OrderItem) Order {
	return Order{
		CustomerID: customerID,
		OrderItems: items,
		Status:     "Pending",
		CreatedAt:  time.Now().Unix(),
	}
}

func (o *Order) TotalItems() int32 {
	var total int32
	for _, item := range o.OrderItems {
		total += item.Quantity
	}
	return total
}

func (o *Order) TotalPrice() float32 {
	var total float32
	for _, item := range o.OrderItems {
		total += item.UnitPrice * float32(item.Quantity)
	}
	return total
}
