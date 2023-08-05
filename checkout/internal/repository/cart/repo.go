package cart

import (
	"context"
	"errors"
	"route256/libs/db/manager"

	sq "github.com/Masterminds/squirrel"
)

const (
	tableNameItems = "items"
	columnUserID   = "user_id"
	columnSKU      = "sku"
	columnCount    = "count"
)

var (
	itemsAllColumns = []string{columnUserID, columnSKU, columnCount}
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
