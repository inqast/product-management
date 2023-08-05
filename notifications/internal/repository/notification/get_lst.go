package notification

import (
	"context"
	"fmt"
	log "route256/libs/logger"
	"route256/libs/mw/tracing"
	"route256/notifications/internal/domain/model"
	"route256/notifications/internal/repository/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/opentracing/opentracing-go"
)

func (r *Repository) GetLst(
	ctx context.Context,
	userID int64,
) ([]*model.Notification, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo/GetLst")
	defer span.Finish()

	span.SetTag("order_id", userID)

	if items, err := r.cache.Range(ctx, fmt.Sprint(userID)); err == nil {
		return mapDomainNotifications(items), nil
	} else {
		log.Warn(err.Error())
	}

	db := r.provider.GetDB(ctx)

	itemsQuery := r.psql.Select(notificationAllColumns...).
		From(tableName).
		Where(sq.Eq{columnUserID: userID}).
		OrderBy(columnCreatedAt)

	rawItemsSQL, args, err := itemsQuery.ToSql()
	if err != nil {
		return nil, tracing.MarkSpanWithError(ctx, fmt.Errorf("build query for get item count: %w", err))
	}

	var itemsSQL []schema.Notification

	err = pgxscan.Select(ctx, db, &itemsSQL, rawItemsSQL, args...)
	if err != nil {
		return nil, tracing.MarkSpanWithError(ctx, fmt.Errorf("exec query get item stocks: %w", err))
	}

	setErr := r.cache.SetLst(ctx, fmt.Sprint(userID), itemsSQL)
	if setErr != nil {
		log.Warnf("error add list to cache: %s", setErr.Error())
	}

	return mapDomainNotifications(itemsSQL), nil
}
