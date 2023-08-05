package stocks

import (
	"context"
	"errors"
	"route256/libs/db/manager"

	sq "github.com/Masterminds/squirrel"
)

const (
	columnWarehouseID = "warehouse_id"
	columnSKU         = "sku"
	columnCount       = "count"
	columnReserved    = "reserved"

	tableNameStocks = "stocks"
)

var (
	itemsAllColumns = []string{columnWarehouseID, columnSKU, columnCount, columnReserved}
	ErrNotFound     = errors.New("not found")
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
