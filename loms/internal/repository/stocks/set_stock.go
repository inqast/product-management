package stocks

import (
	"context"
	"fmt"
	"route256/libs/mw/tracing"

	"github.com/opentracing/opentracing-go"
)

func (r *Repository) SetStock(
	ctx context.Context,
	sku uint32,
	warehouseID int64,
	count uint64,
	reserved uint64,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "stocksRepo/SetStock")
	defer span.Finish()

	span.SetTag("sku", sku)

	db := r.provider.GetDB(ctx)

	query := `
UPDATE stocks
SET "count"=$1, "reserved"=$2
WHERE "warehouse_id" = $3 AND "sku" = $4;
`

	_, err := db.Exec(ctx, query, count, reserved, warehouseID, sku)
	if err != nil {
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("exec stock change: %w", err))
	}

	return nil
}
