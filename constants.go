package sidekiqnq

import redis "gopkg.in/redis.v6"

// Sidekiq defines the struture of a sidekiq object
type Sidekiq struct {
	Namespace     string
	redisPort     string
	redisHost     string
	redisPassword string
	redisDB       int
	RedisClient   *redis.Client
}

// Job defines structure of Sidekiq job accepted in redis
type Job struct {
	Class      string        `json:"class"`
	Args       []interface{} `json:"args"`
	JID        string        `json:"jid"`
	Retry      bool          `json:"retry"`
	EnqueuedAt int64         `json:"enqueued_at"`
}
