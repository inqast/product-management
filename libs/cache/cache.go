package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/redis/go-redis/v9"
	"time"
)

var ErrNotFound = errors.New("not found in cache")

type Cache[Value any] struct {
	realization *redis.Client
	prefix      string
	ttl         time.Duration
}

func New[Value any](keyPrefix string, cache *redis.Client, ttl time.Duration) *Cache[Value] {
	return &Cache[Value]{
		realization: cache,
		prefix:      keyPrefix,
		ttl:         ttl,
	}
}

func (c *Cache[Value]) SetLst(ctx context.Context, key string, values []Value) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cache/SetLst")
	defer span.Finish()

	if len(values) == 0 {
		return errors.New("empty list")
	}

	err := c.Delete(ctx, key)
	if err != nil {
		return fmt.Errorf("failed to clear list: %w", err)
	}

	for _, item := range values {
		err = c.Push(ctx, key, item)
		if err != nil {
			c.Delete(ctx, key)
			return fmt.Errorf("failed to add list: %w", err)
		}
	}

	return nil
}

func (c *Cache[Value]) Push(ctx context.Context, key string, value Value) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cache/Push")
	defer span.Finish()

	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.realization.RPush(ctx, c.prepareKey(key), bytes).Err()
}

func (c *Cache[Value]) Set(ctx context.Context, key string, value Value) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cache/Set")
	defer span.Finish()

	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.realization.Set(ctx, c.prepareKey(key), bytes, c.ttl).Err()
}

func (c *Cache[Value]) Delete(ctx context.Context, key string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cache/Delete")
	defer span.Finish()

	return c.realization.Del(ctx, c.prepareKey(key)).Err()
}

func (c *Cache[Value]) Range(ctx context.Context, key string) ([]Value, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cache/Range")
	defer span.Finish()

	list, err := c.realization.LRange(ctx, c.prepareKey(key), 0, -1).Result()
	if errors.Is(err, redis.Nil) || len(list) == 0 {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	result := make([]Value, len(list))

	for i, item := range list {
		var curItem Value
		err = json.Unmarshal([]byte(item), &curItem)
		if err != nil {
			return nil, err
		}

		result[i] = curItem
	}

	return result, nil
}

func (c *Cache[Output]) prepareKey(key string) string {
	return fmt.Sprintf("%s_%s", c.prefix, key)
}
