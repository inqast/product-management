package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/domain/model"
	"route256/checkout/internal/repository/schema"
	"route256/libs/mw/tracing"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/opentracing/opentracing-go"
)

func (r *Repository) GetItems(
	ctx context.Context,
	userID int64,
) ([]*model.CartItem, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cartRepo/GetItems")
	defer span.Finish()

	span.SetTag("user_id", userID)

	db := r.provider.GetDB(ctx)

	query := r.psql.Select(itemsAllColumns...).
		From(tableNameItems).
		Where(sq.Eq{columnUserID: userID})

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return nil, tracing.MarkSpanWithError(ctx, fmt.Errorf("build query for get items: %w", err))
	}

	var resultSQL []schema.Item

	err = pgxscan.Select(ctx, db, &resultSQL, rawSQL, args...)
	if err != nil {
		return nil, tracing.MarkSpanWithError(ctx, fmt.Errorf("exec query get cart items: %w", err))
	}

	return mapDomainModels(resultSQL), nil
}
