package notification

import (
	"context"
	"route256/libs/db/manager"
	"route256/notifications/internal/repository/schema"

	sq "github.com/Masterminds/squirrel"
)

const (
	columnUserID    = "user_id"
	columnOrderID   = "order_id"
	columnStatus    = "status"
	columnCreatedAt = "created_at"

	tableName = "notifications"
)

var (
	notificationAllColumns = []string{columnUserID, columnOrderID, columnStatus, columnCreatedAt}
)

type Repository struct {
	provider dbProvider
	cache    cache
	psql     sq.StatementBuilderType
}

type dbProvider interface {
	GetDB(ctx context.Context) manager.Querier
}

type cache interface {
	SetLst(ctx context.Context, key string, values []schema.Notification) error
	Range(ctx context.Context, key string) ([]schema.Notification, error)
	Delete(ctx context.Context, key string) error
}

func New(provider dbProvider, cache cache) *Repository {
	return &Repository{
		provider: provider,
		cache:    cache,
		psql:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
