package service

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	Client *redis.Client
}

func InitRedisService() *RedisService {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	return &RedisService{
		Client: rdb,
	}
}

func (s *RedisService) Add(ctx context.Context, key string, value interface{}) (string, error) {
	res, err := s.Client.Set(ctx, key, value, 0).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

func (s *RedisService) Get(ctx context.Context, key string) (string, error) {
	res, err := s.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

func (s *RedisService) Test() string {
	return "foodbar"
}
