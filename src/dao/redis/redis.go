package redisdao

import (
	"fmt"
	"time"

	"github.com/CRL-Studio/AuthServer/src/utils/config"
	"github.com/gomodule/redigo/redis"
)

// RedisDao is an alias of redis.Conn, for which imports redisdao to use
type RedisDao struct {
	conn redis.Conn
}

var pool *redis.Pool

func init() {
	pool = connectionPool()
}

// Redis get redis connection from connection pool
func Redis() RedisDao {
	return RedisDao{
		conn: pool.Get(),
	}
}

func connectionPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool
		MaxIdle: config.Get("redis.max_idle").(int),

		// Maximum number of connections allocated by the pool at a given timeutil.
		MaxActive: config.Get("redis.max_active").(int),

		// Close connections after remaining idle for this duration.
		IdleTimeout: time.Duration(config.Get("redis.idle_timeout").(int)) * time.Millisecond,

		// Dial is an application supplied function for creating and configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", config.Get("redis.host"), config.Get("redis.port")),
				redis.DialDatabase(config.Get("redis.database").(int)),
				redis.DialPassword(config.Get("redis.auth").(string)),
			)

			if err != nil {
				// logger
				return nil, err
			}
			return c, nil
		},

		// PING PONG test
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
