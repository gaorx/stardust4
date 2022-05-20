package sdredis

import (
	"context"
	"github.com/gaorx/stardust4/sderr"
	"github.com/go-redis/redis/v8"
)

func WithContext(client redis.UniversalClient, ctx context.Context) redis.UniversalClient {
	if c1, ok := client.(*redis.Client); ok {
		return c1.WithContext(ctx)
	} else if c1, ok := client.(*redis.Ring); ok {
		return c1.WithContext(ctx)
	} else if c1, ok := client.(*redis.ClusterClient); ok {
		return c1.WithContext(ctx)
	} else {
		panic(sderr.New("sdredis with context error"))
	}
}

func ForEachShards(ctx context.Context, client redis.UniversalClient, action func(context.Context, *redis.Client) error) error {
	if c1, ok := client.(*redis.Client); ok {
		err := action(ctx, c1)
		return sderr.Wrap(err, "sdredis for each shard error")
	} else if c1, ok := client.(*redis.Ring); ok {
		err := c1.ForEachShard(ctx, action)
		return sderr.Wrap(err, "sdredis for each shard error (ring)")
	} else if c1, ok := client.(*redis.ClusterClient); ok {
		err := c1.ForEachShard(ctx, action)
		return sderr.Wrap(err, "sdredis for each shard error (cluster)")
	} else {
		panic(sderr.New("sdredis for each shards error"))
	}
}
