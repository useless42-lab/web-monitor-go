package cache

import (
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	_ "github.com/joho/godotenv/autoload"
)

var redisClient *redis.Pool

func init() {
	// 建立连接池
	redisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,                 //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300 * time.Second, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) {
			redisConn, err := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"))
			if err != nil {
				return nil, err
			}
			return redisConn, err
		},
	}
}

func Set(key string, value string, time int) bool {
	conn := redisClient.Get()
	var err error
	if time < 0 {
		_, err = conn.Do("SET", key, value)
	} else {
		_, err = conn.Do("SET", key, value, "EX", time)
	}
	defer conn.Close()
	if err != nil {
		return false
	} else {
		return true
	}
}

func Get(key string) string {
	conn := redisClient.Get()
	result, err := redis.String(conn.Do("GET", key))
	defer conn.Close()
	if err != nil {
		return ""
	} else {
		return result
	}
}

func Del(key string) {
	conn := redisClient.Get()
	conn.Do("DEL", key)
	defer conn.Close()
}
