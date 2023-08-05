package schema

type StockItem struct {
	WarehouseID int64 `db:"warehouse_id"`
	SKU         int64 `db:"sku"`
	Count       int32 `db:"count"`
	Reserved    int32 `db:"reserved"`
}
