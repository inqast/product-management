package stocks

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

func (r *Repository) GetStocks(
	ctx context.Context,
	sku uint32,
) ([]*model.Stock, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "stocksRepo/GetStocks")
	defer span.Finish()

	span.SetTag("sku", sku)

	db := r.provider.GetDB(ctx)

	query := r.psql.Select(itemsAllColumns...).
		From(tableNameStocks).
		Where(sq.Eq{columnSKU: sku})

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return nil, tracing.MarkSpanWithError(ctx, fmt.Errorf("build query for get stock: %w", err))
	}

	var resultSQL []schema.StockItem

	err = pgxscan.Select(ctx, db, &resultSQL, rawSQL, args...)
	if err != nil {
		return nil, tracing.MarkSpanWithError(ctx, fmt.Errorf("exec query get item stocks: %w", err))
	}

	return mapDomainStocks(resultSQL), nil
}
