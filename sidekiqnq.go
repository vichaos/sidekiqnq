package sidekiqnq

import (
	"encoding/json"
	"log"
	"time"

	"gopkg.in/redis.v5"
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
func (s *Sidekiq) EnqueueJob(queueName string, klass string, args []interface{}) {
	jid, err := generateRandomString(12)
	log.Print(err)
	job := Job{klass, args, jid, true, time.Now().Unix()}
	j, err := json.Marshal(job)

	// log.Print(err)
	s.RedisClient.LPush(queueName, string(j))
}
