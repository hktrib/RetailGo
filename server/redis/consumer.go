package kv

import (
	"fmt"
)

func (cache *Cache) Consumer() {
	subscriber := cache.Client.Subscribe(cache.ctx, fmt.Sprintf("%v", cache.opts.DB))

	defer subscriber.Close()

	for msg := range subscriber.Channel() {
		fmt.Printf("Received key expiry event: %s\n", msg.Payload)

		if msg.Payload == "__keyevent@0__:expired" && msg.Channel == "key" {
			fmt.Printf("Key %s has expired\n", "key")
			// Handle the expired key here.
		}
	}
}
