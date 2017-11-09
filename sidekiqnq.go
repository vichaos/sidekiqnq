package sidekiqnq

import (
	"encoding/json"
	"time"

	"gopkg.in/redis.v6"
)

func (s *Sidekiq) connect() {
	s.RedisClient = redis.NewClient(&redis.Options{
		Addr:     s.redisHost + ":" + s.redisPort,
		Password: s.redisPassword,
		DB:       s.redisDB,
	})
}

// NewSidekiqConnection creates a sidekiq connection
func NewSidekiqConnection(namespace string, redisPort string, redisHost string, redisPassword string, redisDB int) Sidekiq {
	s := Sidekiq{Namespace: namespace, redisPort: redisPort, redisHost: redisHost, redisPassword: redisPassword, redisDB: redisDB}
	s.connect()
	return s
}

// EnqueueJob sends
func (s *Sidekiq) EnqueueJob(queueName string, klass string, args []interface{}) (int64, error) {
	jid, err := generateRandomString(12)
	if err != nil {
		return 0, err
	}
	job := Job{klass, args, jid, true, time.Now().Unix()}
	j, err := json.Marshal(job)
	if err != nil {
		return 0, err
	}
	size, err := s.RedisClient.LPush(queueName, string(j)).Result()
	if err != nil {
		return 0, err
	}
	return size, err
}
