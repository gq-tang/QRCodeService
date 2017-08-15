package model

import (
	"QRCodeService/setting"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	client *redis.Client
}

func NewClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     setting.RedisAddr,
		Password: setting.RedisPwd,
		DB:       setting.RedisDb,
		PoolSize: setting.RedisPoolSize,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Sprint("create redis client error:", err))
	}
	return &RedisClient{client}
}

func (this *RedisClient) IsMember(key string, elem interface{}) bool {
	isMember, err := this.client.SIsMember(key+":name", elem).Result()
	if err != nil {
		return false
	}
	return isMember
}

func (this *RedisClient) SaveQrcode(key string, elem string, expire string) {
	this.client.SAdd(key+":name", elem)
	this.client.LPush(key+":destroy", elem)
	this.client.HSet(key+":expire", elem, expire)
}

func (this *RedisClient) GetExpire(key string, elem string) string {
	expire, _ := this.client.HGet(key+":expire", elem).Result()
	return expire
}

func (this *RedisClient) DelQrcode(key string, elem string) {
	this.client.SRem(key+":name", elem)
	this.client.LRem(key+":destroy", 1, elem)
	this.client.HDel(key+":expire", elem)
}

func (this *RedisClient) GetDestroy(key string) (string, time.Time) {
	val, err := this.client.LRange(key+":destroy", -1, -1).Result()
	if err != nil || len(val) < 1 {
		return "", time.Now().Add(time.Hour * 2)
	}
	expire := this.GetExpire(key+"expire", val[0])
	expiret, _ := time.Parse(time.RFC3339, expire)
	return val[0], expiret
}

func (this *RedisClient) Close() {
	this.client.Close()
}
