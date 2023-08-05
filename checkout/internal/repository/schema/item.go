package schema

type Item struct {
	UserID int64 `db:"user_id"`
	SKU    int64 `db:"sku"`
	Count  int32 `db:"count"`
}
