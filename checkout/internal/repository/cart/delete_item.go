package cart

import (
	"context"
	"fmt"
	"route256/libs/mw/tracing"

	sq "github.com/Masterminds/squirrel"
	"github.com/opentracing/opentracing-go"
)

func (r *Repository) DeleteItem(
	ctx context.Context,
	userID int64,
	sku uint32,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cartRepo/DeleteItem")
	defer span.Finish()

	span.SetTag("user_id", userID)

	db := r.provider.GetDB(ctx)

	query := r.psql.Delete(columnCount).
		From(tableNameItems).
		Where(sq.And{
			sq.Eq{columnUserID: userID},
			sq.Eq{columnSKU: sku},
		})

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("build query for delete item: %w", err))
	}

	result, err := db.Exec(ctx, rawSQL, args...)
	if err != nil {
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("exec delete item: %w", err))
	}
	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
