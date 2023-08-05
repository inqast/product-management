package schema

type Order struct {
	ID     int64 `db:"id"`
	Status int32 `db:"status"`
	UserID int64 `db:"user_id"`
}
