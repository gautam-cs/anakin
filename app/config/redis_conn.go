package config

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	redisV8 "github.com/go-redis/redis/v8"
)

const defRedisDb = 10 //TODO: update

type RedisClient redisV8.Client

var redis *RedisClient

var rMutex = &sync.Mutex{}

func Redis() *RedisClient {
	return NewRedis(defRedisDb)
}

func NewRedis(redisDB int) *RedisClient {
	if redis != nil {
		return redis
	}

	rMutex.Lock()

	defer func() {
		rMutex.Unlock()
	}()

	if redis != nil {
		return redis
	}

	config := redisClientConfig()

	if redisDB < 0 {
		redisDB = defRedisDb
	}

	client := redisV8.NewClient(&redisV8.Options{
		Addr:     config.Address,
		Password: config.Password, // no password set
		DB:       redisDB,         // use default DB
	})

	r := RedisClient(*client)
	redis = &r

	return redis
}

//SetStruct store struct in redis
func (redis *RedisClient) SetStruct(key string, value interface{}, expiration time.Duration) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, e := redis.Set(context.Background(), key, p, expiration).Result()

	return e
}

//GetStruct fetch struct in redis
func (redis *RedisClient) GetStruct(key string, dest interface{}) error {
	data, err := redis.Get(context.Background(), key).Result()

	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), dest)
}

func redisClientConfig() *redisConfig {
	if IsDebugEnv() {
		return appConfig.Redis
	}

	appConfig.Redis = new(redisConfig)
	readDecodedSecret("redis", appConfig.Redis)
	return appConfig.Redis
}
