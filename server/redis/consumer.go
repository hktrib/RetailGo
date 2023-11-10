package kvRedis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (cache *redisCache) Consumer(kvStoreClient *redis.Client) {
	subscriber := kvStoreClient.Subscribe(cache.ctx, fmt.Sprintf("%v", cache.db))

	defer subscriber.Close()

	for msg := range subscriber.Channel() {
		fmt.Printf("Received key expiry event: %s\n", msg.Payload)

		if msg.Payload == "__keyevent@0__:expired" && msg.Channel == "key" {
			fmt.Printf("Key %s has expired\n", "key")
			// Handle the expired key here.
		}
	}
}
