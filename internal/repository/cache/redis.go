package cache

import (
	"context"
)

func (c *client[T]) Get(ctx context.Context, key string) (result T, err error) {
	err = c.Client.Get(ctx, key).Scan(&result)
	return result, err
}
func (c *client[T]) Del(ctx context.Context, keys ...string) error {
	return c.Client.Del(ctx, keys...).Err()
}
func (c *client[T]) Set(ctx context.Context, key string, obj T) error {
	return c.Client.Set(ctx, key, &obj, c.Config.Expire).Err()
}
