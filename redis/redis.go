package redis

import (
	"log"
	"time"

	"AlarmService/g"

	"github.com/garyburd/redigo/redis"
)

var ConnPool *redis.Pool

func InitConnPool() {

	ConnPool = &redis.Pool{
		MaxIdle:     g.Config().RedisMaxIdle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", g.Config().RedisAddr)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: PingRedis,
	}
}

func PingRedis(c redis.Conn, t time.Time) error {
	_, err := c.Do("ping")
	if err != nil {
		log.Println("[ERROR] ping redis fail", err)
	}
	return err
}
