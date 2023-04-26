package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"secureQR/config"
	"time"
)

var Rc *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	rdb := redis.NewClient(config.Conf.ReadRedisOptions())
	Rc = &RedisCache{rdb: rdb}
}

func (rc *RedisCache) Set(ctx context.Context, key, value string, expire time.Duration) error {
	err := rc.rdb.Set(ctx, key, value, expire).Err()
	return err
}

func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := rc.rdb.Get(ctx, key).Result()
	return result, err
}

func (rc *RedisCache) HSet(ctx context.Context, key string, field string, value string) error {
	err := rc.rdb.HSet(ctx, key, field, value).Err()
	return err
}

func (rc *RedisCache) HKeys(ctx context.Context, key string) ([]string, error) {
	result, err := rc.rdb.HKeys(ctx, key).Result()
	return result, err
}

func (rc *RedisCache) ZAdd(ctx context.Context, key string, value string, score float64) error {
	err := rc.rdb.ZAdd(ctx, key, &redis.Z{Member: value, Score: score}).Err()
	return err
}

func (rc *RedisCache) ZRem(ctx context.Context, key string, member ...string) error {
	err := rc.rdb.ZRem(ctx, key, member).Err()
	return err
}

func (rc *RedisCache) ZRange(ctx context.Context, key string, start, end int64) ([]string, error) {
	result, err := rc.rdb.ZRange(ctx, key, start, end).Result()
	return result, err
}
func (rc *RedisCache) ZRevRange(ctx context.Context, key string, start, end int64) ([]string, error) {
	result, err := rc.rdb.ZRevRange(ctx, key, start, end).Result()
	return result, err
}

func (rc *RedisCache) Delete(ctx context.Context, keys ...string) error {
	err := rc.rdb.Del(ctx, keys...).Err()
	return err
}

func (rc *RedisCache) PFAdd(ctx context.Context, key string, value interface{}) error {
	err := rc.rdb.PFAdd(ctx, key, value).Err()
	return err
}

func (rc *RedisCache) PFCount(ctx context.Context, key string) (int64, error) {
	result, err := rc.rdb.PFCount(ctx, key).Result()
	return result, err
}

func (rc *RedisCache) SetBit(ctx context.Context, key string, offset int64, value int) error {
	_, err := rc.rdb.SetBit(ctx, key, offset, value).Result()
	return err
}

func (rc *RedisCache) GetBit(ctx context.Context, key string, memId int64) (int64, error) {
	result, err := rc.rdb.GetBit(ctx, key, memId).Result()
	return result, err
}

func (rc *RedisCache) BitCount(ctx context.Context, key string) (int64, error) {
	result, err := rc.rdb.BitCount(ctx, key, &redis.BitCount{Start: 0, End: -1}).Result()
	return result, err
}

func (rc *RedisCache) Eval(ctx context.Context, script string, keys []string, values ...interface{}) (interface{}, error) {
	result, err := rc.rdb.Eval(ctx, script, keys, values).Result()
	return result, err
}

func (rc *RedisCache) Pipeline() redis.Pipeliner {
	return rc.rdb.Pipeline()
}

func (rc *RedisCache) ErrNil() error {
	return redis.Nil
}
