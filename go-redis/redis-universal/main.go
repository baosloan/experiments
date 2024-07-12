package main

import "github.com/redis/go-redis/v9"

func main() {
	// *redis.Client.
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{":6379"},
	})

	// *redis.ClusterClient.
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{":6379", ":6380"},
	})

	// *redis.FailoverClient.
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      []string{":6379"},
		MasterName: "mymaster",
	})
}
