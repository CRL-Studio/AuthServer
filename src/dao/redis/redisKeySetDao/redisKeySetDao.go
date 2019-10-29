package rediskeysetdao

import (
	redisdao "github.com/CRL-Studio/AuthServer/src/dao/redis"
	"github.com/CRL-Studio/AuthServer/src/utils/keybuilder"
)

func key(prefix string) string {
	return keybuilder.KeySet(prefix)
}

// Add a key into set
func Add(prefix string, member interface{}, ttl int) error {
	redis := redisdao.Redis()
	defer redis.Close()

	key := key(prefix)
	if err := redis.SAdd(key, member); err != nil {
		return err
	}
	if ttl > 0 {
		if err := redis.Expire(key, ttl); err != nil {
			return err
		}
	}
	return nil
}

// GetAll return all keys in set
func GetAll(prefix string) ([]string, error) {
	redis := redisdao.Redis()
	defer redis.Close()

	key := key(prefix)
	return redis.SMembers(key)
}

// Del delete all keys in set
func Del(prefix string) error {
	redis := redisdao.Redis()
	defer redis.Close()

	key := key(prefix)
	keys, err := GetAll(prefix)
	if err != nil {
		return err
	}

	keys = append(keys, key)
	if err := redis.Del(keys); err != nil {
		return err
	}
	return nil
}
