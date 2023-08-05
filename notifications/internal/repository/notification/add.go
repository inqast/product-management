package notification

import (
	"context"
	"fmt"
	"route256/libs/mw/tracing"
	"time"

	"github.com/opentracing/opentracing-go"
)

func (r *Repository) Add(
	ctx context.Context,
	userID int64,
	orderID int64,
	status string,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo/Add")
	defer span.Finish()

	span.SetTag("sku", userID)

	db := r.provider.GetDB(ctx)

	query := r.psql.Insert(tableName).Columns(notificationAllColumns...).
		Values(userID, orderID, status, time.Now())

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("build query for insert item: %w", err))
	}

	_, err = db.Exec(ctx, rawSQL, args...)
	if err != nil {
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("exec insert item: %w", err))
	}

	_ = r.cache.Delete(ctx, fmt.Sprint(userID))

	return nil
}
