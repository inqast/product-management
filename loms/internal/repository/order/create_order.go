package order

import (
	"context"
	"fmt"
	"route256/libs/mw/tracing"
	"route256/loms/internal/domain/model"

	"github.com/opentracing/opentracing-go"
)

func (r *Repository) CreateOrder(
	ctx context.Context,
	userID int64,
) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "orderRepo/CreateOrder")
	defer span.Finish()

	span.SetTag("user_id", userID)

	db := r.provider.GetDB(ctx)

	query := r.psql.Insert(tableNameOrders).Columns(columnStatus, columnUserID).
		Values(model.OrderNew, userID).
		Suffix("RETURNING id")

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return 0, tracing.MarkSpanWithError(ctx, fmt.Errorf("build query for create order: %w", err))
	}

	var res int64
	err = db.QueryRow(ctx, rawSQL, args...).Scan(&res)
	if err != nil {
		return 0, tracing.MarkSpanWithError(ctx, fmt.Errorf("exec create order: %w", err))
	}

	return res, nil

}
