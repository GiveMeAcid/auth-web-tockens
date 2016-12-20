package redis

import "github.com/garyburd/redigo/redis"

type RedisCli struct {
	conn redis.Conn
}

func GetInstance() *RedisCli {
	return
}

func (redisCli *RedisCli) SetValue(key string, value string, expiration ...interface{}) error {
	return
}