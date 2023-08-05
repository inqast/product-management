package schema

import (
	"time"
)

type Notification struct {
	UserID    int64     `db:"user_id" json:"user_id"`
	OrderID   int64     `db:"order_id" json:"order_id"`
	Status    string    `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
