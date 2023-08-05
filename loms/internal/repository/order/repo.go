package order

import (
	"context"
	"route256/libs/db/manager"

	sq "github.com/Masterminds/squirrel"
)

const (
	columnID     = "id"
	columnStatus = "status"
	columnUserID = "user_id"

	columnOrderID = "order_id"
	columnSKU     = "sku"
	columnCount   = "count"

	tableNameOrders = "orders"
	tableNameItems  = "items"
)

var (
	orderAllColumns = []string{columnID, columnStatus, columnUserID}
	itemsAllColumns = []string{columnOrderID, columnSKU, columnCount}
)

type Repository struct {
	provider dbProvider
	psql     sq.StatementBuilderType
}

type dbProvider interface {
	GetDB(ctx context.Context) manager.Querier
}

func New(provider dbProvider) *Repository {
	return &Repository{
		provider: provider,
		psql:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
