package order

import (
	"context"
	"fmt"
	"route256/libs/mw/tracing"
	"route256/loms/internal/domain/model"
	"route256/loms/internal/repository/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/opentracing/opentracing-go"
)

func (r *Repository) GetOrder(
	ctx context.Context,
	orderID int64,
) (*model.Order, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "orderRepo/GetOrder")
	defer span.Finish()

	span.SetTag("order_id", orderID)

	db := r.provider.GetDB(ctx)

	orderQuery := r.psql.Select(orderAllColumns...).
		From(tableNameOrders).
		Where(sq.Eq{columnID: orderID})

	rawOrderSQL, args, err := orderQuery.ToSql()
	if err != nil {
		return nil, tracing.MarkSpanWithError(ctx, fmt.Errorf("build query for get order: %w", err))
	}

	var orderSQL schema.Order

	err = pgxscan.Get(ctx, db, &orderSQL, rawOrderSQL, args...)
	if err != nil {
		return nil, tracing.MarkSpanWithError(ctx, fmt.Errorf("exec query for get order: %w", err))
	}

	itemsQuery := r.psql.Select(itemsAllColumns...).
		From(tableNameItems).
		Where(sq.Eq{columnOrderID: orderID})

	rawItemsSQL, args, err := itemsQuery.ToSql()
	if err != nil {
		return nil, tracing.MarkSpanWithError(ctx, fmt.Errorf("build query for get item count: %w", err))
	}

	var itemsSQL []schema.Item

	err = pgxscan.Select(ctx, db, &itemsSQL, rawItemsSQL, args...)
	if err != nil {
		return nil, tracing.MarkSpanWithError(ctx, fmt.Errorf("exec query get item stocks: %w", err))
	}

	return mapDomainOrder(orderSQL, itemsSQL), nil
}
