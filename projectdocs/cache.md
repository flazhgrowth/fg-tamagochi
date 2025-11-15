# Cache
Cache by default use redis. On app initialization, the config is only a boolean. Set to true if you want to enable cache on your app. You can define the necessary value for initializing vault in `vault.json` file.

## Secret value in vault.json
Take a look at vault.json below:
```
{
    "database": {
        "driver": "postgres",
        "reader_dsn": "",
        "writer_dsn": ""
    },
    "cache": {
        "host": "localhost",
        "port": "6379",
        "db": 0,
        "password": ""
    }
}
```
Fill out the cache object when you set `IsUseCache` true on app initialization.

## Methods
Take a look at Cache interface definition below:
```
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
```
