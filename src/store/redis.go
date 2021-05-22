package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type MetricStore struct {
	redis    *redis.Client
	cacheKey string
	duration time.Duration
}

func New(instance *redis.Client, cacheKey string, duration time.Duration) MetricStore {
	return MetricStore{redis: instance, cacheKey: cacheKey, duration: duration}
}

func (m *MetricStore) Set(key string, val string) {
	var value interface{} = val
	m.redis.Set(context.Background(),
		m.cacheKey+":"+key,
		value,
		m.duration)
}

func (m *MetricStore) ReadAll() []string {
	keysSliceCmd := m.redis.Keys(context.Background(), m.cacheKey+":"+"*")
	keys, err := keysSliceCmd.Result()

	var results []string
	if err != nil || len(keys) == 0 {
		return results
	}

	valuesSliceCmd := m.redis.MGet(context.Background(), keys...)
	for _, val := range valuesSliceCmd.Val() {
		results = append(results, fmt.Sprintf("%v", val))
	}

	return results
}
