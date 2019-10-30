package redisverificationdao

import (
	redisdao "github.com/CRL-Studio/AuthServer/src/dao/redis"
	"github.com/CRL-Studio/AuthServer/src/utils/config"
	"github.com/CRL-Studio/AuthServer/src/utils/keybuilder"
)

func key(account string) string {
	return keybuilder.Verification(account)
}

// New a verification cache in redis
func New(account, verification string) {
	redis := redisdao.Redis()
	defer redis.Close()

	key := key(account)
	if err := redis.SetEx(key, config.Get("redis.verification_active").(int), verification); err != nil {
		panic(err)
	}
}

// Del delete a notify cache in redis
func Del(account string) {
	redis := redisdao.Redis()
	defer redis.Close()

	key := key(account)
	if err := redis.HDel(key, account); err != nil {
		panic(err)
	}
}

// Get notify url from redis
func Get(account string) (string, error) {
	redis := redisdao.Redis()
	defer redis.Close()

	key := key(account)
	url, err := redis.Get(key)

	return url, err
}
