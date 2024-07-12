package main

import "github.com/redis/go-redis/v9"

func main() {
	rdb := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			// shardName => host:port
			"shard1": "localhost:7000",
			"shard2": "localhost:7001",
			"shard3": "localhost:7002",
		},
	})
}
