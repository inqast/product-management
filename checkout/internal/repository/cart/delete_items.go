package cart

import (
	"context"
	"fmt"
	"route256/libs/mw/tracing"

	sq "github.com/Masterminds/squirrel"
	"github.com/opentracing/opentracing-go"
)

func (r *Repository) DeleteItems(
	ctx context.Context,
	userID int64,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cartRepo/DeleteItems")
	defer span.Finish()

	span.SetTag("user_id", userID)

	db := r.provider.GetDB(ctx)

	query := r.psql.Delete(columnCount).
		From(tableNameItems).
		Where(sq.Eq{columnUserID: userID})

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("build query for delete items: %w", err))
	}

	result, err := db.Exec(ctx, rawSQL, args...)
	if err != nil {
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("exec delete items: %w", err))
	}
	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
