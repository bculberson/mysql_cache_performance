package mysql_cache_performance

import (
	"testing"

	redis "gopkg.in/redis.v3"
)

func BenchmarkRedisGet(b *testing.B) {
	key := randSeq(30)

	client := redis.NewClient(&redis.Options{Addr: "redis:6379"})

	err := client.Set(key, randSeq(30), 0).Err()
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		_, err := client.Get(key).Result()
		if err != nil {
			panic(err)
		}
	}
}

