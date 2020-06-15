package Redis

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"time"
)

var maxIdle, maxOpen, maxLifetime = 10, 0, 30

var redisPool *redis.Pool

func CreatePool(addr string, db int, password string) (err error) {
	redisPool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxOpen,
		IdleTimeout: time.Duration(maxLifetime) * time.Minute,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, redis.DialDatabase(db), redis.DialPassword(password))
		},
	}
	conn := Conn()
	defer conn.Close()

	if r, _ := redis.String(conn.Do("PING")); r != "PONG" {
		err = errors.New("redis connect failed.")
	}
	return
}

func Conn() redis.Conn {
	return redisPool.Get()
}

func ClosePool() {
	if redisPool != nil {
		redisPool.Close()
	}
}
