package order

import (
	"context"
	"fmt"
	"route256/libs/mw/tracing"

	"github.com/opentracing/opentracing-go"
)

func (r *Repository) AddOrderItem(
	ctx context.Context,
	orderID int64,
	sku uint32,
	count uint16,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "orderRepo/AddOrderItem")
	defer span.Finish()

	span.SetTag("sku", sku)

	db := r.provider.GetDB(ctx)

	query := r.psql.Insert(tableNameItems).Columns(itemsAllColumns...).
		Values(orderID, sku, count)

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("build query for insert item: %w", err))
	}

	_, err = db.Exec(ctx, rawSQL, args...)
	if err != nil {
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("exec insert item: %w", err))
	}

	return nil
}
