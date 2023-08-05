package cart

import (
	"context"
	"fmt"
	"route256/libs/mw/tracing"

	"github.com/opentracing/opentracing-go"
)

func (r *Repository) SetItemCount(
	ctx context.Context,
	userID int64,
	sku uint32,
	count uint16,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cartRepo/SetItemCount")
	defer span.Finish()

	span.SetTag("user_id", userID)

	db := r.provider.GetDB(ctx)

	query := `
INSERT INTO items("user_id", "sku", "count") VALUES 
    ($1, $2, $3)
ON CONFLICT("user_id", "sku") DO UPDATE 
	SET count=excluded.count;
`

	_, err := db.Exec(ctx, query, userID, sku, count)
	if err != nil {
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("exec insert cart: %w", err))
	}

	return nil
}
