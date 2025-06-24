package cache

import (
	"context"
	"encoding/json"
	"time"
)

// Get get data and unmarshal it to dest. Argument dest must be pointer to val
func (impl *CacheImpl) Get(ctx context.Context, key string, dest any) error {
	raw, err := impl.client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, dest)
}

// SetWithExpire save data to cache with defined duration
func (impl *CacheImpl) SetWithExpire(ctx context.Context, key string, data any, dur time.Duration) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return impl.client.Set(ctx, key, dataBytes, dur).Err()
}

// SetWithoutExpire save data to cache without expiration set
func (impl *CacheImpl) SetWithoutExpire(ctx context.Context, key string, data any) error {
	return impl.SetWithExpire(ctx, key, data, 0)
}

// SAdd set values to a tag
func (impl *CacheImpl) SAdd(ctx context.Context, tag string, members ...any) error {
	return impl.client.SAdd(ctx, tag, members...).Err()
}

// SMembers gets all the values of a tag
func (impl *CacheImpl) SMembers(ctx context.Context, tag string) ([]string, error) {
	return impl.client.SMembers(ctx, tag).Result()
}

// SIsMember determine if a member is belong to a tag or not
func (impl *CacheImpl) SIsMember(ctx context.Context, tag string, member any) (bool, error) {
	return impl.client.SIsMember(ctx, tag, member).Result()
}

// TTL gets the time to live (ttl) for a key
func (impl *CacheImpl) TTL(ctx context.Context, key string) (time.Duration, error) {
	return impl.client.TTL(ctx, key).Result()
}

// Del delete data based on its keys
func (impl *CacheImpl) Del(ctx context.Context, keys ...string) error {
	return impl.client.Del(ctx, keys...).Err()
}

// InvalidateSMembers invalidate all the member of a tag. Delete tag once all members invalidated
func (impl *CacheImpl) InvalidateSMembers(ctx context.Context, tag string) (err error) {
	keys, err := impl.SMembers(ctx, tag)
	if err != nil {
		return err
	}
	if len(keys) < 1 {
		return nil
	}
	keys = append(keys, tag)

	if err = impl.PipelineDel(ctx, keys); err != nil {
		return err
	}

	return nil
}

// PipelineDel delete all given keys using pipeline, instead of manually delete it one by one. Use pipelinetx by default
func (impl *CacheImpl) PipelineDel(ctx context.Context, keys []string) error {
	pipeline := impl.client.TxPipeline()

	if err := pipeline.Del(ctx, keys...).Err(); err != nil {
		return err
	}

	if _, err := pipeline.Exec(ctx); err != nil {
		return err
	}

	return nil
}
