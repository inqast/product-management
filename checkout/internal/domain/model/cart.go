package model

type Cart struct {
	Items      []*CartItem
	TotalPrice uint32
}

type CartItem struct {
	SKU   uint32
	Count uint16
	Name  string
	Price uint32
}
