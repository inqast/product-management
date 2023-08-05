package order

import (
	"context"
	"fmt"
	"route256/libs/mw/tracing"
	"route256/loms/internal/domain/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/opentracing/opentracing-go"
)

func (r *Repository) SetOrderStatus(
	ctx context.Context,
	orderID int64,
	status model.OrderStatus,
) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "orderRepo/SetOrderStatus")
	defer span.Finish()

	span.SetTag("order_id", orderID)

	db := r.provider.GetDB(ctx)

	query := r.psql.Update(tableNameOrders).
		Where(sq.Eq{columnID: orderID}).
		Set(columnStatus, status)

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("build query for order status change: %w", err))
	}

	_, err = db.Exec(ctx, rawSQL, args...)
	if err != nil {
		return tracing.MarkSpanWithError(ctx, fmt.Errorf("exec order status change: %w", err))
	}

	return nil
}
