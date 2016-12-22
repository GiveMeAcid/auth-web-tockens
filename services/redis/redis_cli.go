package redis

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"log"
	"github.com/auth-web-tokens/services/config"
)

type RedisCli struct {
	conn redis.Conn
}

var (
	connectionPool  *redis.Pool = nil
	ExpiredKeysChan = make(chan string)
)

func init() {
	connectionPool = newPool()
}

func GetInstance() *RedisCli {
	instanceRedisCli := new(RedisCli)
	instanceRedisCli.conn = connectionPool.Get()
	if err := instanceRedisCli.conn.Err(); err != nil {
		log.Fatalln("Error connection to redis Server, error: ", err)
	}
	return instanceRedisCli
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        dial,
	}
}

func dial() (redis.Conn, error) {
	redisAddress := config.Config.RedisAddress
	if redisAddress == "" {
		log.Println("You did not provide redis_address, using default address instead")
		redisAddress = ":6379"
	}
	conn, err := redis.Dial("tcp", redisAddress)
	if err != nil {
		return nil, err
	}

	if config.Config.RedisAuth != "" {
		if _, err := conn.Do("AUTH", config.Config.RedisAuth); err != nil {
			conn.Close()
			return nil, err
		}
	}
	return conn, nil
}

func (redisCli *RedisCli) SetValue(key string, value string, expiration ...interface{}) error {
	defer redisCli.conn.Close()
	_, err := redisCli.conn.Do("SET", key, value)

	if err == nil && expiration != nil {
		redisCli.conn.Do("EXPIRE", key, expiration[0])
	}

	return err
}

func (redisCli *RedisCli) GetValue(key string) (interface{}, error) {
	defer redisCli.conn.Close()
	return redisCli.conn.Do("GET", key)
}