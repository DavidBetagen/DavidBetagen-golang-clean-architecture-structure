package redis

import (
	"context"
	"crypto/tls"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
)

func NewRedis() (redis.UniversalClient, error) {
	redis, _ := NewRedisClusterClient()
	// redis, _ := NewRedisClient()

	// Set a timeout of 5 seconds
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// if err := redis.Ping(ctx).Err(); err != nil {
	// 	log.Error("Cannot ping to Redis: ", err.Error())
	// 	return nil, err
	// }

	return redis, nil
}

func NewRedisClient() (*redis.Client, error) {
	REDIS_CLIENT_ADDRESS := os.Getenv("REDIS_CLIENT_ADDRESS")
	REDIS_CLIENT_PASSWORD := os.Getenv("REDIS_CLIENT_PASSWORD")
	REDIS_CLIENT_DB, _ := strconv.Atoi(os.Getenv("REDIS_CLIENT_DB"))

	return redis.NewClient(&redis.Options{
		Addr:     REDIS_CLIENT_ADDRESS,
		Password: REDIS_CLIENT_PASSWORD,
		DB:       REDIS_CLIENT_DB, // Default DB is 0
	}), nil
}

func NewRedisClusterClient() (*redis.ClusterClient, error) {
	REDIS_SSL := strings.EqualFold(os.Getenv("REDIS_SSL"), "true")
	REDIS_CLUSTER_MAX_REDIRECTS, _ := strconv.Atoi(os.Getenv("REDIS_CLUSTER_MAX_REDIRECTS"))
	REDIS_CLUSTER_NODE := strings.Split(os.Getenv("REDIS_CLUSTER_NODE"), ",")
	REDIS_CLUSTER_USERNAME := os.Getenv("REDIS_CLUSTER_USERNAME")
	REDIS_CLUSTER_PASSWORD := os.Getenv("REDIS_CLUSTER_PASSWORD")

	// Create a custom TLS configuration for SSL
	tlsConfig := &tls.Config{
		InsecureSkipVerify: REDIS_SSL, // You may want to set this to false in a production environment
	}

	// Create a Redis Cluster client with custom options
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        REDIS_CLUSTER_NODE,
		TLSConfig:    tlsConfig,                   // Enable SSL/TLS
		MaxRedirects: REDIS_CLUSTER_MAX_REDIRECTS, // Maximum number of redirections (default is 😎
		Username:     REDIS_CLUSTER_USERNAME,
		Password:     REDIS_CLUSTER_PASSWORD,
	}), nil
}

func Subscribe(channel string) {
	redis, err := NewRedis()
	if err != nil {
		log.Error(err.Error())
	}

	subscriber := redis.Subscribe(context.Background(), channel)

	for {
		msg, err := subscriber.ReceiveMessage(context.Background())
		if err != nil {
			log.Error(err.Error())
		}

		log.Info("Lineinsider-label:: ##Redis (Subscribe) received message from " + msg.Channel)
	}
}
