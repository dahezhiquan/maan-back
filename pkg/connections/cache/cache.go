package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key, value string, expire time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	HSet(ctx context.Context, key string, field string, value string) error
	HKeys(ctx context.Context, key string) ([]string, error)
	ZAdd(ctx context.Context, key string, value string, score float64) error
	ZRem(ctx context.Context, key string, member ...string) error
	ZRange(ctx context.Context, key string, start, end int64) ([]string, error)
	ZRevRange(ctx context.Context, key string, start, end int64) ([]string, error)
	SetBit(ctx context.Context, key string, offset int64, value int) error
	GetBit(ctx context.Context, key string, memId int64) (int64, error)
	BitCount(ctx context.Context, key string) (int64, error)
	PFAdd(ctx context.Context, key string, value interface{}) error
	PFCount(ctx context.Context, key string) (int64, error)
	Eval(ctx context.Context, script string, key []string, value ...interface{}) (interface{}, error)
	Pipeline() redis.Pipeliner
	ErrNil() error
}
