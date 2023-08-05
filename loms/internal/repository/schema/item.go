package schema

type Item struct {
	OrderID int64 `db:"order_id"`
	SKU     int64 `db:"sku"`
	Count   int32 `db:"count"`
}
