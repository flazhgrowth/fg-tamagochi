package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type Cache interface {
	// Get get data and unmarshal it to dest. Argument dest must be pointer to val
	Get(ctx context.Context, key string, dest any) error

	// SetWithExpire save data to cache with defined duration
	SetWithExpire(ctx context.Context, key string, data any, dur time.Duration) error

	// SetWithoutExpire save data to cache without expiration set
	SetWithoutExpire(ctx context.Context, key string, data any) error

	// SAdd set values to a tag
	SAdd(ctx context.Context, tag string, members ...any) error

	// SMembers gets all the values of a tag
	SMembers(ctx context.Context, tag string) ([]string, error)

	// SIsMember determine if a member is belong to a tag or not
	SIsMember(ctx context.Context, tag string, member any) (bool, error)

	// TTL gets the time to live (ttl) for a key
	TTL(ctx context.Context, key string) (time.Duration, error)

	// Del delete data based on its keys
	Del(ctx context.Context, keys ...string) error

	// InvalidateSMembers invalidate all the member of a tag. Delete tag once all members invalidated
	InvalidateSMembers(ctx context.Context, tag string) (err error)

	// PipelineDel delete all given keys using pipeline, instead of manually delete it one by one
	PipelineDel(ctx context.Context, keys []string) error
}

type CacheImpl struct {
	client *redis.Client
}

// Instantiate a new cache facade
func New(cfg CacheConfig) Cache {
	opt := &redis.Options{
		Addr:        cfg.Addr,
		Password:    cfg.Password,
		DB:          cfg.DB,
		DialTimeout: cfg.DialTimeout,
		ReadTimeout: cfg.ReadTimeout,
	}
	client := redis.NewClient(opt)
	log.Info().Msgf("ping redis error status: %s", client.Ping(context.Background()).Err().Error())

	return &CacheImpl{
		client: client,
	}
}

// Instantiate a new cache facade. Panic when the ping attempt return in error
func MustNew(cfg CacheConfig) Cache {
	opt := &redis.Options{
		Addr:        cfg.Addr,
		Password:    cfg.Password,
		DB:          cfg.DB,
		DialTimeout: cfg.DialTimeout,
		ReadTimeout: cfg.ReadTimeout,
	}
	client := redis.NewClient(opt)
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return &CacheImpl{
		client: client,
	}
}
