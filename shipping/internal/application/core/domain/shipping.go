package domain

type Item struct {
	ProductCode string
	Quantity    int32
}

type Shipping struct {
	OrderID int64
	Items   []Item
}

func (s *Shipping) TotalUnits() int32 {
	var total int32
	for _, i := range s.Items {
		total += i.Quantity
	}
	return total
}

func (s *Shipping) DeliveryDays() int32 {
	total := s.TotalUnits()
	days := int32(1)
	days += total / 5
	return days
}
