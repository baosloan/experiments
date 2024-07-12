package main

import "github.com/redis/go-redis/v9"

func main() {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master-name",
		SentinelAddrs: []string{":9126", ":9127", ":9128"},
	})
}
