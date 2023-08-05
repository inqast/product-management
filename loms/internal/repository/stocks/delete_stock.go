package stocks

import (
	"context"
	"fmt"
	"route256/libs/mw/tracing"

	sq "github.com/Masterminds/squirrel"
	"github.com/opentracing/opentracing-go"
)

func (r *Repository) DeleteStock(
	ctx context.Context,
	sku uint32,
	warehouseID int64,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "stocksRepo/DeleteStock")
	defer span.Finish()

	span.SetTag("sku", sku)

	db := r.provider.GetDB(ctx)

	query := r.psql.Delete(columnCount).
		From(tableNameStocks).
		Where(sq.And{
			sq.Eq{columnSKU: sku},
			sq.Eq{columnWarehouseID: warehouseID},
		})

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("build query for delete stock: %w", err))
	}

	result, err := db.Exec(ctx, rawSQL, args...)
	if err != nil {
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("exec delete stock: %w", err))
	}
	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
