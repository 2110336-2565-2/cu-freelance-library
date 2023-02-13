package gosdk

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{client: client}
}

func (r *RedisRepository) SaveCache(key string, value interface{}, ttl int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	v, err := json.Marshal(value)
	if err != nil {
		return
	}

	return r.client.Set(ctx, key, v, time.Duration(ttl)*time.Second).Err()
}

func (r *RedisRepository) SaveHashCache(key string, field string, value string, ttl int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := r.client.HSet(ctx, key, field, value).Err(); err != nil {
		return err
	}

	return r.client.Expire(ctx, key, time.Duration(ttl)*time.Second).Err()
}

func (r *RedisRepository) SaveAllHashCache(key string, value map[string]string, ttl int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := r.client.HSet(ctx, key, value).Err(); err != nil {
		return err
	}

	return r.client.Expire(ctx, key, time.Duration(ttl)*time.Second).Err()
}

func (r *RedisRepository) GetHashCache(key string, field string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return r.client.HGet(ctx, key, field).Result()
}

func (r *RedisRepository) GetAllHashCache(key string) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return r.client.HGetAll(ctx, key).Result()
}

func (r *RedisRepository) GetCache(key string, value interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	v, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return
	}

	return json.Unmarshal([]byte(v), value)
}

func (r *RedisRepository) RemoveCache(key string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = r.client.Del(ctx, key).Result()
	return err
}

func (r *RedisRepository) SetExpire(key string, ttl int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return r.client.Expire(ctx, key, time.Duration(ttl)*time.Second).Err()
}
