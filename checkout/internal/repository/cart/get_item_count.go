package cart

import (
	"context"
	"fmt"
	"route256/libs/mw/tracing"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/opentracing/opentracing-go"
)

func (r *Repository) GetItemCount(
	ctx context.Context,
	userID int64,
	SKU uint32,
) (uint16, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cartRepo/GetItemCount")
	defer span.Finish()

	span.SetTag("user_id", userID)

	db := r.provider.GetDB(ctx)

	query := r.psql.Select("count").
		From(tableNameItems).
		Where(sq.And{
			sq.Eq{columnUserID: userID},
			sq.Eq{columnSKU: SKU},
		})

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return 0, tracing.MarkSpanWithError(ctx, fmt.Errorf("build query for get item count: %w", err))
	}

	var count uint16
	err = pgxscan.Get(ctx, db, &count, rawSQL, args...)
	if err != nil && !pgxscan.NotFound(err) {
		return 0, tracing.MarkSpanWithError(ctx, fmt.Errorf("exec query for get item count: %w", err))
	}

	return count, nil
}
