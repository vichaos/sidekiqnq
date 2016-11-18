## sidekiqnq (sidekiq-enqueue)
**sidekiqnq** is a mini lib that creates jobs in sidekiq protocol and push them to redis. This enables your blazing fast go web framework to work with sidekiq workers written in ruby. We built it as we are separating our monolithic Rails app into smaller pieces. 

##### Installations:
```
go get github.com/dgm59/sidekiqnq

```

##### How to use:
```go
# Create a redis connection:
s := sidekiqnq.NewSidekiqConnection(redisNamespace, redisPort, redisHost, redisPassword, redisDB)

# Enqueue a job to redis:
s.EnqueueJob("sidekiq:queue:default", "RecordWorker", []interface{}{13})
```

