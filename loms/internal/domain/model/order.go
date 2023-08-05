package model

type Order struct {
	ID     int64
	Status OrderStatus
	UserID int64
	Items  []*OrderItem
}

type OrderItem struct {
	SKU   uint32
	Count uint16
}

type OrderStatus int

const (
	OrderNew OrderStatus = iota
	OrderAwaitingPayment
	OrderFailed
	OrderPayed
	OrderCancelled
)
